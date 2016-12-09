package hmac

import (
	"encoding/hex"
	"testing"
)

func TestEmpty(t *testing.T) {
	o := HmacSha1([]byte(""), []byte(""))
	e := "fbdb1d1b18aa6c08324b7d64b71fb76370690e1d"
	if !(hex.EncodeToString(o) == e) {
		t.Fail()
	}
}

func Test(t *testing.T) {

	o := HmacSha1([]byte("key"), []byte("The quick brown fox jumps over the lazy dog"))
	e := "de7c9b85b8b78aa6bc8a7a36f70a90701c9db4d9"
	if !(hex.EncodeToString(o) == e) {
		t.Fail()
	}
}
