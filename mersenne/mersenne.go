package mersenne

var MIDDLEWORDSIZE int = 624

type Mersenne struct {
	w, n, m, r, a, u, d, s, b, t, c, l, f uint32
	index                                 uint32
	state                                 []uint32
}

// parameters from
// http://en.cppreference.com/w/cpp/numeric/random/mersenne_twister_engine
func NewMersenne19937(seed uint32) *Mersenne {
	n := uint32(624)
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
		state: make([]uint32, n), // ''
	}

	mt.state[0] = seed
	for i := uint32(1); i < n; i++ {
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
	return y
}

func (mt *Mersenne) temper(y uint32) uint32 {
    y = y ^ ((y >> mt.u) & mt.d)
    y = y ^ ((y << mt.s) & mt.b)
    y = y ^ ((y << mt.t) & mt.c)
    y = y ^ (y >> mt.l)
    return y
}

func (mt *Mersenne) twist() {
	for i := uint32(0); i < mt.n; i++ {
		lower_mask := (uint32(1) << mt.r) - 1
		upper_mask := ^lower_mask
		y := (mt.state[i] & upper_mask) + (mt.state[(i+1)%mt.n] ^ lower_mask)
		mt.state[i] = mt.state[(i+mt.m)%mt.n] ^ (y>>1)

		if y%2 != 0 {
			mt.state[i] = mt.state[i] ^ mt.a
		}
	}
    mt.index = 0
}
