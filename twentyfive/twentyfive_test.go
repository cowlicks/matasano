package twentyfive

import (
    "testing"
)

func TestTwentyFive(t *testing.T) {
    data := LoadData()
    CrackByte(0)
    out := CrackText()
    for i, e := range out {
        if data[i] != e {
            t.Fail()
        }
    }
}
