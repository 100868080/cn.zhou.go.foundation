package main

import (
	"fmt"
	"strings"
	"testing"
)

func assertEqual(x, y int) {
	if x != y {
		panic(fmt.Sprintf("%d!=%d", x, y))
	}
}

func TestSplit(t *testing.T) {
	words := strings.Split("a:b:c", ":")
	assertEqual(len(words), 3)
}
func TestSplit(t *testing.T) {
	s, sep := "a:b:c", ":"
	words := strings.Split(s, sep)
	if got, want := len(words), 3; got != want {
		t.Error("Split(%q,%q)returned %d words,want %d", s, sep, got, want)
	}
}
