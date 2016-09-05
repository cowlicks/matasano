package sixteen

import (
    "fmt"
    "errors"
    "bytes"
    "encoding/binary"
    "crypto/aes"
    "crypto/cipher"
    "../xor"
)

var key = []byte("YELLOW SUBMARINE")
var err error
var CTRError = errors.New("CTR error")
var bs = 16

func P(in ...interface{}) {
    fmt.Println(in)
}

func NonceBytes(nonce int) []byte {
    unonce64 := uint64(nonce)
    counter_buff := new(bytes.Buffer)
    counter_buff.Write(make([]byte, 8))
    err = binary.Write(counter_buff, binary.LittleEndian, unonce64)
    if err != nil {
        panic("write error")
    }
    return counter_buff.Bytes()
}

func CTRStep(block cipher.Block, nonce int) []byte {
    ctrbytes := NonceBytes(nonce)
    block.Encrypt(ctrbytes, ctrbytes)
    return ctrbytes
}

func EncryptCTR(plaintext []byte, nonce int) ([]byte, error) {
    ptlen := len(plaintext)
    block, _ := aes.NewCipher(key)
    bs := block.BlockSize()

    out := new(bytes.Buffer)
    steps := ptlen/bs
    last_block_size := bs
    if ptlen % bs != 0 {
        steps += 1
        last_block_size = ptlen % bs
    }
    for i := 0; i < steps; i++ {
        ctrbytes := CTRStep(block, nonce + i)

        ct_block := make([]byte, bs)
        var xorerr error
        if i == steps - 1 {
            ct_block, xorerr = xor.Xor(plaintext[bs*i:], ctrbytes[:last_block_size])
        } else {
            ct_block, xorerr = xor.Xor(plaintext[bs*i:bs*(i + 1)], ctrbytes)
        }
        if xorerr != nil {
            panic("xor error")
        }
        out.Write(ct_block)
    }
    return out.Bytes(), err
}

func DecryptCTR(ciphertext []byte, nonce int) ([]byte, error) {
    return EncryptCTR(ciphertext, nonce)
}
