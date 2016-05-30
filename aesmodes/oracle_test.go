package aesmodes

import (
    "testing"
    "../util"
)

func TestOracle(t * testing.T) {
    MakeKey()
    c, _ := EncryptionOracle(make([]byte, 5))
    util.P(c)
}
