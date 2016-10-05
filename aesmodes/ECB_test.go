package aesmodes

import (
	"../util"
	"encoding/base64"
	"io/ioutil"
	"testing"
)

func TestECB(t *testing.T) {
	key := []byte("YELLOW SUBMARINE")

	file_bytes, _ := ioutil.ReadFile("../data/data7.txt")

	decodelen := base64.StdEncoding.DecodedLen(len(file_bytes))
	ciphertext := make([]byte, decodelen)
	ctlen, _ := base64.StdEncoding.Decode(ciphertext, file_bytes)
	ciphertext = ciphertext[:ctlen]

	plaintext, err := DecryptECB(key, ciphertext)
	if err != nil {
		t.Fatal()
	}
	exp := []byte("I'm back and I'm ringin' the bell")
	if !util.ByteEq(plaintext[:len(exp)], exp) {
		t.Fatal()
	}

	new_ciphertext, err := EncryptECB(key, plaintext)
	if err != nil {
		t.Fatal()
	}
	if !util.ByteEq(new_ciphertext, ciphertext) {
		t.Fatal()
	}
}
