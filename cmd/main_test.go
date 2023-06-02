package main

import (
	"testing"
)

func TestHelloWorld(t *testing.T) {
	expected := "Hello world"
	if actual := HelloWorld(); actual != expected {
		t.Errorf("HelloWorld isn't right, gave %q but it should have been %q", actual, expected)
	}
}
