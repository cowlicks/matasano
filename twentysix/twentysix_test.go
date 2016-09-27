package twentysix

import (
    "../util"
    "testing"
)

func TestMe(t *testing.T) {
    ct := Encrypt(string(Username))
    BitFlip(ct)
    util.P(string(Decrypt(ct)))
}
