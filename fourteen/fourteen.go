/* The trick to this is that we have to find the length of the random prefix
* that is added. Once we have that we can pad our input so that prefix plus
* some of our input is an even number of blocks. Then we basically have
* challenge twelve.
*/
package fourteen

import (
    "fmt"
    "encoding/base64"
    "../aesmodes"
    "../util"
)

var key, _ = aesmodes.MakeKey()
var prefix = aesmodes.RandBytes(100)
var Target, _  = base64.StdEncoding.DecodeString("Um9sbGluJyBpbiBteSA1LjAKV2l0aCBteSByYWctdG9wIGRvd24gc28gbXkgaGFpciBjYW4gYmxvdwpUaGUgZ2lybGllcyBvbiBzdGFuZGJ5IHdhdmluZyBqdXN0IHRvIHNheSBoaQpEaWQgeW91IHN0b3A/IE5vLCBJIGp1c3QgZHJvdmUgYnkK")

func P(a interface{}) {
    fmt.Println(a)
}

func Encryptor(plaintext []byte) []byte {
    input := append(prefix, append(plaintext, Target...)...)
    out, _ := aesmodes.EncryptECB(key, input)
    return out
}

func FindBlockSize(encryptor func([]byte) []byte) int {
    bs := 0
    init := len(encryptor(make([]byte, 0)))
    for padlen := 0; bs == 0; padlen++ {
        longpad := make([]byte, padlen)
        final := len(encryptor(longpad))
        bs = final - init
    }
    return bs
}

func FindPrefixPlusTargetLen(encryptor func([]byte) []byte) int {
    bs := 0
    init := len(encryptor(make([]byte, 0)))
    for padlen := 0; bs == 0; padlen++ {
        longpad := make([]byte, padlen)
        final := len(encryptor(longpad))
        if (final - init) != 0 {
            return init - padlen
        }
    }
    return 0
}

// Return the index of the differing block. -1 if same
func FindBlockDiff(a, b []byte, bs int) int {
    if util.ByteEq(a, b) {
        return -1
    }
    for i := 0; i < (len(a) / bs); i++ {
        start := i*bs
        stop := start + bs
        if !util.ByteEq(a[start:stop], b[start:stop]) {
            return i
        }
    }
    return -1
}

func FindPrefixLen(encryptor func([]byte) []byte) int {
    bs := FindBlockSize(encryptor)
    init :=encryptor(make([]byte, 0))

    ac_start_block_i := FindBlockDiff(init, encryptor(make([]byte, 1)), bs)
    start := ac_start_block_i * bs
    stop := start + bs

    aclen := 0
    for {
        old := encryptor(make([]byte, aclen))
        next := encryptor(make([]byte, aclen + 1))
        if util.ByteEq(old[start:stop], next[start:stop]) {
            break
        }
        aclen++
    }
    return bs*(ac_start_block_i + 1) - aclen
}

func FindTargetLen(encryptor func([]byte) []byte) int {
    tot := FindPrefixPlusTargetLen(encryptor)
    return tot - FindPrefixLen(encryptor)
}
