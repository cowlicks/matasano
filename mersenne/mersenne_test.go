package mersenne

import (
	"testing"
    //"../util"
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
    //mt.Next()
//    copied_mt := CloneMT19937(mt)
    CloneMT19937(mt)
 //   util.P(copied_mt.Next())
}
