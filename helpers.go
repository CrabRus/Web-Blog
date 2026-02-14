package main

import (
	"strings"
	"testing"
)

func AssertEqual(t *testing.T, expected, actual interface{}) {
	t.Helper()
	if expected != actual {
		t.Errorf("Expected %v, got %v", expected, actual)
	}
}

func AssertNoError(t *testing.T, err error) {
	t.Helper()
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}
}

func AssertError(t *testing.T, err error, expectedMsg string) {
	t.Helper()
	if err == nil {
		t.Fatalf("Expected error but got none")
	}
	if !strings.Contains(err.Error(), expectedMsg) {
		t.Errorf("Expected error containing %q, got %q", expectedMsg, err.Error())
	}
}
