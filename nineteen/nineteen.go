package sixteen

import (
    "fmt"
    "io/ioutil"
    "bytes"
    "encoding/base64"
    //"../xor"
)

func P(stuff ...interface{}) {
    fmt.Println(stuff)
}

func ReadFile(name string) [][]byte {
    fn := "data.txt"
    data, _ := ioutil.ReadFile(fn)
    lines := bytes.Split(data, []byte("\n"))
    lines = lines[:len(lines) - 1]
    out := make([][]byte, len(lines))
    for i := range lines {
        out[i] = make([]byte, base64.StdEncoding.DecodedLen(len(lines[i])))
        l, _ := base64.StdEncoding.Decode(out[i], lines[i])
        out[i] = out[i][:l]
    }
    return out
}

func MinLenInArray(arr [][]byte) int {
    min := -1
    for i := range arr {
        length := len(arr[i])
        if min == -1 || length < min {
            min = length
        }
    }
    return min
}

func MaxLenInArray(arr [][]byte) int {
    var max int
    for i := range arr {
        length := len(arr[i])
        if max == -1 || length > max {
            max = length
        }
    }
    return max
}

func GetColumns(arr [][]byte) [][]byte {
    ncols := MaxLenInArray(arr)
    cols := make([][]byte, ncols)
    for i := range cols {
        cols[i] = make([]byte, 0)
    }
    for _, row := range arr {
        for coli := range row {
            cols[coli] = append(cols[coli], row[coli])
        }
    }
    return cols
}

func column_freq_winner(cf map[byte]int) (byte, int) {
    var max int
    var out_byte byte

    for k, v := range cf {
        if v > max {
            max = v
            out_byte = k
        }
    }
    return out_byte, max
}

func MostFrequent(arr [][]byte) ([]byte, []int) {
    max := MaxLenInArray(arr)
    freq_bytes := make([]byte, max)
    freq_counts := make([]int, max)

    for i := 0; i < max; i++ {
        column_freq := make(map[byte]int)
        for _, row := range arr {
            if i < len(row) {
                column_freq[row[i]] += 1
            }
        }
        freq_bytes[i], freq_counts[i] = column_freq_winner(column_freq)
    }
    return freq_bytes, freq_counts
}
