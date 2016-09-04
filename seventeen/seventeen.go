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

func Pb(in [][]byte) {
    for i := range in {
        P(in[i])
    }
}

func Oracle(input []byte) bool {
    _, err := aesmodes.DecryptCBC(key, input)
    if err == padding.InvalidPad {
        return false  // bad pad
    }
    return true  // good pad
}

func GetBlocks(ct []byte) [][]byte {
    out := make([][]byte, len(ct)/bs)
    for i := 0; i < len(ct) / bs; i++ {
        out[i] = ct[bs*i:bs*(i + 1)]
    }
    return out
}

func CombineBlocks(blocks [][]byte) []byte {
    ctlen := 0
    for i := range blocks {
        ctlen += len(blocks[i])
    }
    out := make([]byte, ctlen)
    for i := range blocks {
        copy(out[bs*i:bs*(i + 1)], blocks[i][:])
    }
    return out
}

func OneByte() {
    ct := Encrypt()
    //ctlen := len(ct)
    blocks := GetBlocks(ct)
    c1 := blocks[1]
    //c2 := blocks[2]
    for guess := 0; guess < 256; guess++ {
        c1[bs - 1] = c1[bs - 1] ^ byte(1) ^ byte(guess)
        if Oracle(CombineBlocks(blocks[:3])) {
            P("got it")
            P(guess)
        }
    }
}

func OneBlock(blocknum int) []byte {
    ct := Encrypt()
    blocks := GetBlocks(ct)
    c1 := blocks[blocknum]
    out := make([]byte, bs)
    orig := make([]byte, bs)
    copy(orig, c1)
    for bite := bs - 1; bite >= 0; bite-- {
        copy(c1, orig)
        // add pad
        for i := bite; i < bs; i++ {
            c1[i] = c1[i] ^ byte(bs - bite)
        }

        // add already found bytes
        for i := bite + 1; i < bs; i++ {
            c1[i] = c1[i] ^ out[i]
        }

        before_guess := make([]byte, bs)
        copy(before_guess, c1)
        for guess := 0; guess < 256; guess++ {
            copy(c1, before_guess)
            c1[bite] = c1[bite] ^ byte(guess)
            if Oracle(CombineBlocks(blocks[:blocknum + 2])) {
                out[bite] = byte(guess)
                break
            }
        }
    }
    P(string(out))
    return out
}

func OneBlock2(blocknum int, blocks [][]byte) []byte {
    block := blocks[blocknum]

    out_block := make([]byte, bs)

    orig_block := make([]byte, bs)
    copy(orig_block, block)

    for bite := bs - 1; bite >= 0; bite-- {
        // add pad
        for i := bite; i < bs; i++ {
            block[i] = block[i] ^ byte(bs - bite)
        }
        // add already found bytes
        for i := bite + 1; i < bs; i++ {
            block[i] = block[i] ^ out_block[i]
        }

        before_guess := make([]byte, bs)
        copy(before_guess, block)
        for guess := 0; guess < 256; guess++ {
            copy(block, before_guess)
            block[bite] = block[bite] ^ byte(guess)
            if Oracle(CombineBlocks(blocks[:blocknum + 2])) {
                P("gotem")
                P(bite)
                P(guess)
                out_block[bite] = byte(guess)
                break
            }
            copy(block, before_guess)
        }
        copy(block, orig_block)
    }
    return out_block
}

func AllBlocks() {
    ct := Encrypt()
    blocks := GetBlocks(ct)
    out := make([]byte, len(ct))

    for i := range blocks {
        if i == 3 || i == 1 || i == 0 {
            continue
        }
        P(i)
        out_block := OneBlock2(i, blocks)
        P(string(out_block))
        copy(out[bs*i:bs*(i + 1)], out_block)
    }
}
