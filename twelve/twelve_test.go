package twelve

import (
    "fmt"
    "testing"
    //"../aesmodes"
)

func TestTwelve(t * testing.T) {
    fmt.Println("start test\n")

    // test encryptor
    Encryptor([]byte("socks"))

    // test get blocksize
    bs := FindBlockSize(Encryptor)
    if bs != 16 {
        t.Fatal()
    }

    // test check ecb
    if !CheckEncryptorECB(Encryptor) {
        t.Fatal()
    }

    tlen := FindTargetSize(Encryptor)
    if tlen != len(Target) {
        t.Fatal()
    }

    OneShort()

    p := OneBlock()
    fmt.Println(string(p))

    q := WholeThing(Encryptor)
    fmt.Println(string(q))
}
