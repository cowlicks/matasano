package padding_oracle

import (
	"../aesmodes"
	"../padding"
	"../util"
)

var key, _ = aesmodes.MakeKey()

var bs = 16

func Encrypt(data []byte) []byte {
	ct, _ := aesmodes.EncryptCBC(key, data)
	return ct
}

func Pb(in [][]byte) {
	for i := range in {
		util.P(in[i])
	}
}

func Oracle(input []byte) bool {
	_, err := aesmodes.DecryptCBC(key, input)
	if err == padding.InvalidPad {
		return false // bad pad
	}
	return true // good pad
}

func GetBlocks(ct []byte) [][]byte {
	out := make([][]byte, len(ct)/bs)
	for i := 0; i < len(ct)/bs; i++ {
		out[i] = ct[bs*i : bs*(i+1)]
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
		copy(out[bs*i:bs*(i+1)], blocks[i])
	}
	return out
}

func OneByte(bite int, blocknum int, solved []byte, blocks [][]byte) int {
	block := blocks[blocknum]
	orig_block := make([]byte, bs)
	copy(orig_block, block)

	for i := bite + 1; i < bs; i++ {
		block[i] = block[i] ^ solved[i]
	}
	for i := bite; i < bs; i++ {
		block[i] = block[i] ^ byte(bs-bite)
	}

	before_guess := make([]byte, bs)
	copy(before_guess, block)
	for guess := 0; guess < 256; guess++ {
		block[bite] = block[bite] ^ byte(guess)

		if Oracle(CombineBlocks(blocks[:blocknum+2])) {
			copy(block, orig_block)
			return guess
		}
		copy(block, before_guess)
	}
	copy(block, orig_block)
	panic("missed byte")
	return -1
}

func OneBlock(blocknum int, blocks [][]byte) []byte {
	block := blocks[blocknum]

	out_block := make([]byte, bs)

	orig_block := make([]byte, bs)
	copy(orig_block, block)

	for bite := bs - 1; bite >= 0; bite-- {
		ptbyte := OneByte(bite, blocknum, out_block, blocks)
		out_block[bite] = byte(ptbyte)
	}
	return out_block
}

func LastBlock(blocks [][]byte) []byte {
	block := blocks[len(blocks)-2]

	out := make([]byte, bs)

	orig_block := make([]byte, bs)
	copy(orig_block, block)

	result := 1

	block[bs-1] = block[bs-1] ^ byte(1)
	pre_guess := make([]byte, bs)
	copy(pre_guess, block)
	for guess := 1; guess <= bs; guess++ {
		copy(block, pre_guess)
		if guess == 1 {
			continue
		}
		block[bs-1] = block[bs-1] ^ byte(guess)
		if Oracle(CombineBlocks(blocks)) {
			result = guess
			break
		}
	}

	if result == 0 {
		return out
	}
	for i := 0; i < result; i++ {
		out[bs-1-i] = byte(result)
	}

	copy(block, orig_block)
	for bite := bs - result - 1; bite >= 0; bite-- {
		ptbyte := OneByte(bite, len(blocks)-2, out, blocks)
		out[bite] = byte(ptbyte)

	}
	return out
}

func PaddingOracleDecrypt(ct []byte) []byte {
	blocks := GetBlocks(ct)
	out := make([]byte, len(ct)-bs)

	for i := range blocks {
		var out_block []byte
		if i == len(blocks)-2 {
			out_block = LastBlock(blocks)
			var err error
			out_block, err = padding.UnPad(bs, out_block)
			if err != nil {
				panic("???")
			}
		} else if i == len(blocks)-1 {
			continue
		} else {
			out_block = OneBlock(i, blocks)
		}
		out = append(out, out_block...)
	}
	return out
}
