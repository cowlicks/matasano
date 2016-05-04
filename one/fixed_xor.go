package one

import (
    "encoding/hex"
)

func BinXor(b1 []byte, b2 []byte) []byte {
    out := make([]byte, len(b1))
    for i := range b1 {
        out[i] = b1[i] ^ b2[i]
        }
    return out
}

func Xor(s1 string, s2 string) string {
    b1, _ := hex.DecodeString(s1)
    b2, _ := hex.DecodeString(s2)
    bin_out := BinXor(b1, b2)
    hex_out := hex.EncodeToString(bin_out)
    return hex_out
}
