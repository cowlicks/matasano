package one

import (
    "encoding/hex"
    "bytes"
)

func HexToBin(s string) ([]byte, error) {
    return hex.DecodeString(s)
}

func base64map(i byte) byte {
    var codes = "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789+/="
    return codes[i]
}

func three_bytes_to_four_base64(b []byte) []byte {
        bone := b[0]
        btwo := b[1]
        bthree := b[2]

        out := []byte{
                ((bone & (255 - 3)) >> 2),
                ((bone & 3) << 4) + ((btwo & (255 - 15)) >> 4),
                ((btwo & 15) << 2) + ((bthree & 192) >> 6),
                bthree & (255 - 192)}

        for i := 0; i < 4; i++ {
            out[i] = base64map(out[i])
        }
        return out
}


func HexToBase64(s string) (string, error) {
    // (128 64 32 16 8 4 2 1)
    var buffer bytes.Buffer
    b, err := HexToBin(s)

    for i := 0; i < len(b); i += 3 {
        buffer.Write(three_bytes_to_four_base64(b[i:i+3]))
    }
    return buffer.String(), err
}
