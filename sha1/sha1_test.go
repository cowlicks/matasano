package sha1

import (
	"crypto/sha1"
	"encoding/hex"
	"bytes"
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

func TestBadMac(t *testing.T) {
	key := []byte("lol")
	msg := []byte("The Mendez Trial Documentary")
	mac := BadMac(key, msg)
	if !VerifyBadMac(key, msg, mac) {
		t.Fatal()
	}
}

func TestForge(t *testing.T) {
	msg := []byte("The Mendez Trial Documentary")
	mac := GetMac(msg)
	suffix := []byte(";admin=true")
	newmsg, newmac, err := ExtendMac(msg, mac, []byte(";admin=true"))
	if err != nil {
		t.Fatal(err)
	}
	if !bytes.Contains(newmsg, suffix) {
		t.Fatal()
	}
	if !CheckMac(newmsg, newmac) {
		t.Fatal()
	}
}
