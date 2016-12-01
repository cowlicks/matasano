package sha1

import (
	"errors"
)

func ExtendMac(msg, mac, suffix []byte) (newmsg, newmac []byte, err error) {
	for i := 0; i < 512; i++ {
		dummykey := make([]byte, i)
		gluepad := MDPad(append(dummykey, msg...))
		endpad := MDPad(append(dummykey, append(msg, append(gluepad, suffix...)...)...))

		h0, h1, h2, h3, h4 := GetRegisters(mac)

		newmac := Extend(append(suffix, endpad...), h0, h1, h2, h3, h4)

		newmsg := append(msg, append(gluepad, suffix...)...)
		if CheckMac(newmsg, newmac) {
			return newmsg, newmac, nil
		}
	}
	return nil, nil, errors.New("failed")
}
