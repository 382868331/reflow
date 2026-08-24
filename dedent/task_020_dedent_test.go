package dedent

import "testing"

func TestTaskReflow020Primary(t *testing.T) {
 input := "    alpha\n  beta"
 if got := String(input); got != "  alpha\nbeta" {
  t.Fatalf("dedent=%q", got)
 }
}
