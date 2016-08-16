package twelve

import (
    "encoding/base64"
    "../aesmodes"
)

var key, _ = aesmodes.MakeKey()
var appendme, decodeerr  = base64.StdEncoding.DecodeString("Um9sbGluJyBpbiBteSA1LjAKV2l0aCBteSByYWctdG9wIGRvd24gc28gbXkgaGFpciBjYW4gYmxvdwpUaGUgZ2lybGllcyBvbiBzdGFuZGJ5IHdhdmluZyBqdXN0IHRvIHNheSBoaQpEaWQgeW91IHN0b3A/IE5vLCBJIGp1c3QgZHJvdmUgYnkK")

func Encryptor(plaintext []byte, key []byte) []byte {
    input := append(plaintext, appendme...)
    out, _ := aesmodes.EncryptECB(key, input)
    return out
}

func KeyedEncryptor(plaintext []byte) []byte {
    return Encryptor(plaintext, key)
}

func FindBlockSize(kencryptor func([]byte) []byte) int {
    bs := 0
    for padlen := 0; bs == 0; padlen++ {
        init := len(kencryptor(make([]byte, 0)))
        longpad := make([]byte, padlen)
        final := len(kencryptor(longpad))
        bs = final - init
    }
    return bs
}

func CheckEncryptorECB(kencryptor func([]byte) []byte) bool {
    bs := FindBlockSize(kencryptor) // maybe make this an arg
    block := make([]byte, bs)
    ct := kencryptor(append(block, block...))
    for i := 0; i < bs; i++ {
        if ct[i] != ct[i + bs] {
            return false
        }
    }
    return true
}

func MakeDict(bs int) map[string][]byte {
    out := make(map[string][]byte)
    for i := 0; i < 256; i++ {
        block := make([]byte, bs)
        block[bs - 1] = byte(i)
        ct := KeyedEncryptor(block)
        out[string(block[:bs])] = ct[:bs]
    }
    return out
}
