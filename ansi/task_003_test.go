package ansi

import (
	"bytes"
	"testing"
)

func TestTask003ResetClearsRememberedStyle(t *testing.T) {
	var out bytes.Buffer
	w := Writer{Forward: &out}
	if _, err := w.Write([]byte("\x1b[31mred\x1b[0m")); err != nil {
		t.Fatal(err)
	}
	if w.LastSequence() != "" {
		t.Fatalf("last sequence=%q", w.LastSequence())
	}
	w.RestoreAnsi()
	if got := out.String(); got != "\x1b[31mred\x1b[0m" {
		t.Fatalf("output=%q", got)
	}
	if _, err := w.Write([]byte("\x1b[1mhot\x1b[0m")); err != nil {
		t.Fatal(err)
	}
	if w.LastSequence() != "" {
		t.Fatalf("second reset left=%q", w.LastSequence())
	}
}
