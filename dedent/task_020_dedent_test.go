package dedent

import "testing"

func TestTaskReflow020Primary(t *testing.T) {
 input := "    alpha\n  beta"
 if got := String(input); got != "  alpha\nbeta" {
  t.Fatalf("dedent=%q", got)
 }
}

func TestTaskReflow020Boundary(t *testing.T) {
 input := "		left\n	right"
 if got := String(input); got != "	left\nright" {
  t.Fatalf("tab dedent=%q", got)
 }
}
