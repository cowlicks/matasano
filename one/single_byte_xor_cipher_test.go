package one

import (
    "testing"
    "encoding/hex"
)

var in = "1b37373331363f78151b7f2b783431333d78397828372d363c78373e783a393b3736"

func TestCipher(t * testing.T) {
    cc := MakeCharCount(in)
    xorde := cc.order[0]
    otp, _ := Xor(xorde, hex.EncodeToString([]byte{' '}))
    password := ""
    for i := 0; i < len(in); i += 2 {
        password += otp
    }
    out_hex, _ := Xor(password, in)
    out_bytes, _ := hex.DecodeString(out_hex)
    out_string := string(out_bytes)
    if out_string != "Cooking MC's like a pound of bacon" {
        t.Fatalf(out_string)
    }
}
