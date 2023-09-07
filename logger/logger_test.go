package logger

import "testing"

func TestNewLogger(t *testing.T) {
	_, err := New()
	if err != nil {
		t.Fatalf("failed to create a logger: %s", err.Error())
	}
}
