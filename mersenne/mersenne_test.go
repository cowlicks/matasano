package mersenne

import (
	"testing"
    u "../util"
)

/*
func TestMersenne(t *testing.T) {
	mt := NewMersenne19937(5489)
	for i := 0; i < 3; i++ {
		util.P(mt.Next())
	}
}

func TestTwentyTwo(t *testing.T) {
	MTFromTime()
//    Crack(MakeMTNumberAtRandomTime())
}
*/

func TestCloneMT(t *testing.T) {
    mt := NewMersenne19937(5489)
    mt_data := make([]uint32, mt.n)
    for i := uint32(0); i < mt.n; i++ {
        mt_data[i] = mt.Next()
    }
    copied_mt := CloneMT19937(mt_data)
    new_data := make([]uint32, copied_mt.n)
    for i := uint32(0); i < copied_mt.n; i++ {
        new_data[i] = copied_mt.Next()
    }
    u.P(mt_data[:5])
    u.P(new_data[:5])
}
