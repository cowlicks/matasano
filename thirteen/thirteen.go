package thirteen
/*
The trick to this:
Get a ciphertext block for just 'admin' with the proper padding bytes (11
elevens).  We get this by putting it in the email field with some padding in
front of it so it gets it own block.

Once we have that block, we want to put replace the last block of some other
ciphertext with it. We want this new ciphertext to line so that the "&role="
ends at the end of a block, so we can put our admin block at the end.
*/

import (
    "fmt"
    "bytes"
    "strings"
    "../aesmodes"
)

var usercount = 0
var key, _ = aesmodes.MakeKey()

func P(p []byte) {
    fmt.Println(string(p))
}

func Encryptor(plaintext []byte) []byte {
    out, _ := aesmodes.EncryptECB(key, plaintext)
    return out
}

func Decryptor(ciphertext []byte) []byte {
    out, _ := aesmodes.DecryptECB(key, ciphertext)
    return out
}

func Printer(b []byte) []byte {
    var buf bytes.Buffer

    buf.WriteString("{\n")

    fields := bytes.Split(b, []byte("&"))
    for _, field := range fields {
        keyandval := bytes.Split(field, []byte("="))
        buf.WriteString("\t'")
        buf.Write(keyandval[0])
        buf.WriteString("': '")
        buf.Write(keyandval[1])
        buf.WriteString("'\n")
    }
    buf.WriteString("}")
    fmt.Println(buf.String())
    return buf.Bytes()
}

func ProfileFor(s string) []byte {
    var buf bytes.Buffer
    if strings.Index(s, "&") != -1 {
        panic("input contains &")
    }
    if strings.Index(s, "=") != -1 {
        panic("input contains =")
    }
    buf.WriteString("email=")
    buf.WriteString(s)
    buf.WriteString("&uid=")
    buf.WriteString(fmt.Sprintf("%d", usercount))
    usercount++
    buf.WriteString("&role=user")
    return Encryptor(buf.Bytes())
}

func Escalate() []byte {
    // this could be cleaner
    bs := 16 // blocksize
    beforelen := len("email=")
    adminlen := len("admin")
    email_prepadlen := bs - beforelen
    adminblock := make([]byte, 16)
    copy(adminblock, []byte("admin"))
    for i := 5; i < bs; i++ {
        adminblock[i] = byte(bs - adminlen)
    }
    makeadminemail := append(make([]byte, email_prepadlen), adminblock...)

    ctwithadmin := ProfileFor(string(makeadminemail))
    ctadminblock := ctwithadmin[bs : bs*2]

    midlen := len("&uid=0&role=")
    emaillen := bs - ((midlen + beforelen) % bs)

    var emailpadbuf bytes.Buffer
    for i := 0; i < emaillen; i++ {
        // put whatever in this emailpadbuf as long as it is
        // n * bs + emaillen bytes long
        emailpadbuf.WriteString("Y")
    }
    emailpadding := emailpadbuf.Bytes()

    out := ProfileFor(string(emailpadding))
    outlen := len(out)
    copy(out[outlen - bs:outlen], ctadminblock)

    return out
}

