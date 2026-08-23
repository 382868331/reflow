package indent

import "testing"

func TestTask005NewlineResetsIndentState(t *testing.T) {
	w := NewWriter(2, nil)
	w.Write([]byte("one\n"))
	w.Write([]byte("two\nthree"))
	if got := w.String(); got != "  one\n  two\n  three" {
		t.Fatalf("got=%q", got)
	}
	w.Write([]byte("\nfour"))
	if got := w.String(); got != "  one\n  two\n  three\n  four" {
		t.Fatalf("continued=%q", got)
	}
}
