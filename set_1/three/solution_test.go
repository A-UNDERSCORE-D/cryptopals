package three

import (
	"sort"
	"testing"

	"awesome-dragon.science/go/cryptopals/set_1/letterFrequency"
	"awesome-dragon.science/go/cryptopals/set_1/one"
)

const target = "1b37373331363f78151b7f2b783431333d78397828372d363c78373e783a393b3736"

func TestSingleByteXOR(t *testing.T) {
	decoded := one.MustDecodeHex(target)
	pairs := []struct {
		f float64
		s string
	}{}

	for _, v := range BruteSingleByteXOR(decoded) {
		pairs = append(pairs, struct {
			f float64
			s string
		}{
			f: letterFrequency.Frequency(string(v)),
			s: string(v),
		})
	}

	sort.Slice(pairs, func(i, j int) bool { return pairs[i].f > pairs[j].f })
	if pairs[0].s != "Cooking MC's like a pound of bacon" {
		t.FailNow()
	}
}
