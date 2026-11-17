package main

import "testing"

func TestHello(t *testing.T) {
	result := hello()

	if result != "Hello world" {
		t.Errorf("expected Hello world, got %s", result)
	}
}
