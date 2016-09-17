package twentyfive

import (
    "io/ioutil"
    "encoding/base64"
    "../mersenne"
)

var data = LoadData()
var key = []byte("yo")
var ciphertext = mersenne.Encrypt(key, data)

func LoadData() []byte {
    filename := "data.txt"
    filebytes, _ := ioutil.ReadFile(filename)
    data := make([]byte, base64.StdEncoding.DecodedLen(len(filebytes)))
    datalen, _ := base64.StdEncoding.Decode(data, filebytes)
    data = data[:datalen]
    return data
}
func EditByte(newbyte byte, offset int) []byte {
    if offset > len(ciphertext) {
        panic("offset bad")
    }
    return mersenne.Edit(key, ciphertext, newbyte, offset)
}

func Edit(newbytes []byte, offset int) []byte {
    out := make([]byte, len(ciphertext))
    for i := range newbytes {
        out = EditByte(newbytes[i], offset + i)
    }
    return out
}

func CrackByte(offset int) (byte, byte) {
    out := EditByte(byte('a'), offset)
    streambyte := out[offset] ^ byte('a')
    ptbyte := ciphertext[offset] ^ streambyte
    return ptbyte, streambyte
}

func CrackText() []byte {
    out := make([]byte, len(ciphertext))
    for i := range ciphertext {
        ptbyte, _ := CrackByte(i)
        out[i] = ptbyte
    }
    return out
}
