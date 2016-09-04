package sixteen

import (
    "testing"
)

func TestSeventeen(t * testing.T) {
    ct := Encrypt()
    blocks := GetBlocks(ct)
    Pb(blocks)
    OneByte()
    //OneBlock(0)
    //OneBlock(1)
    P("\n")
    //OneBlock(2)
    //OneBlock(3)
    AllBlocks()
}
