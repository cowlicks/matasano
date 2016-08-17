package twelve

import (
    "fmt"
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

func MakePadDict(pad []byte) map[string]byte {
    out := make(map[string]byte)
    padlen := len(pad)

    for i := 0; i < 256; i++ {
        b := byte(i)

        block := make([]byte, padlen + 1)
        copy(block[:padlen], pad)
        block[padlen] = b
        ct := KeyedEncryptor(block)
        out[string(ct[:padlen + 1])] = b
    }
    return out
}

func MakeCTDict(bs int) map[int][]byte {
    out := make(map[int][]byte)

    for i := 0; i < bs; i++ {
        out[i] = KeyedEncryptor(make([]byte, i))
    }
    return out
}


func OneShort() byte {
    bs := FindBlockSize(KeyedEncryptor)
    ctdict := MakeCTDict(bs)

    pad := make([]byte, bs - 1)
    paddict := MakePadDict(pad)

    res := paddict[string(ctdict[bs - 1])]

    return res
}

func OneBlock() []byte {
    bs := FindBlockSize(KeyedEncryptor)
    padlen := bs - 1
    ctdict := MakeCTDict(bs)

    pad := make([]byte, padlen)

    for i := 0; i < bs - 1; i++ {
        paddict := MakePadDict(pad)
        copy(pad[:padlen - 1], pad[1:])
        if val, ok := ctdict[bs - 1 - i]; ok {
            if r, ok2 := paddict[string(val)[:bs]]; ok2 {
                pad[padlen - 1] = r
            } else {
                panic("Error matiching ciphertext")
            }
        } else {
            panic("block index error")
        }
    }
    return pad
}


func WholeThing() []byte {
    ptlen := len(appendme)
    bs := FindBlockSize(KeyedEncryptor)
    ctdict := MakeCTDict(bs)

    pad := make([]byte, bs - 1)

    for i := 0; i < ptlen; i++ {
        bn := i / bs
        if (i != 0) && (0 == i % (bs - 1)) {
            fmt.Println("appending to pad")
            pad = append(make([]byte, bs), pad...)
        }
        paddict := MakePadDict(pad)
        fmt.Println(string(pad))
        copy(pad[:len(pad) - 1], pad[1:])
        if val, ok := ctdict[bs - 1 - (i % bs)]; ok {
            tomatch := string(val)[:bs * (bn + 1)]
            if r, ok2 := paddict[tomatch]; ok2 {
                pad[len(pad) - 1] = r
            } else {
                panic("Error matiching ciphertext")
            }
        } else {
            panic("block index error")
        }
    }
    return pad
}
