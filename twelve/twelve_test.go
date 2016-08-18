package twelve

import (
    "fmt"
    "testing"
    "../aesmodes"
)

func TestOracle(t * testing.T) {
    fmt.Println("start test\n")
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

    OneShort()

    p := OneBlock()
    fmt.Println(string(p))

    q := WholeThing()
    fmt.Println(string(q))
}
