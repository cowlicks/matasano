package mersenne

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
		a:     2567483615,      // coefficients of twist matrix
		d:     4294967295,      // --
		b:     2636928640,      //   \
		c:     4022730752,      //    -- bit shifting masks
		l:     18,              //   /
		f:     1812433253,      // --
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
	y = y ^ ((y >> mt.u) & mt.d)
	y = y ^ ((y << mt.s) & mt.b)
	y = y ^ ((y << mt.t) & mt.c)
	y = y ^ (y >> mt.l)
	mt.index++
	return uint32(y & (uint(^uint32(0))))
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
		mt.index = 0
	}
}

func (mt *Mersenne) untwist() {
    for i := uint(mt.n - 1); i <= 0; i-- {
        x := mt.state[i] ^ mt.a

    }
}
