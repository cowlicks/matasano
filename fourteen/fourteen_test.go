package fourteen

import (
    "fmt"
    "testing"
    "../twelve"
)

func TestFourteen(t * testing.T) {
    fmt.Println("start test\n")
    // test encryptor
    Encryptor([]byte("manifolds"))

    // test get blocksize
    bs := FindBlockSize(Encryptor)
    if bs != 16 {
        t.Fatal()
    }

    a := FindPrefixPlusTargetLen(Encryptor)
    if !(a > 0) {
        t.Fatal()
    }

    tlen := FindTargetLen(Encryptor)
    if tlen != len(Target) {
        t.Fatal()
    }
    EncryptorClipper(Encryptor, []byte("power"))

    pt := twelve.WholeThing(ClippedEncryptor)
    P(string(pt))
}
