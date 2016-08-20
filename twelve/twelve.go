package twelve

import (
    "encoding/base64"
    "../aesmodes"
)

var key, _ = aesmodes.MakeKey()
var Target, decodeerr  = base64.StdEncoding.DecodeString("Um9sbGluJyBpbiBteSA1LjAKV2l0aCBteSByYWctdG9wIGRvd24gc28gbXkgaGFpciBjYW4gYmxvdwpUaGUgZ2lybGllcyBvbiBzdGFuZGJ5IHdhdmluZyBqdXN0IHRvIHNheSBoaQpEaWQgeW91IHN0b3A/IE5vLCBJIGp1c3QgZHJvdmUgYnkK")

func Encryptor(plaintext []byte) []byte {
    input := append(plaintext, Target...)
    out, _ := aesmodes.EncryptECB(key, input)
    return out
}

func finder(encryptor func([]byte) []byte) (int, int) {
    bs := 0
    padlen := 0
    init := len(encryptor(make([]byte, 0)))
    for ; bs == 0; padlen++ {
        longpad := make([]byte, padlen)
        final := len(encryptor(longpad))
        bs = final - init
    }
    return bs, init - (padlen - 1)
}

func FindBlockSize(encryptor func([]byte) []byte) int {
    bs, _ := finder(encryptor)
    return bs
}

func FindTargetSize(encryptor func([]byte) []byte) int {
    _, tlen := finder(encryptor)
    return tlen
}

func CheckEncryptorECB(encryptor func([]byte) []byte) bool {
    bs := FindBlockSize(encryptor) // maybe make this an arg
    block := make([]byte, bs)
    ct := encryptor(append(block, block...))
    for i := 0; i < bs; i++ {
        if ct[i] != ct[i + bs] {
            return false
        }
    }
    return true
}

func MakePadDict(bs int, pad []byte) map[string]byte {
    out := make(map[string]byte)
    padlen := len(pad)

    for i := 0; i < 256; i++ {
        b := byte(i)

        block := make([]byte, padlen + 1)
        copy(block[:padlen], pad)
        block[padlen] = b
        ct := Encryptor(block)
        out[string(ct[padlen + 1 - bs : padlen + 1])] = b
    }
    return out
}

func MakeCTDict(bs int) map[int][]byte {
    out := make(map[int][]byte)

    for i := 0; i < bs; i++ {
        out[i] = Encryptor(make([]byte, i))
    }
    return out
}


func OneShort() byte {
    bs := FindBlockSize(Encryptor)
    ctdict := MakeCTDict(bs)

    pad := make([]byte, bs - 1)
    paddict := MakePadDict(bs, pad)

    res := paddict[string(ctdict[bs - 1])]

    return res
}

func OneBlock() []byte {
    bs := FindBlockSize(Encryptor)
    padlen := bs - 1
    ctdict := MakeCTDict(bs)

    pad := make([]byte, padlen)

    for i := 0; i < bs - 1; i++ {
        paddict := MakePadDict(bs, pad)
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


func WholeThing(encryptor func([]byte) []byte) []byte {
    ptlen := FindTargetSize(encryptor)
    bs := FindBlockSize(encryptor)
    ctdict := MakeCTDict(bs)

    pad := make([]byte, bs - 1)

    for i := 0; i < ptlen; i++ {
        bn := i / bs
        if (i != 0) && (0 == i % (bs - 1)) {
            pad = append(make([]byte, bs), pad...)
        }
        paddict := MakePadDict(bs, pad)
        copy(pad[:len(pad) - 1], pad[1:])
        if val, ok := ctdict[bs - 1 - (i % bs)]; ok {
            tomatch := string(val)[bs*bn : bs * (bn + 1)]
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
