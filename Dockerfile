ARG GO_VERSION=1.24-alpine

FROM --platform=$BUILDPLATFORM golang:${GO_VERSION} AS builder
WORKDIR /src
RUN apk --update add ca-certificates

# This is the architecture you're building for, which is passed in by the builder.
# Placing it here allows the previous steps to be cached across architectures.
ARG TARGETARCH
ARG GOARCH=$TARGETARCH

ARG CGO_ENABLED=0
ARG GOCACHE=/root/.cache/go-build
ARG GOMODCACHE=/root/.cache/go-mod

# Dependency management
COPY go.mod go.sum ./

# Download dependencies with cache mount
RUN --mount=type=cache,target=$GOMODCACHE \
    go mod download -x

# Copy source code
COPY . .

# Run vet and test with cache mounts
RUN --mount=type=cache,target=$GOCACHE \
    --mount=type=cache,target=$GOMODCACHE \
    go vet ./... && go test ./...

# Build the application with cache mounts
RUN --mount=type=cache,target=$GOCACHE \
    --mount=type=cache,target=$GOMODCACHE \
    go build -o /bin/app ./cmd/app

# Create a non-privileged user that the app will run under.
ARG UID=10001
RUN adduser \
    --disabled-password \
    --gecos "" \
    --home "/nonexistent" \
    --shell "/sbin/nologin" \
    --no-create-home \
    --uid "${UID}" \
    nonroot

FROM scratch AS final
LABEL org.opencontainers.image.source=https://github.com/atrakic/go-project-template.git
COPY --from=builder /etc/passwd /etc/passwd
COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/ca-certificates.crt
COPY --from=builder /bin/app /app
USER nonroot
ENTRYPOINT ["/app"]
