package one

import (
    "errors"
    "encoding/hex"
)

func BinXor(b1 []byte, b2 []byte) []byte {
    out := make([]byte, len(b1))
    for i := range b1 {
        out[i] = b1[i] ^ b2[i]
        }
    return out
}

func Xor(s1 string, s2 string) (string, error) {
    b1, e1 := hex.DecodeString(s1)
    if e1 != nil {
        return "", errors.New("String 1 badsize")
    }

    b2, e2 := hex.DecodeString(s2)
    if e2 != nil {
        return "", errors.New("String 2 badsize")
    }

    bin_out := BinXor(b1, b2)
    hex_out := hex.EncodeToString(bin_out)
    return hex_out, nil
}
