package padding

import (
    "../util"
    "errors"
)

var InvalidPad = errors.New("Invalid Pad")

func Pad(bs int, data []byte) ([]byte, error) {
    if !((0 < bs) &&(bs < 256)) {
        return nil, errors.New("Invalid blocksize, must be between 0 - 256")
    }
    ldata := len(data)
    nblocks := (ldata / bs) + 1
    lpad := (bs * nblocks) - ldata
    if (ldata % bs) == 0 {
        lpad = bs
    }

    out := make([]byte, nblocks * bs)
    copy(out, data)
    for i := ldata; i < (ldata + lpad); i++ {
        out[i] = uint8(lpad)
    }

    return out, nil
}

func UnPad(bs int, data []byte) ([]byte, error) {
    errout := make([]byte, 0)
    if !((0 < bs) &&(bs < 256)) {
        return errout, errors.New("Invalid blocksize, must be between 0 - 256")
    }
    datalen := len(data)
    padval := int(data[datalen - 1])
    if padval > datalen {
        return errout, InvalidPad
    }
    exp_pad := make([]byte, padval)
    if padval == 0 {
        return errout, InvalidPad
    }
    for i := range exp_pad {
        exp_pad[i] = uint8(padval)
    }
    if !util.ByteEq(data[datalen - padval:], exp_pad) {
        return errout, InvalidPad
    }
    unpadded := make([]byte, datalen - padval)
    copy(unpadded, data[:datalen - padval])
    return unpadded, nil
}
