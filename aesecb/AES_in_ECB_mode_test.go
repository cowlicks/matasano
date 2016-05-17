package aesecb

import (
//    "crypto/cipher"
    "testing"
    "io/ioutil"
    "encoding/base64"
)

func TestThis(t * testing.T) {
    key := []byte("YELLOW SUBMARINE")

    file_bytes, _ := ioutil.ReadFile("../data/data7.txt")

    decodelen := base64.StdEncoding.DecodedLen(len(file_bytes))
    ciphertext := make([]byte, decodelen)
    ctlen, _ := base64.StdEncoding.Decode(ciphertext, file_bytes)
    ciphertext = ciphertext[:ctlen]

    plaintext, err := DecryptAESECB(key, ciphertext)
    if err != nil {
        t.Fatal()
    }
    exp := "I'm back and I'm ringin' the bell"
    if string(plaintext[:len(exp)]) != exp {
        t.Fatal()
    }
}
