package hex_to_base64

import (
    "testing"
    "fmt"
)

var input = "49276d206b696c6c696e6720796f757220627261696e206c696b65206120706f69736f6e6f7573206d757368726f6f6d"

var output = "SSdtIGtpbGxpbmcgeW91ciBicmFpbiBsaWtlIGEgcG9pc29ub3VzIG11c2hyb29t"
func TestHexToBase64(t * testing.T) {
    if _, err := HexToBin("ab"); err != nil {
        t.Fatalf("HexToBin failed with %!s", err)
    }

    if res, _ := HexToBase64(input); res != output {
        fmt.Println(res)
        t.Fatalf("HexToBase64 failed")
    }
}
