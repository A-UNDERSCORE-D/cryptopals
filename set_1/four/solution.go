package four

import (
	"awesome-dragon.science/go/cryptopals/set_1/letterFrequency"
	"awesome-dragon.science/go/cryptopals/set_1/three"
)

func IsSingleCharXOR(data []byte) ([]byte, bool) {
	for _, v := range three.BruteSingleByteXOR(data) {
		if letterFrequency.Frequency(string(v)) > 90 {
			return v, true
		}
	}

	return nil, false
}
