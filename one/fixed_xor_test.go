package one

import (
    "testing"
)

var in1 = "1c0111001f010100061a024b53535009181c"
var in2 = "686974207468652062756c6c277320657965"
var out = "746865206b696420646f6e277420706c6179"

func TestXor(t * testing.T) {
    val := Xor(in1, in2)
    if val != out {
        t.Fatalf("Output incorrect, got %s", val)
    }
}
