package aesmodes

import (
	"../padding"
	"crypto/aes"
)

func DecryptECB(key, ciphertext []byte) ([]byte, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return []byte(""), err
	}
	blocksize := block.BlockSize()
	ctsize := len(ciphertext)

	nblocks := ctsize / blocksize
	if ctsize%blocksize != 0 {
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
	plaintext, _ = padding.UnPad(blocksize, plaintext)
	return plaintext, nil
}

func EncryptECB(key, plaintext []byte) ([]byte, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return []byte(""), err
	}
	blocksize := block.BlockSize()
	plaintext, _ = padding.Pad(blocksize, plaintext)
	ptsize := len(plaintext)

	nblocks := ptsize / blocksize
	if ptsize%blocksize != 0 {
		nblocks += 1
	}

	ciphertext := make([]byte, ptsize)

	for i := 0; i < nblocks; i++ {
		ll := blocksize * i
		ul := blocksize * (i + 1)
		if ul > ptsize {
			ul = ptsize
		}

		ctblock := ciphertext[ll:ul]
		ptblock := plaintext[ll:ul]
		block.Encrypt(ctblock, ptblock)
	}
	return ciphertext, nil
}
