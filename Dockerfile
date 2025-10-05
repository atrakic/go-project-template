FROM golang:1.24 AS builder
WORKDIR /go/src/app
COPY . .

RUN useradd -u 1001 nonroot

RUN go mod download
RUN go vet -v
RUN go test -v

RUN CGO_ENABLED=0 go build -o /go/bin/app

FROM scratch AS final
COPY --from=builder /etc/passwd /etc/passwd
COPY --from=builder /go/bin/app /
USER nonroot
ENTRYPOINT ["/app"]
