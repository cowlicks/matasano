package sha1

import (
	"encoding/hex"
	"testing"
)

func Test(t *testing.T) {
	o := Sha1([]byte("butts"))
	expected := "cd89a20adde7a608f3331e71c37bdfa087bacbf3"
	if !(hex.EncodeToString(o) == expected) {
		t.Fail()
	}
}
