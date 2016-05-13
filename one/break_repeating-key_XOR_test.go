package one
import (
    "testing"
    "encoding/hex"
)
func testEq(a, b []byte) bool {

    if a == nil && b == nil {
        return true;
    }

    if a == nil || b == nil {
        return false;
    }

    if len(a) != len(b) {
        return false
    }

    for i := range a {
        if a[i] != b[i] {
            return false
        }
    }

    return true
}

func TestXor(t * testing.T) {
    in1, _ := hex.DecodeString("1c0111001f010100061a024b53535009181c")
    in2, _ := hex.DecodeString("686974207468652062756c6c277320657965")
    out, _ := hex.DecodeString("746865206b696420646f6e277420706c6179")

    res, err := Xor(in1, in2)

    if err != nil {
        t.Error(err)
    }

    if !testEq(res, out) {
        t.Fatal("fail")
    }
}

func TestHamming(t * testing.T) {
    in1 := []byte("this is a test")
    in2 := []byte("wokka wokka!!!")
    out := 37

    res, err := HammingDistance(in1, in2)
    if err != nil {
        t.Error(err)
    }

    if res != out {
        t.Fatal(res, out)
    }
}
