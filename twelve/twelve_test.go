package twelve

import (
    "fmt"
    "testing"
    "../aesmodes"
)

func TestOracle(t * testing.T) {
    // test encryptor
    key, _ := aesmodes.MakeKey()
    Encryptor([]byte("manifolds"), key)

    // test keyedencryptor
    KeyedEncryptor([]byte("socks"))

    // test get blocksize
    bs := FindBlockSize(KeyedEncryptor)
    if bs != 16 {
        t.Fatal()
    }

    // test check ecb
    if !CheckEncryptorECB(KeyedEncryptor) {
        t.Fatal()
    }

    // test make dict
    yo := MakeDict(16)
    for k, v := range yo {
        fmt.Println(k)
        fmt.Println(v)
    }
}
