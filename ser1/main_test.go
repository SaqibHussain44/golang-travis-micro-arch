package main

import (
	"testing"
)

func TestGetTemp(t *testing.T) {
	x := GetTemp()
	if x != 100 {
		t.Error("Expected 1 got", x)
	}
}