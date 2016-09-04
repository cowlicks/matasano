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

func P(in ...interface{}) {
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
        copy(out[bs*i:bs*(i + 1)], blocks[i])
    }
    return out
}

func OneBlock(blocknum int, blocks [][]byte) []byte {
    block := blocks[blocknum]

    out_block := make([]byte, bs)

    orig_block := make([]byte, bs)
    copy(orig_block, block)

    for bite := bs - 1; bite >= 0; bite-- {
        // add already found bytes
        for i := bite + 1; i < bs; i++ {
            block[i] = block[i] ^ out_block[i]
        }
        // add pad
        pad := bs - bite
        for i := bite; i < bs; i++ {
            block[i] = block[i] ^ byte(pad)
        }

        before_guess := make([]byte, bs)
        copy(before_guess, block)
        for guess := 0; guess < 256; guess++ {

            block[bite] = block[bite] ^ byte(guess)

            if Oracle(CombineBlocks(blocks[:blocknum + 2])) {
                out_block[bite] = byte(guess)
                break
            }
            copy(block, before_guess)
        }
        copy(block, orig_block)
    }
    return out_block
}

func LastBlock(blocks [][]byte) []byte {
    block := blocks[len(blocks) - 2]

    out := make([]byte, bs)

    orig_block := make([]byte, bs)
    copy(orig_block, block)

    pad := 1
    result := pad

    block[bs - 1] = block[bs - 1] ^ byte(pad)
    for guess := 0; guess < bs; guess++ {
        if guess == 1 {
            continue
        }
        block[bs - 1] = block[bs - 1] ^ byte(guess)
        if Oracle(CombineBlocks(blocks)) {
            result = guess
            break
        }
    }

    if result == 0 {
        P("padd is zero")
        return out
    }

    for i := 0; i < result; i++ {
        out[bs - 1 - i] = byte(result)
    }

    copy(block, orig_block)
    for bite := bs - result - 1; bite >= 0; bite-- {
        // add already found bytes
        for i := bite + 1; i < bs; i++ {
            block[i] = block[i] ^ out[i]
        }
        // add pad
        pad := bs - bite
        P("pad", pad)
        for i := bite; i < bs; i++ {
            block[i] = block[i] ^ byte(pad)
        }

        before_guess := make([]byte, bs)
        copy(before_guess, block)
        for guess := 0; guess < 256; guess++ {

            block[bite] = block[bite] ^ byte(guess)

            if Oracle(CombineBlocks(blocks)) {
                out[bite] = byte(guess)
                P("guess", string(block[bite]))
                break
            }
            copy(block, before_guess)
        }
        copy(block, orig_block)
    }
    P(out)
    return out
}

func AllBlocks() {
    ct := Encrypt()
    blocks := GetBlocks(ct)
    out := make([]byte, len(ct))

    Pb(blocks)
    for i := range blocks {
        out_block := make([]byte, bs)
        if i == len(blocks) - 2 {
            out_block = LastBlock(blocks)
        } else if i == len(blocks) - 1 {
            continue
        } else {
            out_block = OneBlock(i, blocks)
        }
        copy(out[bs*i:bs*(i + 1)], out_block)
    }
}
