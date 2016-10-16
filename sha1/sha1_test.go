package sha1

import (
	"crypto/sha1"
	"encoding/hex"
	"testing"
)

func Test(t *testing.T) {
	o := Sha1([]byte("butts"))
	expected := "cd89a20adde7a608f3331e71c37bdfa087bacbf3"
	if !(hex.EncodeToString(o) == expected) {
		t.Fail()
	}

	a := sha1.Sum([]byte("butts"))
	for i := 0; i < len(a); i++ {
		if a[i] != o[i] {
			t.Fail()
		}
	}

}
