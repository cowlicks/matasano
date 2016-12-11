package dh

import (
	"testing"
)

func Test(t *testing.T) {
	a := Rand(g)
	b := Rand(g)
	A := Pow(g, a, p)
	B := Pow(g, b, p)
	r := Pow(A, b, p)
	s := Pow(B, a, p)
	if r.Cmp(s) != 0 {
		t.Fail()
	}
}
