package sha1

func BadMac(key, message []byte) []byte {
	return Sha1(append(key, message...))
}
