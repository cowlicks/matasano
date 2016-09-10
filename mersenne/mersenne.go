package mersenne

import (
    "strconv"
    "fmt"
    "../util"
)

func PrintBin(a uint) {
    fmt.Println(strconv.FormatUint(uint64(a), 2))
}

type Mersenne struct {
	w, n, m, r, a, u, d, s, b, t, c, l, f uint
	index                                 uint
	state                                 []uint
}

// parameters from
// http://en.cppreference.com/w/cpp/numeric/random/mersenne_twister_engine
func NewMersenne19937(seed uint) *Mersenne {
	n := uint(624)
	mt := Mersenne{
		w:     32,              // word size
		n:     n,               // degree of recurrence
		m:     397,             // middle word
		r:     31,              // seperation point of one word
		a:     0x9908B0DF,      // coefficients of twist matrix (wtf wikipedia)
        u:     11,              // -- 
        s:     7,               //   \__ bit shifts
        t:     15,              //   /
		l:     18,              // --
		d:     0xFFFFFFFF,      // --
		b:     0x9D2C5680,      //   \__ bit masks
		c:     0xEFC60000,      // __/
		f:     1812433253,      // initialization constant
		index: n,               // state holders
		state: make([]uint, n), // ''
	}

	mt.state[0] = seed
	for i := uint(1); i < n; i++ {
		mt.state[i] = mt.f*(mt.state[i-1]^(mt.state[i-1]>>(mt.w-2))) + i
	}
	return &mt
}

func (mt *Mersenne) Next() uint32 {
	if mt.index >= mt.n {
		mt.twist()
	}

	y := mt.state[mt.index]
    y = mt.temper(y)
	mt.index++
	return uint32(y & (uint(^uint32(0))))
}

func (mt *Mersenne) temper(y uint) uint {
    y = y ^ ((y >> mt.u) & mt.d)
    y = y ^ ((y << mt.s) & mt.b)
    y = y ^ ((y << mt.t) & mt.c)
    y = y ^ (y >> mt.l)
    return y
}

func (mt *Mersenne) twist() {
	for i := uint(0); i < mt.n; i++ {
		lower_mask := (uint(1) << mt.r) - 1
		upper_mask := ^lower_mask
		y := (mt.state[i] & upper_mask) + (mt.state[(i+1)%mt.n] ^ lower_mask)
		mt.state[i] = mt.state[(i+mt.m)%mt.n] ^ (y>>1)

		if y%2 != 0 {
			mt.state[i] = mt.state[i] ^ mt.a
		}
	}
    util.P(mt.state[:5])
    mt.index = 0
}

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
func untemper(y uint) uint {
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

func untemper2(y uint) uint {
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

func CloneMT19937(mt *Mersenne) *Mersenne {
    mtout := make([]uint, mt.n)
    state_copy := make([]uint, mt.n)

    for ;mt.index != 624; {
        mt.Next()
    }
    for i := uint(0); i < mt.n; i++ {
        mtout[i] = uint(mt.Next())
    }
    util.P(mt.index)
    for i := uint(0); i < mt.n; i++ {
        state_copy[i] = untemper(mtout[i])
    }
	copied_mt := Mersenne{
		w:     32,              // word size
		n:     mt.n,               // degree of recurrence
		m:     397,             // middle word
		r:     31,              // seperation point of one word
		a:     0x9908B0DF,      // coefficients of twist matrix (wtf wikipedia)
        u:     11,              // -- 
        s:     7,               //   \__ bit shifts
        t:     15,              //   /
		l:     18,              // --
		d:     0xFFFFFFFF,      // --
		b:     0x9D2C5680,      //   \__ bit masks
		c:     0xEFC60000,      // __/
		f:     1812433253,      // initialization constant
		index: 0,               // state holders
		state: make([]uint, mt.n), // ''
	}
    copy(copied_mt.state, state_copy)
    util.P(copied_mt.state[:5])
    copied_mt.index = 0
    return &copied_mt
}
