package aesmodes

import (
    "testing"
)

func TestOracle(t * testing.T) {
    MakeKey()
    c, _ := OracleEncryptor(make([]byte, 300))
    OracleDetector(c)
}
