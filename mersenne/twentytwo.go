package mersenne

import (
    "../util"
    "time"
    "math/rand"
)

func MTFromTime() *Mersenne {
    time := uint(time.Now().Unix())
    return NewMersenne19937(time)
}

func MakeMTNumberAtRandomTime() uint32 {
    wait_before := rand.Intn(300)
    wait_after := rand.Intn(300)

    time.Sleep(time.Duration(wait_before) * time.Second)
    mt := MTFromTime()
    time.Sleep(time.Duration(wait_after) * time.Second)

    return mt.Next()
}

func Crack(input uint32) int {
    window_size := 601
    now := int(time.Now().Unix())
    for i := now - window_size; i < now; i++ {
        mt := NewMersenne19937(uint(i))
        out := mt.Next()
        if out == input {
            util.P(i)
            return i
        }
    }
    return -1
}
