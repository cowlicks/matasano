package dh

import (
	"math/big"
)

func mod(a *big.Int) *big.Int {
	return new(big.Int).Mod(a, PNist)
}

func pow(base, exp *big.Int) *big.Int {
	return mod(new(big.Int).Exp(base, exp, PNist))
}

func mul(a, b *big.Int) *big.Int {
	return mod(new(big.Int).Mul(a, b))
}

func add(a, b *big.Int) *big.Int {
	return mod(new(big.Int).Add(a, b))
}

func sub(a, b *big.Int) *big.Int {
	return mod(new(big.Int).Sub(a, b))
}

func fromBytes(b []byte) *big.Int {
	return mod(new(big.Int).SetBytes(b))
}

type both struct {
	salt *big.Int
	x    *big.Int
	N    *big.Int
	g    *big.Int
	k    *big.Int
	u    *big.Int

	A        *big.Int
	B        *big.Int
	S        *big.Int
	K        []byte
	password []byte
}

type Server struct {
	*both
	v *big.Int
	b *big.Int
}

type Client struct {
	*both
	a *big.Int
}

func NewSession(password []byte) (*Client, *Server) {
	bclient := &both{salt: Rand(PNist), N: PNist, g: big.NewInt(2), k: big.NewInt(3), password: password}
	xH := HashBytes(append(bclient.salt.Bytes(), bclient.password...))
	bclient.x = fromBytes(xH)
	//bserver := &both{}
	//*bserver = *bclient
	//return &Client{both: bclient}, NewServer(bserver)
	return &Client{both: bclient}, NewServer(bclient)
}

// make v = g**x
func NewServer(both *both) *Server {
	xH := HashBytes(append(both.salt.Bytes(), both.password...))
	x := fromBytes(xH)
	v := pow(both.g, x)
	return &Server{both: both, v: v}
}

// make & send A = g**a
func (c *Client) Send() *big.Int {
	c.a = Rand(c.N)
	c.A = pow(c.g, c.a)
	return c.A
}

func (s *Server) Receive(A *big.Int) {
	s.A = A
}

// make & send B = g**b
func (s *Server) Send() *big.Int {
	s.b = Rand(s.N)
	s.B = add(mul(s.k, s.v), pow(s.g, s.b))
	return s.B
}

func (c *Client) Receive(B *big.Int) {
	c.B = B
}

func (b *both) GetU() {
	uH := HashBytes(append(b.A.Bytes(), b.B.Bytes()...))
	b.u = fromBytes(uH)
}

func (c *Client) GetK() {
	c.GetU()
	base := sub(c.B, mul(c.k, pow(c.g, c.x)))
	exp := add(c.a, mul(c.u, c.x))
	c.S = pow(base, exp)
	c.K = HashBytes(c.S.Bytes())
}

func (s *Server) GetK() {
	s.GetU()
	s.S = pow(mul(s.A, pow(s.v, s.u)), s.b)
	s.K = HashBytes(s.S.Bytes())
}
