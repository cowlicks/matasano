package sixteen

import (
    "testing"
)

func TestSeventeen(t * testing.T) {
    ct := Encrypt()
    blocks := GetBlocks(ct)
    P(Oracle(CombineBlocks(blocks)))
    Pb(blocks)
    P("\n")
    //OneBlock(2)
    //OneBlock(3)
    AllBlocks()
}
