package mersenne

import (
	"testing"
)

func TestMersenne(t *testing.T) {
	mt := NewMersenne19937(5489)
	for i := 0; i < 3; i++ {
		mt.Next()
	}
}

func TestCloneMT(t *testing.T) {
    mt := NewMersenne19937(5489)
    mt_data := make([]uint32, mt.n)
    for i := uint32(0); i < mt.n; i++ {
        mt_data[i] = mt.Next()
    }
    copied_mt := CloneMT19937(mt_data)
    new_data := make([]uint32, copied_mt.n)
    for i := uint32(0); i < copied_mt.n; i++ {
        new_data[i] = copied_mt.Next()
    }
    for i := uint32(0); i < mt.n; i++ {
        if new_data[i] != mt_data[i] {
            t.Fail()
        }
    }
}

func TestShittyStreamCipher(t *testing.T) {
    pt := []byte("This is my plaintext")
    key := []byte("y0")
    ct := Encrypt(key, pt)
    decrypted := Decrypt(key, ct)
    for i:= range pt {
        if pt[i] != decrypted[i] {
            t.Fail()
        }
    }
}
