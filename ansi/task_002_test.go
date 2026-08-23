package ansi

import "testing"

func TestTask002PrintableWidthExcludesTerminators(t *testing.T) {
	for s, w := range map[string]int{"\x1b[31mred\x1b[0m": 3, "\x1b[1m红\x1b[0m": 2, "plain": 5} {
		if got := PrintableRuneWidth(s); got != w {
			t.Fatalf("%q width=%d want=%d", s, got, w)
		}
	}
	var b Buffer
	b.WriteString("\x1b[32mok\x1b[0m")
	if b.PrintableRuneWidth() != 2 {
		t.Fatalf("buffer width=%d", b.PrintableRuneWidth())
	}
}
