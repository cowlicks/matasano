package sixteen

import (
    "bytes"
    "net/url"
    "../aesmodes"
)

var key, _ = aesmodes.MakeKey()

var prefix = []byte("comment1=cooking%20MCs;userdata=")
var suffix = []byte(";comment2=%20like%20a%20pound%20of%20bacon")

func Encrypt(input string) []byte {
    escaped := []byte(url.QueryEscape(input))
    plaintext := append(prefix, append(escaped, suffix...)...)
    ct, _ := aesmodes.EncryptCBC(key, plaintext)
    return ct
}

func Decrypt(ct []byte) []byte {
    pt, _ := aesmodes.DecryptCBC(key, ct)
    return pt
}

func MakeAdmin() []byte {
    bs := 16
    plen := len(prefix)
    magic_userdata_len := bs - (plen % bs)
    magic_admin := "0admin0true"
    magic_userdata_bytes := make([]byte, magic_userdata_len)
    for i := 0; i < magic_userdata_len; i++ {
        magic_userdata_bytes[i] = byte('a')
    }
    magic_userdata := string(magic_userdata_bytes)
    magic_input := magic_userdata + magic_admin

    ct := Encrypt(magic_input)
    pt := Decrypt(ct)
    i := bytes.Index(pt, []byte(magic_admin))

    ct[i] = byte('0') ^ byte(';') ^ ct[i]
    ct[i + 6] =byte('0') ^ byte('=') ^ ct[i + 6]
    return ct
}

// No URL parsing here because the scrambled block ruins it.
func IsAdmin(ciphertext []byte) bool {
    plaintext_url := Decrypt(ciphertext)
    return bytes.Contains(plaintext_url, []byte(";admin=true"))
}
