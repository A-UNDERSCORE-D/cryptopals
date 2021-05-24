package three

import (
	"bytes"

	"awesome-dragon.science/go/cryptopals/set_1/two"
)

func SingleByteXOR(b byte, plaintext []byte) []byte {
	toUse := bytes.Repeat([]byte{b}, len(plaintext))
	res, err := two.FixedXOR(toUse, plaintext)
	if err != nil {
		panic(err)
	}
	return res
}

func BruteSingleByteXOR(plaintext []byte) [][]byte {
	out := make([][]byte, 256)
	for i := 0; i < 256; i++ {
		out[i] = SingleByteXOR(byte(i), plaintext)
	}
	return out
}
