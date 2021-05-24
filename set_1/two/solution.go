package two

import (
	"errors"
)

func FixedXOR(buf1, buf2 []byte) ([]byte, error) {
	if len(buf1) != len(buf2) {
		return nil, errors.New("buffer sizes do not match")
	}

	res := make([]byte, len(buf1))
	for i := range buf1 {
		res[i] = buf1[i] ^ buf2[i]
	}

	return res, nil
}
