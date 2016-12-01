/*
Implementation of SHA-1 ported from:
https://en.wikipedia.org/wiki/SHA-1#SHA-1_pseudocode
https://github.com/ajalt/python-sha1

I should really make a struct for the h0 - h4 registers to clean this up
*/
package sha1

import (
	"encoding/binary"
)

func left_rotate(n, b uint32) uint32 {
	return ((n << b) | (n >> (32 - b))) & 0xffffffff
}

func MDPad(message []byte) []byte {
	original_byte_len := uint64(len(message))
	original_bit_len := original_byte_len * 8

	// append the bit '1' to the message
	out := []byte{0x80}

	// append 0 <= k < 512 bits '0', so that the resulting message length (in bits)
	// is congruent to 448 (mod 512)
	out = append(out, make([]byte, ((56-(original_byte_len+1)%64)%64))...)

	// append length of message (before pre-processing), in bits, as 64-bit big-endian integer
	len_pad := make([]byte, 8)
	binary.BigEndian.PutUint64(len_pad, original_bit_len)
	out = append(out, len_pad...)
	return out
}

func HashChunk(chunk []byte, h0, h1, h2, h3, h4 uint32) (o0, o1, o2, o3, o4 uint32) {
	w := make([]uint32, 80)
	// break chunk into sixteen 32-bit big-endian words w[i]
	for j := 0; j < 16; j++ {
		w[j] = binary.BigEndian.Uint32(chunk[(j*4) : (j*4)+4])
	}
	// Extend the sixteen 32-bit words into eighty 32-bit words:
	for j := 16; j < 80; j++ {
		w[j] = left_rotate(w[j-3]^w[j-8]^w[j-14]^w[j-16], 1)
	}

	// Initialize hash value for this chunk:
	a := h0
	b := h1
	c := h2
	d := h3
	e := h4

	for i := 0; i < 80; i++ {
		var f, k uint32
		switch {
		case 0 <= i && i <= 19:
			f = (b & c) | ((^b) & d)
			k = 0x5A827999
		case 20 <= i && i <= 39:
			f = b ^ c ^ d
			k = 0x6ED9EBA1
		case 40 <= i && i <= 59:
			f = (b & c) | (b & d) | (c & d)
			k = 0x8F1BBCDC
		case 60 <= i && i <= 79:
			f = b ^ c ^ d
			k = 0xCA62C1D6
		}

		temp := left_rotate(a, 5) + f + e + k + w[i]
		e = d
		d = c
		c = left_rotate(b, 30)
		b = a
		a = temp
	}

	// Add this chunk's hash to result so far:
	o0 = (h0 + a) & 0xffffffff
	o1 = (h1 + b) & 0xffffffff
	o2 = (h2 + c) & 0xffffffff
	o3 = (h3 + d) & 0xffffffff
	o4 = (h4 + e) & 0xffffffff
	return
}


func Sha1(message []byte) []byte {
	var h0, h1, h2, h3, h4 uint32
	h0 = 0x67452301
	h1 = 0xEFCDAB89
	h2 = 0x98BADCFE
	h3 = 0x10325476
	h4 = 0xC3D2E1F0

	// pad the message
	message = append(message, MDPad(message)...)

	// Process the message in successive 512-bit chunks:
	// break message into 512-bit chunks
	for i := 0; i < len(message); i += 64 {
		h0, h1, h2, h3, h4 = HashChunk(message[i:i+16], h0, h1, h2, h3, h4)
	}

	// Produce the final hash value (big-endian):
	out := make([]byte, 20)
	binary.BigEndian.PutUint32(out[0:4], h0)
	binary.BigEndian.PutUint32(out[4:8], h1)
	binary.BigEndian.PutUint32(out[8:12], h2)
	binary.BigEndian.PutUint32(out[12:16], h3)
	binary.BigEndian.PutUint32(out[16:20], h4)
	return out
}

// expects len(addme) % 512 == 0
func Extend(addme []byte, h0, h1, h2, h3, h4 uint32) []byte {
	// Process the message in successive 512-bit chunks:
	// break message into 512-bit chunks
	for i := 0; i < len(addme); i += 64 {
		h0, h1, h2, h3, h4 = HashChunk(addme[i:i+16], h0, h1, h2, h3, h4)
	}

	// Produce the final hash value (big-endian):
	out := make([]byte, 20)
	binary.BigEndian.PutUint32(out[0:4], h0)
	binary.BigEndian.PutUint32(out[4:8], h1)
	binary.BigEndian.PutUint32(out[8:12], h2)
	binary.BigEndian.PutUint32(out[12:16], h3)
	binary.BigEndian.PutUint32(out[16:20], h4)
	return out
}

func GetRegisters(sha1out []byte) (h0, h1, h2, h3, h4 uint32) {
	h0 = binary.BigEndian.Uint32(sha1out[0:4])
	h1 = binary.BigEndian.Uint32(sha1out[4:8])
	h2 = binary.BigEndian.Uint32(sha1out[8:12])
	h3 = binary.BigEndian.Uint32(sha1out[12:16])
	h4 = binary.BigEndian.Uint32(sha1out[16:20])
	return
}
