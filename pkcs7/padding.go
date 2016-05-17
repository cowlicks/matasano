package padding

import (
    "errors"
)

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
