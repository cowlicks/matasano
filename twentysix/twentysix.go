package twentysix

import (
    //"../util"
    "net/url"
    "../mersenne"
)

var key = []byte("COLDBREW")

var imadmin =  []byte(";admin=true")
var Username = []byte("realusernameClobbbberMe")
var prefix = []byte("comment1=cooking%20MCs;userdata=")
var suffix = []byte(";comment2=%20like%20a%20pound%20of%20bacon")

func Encrypt(input string) []byte {
    escaped := []byte(url.QueryEscape(input))
    plaintext := append(prefix, append(escaped, suffix...)...)
    ct := mersenne.Encrypt(key, plaintext)
    return ct
}

func BitFlip(ct []byte) {
    start := len(prefix) + len(Username) - 1
    username_index := len(Username) - 1
    admin_index := len(imadmin) - 1
    for i := start; i > start - len(imadmin); i-- {
        ct[i] = ct[i] ^ Username[username_index] ^ imadmin[admin_index]
        username_index--
        admin_index--
    }
}

func Decrypt(ct []byte) []byte {
    return mersenne.Decrypt(key, ct)
}
