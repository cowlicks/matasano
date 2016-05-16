package one

import (
    "crypto/aes"
)

func DecryptAESECB(key, ciphertext []byte) ([]byte, error) {
    block, err := aes.NewCipher(key)
    if err != nil {
        return []byte(""), err
    }
    blocksize := block.BlockSize()
    ctsize := len(ciphertext)

    nblocks := ctsize/blocksize
    if ctsize % blocksize != 0 {
        nblocks += 1
    }

    plaintext := make([]byte, ctsize)

    for i := 0; i < nblocks; i++ {
        ll := blocksize * i
        ul := blocksize * (i + 1)
        if ul > ctsize {
            ul = ctsize
        }

        ptblock := plaintext[ll:ul]
        ctblock := ciphertext[ll:ul]
        block.Decrypt(ptblock, ctblock)
    }
    return plaintext, nil
}
