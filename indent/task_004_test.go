package indent

import "testing"

func TestTask004FirstAndLaterLinesIndented(t *testing.T) {
	if got := String("one\ntwo", 2); got != "  one\n  two" {
		t.Fatalf("got=%q", got)
	}
	if got := String("x", 1); got != " x" {
		t.Fatalf("single=%q", got)
	}
	if got := String("a\nb\nc", 1); got != " a\n b\n c" {
		t.Fatalf("three=%q", got)
	}
}
