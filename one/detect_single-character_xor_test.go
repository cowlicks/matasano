package one

import (
    "encoding/hex"
    "testing"
    "strings"
    "io/ioutil"
)


func TestSingleByteXor(t * testing.T) {
    file_bytes, err := ioutil.ReadFile("data4.txt")
    Check(err)

    file_strings := strings.Split(string(file_bytes), "\n")
    char_counts := make([]CharCount, len(file_strings))
    var max_count int
    var max_i int
    for i, s := range(file_strings) {
        cc := *MakeCharCount(s)
        if (len(cc.order) > 0) && (cc.mapping[cc.order[0]] >= max_count) {
                max_count = cc.mapping[cc.order[0]]
                max_i = i
            }
        char_counts[i] = cc
    }
    out, _ := hex.DecodeString(UnXor(file_strings[max_i]))
    exp := "Now that the party is jumping\n"
    if string(out) != exp {
        t.Fatalf("Output incorrect, got %s", out)
    }
}
