package mersenne


func Encrypt(key, plaintext []byte) []byte {
    var intkey uint32
    for i := range key {
        intkey = (intkey << 32) + uint32(key[i])
    }
    out := make([]byte, len(plaintext))
    mt := NewMersenne19937(intkey)
    for i, val := range plaintext {
        out[i] = byte(uint32(val) ^ mt.Next())
    }
    return out
}

func Decrypt(key, ciphertext []byte) []byte {
    return Encrypt(key, ciphertext)
}
