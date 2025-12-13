package greeting

import "testing"

func TestHello(t *testing.T) {
	want := "Hello, world."
	got := Hello()

	if got != want {
		t.Errorf("Hello() = %q, want %q", got, want)
	}
}

func TestHelloWithName(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected string
	}{
		{
			name:     "empty name",
			input:    "",
			expected: "Hello, world.",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := HelloWithName(tt.input)
			if got != tt.expected {
				t.Errorf("HelloWithName(%q) = %q, want %q", tt.input, got, tt.expected)
			}
		})
	}
}
