package twelve

import (
//    "fmt"
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

func MakeDict(bs int) map[string]byte {
    out := make(map[string]byte)
    for i := 0; i < 256; i++ {
        b := byte(i)
        block := make([]byte, bs)
        block[bs - 1] = b
        ct := KeyedEncryptor(block)
        out[string(ct[:bs])] = b
    }
    return out
}

func MakeDictWithPad(bs int, pad []byte) map[string]byte {
    out := make(map[string]byte)
    padlen := len(pad)

    for i := 0; i < 256; i++ {
        b := byte(i)

        block := make([]byte, padlen + 1)
        copy(block[:padlen], pad)
        block[padlen] = b
        ct := KeyedEncryptor(block)
        out[string(ct[:bs])] = b
    }
    return out
}

func MakeCTDict(bs int) map[int][]byte {
    out := make(map[int][]byte)

    for i := 0; i < bs - 1; i++ {
        out[i] = KeyedEncryptor(make([]byte, i))
    }
    return out
}


func OneShort() byte {
    bs := FindBlockSize(KeyedEncryptor)
    dict := MakeDict(bs)

    short := make([]byte, bs - 1)
    ct := KeyedEncryptor(short)
    return dict[string(ct[bs:])]
}

/*
func All() []byte {
    bs := FindBlockSize(KeyedEncryptor)
    ptlen := len(appendme)
    initct := KeyedEncryptor(make([]byte, 0))
    initctlen := len(initct)

    out := make([]byte, ptlen)
    buf := make([]byte, bs)

    for i := 0; i < ptlen; i++ {
        dict := MakeDictWithPad(bs, buf[:bs - 1])
        ct := KeyedEncryptor(buf[:bs - 1])
        res := dict[string(ct[:bs])]
        out[i] = res
        buf[bs - 1] = res
        fmt.Println(string(ct))
        copy(buf[:bs - 1], buf[1:])
    }
    return out
}
*/
