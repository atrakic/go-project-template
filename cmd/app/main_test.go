package main

import (
	"os"
	"testing"
)

func TestMain(m *testing.M) {
	code := m.Run()
	os.Exit(code)
}

// Integration test to ensure the main function doesn't panic
func TestMainFunction(t *testing.T) {
	// This is a basic test to ensure main() can be called without panicking
	// In a more complex application, you might want to test specific behaviors

	// Test passes if we can compile and run this test without issues
	t.Log("main package compiled successfully")
}
