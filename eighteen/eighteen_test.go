package sixteen

import (
    "testing"
)

func TestEighteen(t * testing.T) {
    data := []byte("yellow submarine")
    ct, _ := EncryptCTR(data, 0)
    P(ct)
    pt, _ := DecryptCTR(ct, 0)
    P(string(pt))
}
