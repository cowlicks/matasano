package sha1

var KEY = []byte("super secret")

func BadMac(key, message []byte) []byte {
	return Sha1(append(key, message...))
}

func VerifyBadMac(key, message, mac []byte) bool {
	expected := BadMac(key, message)
	out := true
	for i := 0; i < len(expected); i++ {
		if expected[i] != mac[i] {
			out = false
		}
	}
	return out
}

func CheckMac(message, mac []byte) bool {
	return VerifyBadMac(KEY, message, mac)
}

func GetMac(message []byte) []byte {
	return BadMac(KEY, message)
}
