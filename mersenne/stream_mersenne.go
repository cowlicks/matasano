package mersenne

func key_byte_to_uint32(data []byte) uint32 {
    var out uint32
    for i := range data {
        out = (out << 8) + uint32(data[i])
    }
    return out
}

func Encrypt(key, plaintext []byte) []byte {
    intkey := key_byte_to_uint32(key)
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

func Edit(key, ciphertext []byte, newbyte byte, offset int) []byte {
    out := make([]byte, len(ciphertext))
    copy(out, ciphertext)

    intkey := key_byte_to_uint32(key)
    mt := NewMersenne19937(intkey)
    for i := 0; i < offset; i++ {
        mt.Next()
    }
    xorbyte := byte(mt.Next())
    plaintextbyte := ciphertext[offset] ^ xorbyte

    out[offset] = ciphertext[offset] ^ plaintextbyte ^ newbyte
    return out
}
