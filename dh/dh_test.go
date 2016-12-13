package dh

import (
	"../aesmodes"
	"testing"
//	"math/big"
)

func Test(t *testing.T) {
	a := Rand(g)
	b := Rand(g)
	A := Pow(g, a, p)
	B := Pow(g, b, p)
	r := Pow(A, b, p)
	s := Pow(B, a, p)
	if r.Cmp(s) != 0 {
		t.Fail()
	}
}

func TestMITM(t *testing.T) {
	alice := NewPerson()
	bob := NewPerson()

	//A := alice.Send()
	bob.Receive(p)

	//B := bob.Send()
	alice.Receive(p)

	secreta := Hash(Pow(alice.Send(), p, p))[:16]
	secretb := Hash(Pow(bob.Send(), p, p))[:16]

	//act := alice.Encrypt([]byte("tell me about the library"))
	msga := []byte("Tell me about the library")
	msgb := []byte("Hello Alexa")
	cta := alice.Encrypt(msga)
	ctb := bob.Encrypt(msgb)

	resa, err := aesmodes.DecryptCBC(secreta, cta)
	resb, err := aesmodes.DecryptCBC(secretb, ctb)
	if err != nil {
		t.Fail()
	}

	for i, _ := range resa {
		if resa[i] != msga[i] {
			t.Fail()
		}
	}

	for i, _ := range resb {
		if resb[i] != msgb[i] {
			t.Fail()
		}
	}
}


func TestSRP(t *testing.T) {
	client, server := NewSession([]byte("waddup"))
	server.Receive(client.Send())
	client.Receive(server.Send())

	client.GetU()
	server.GetU()

	client.GetK()
	server.GetK()
}
