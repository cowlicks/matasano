package sha1

import (
	"bytes"
	"crypto/sha1"
	"encoding/hex"
	"testing"
)

func Test(t *testing.T) {
	o := Sha1([]byte("butts"))
	expected := "cd89a20adde7a608f3331e71c37bdfa087bacbf3"
	if hex.EncodeToString(o) != expected {
		t.Fail()
	}

	a := sha1.Sum([]byte("butts"))
	if !bytes.Equal(o, a[:]) {
		t.Fail()
	}
}
