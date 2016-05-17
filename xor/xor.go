package xor

import (
    "errors"
)

func Xor(a, b []byte) ([]byte, error) {
	out := make([]byte, len(a))
	if len(a) != len(b) {
		return out, errors.New("inputs different length")
	}

	for i, v := range a {
		out[i] = v ^ b[i]
	}
	return out, nil
}

func VectorXor(a, b []byte) []byte {
    var short []byte
    var long []byte
    if len(a) <= len(b) {
        short = a
        long = b
    } else {
        short = b
        long = a
    }

	ls := len(short)
	ll := len(long)
	nshorts := ll / ls
	if ll%ls != 0 {
		nshorts += 1
	}
	out := make([]byte, ll)
	for i := 0; i < nshorts; i++ {
		if ls*(i+1) <= ll {
			copy(out[ls*i:ls*(i+1)], short)
		} else {
			copy(out[ls*i:ll], short[:ll%ls])
		}
	}
	out, _ = Xor(out, long)
	return out
}

