package main

import "testing"

func TestSoma(t *testing.T) {
	total := Soma(5, 5)
	if total != 10 {
		t.Errorf("Soma was incorrect, got: %d, want: %d.", total, 10)
	}
}
