package uuid

import (
	"testing"
)

func TestGenerate(t *testing.T) {
	id, err := Generate()
	if err != nil {
		t.Fatalf("Generate() returned error: %v", err)
	}

	if id == "" {
		t.Error("Generate() returned empty string")
	}

	// Basic UUID format check (36 characters with hyphens)
	if len(id) != 36 {
		t.Errorf("Generate() returned UUID with incorrect length: got %d, want 36", len(id))
	}
}

func TestMustGenerate(t *testing.T) {
	// Test that MustGenerate doesn't panic under normal circumstances
	defer func() {
		if r := recover(); r != nil {
			t.Errorf("MustGenerate() panicked: %v", r)
		}
	}()

	id := MustGenerate()
	if id == "" {
		t.Error("MustGenerate() returned empty string")
	}

	// Basic UUID format check
	if len(id) != 36 {
		t.Errorf("MustGenerate() returned UUID with incorrect length: got %d, want 36", len(id))
	}
}

func TestGenerateUniqueness(t *testing.T) {
	// Generate multiple UUIDs and ensure they're unique
	uuids := make(map[string]bool)

	for i := 0; i < 100; i++ {
		id, err := Generate()
		if err != nil {
			t.Fatalf("Generate() returned error on iteration %d: %v", i, err)
		}

		if uuids[id] {
			t.Errorf("Generate() returned duplicate UUID: %s", id)
		}
		uuids[id] = true
	}
}
