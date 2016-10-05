package mersenne

/*
These numbers are a mystery to me.

The operations in the temper function, actually represent a matrix
multiplication with a vector.  The matrix is a constant tempering matrix, the
vector is the uint (each bit is an element in the vector.  Somehow this
operation is reduced to the bitshifting and &'ing in the temper function. Idk
how this translation happens.  But the tempering matrix is invertible, so the
result of the temper function can be multiplied by the inverse of the tempering
matrix to get the original value. This untemper function apparently does this.
Its values are magic to me and they are based on the constants for MT19937.
This was taken from:

https://www.randombit.net/bitbashing/2009/07/21/inverting_mt19937_tempering.html#

The matrices themselves and another untempering formula can be found here:

https://gist.github.com/oupo/ce045423a15395d31d3c
*/
func untemper(y uint32) uint32 {
	y = y ^ (y >> 18)
	y = y ^ ((y << 15) & 0xEFC60000)
	y = y ^ ((y << 7) & 0x1680)
	y = y ^ ((y << 7) & 0xC4000)
	y = y ^ ((y << 7) & 0xD200000)
	y = y ^ ((y << 7) & 0x90000000)
	y = y ^ ((y >> 11) & 0xFFC00000)
	y = y ^ ((y >> 11) & 0x3FF800)
	y = y ^ ((y >> 11) & 0x7FF)
	return y
}

func CloneMT19937(mt_data []uint32) *Mersenne {
	n := uint32(MIDDLEWORDSIZE)
	state_copy := make([]uint32, n)

	for i := uint32(0); i < n; i++ {
		state_copy[i] = untemper(mt_data[i])
	}
	copied_mt := Mersenne{
		w:     32,                // word size
		n:     n,                 // degree of recurrence
		m:     397,               // middle word
		r:     31,                // seperation point of one word
		a:     0x9908B0DF,        // coefficients of twist matrix (wtf wikipedia)
		u:     11,                // --
		s:     7,                 //   \__ bit shifts
		t:     15,                //   /
		l:     18,                // --
		d:     0xFFFFFFFF,        // --
		b:     0x9D2C5680,        //   \__ bit masks
		c:     0xEFC60000,        // __/
		f:     1812433253,        // initialization constant
		index: 0,                 // state holders
		state: make([]uint32, n), // ''
	}
	copied_mt.index = 0
	copy(copied_mt.state, state_copy)
	return &copied_mt
}
