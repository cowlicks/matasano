package aesmodes

import (
	"bytes"
	"encoding/base64"
	"io/ioutil"
	"testing"
)

func TestDetectECB(t *testing.T) {
	file_bytes, _ := ioutil.ReadFile("../data/data8.txt")
	encoded := bytes.Split(file_bytes, []byte("\n"))
	ciphertexts := make([][]byte, len(encoded))
	for i, e := range encoded {
		d, _ := base64.StdEncoding.DecodeString(string(e))
		ciphertexts[i] = d
	}

	blocksize := 16
	ecb_ct := 0
	for i, ct := range ciphertexts {
		if DetectECB(blocksize, ct) {
			ecb_ct = i
		}
	}
	if ecb_ct != 132 {
		t.Fatal()
	}
}
