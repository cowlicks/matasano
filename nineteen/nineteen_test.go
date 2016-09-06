package sixteen

import (
    "testing"
    "../xor"
)

func TestEighteen(t * testing.T) {
    fn := "data.txt"
    data := ReadFile(fn)
    fb, _ := MostFrequent(data)
    space_arr := make([]byte, len(fb))
    for i := range space_arr {
        space_arr[i] = byte(' ')
    }
    _, _ = xor.Xor(space_arr, fb)
    xor.Xor(space_arr, fb)
}
