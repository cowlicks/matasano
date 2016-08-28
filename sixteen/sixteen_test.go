package sixteen

import (
    "testing"
)

func TestTwelve(t * testing.T) {
    out := IsAdmin(MakeAdmin())
    if out != true {
        t.Fatal()
    }
}
