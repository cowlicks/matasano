package dh

import (
	"math/big"
)

type both struct {
	salt *big.Int
	x *big.Int
	N *big.Int
	g *big.Int
	k *big.Int
	u *big.Int
	a *big.Int
	b *big.Int
	A *big.Int
	B *big.Int
	S *big.Int
	K []byte
	password []byte
}

type Server struct {
	*both
	v *big.Int
	b *big.Int
}

type Client struct {
	*both
}


func NewSession(password []byte) (*Client, *Server) {
	bc := &both{salt: Rand(PNist), N: PNist, g: big.NewInt(2), k: big.NewInt(3), password: password}
	return &Client{both: bc}, NewServer(bc)
}

func NewServer(both *both) *Server {
	xH := HashBytes(append(both.salt.Bytes(), both.password...))
	x := new(big.Int).SetBytes(xH)
	v := Pow(both.g, x, both.N)
	return &Server{both: both, v: v}
}

func (c *Client) Send() *big.Int {
	c.a = Rand(c.N)
	c.A = Pow(c.g, c.a, c.N)
	return c.A
}

func (s *Server) Receive(A *big.Int) {
	s.A = A
}

func (s *Server) Send() (*big.Int, *big.Int) {
	s.b = Rand(s.N)
	s.B = Add(Mul(s.k, s.v, s.N), Pow(s.g, s.b, s.N), s.N)
	return s.salt, s.B
}

func (c *Client) Receive(salt, B *big.Int) {
	c.salt = salt
	c.B = B
}

func (b *both) GetU() {
	uH := HashBytes(append(b.A.Bytes(), b.B.Bytes()...))
	b.u = new(big.Int).SetBytes(uH)
}

func (c *Client) GetK() {
	xH := HashBytes(append(c.salt.Bytes(), c.password...))
	c.x = new(big.Int).SetBytes(xH)
	base := Sub(c.B, Mul(c.k, Pow(c.g, c.x, c.N), c.N), c.N)
	exp := Add(c.a, Mul(c.u, c.x, c.N), c.N)
	c.S = Pow(base, exp, c.N)
	c.K = HashBytes(c.S.Bytes())
	P(c.S)
}

func (s *Server) GetK() {
	s.S = Pow(Mul(s.A, Pow(s.v, s.u, s.N), s.N), s.b, s.N)
	s.K = HashBytes(s.S.Bytes())
	P(s.S)
}

/*
func Test36(t *testing.T) {
	// C & S
	k := big.NewInt(3)
	password := []byte("password")
	salt := Rand(p)
	// S
	xHserver := HashBytes(append(salt.Bytes(), password...))
	xserver := new(big.Int).SetBytes(xHserver)
	v := Pow(g, xserver, p)

	// C -> S
	a := Rand(p)
	A := Pow(g, a, p)

	// S -> C
	b := Rand(p)
	B := Add(Mul(k, v, p), Pow(g, b, p), p)

	// S, C
	uH := HashBytes(append(A.Bytes(), B.Bytes()...))
	u := new(big.Int).SetBytes(uH)

	// C
	// gets xH and convert to big.Int x
	xHclient := HashBytes(append(salt.Bytes(), password...))
	xclient := new(big.Int).SetBytes(xHclient)
	Sclient := Pow(Sub(B, Mul(k, Pow(g, xclient, p), p), p), Add(a, Mul(u, xclient, p), p), p)
	Kclient := Hash(Sclient)

	// S
	Sserver := Pow(Mul(A, Pow(v, u, p), p), b, p)
	Kserver := Hash(Sserver)

	// C -> S
	for i, _ := range Kclient {
		if Kclient[i] != Kserver[i] {
			t.Fail()
		}
	}
}
*/
