package one

import (
    // keep hex out of here and in tests
    "fmt"
    "math"
    "errors"
)

func Xor(a, b []byte) ([]byte, error) {
    out := make([]byte, len(a))
    if len(a) != len(b) {
        return out, errors.New("inputs different length")
    }

    for i, v := range a {
        out[i] = v ^ b[i]
    }
    return out, nil
}

func countOnesInByte(a byte) int {
    pows := []byte{1, 2, 4, 8, 16, 32, 64, 128}
    out := 0
    for _, p := range pows {
        if b := p & a; b == p {
            out += 1
        }
    }
    return out
}

func HammingDistance(a, b []byte) (int, error) {
    out := 0
    if len(a) != len(b) {
        return 0, errors.New("inputs different length")
    }

    difs, _ := Xor(a, b)
    for _, v := range difs {
        out += countOnesInByte(v)
    }
    return out, nil
}

func ShannonEntropy(bytes []byte) float64 {
    l := float64(len(bytes))
    count := make(map[int]int)
    for b := range bytes {
        count[b] += 1
    }
    out := 0.0
    for _, v := range count {
        freq := float64(v) / l
        out -= freq * math.Log2(freq)
    }
    return out
}
