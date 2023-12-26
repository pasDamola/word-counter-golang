package main

import (
	"bytes"
	"testing"
)


func TestCountWords(t *testing.T) {
	b := bytes.NewBufferString("word1 word2 word3 word4\n")

	want := 4

	got := count(b, false, false)

	if got != want {
		t.Errorf("Expected %d, got %d instead.\n", want, got)
	}
}

// TestCountLines tests the count function set to count lines
func TestCountLines(t *testing.T) {
	b := bytes.NewBufferString("word1 word2 word3\nline2\nline3 word1")

	want := 3

	got := count(b, true, false)
	if got != want {
	t.Errorf("Expected %d, got %d instead.\n", want, got)
	}
}

func TestCountBytes(t *testing.T) {
	b := bytes.NewBufferString("word1 word2 word3\nline2\nline3 word1")

	want := 35

	got := count(b, false, true)
	if got != want {
	t.Errorf("Expected %d, got %d instead.\n", want, got)
	}
}