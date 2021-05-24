package four

import (
	"bytes"
	"os"
	"testing"

	"awesome-dragon.science/go/cryptopals/set_1/one"
)

func TestIsSingleCharXOR(t *testing.T) {
	lines, err := os.ReadFile("./input.txt")
	if err != nil {
		panic(err)
	}
	split := bytes.Split(lines, []byte{'\n'})
	good := []string{}

	for _, line := range split {
		decoded := one.MustDecodeHex(string(line))
		res, ok := IsSingleCharXOR(decoded)
		if ok {
			good = append(good, string(res))
		}
	}

	if len(good) != 1 {
		t.Fatal("Unexpected number of results", good)
	}

	if good[0] != "Now that the party is jumping\n" {
		t.Fatal("Unexpected result", good[0], good)
	}
}
