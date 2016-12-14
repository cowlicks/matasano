package dh

import (
	"../aesmodes"
	"crypto/rand"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"math/big"
)

func P(i ...interface{}) {
	fmt.Println(i)
}

var p *big.Int
var PNist *big.Int
var g = big.NewInt(2)

func init() {
	phex := `ffffffffffffffffc90fdaa22168c234c4c6628b80dc1cd129024e088a67cc74020bbea63b139b22514a08798e3404ddef9519b3cd3a431b302b0a6df25f14374fe1356d6d51c245e485b576625e7ec6f44c42e9a637ed6b0bff5cb6f406b7edee386bfb5a899fa5ae9f24117c4b1fe649286651ece45b3dc2007cb8a163bf0598da48361c55d39a69163fa8fd24cf5f83655d23dca3ad961c62f356208552bb9ed529077096966d670c354e4abc9804f1746c08ca237327ffffffffffffffff`
	pbytes, err := hex.DecodeString(phex)
	if err != nil {
		panic(err)
	}
	p = new(big.Int).SetBytes(pbytes)
	PNist = p
}

func Pow(base, exp, mod *big.Int) *big.Int {
	return new(big.Int).Exp(base, exp, mod)
}

func Mul(a, b, mod *big.Int) *big.Int {
	return new(big.Int).Mod(new(big.Int).Mul(a, b), mod)
}

func Add(a, b, mod *big.Int) *big.Int {
	return new(big.Int).Mod(new(big.Int).Add(a, b), mod)
}

func Sub(a, b, mod *big.Int) *big.Int {
	return new(big.Int).Mod(new(big.Int).Sub(a, b), mod)
}

func Div(a, b, mod *big.Int) *big.Int {
	return Mul(a, new(big.Int).ModInverse(b, mod), mod)
}

func Rand(mod *big.Int) *big.Int {
	a, err := rand.Int(rand.Reader, mod)
	if err != nil {
		panic(err)
	}

	return a
}

func Hash(k *big.Int) []byte {
	return HashBytes(k.Bytes())
}

func HashBytes(b []byte) []byte {
	hasher := sha256.New()
	return hasher.Sum(b)
}

type Person struct {
	p      *big.Int
	g      *big.Int
	self   *big.Int
	other  *big.Int
	secret *big.Int
	key    []byte
}

func NewPerson() *Person {
	self := Rand(p)
	return &Person{p: p, g: g, self: self}
}

func (pers *Person) Send() *big.Int {
	return Pow(pers.g, pers.self, pers.p)
}

func (pers *Person) Receive(b *big.Int) {
	pers.other = b
	pers.secret = Pow(Pow(pers.g, pers.self, pers.p), pers.other, pers.p)
	pers.key = Hash(pers.secret)[:16]
}

func (pers *Person) Encrypt(msg []byte) []byte {
	iv := make([]byte, 16)
	rand.Reader.Read(iv)
	ct, err := aesmodes.EncryptCBCWithIV(pers.key, iv, msg)
	if err != nil {
		panic(err)
	}
	return ct
}
