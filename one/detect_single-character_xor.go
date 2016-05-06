package one

import (
    "fmt"
    "encoding/hex"
)

func Check(e error) {
    if e != nil {
        panic(e)
    }
}

func CalcXorByte(s string) string {
    cc := *MakeCharCount(s)
    hex_max := cc.order[0]
    hex_space := hex.EncodeToString([]byte(" "))
    hex_x, _ := Xor(hex_space, hex_max)
    return hex_x
}

func XorWithSingleByte(s string, char string) string {
    // the 'byte' should be 2 char hex string
    slen := len(s)
    char_arr := ""
    for i := 0; i < slen; i += 2 {
        char_arr += char
    }
    out, err := Xor(s, char_arr)
    Check(err)
    return out
}

func UnXor(s string) string {
    xval := CalcXorByte(s)
    return XorWithSingleByte(s, xval)
}


func PrintHex(s string) {
    b, err := hex.DecodeString(s)
    Check(err)
    fmt.Println(string(b))
}

