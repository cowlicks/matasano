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
    MakeDict(16)

    MakeDictWithPad(16, make([]byte, 70))

    lttr := OneShort()
    fmt.Println(lttr)

    MakeCTDict(16)

    //All()
}
