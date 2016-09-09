package mersenne

import (
	"testing"
)

func TestMersenne(t *testing.T) {
	mt := NewMersenne19937(666)
	for i := 0; i < 9999; i++ {
		mt.Next()
	}
}

func TestTwentyTwo(t *testing.T) {
	MTFromTime()
    Crack(MakeMTNumberAtRandomTime())
}
