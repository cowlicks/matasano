package one

import (
	// keep hex out of here and in tests
    "../xor"
	"errors"
	"fmt"
	"sort"
)

func P(b interface{}) {
	fmt.Println(b)
}

func countOnesInByte(a byte) int {
	pows := []byte{1, 2, 4, 8, 16, 32, 64, 128}
	out := 0
	for _, p := range pows {
		if b := p & a; b == p {
			out += 1
		}
	}
	return out
}

func HammingDistance(a, b []byte) (int, error) {
	out := 0
	if len(a) != len(b) {
		return 0, errors.New("inputs different length")
	}

	difs, _ := xor.Xor(a, b)
	for _, v := range difs {
		out += countOnesInByte(v)
	}
	return out, nil
}

func HammingsOfData(b []byte) (hammings []float64, keysizes []int) {
	keysizes = make([]int, 38)
	// possible keysizes
	for i := range keysizes {
		keysizes[i] = i + 2
	}

	hammings = make([]float64, 38)

	fs := len(b) // filesize
	for ki, ks := range keysizes {
		nkb := (fs / ks) - 1 // # of ks blocks
		hamavg := 0.0
		hamsum := 0.0

		for i := 0; i < nkb; i++ {
			s1 := b[ks*i : ks*(i+1)]
			s2 := b[ks*(i+1) : ks*(i+2)]
			tmpham, _ := HammingDistance(s1, s2)
			hamsum += float64(tmpham)
		}
		hamavg = hamsum / float64(ks)
		hamavg = hamavg / float64(nkb-1)
		hammings[ki] = hamavg
	}

	sort.Sort(&IndexSort{hammings, keysizes})
	return hammings, keysizes
}

type IndexSort struct {
	sort.Float64Slice
	idxs []int
}

func (is *IndexSort) Swap(i, j int) {
	is.Float64Slice[i], is.Float64Slice[j] = is.Float64Slice[j], is.Float64Slice[i]
	is.idxs[i], is.idxs[j] = is.idxs[j], is.idxs[i]
}

func GetKeySizeBlocks(b []byte, ks int) [][]byte {
	nblocks := len(b) / ks
	if (len(b) % ks) != 0 {
		nblocks += 1
	}
	blocks := make([][]byte, nblocks)
	for i := 0; i < nblocks; i++ {
		if ks*(i+1) < len(b) {
			blocks[i] = make([]byte, ks)
			copy(blocks[i], b[i*ks:ks*(i+1)])
		} else {
			blocks[i] = make([]byte, len(b)%ks)
			copy(blocks[i], b[i*ks:len(b)])
		}
	}
	return blocks
}

func TransposeBlocks(blocks [][]byte) [][]byte {
	nblocks := len(blocks)
	ks := len(blocks[0])
	tblocks := make([][]byte, ks)
	for i := range tblocks {
		if i < len(blocks[nblocks-1]) {
			tblocks[i] = make([]byte, nblocks)
		} else {
			tblocks[i] = make([]byte, nblocks-1)
		}
		for j := range blocks {
			if j < len(tblocks[i]) {
				tblocks[i][j] = blocks[j][i]
			}
		}
	}
	return tblocks
}

type ByteCounter struct {
	Mapping map[byte]int
	Order   []byte
}

func (c *ByteCounter) Len() int {
	return len(c.Mapping)
}

func (c *ByteCounter) Less(i, j int) bool {
	return c.Mapping[c.Order[i]] > c.Mapping[c.Order[j]]
}

func (c *ByteCounter) Swap(i, j int) {
	c.Order[j], c.Order[i] = c.Order[i], c.Order[j]
}

func CountBytes(b []byte) *ByteCounter {
	out := new(ByteCounter)
	out.Mapping = make(map[byte]int)
	for _, v := range b {
		out.Mapping[v] += 1
	}
	out.Order = make([]byte, len(out.Mapping))
	i := 0
	for k := range out.Mapping {
		out.Order[i] = k
		i++
	}
	sort.Sort(out)
	return out
}

func isPrintable(b byte) bool {
	if (127 >= b) && (b >= 32) {
		return true
	}
	if b == 10 { // newline
		return true
	} else {
		return false
	}
}

func AllPrintable(arr []byte) bool {
	for _, v := range arr {
		if !isPrintable(v) {
			return false
		}
	}
	return true
}

func FindXorChar(b []byte) (byte, error) {
	common_chars := []byte(" aoeuidhtns")
	count := CountBytes(b)

	for _, u := range count.Order {
		for _, v := range common_chars {
			x := u ^ v
			xord := xor.VectorXor([]byte{x}, b)
			if AllPrintable(xord) && isPrintable(x) {
				return x, nil
			}
		}
	}
	return uint8(0), errors.New("xor char not found")
}

func BuildKey(ks int, tblocks [][]byte) ([]byte, error) {
	key := make([]byte, ks)
	for i := range tblocks {
		x, err := FindXorChar(tblocks[i])
		key[i] = x
		if err != nil {
			return key, err
		}
	}
	return key, nil
}

func makeBlocks(ks int, data []byte) [][]byte {
	blocks := GetKeySizeBlocks(data, ks)
	tblocks := TransposeBlocks(blocks)
	return tblocks
}

func CrackVigenere(data []byte) ([]byte, []byte, error) {
	var output []byte
	var k []byte
	_, keysizes := HammingsOfData(data)

	for _, ks := range keysizes {
		blocks := makeBlocks(ks, data)
		k, err := BuildKey(ks, blocks)
		if err == nil {
			return xor.VectorXor(k, data), k, nil
		}
	}
	return output, k, errors.New("couldn't crack it")
}
