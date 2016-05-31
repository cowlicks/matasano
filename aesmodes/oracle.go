package aesmodes

import (
    crand "crypto/rand"
    mrand "math/rand"
)

func MakeKey() ([]byte, error) {
    key := make([]byte, 16)
    _, err := crand.Read(key)
    return key, err
}

func concat(b ...[]byte) []byte {
    var out []byte
    for _, v := range b {
        out = append(out, v...)
    }
    return out
}

func OracleEncryptor(data []byte) ([]byte, error) {
    var ct []byte
    prefix := make([]byte, mrand.Intn(6) + 5)
    sufix := make([]byte, mrand.Intn(6) + 5)
    padded := concat(prefix, data, sufix)

    key, err := MakeKey()
    if err != nil {
        return padded, err
    }
    encrypt_mode := mrand.Intn(2)
    if encrypt_mode == 0 {
        ct, err = EncryptCBC(key, padded)
    } else {
        ct, err = EncryptECB(key, padded)
    }
    return ct, err
}

func OracleDetector(data []byte) string {
    if DetectECB(16, data) {
        return "ECB"
    }
    return "CBC"
}
