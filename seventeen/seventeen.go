package sixteen

import (
    "encoding/base64"
    //"bytes"
    "fmt"
    "../padding"
    "../aesmodes"
)

var key, _ = aesmodes.MakeKey()

var data, _ = base64.StdEncoding.DecodeString("MDAwMDAwTm93IHRoYXQgdGhlIHBhcnR5IGlzIGp1bXBpbmc=")

var bs = 16

func P(in interface{}) {
    fmt.Println(in)
}

func Encrypt() []byte {
    ct, _ := aesmodes.EncryptCBC(key, data)
    return ct
}

func Oracle(input []byte) bool {
    _, err := aesmodes.DecryptCBC(key, input)
    if err == padding.InvalidPad {
        return false  // bad pad
    }
    return true  // good pad
}

func LastByte() {
    ct := Encrypt()
    ctlen := len(ct)
    ci := ctlen - (2*bs)
    cibyte := ct[ci]
    for i := 0; i < 256; i++ {
        newbyte := cibyte ^ byte(i) ^ byte(1)
        ct[ci] = newbyte
        if Oracle(ct[:ci +  bs]) {
            P("gotit")
            P(i)
        }
    }
}
