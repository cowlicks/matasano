package thirteen

import (
    "fmt"
    "testing"
//    "../aesmodes"
)

func TestThirteen(t * testing.T) {
    fmt.Println("start test\n")
    ctprof := ProfileFor("blake.a.griffith@gmail.com")
    Decryptor(ctprof)
    escct := Escalate()
    esc := Decryptor(escct)
    Printer(esc)
}
