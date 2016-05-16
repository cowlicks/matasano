package one

import (
	"bytes"
	"encoding/base64"
	"encoding/hex"
	"io/ioutil"
	"testing"
)

func testEq(a, b []byte) bool {

	if a == nil && b == nil {
		return true
	}

	if a == nil || b == nil {
		return false
	}

	if len(a) != len(b) {
		return false
	}

	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}

	return true
}

func assert(a, b []byte) {
	x := testEq(a, b)
	if !x {
		panic("assert fail")
	}
}

func TestXor(t *testing.T) {
	in1, _ := hex.DecodeString("1c0111001f010100061a024b53535009181c")
	in2, _ := hex.DecodeString("686974207468652062756c6c277320657965")
	out, _ := hex.DecodeString("746865206b696420646f6e277420706c6179")

	res, err := Xor(in1, in2)

	if err != nil {
		t.Error(err)
	}

	if !testEq(res, out) {
		t.Fatal("fail")
	}
}

func TestHamming(t *testing.T) {
	in1 := []byte("this is a test")
	in2 := []byte("wokka wokka!!!")
	out := 37

	res, err := HammingDistance(in1, in2)
	if err != nil {
		t.Error(err)
	}

	if res != out {
		t.Fatal(res, out)
	}
}

/*
func TestKeySizes(t * testing.T) {
    file_bytes, _ := ioutil.ReadFile("data6.txt")
    // strip \n
    file_bytes = bytes.Replace(file_bytes, []byte("\n"), []byte(""), -1)
    FindKeySize(file_bytes)
}
*/

func TestAllPrintable(t *testing.T) {
	in := []byte("All thesee chars are printabel")
	if !AllPrintable(in) {
		t.Fatal()
	}
}

func TestMyVig(t *testing.T) {
    var plaintext_n []byte = []byte(`
    Lol this songs is funny
    As he came into the window
    Was a sound of a crescendo
    He came into her apartment
    He left the bloodstains on the carpet
    She was sitting at the table
    He could see she was unable
    So she ran into the bedroom
    She was struck down
    It was her doom
    Annie, are you OK
    Are you OK
    Are you OK, Annie
    Annie, are you OK
    You OK
    Are you OK, Annie
    Annie, are you OK
    You OK
    You OK, Annie
    Annie, are you OK
    You OK
    Are you OK, Annie
    Annie, are you OK
    Will you tell us that you're OK
    There's a sign at the window
    That he struck you
    A crescendo, Annie
    He came into your apartment
    He left the bloodstains on the carpet
    Then you ran into the bedroom
    You were struck down
    It was your doom
    Annie, are you OK
    You OK
    Are you OK, Annie
    Annie, are you OK
    You OK
    Are you OK, Annie
    Annie, are you OK
    You OK
    Are you OK, Annie
    You've been hit by
    You've been struck by
    A smooth criminal
    So they came into the outway
    It was Sunday
    What a black day
    I could feel your salutation
    Sounding heartbeats
    Intimidations
    Annie, are you OK
    You OK
    Are you OK, Annie
    Annie, are you OK
    You OK
    Are you OK, Annie
    Annie, are you OK
    You OK
    Are you OK, Annie
    Annie, are you OK
    You OK
    Are you OK, Annie
    Annie, are you OK
    Will you tell us that you're OK
    There's a sign at the window
    That he struck you
    A crescendo, Annie
    He came into your apartment
    He left the bloodstains on the carpet
    Then you ran into the bedroom
    You were struck down
    It was your doom
    Annie, are you OK
    You OK
    Are you OK, Annie
    You've been hit by
    You've been struck by
    A smooth criminal
    Annie, are you OK
    Will you tell us that you're OK
    There's a sign at the window
    That he struck you
    A crescendo, Annie
    He came into your apartment
    He left the bloodstains on the carpet
    Then you ran into the bedroom
    You were struck down
    It was your doom
    Annie, are you OK
    You OK
    Are you OK, Annie
    Annie, are you OK
    You OK
    Are you OK, Annie
    Annie, are you OK
    You OK
    Are you OK, Annie
    Annie, are you OK
    You OK
    Are you OK, Annie
    As he came into the window
    Was a sound of a crescendo
    He came into her apartment
    He left the bloodstains on the carpet
    She was sitting at the table
    He could see she was unable
    So she ran into the bedroom
    She was struck down
    It was her doom
    Annie, are you OK
    Are you OK
    Are you OK, Annie
    Annie, are you OK
    lol`)
    var plaintext []byte = bytes.Replace(plaintext_n, []byte("\n"), []byte(""), -1)

    var key []byte = []byte("Alien ant farm")
    var ciphertext []byte = VectorXor(key, plaintext)

	pt1 := []byte("foobarqux")
	key1 := []byte("xofoo")
	ct1 := VectorXor(key1, pt1)
	res1 := VectorXor(key1, ct1)
	if string(res1) != string(pt1) {
		t.Fatal()
	}

	res2 := VectorXor(key, ciphertext)
	if string(res2) != string(plaintext) {
		t.Fatal()
	}

	res3 := VectorXor(key, ciphertext[:len(key)])
	if !testEq(res3, plaintext[:len(key)]) {
		t.Fatal()
	}

	ct2 := make([]byte, 20)
	for i := range ct2 {
		ct2[i] = plaintext[i]
		ct2[i] ^= uint8('a')
	}
	res4 := VectorXor([]byte{'a'}, ct2)
	if !testEq(res4, plaintext[:20]) {
		t.Fatal()
	}

	r, _, _ := CrackVigenere(ciphertext)
	if !testEq(r, plaintext) {
		t.Fatal()
	}
}

func TestCryptopalskVigenere(t *testing.T) {
	file_bytes, _ := ioutil.ReadFile("data6.txt")
	file_bytes, err := base64.StdEncoding.DecodeString(string(file_bytes))
	if err != nil {
		t.Fatal()
	}

	_, k, err := CrackVigenere(file_bytes)
	if err != nil {
		t.Fatal()
	}

	if string(k) != "Terminator X: Bring the noise" {
		t.Fatal()
	}
}
