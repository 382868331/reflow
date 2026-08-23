package ansi
import "testing"
func TestTask001UppercaseTerminatorBoundary(t *testing.T){
	for _,r:=range[]rune{'@','A','Z','a','z'}{if !IsTerminator(r){t.Fatalf("%q rejected",r)}}
	for _,r:=range[]rune{'?','[','`','{'}{if IsTerminator(r){t.Fatalf("%q accepted",r)}}
}
