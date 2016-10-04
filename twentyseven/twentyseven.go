package thepersonuwer

import (
    "../xor"
    "../padding"
    "errors"
    "crypto/aes"
)

// Pads input with pkcs #7 to 16 byte blocksize.
// key is used as the IV and IV is not included.
func EncryptCBC(key, plaintext []byte) ([]byte, error) {
    errout := []byte("")
    block, err := aes.NewCipher(key)
    if err != nil {
        return errout, err
    }
    bs := block.BlockSize()
    plaintext, err = padding.Pad(bs, plaintext)
    if err != nil {
        return errout, err
    }

    iv := key
    ciphertext := make([]byte, len(plaintext))

    nblocks := len(plaintext)/bs

    xorblock, err := xor.Xor(iv, plaintext[:bs])
    if err != nil {
        return errout, err
    }
    for i := 1; i < nblocks; i++ {
        ciphertextblock := ciphertext[bs*i:bs*(i + 1)]
        block.Encrypt(ciphertextblock, xorblock)
        if (i != nblocks - 1) {
            plaintextblock := plaintext[bs*i: bs*(i + 1)]
            xorblock, err = xor.Xor(ciphertextblock, plaintextblock)
            if err != nil {
                return errout, err
            }
        }
    }
    return ciphertext, nil
}

// Inverse of EncryptCBC.
func DecryptCBC(key, ciphertext []byte) ([]byte, error) {
    errout := []byte("")
    block, err := aes.NewCipher(key)
    if err != nil {
        return errout, err
    }
    bs := block.BlockSize()

    iv := key

    ctsize := len(ciphertext)
    nblocks := ctsize / bs

    if ctsize % bs != 0 {
        return errout, errors.New("Ciphertext not an integer multiple of blocksize")
    }

    xorblock := make([]byte, bs)
    decblock := make([]byte, bs)
    plaintext := make([]byte, ctsize)

    for i := 0; i < nblocks ; i++ {
        block.Decrypt(decblock, ciphertext[bs*i:bs*(i + 1)])
        if i == 0 {
            xorblock, err = xor.Xor(iv, decblock)
        } else {
            xorblock, err = xor.Xor(ciphertext[bs*(i - 1):bs*i], decblock)
        }
        if err != nil {
            return errout, err
        }
        copy(plaintext[bs*i:bs*(i + 1)], xorblock)
    }
    plaintext, err = padding.UnPad(bs, plaintext)
    if err != nil {
        return errout, err
    }
    return plaintext, nil
}
