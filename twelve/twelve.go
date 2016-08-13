package main

import (
    "fmt"
    "encoding/base64"
    "crypto/aes"
    "../aesmodes"
)

var key, _ = aesmodes.MakeKey()
var appendme, decodeerr  = base64.StdEncoding.DecodeString("Um9sbGluJyBpbiBteSA1LjAKV2l0aCBteSByYWctdG9wIGRvd24gc28gbXkgaGFpciBjYW4gYmxvdwpUaGUgZ2lybGllcyBvbiBzdGFuZGJ5IHdhdmluZyBqdXN0IHRvIHNheSBoaQpEaWQgeW91IHN0b3A/IE5vLCBJIGp1c3QgZHJvdmUgYnkK")

var aesblock, err = aes.NewCipher(key)

func P(b interface{}) {
	fmt.Println(b)
}

func Encryptor(plaintext []byte, key []byte) []byte {
    input := append(plaintext, appendme...)
    out, _ := aesmodes.EncryptECB(key, input)
    return out
}

func EncryptorInspector(encryptor func([]byte, []byte) []byte, key []byte) (int, int) {
    /*
    This gets two useful pieces of info about the encrypto, the blocksize.
    And the smallest size of input to make the output get an extra block.
    */
    inputsize := 0
    init_output_size := len(encryptor(make([]byte, inputsize), key))
    new_output_size := init_output_size
    for ;init_output_size == new_output_size;{
        inputsize++
        new_output_size = len(encryptor(make([]byte, inputsize), key))
    }
    blocksize := new_output_size - init_output_size
    appendsize := init_output_size - inputsize
    return blocksize, inputsize
}

func makeByteDictionary(prefix []byte) map[[16]byte][]byte {
    if len(prefix) < blocksize - 1 {
        padsize = blocksize - 1 - len(prefix)
        tmp = make([]byte, blocksize)
        prefix = append(make([]byte, padsize), prefix...)
    }

    out := make(map[[16]byte][]byte)
    for i := 0; i < 256; i++ {
        justshort := make([]byte, blocksize)
        ctblock := make([]byte, blocksize)
        justshort[blocksize - 1] = i
        aesblock.Encrypt(ctblock, justshort)
        out[ctblock] = justshort
    }
    return out
}

/*
func getFirstByte(inputsize) byte {
    justshort = make([]byte, blocksize - 1)
    ct = Encryptor(justshort, key)
    ct[0] = 
*/

func main() {
    blocksize, inputsize := EncryptorInspector(Encryptor, key)
    /* get first byte */

}
