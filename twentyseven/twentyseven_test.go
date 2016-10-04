package thepersonuwer

import (
    "../util"
    "testing"
)

func TestMe(t *testing.T) {
	text := []byte("COLDBREWCOFFEESS")
	pt := append(text, append(text, text...)...)
	key := []byte("YELLOWSUBMARINES")
    ct, _ := EncryptCBC(key, pt)
	out, er2 := DecryptCBC(key, ct)
	util.P(string(pt))
	util.P(ct)
	util.P(er2)
	util.P(string(out))
}
