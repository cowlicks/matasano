package one

import (
    "encoding/hex"
    "bytes"
)

func ByteXor(a []byte, b []byte) []byte {
    var buff bytes.Buffer
    if len(a) != len(b) {
        panic("arrays are not same length")
    }
    for i, v := range a {
        buff.Write([]byte{v ^ b[i]})
    }
    return buff.Bytes()
}

func VectorXor(short []byte, long []byte) []byte {
    var buff bytes.Buffer
    lshort := len(short)
    llong := len(long)
    for i := 0; i < llong; {
        for j := 0; j < lshort && i < llong; {
            buff.Write([]byte{short[j] ^ long[i]})
            j++
            i++
        }
    }
    return buff.Bytes()
}

func RepeatedKeyXor(s string, k string) string {
    return hex.EncodeToString(VectorXor([]byte(k), []byte(s)))
}
