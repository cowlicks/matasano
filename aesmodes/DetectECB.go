package aesmodes

func DetectECB(blocksize int, ciphertext []byte) bool {
	counter := make(map[string]int)

	ctlen := len(ciphertext)
	nblocks := ctlen / blocksize
	if ctlen%blocksize != 0 {
		nblocks += 1
	}

	for i := 0; i < nblocks; i++ {
		ll := blocksize * i
		ul := blocksize * (i + 1)
		if ul > ctlen {
			ul = ctlen
		}

		k := string(ciphertext[ll:ul])
		counter[k] += 1
		if counter[k] > 1 {
			return true
		}
	}
	return false
}
