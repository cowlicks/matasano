package hmac

import (
	"../sha1"
	"../xor"
	"fmt"
)

func P(i ...interface{}) {
	fmt.Println(i)
}

// HMAC(K, m) = H((K' ⊕ opad) || H((K' ⊕ ipad) || m))
func HmacSha1(key, data []byte) []byte {
	sha1len := 64
	if len(key) > sha1len {
		key = sha1.Sha1(key)
	}
	if len(key) < sha1len {
		key = append(key, make([]byte, sha1len-len(key))...)
	}
	opad := make([]byte, sha1len)
	ipad := make([]byte, sha1len)
	for i := 0; i < sha1len; i++ {
		opad[i] = 0x5c
		ipad[i] = 0x36
	}
	kopad, _ := xor.Xor(key, opad)
	kipad, _ := xor.Xor(key, ipad)
	return sha1.Sha1(append(kopad, sha1.Sha1(append(kipad, data...))...))
}
