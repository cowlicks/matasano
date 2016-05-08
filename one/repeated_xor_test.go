package one

import (
    "testing"
)

func TestRepeatedXor(t * testing.T) {
    var in1 = "Burning 'em, if you ain't quick and nimble\nI go crazy when I hear a cymbal"
    var key = "ICE"
    exp1 := "0b3637272a2b2e63622c2e69692a23693a2a3c6324202d623d63343c2a26226324272765272a282b2f20430a652e2c652a3124333a653e2b2027630c692b20283165286326302e27282f"

    res1 := RepeatedKeyXor(in1, key)
    if res1 != exp1 {
        t.Fatalf("output wrong, got %s", res1)
    }

}
