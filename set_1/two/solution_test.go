package two

import (
	"encoding/hex"
	"testing"

	"awesome-dragon.science/go/cryptopals/set_1/one"
)

const (
	testStr1 = "1c0111001f010100061a024b53535009181c"
	testStr2 = "686974207468652062756c6c277320657965"
	result   = "746865206b696420646f6e277420706c6179"
)

func TestFixedXOR(t *testing.T) {
	res, err := FixedXOR(one.MustDecodeHex(testStr1), one.MustDecodeHex(testStr2))
	if err != nil {
		t.Fatal(err)
	}

	if strRes := hex.EncodeToString(res); strRes != result {
		t.Fatalf("%q != %q", strRes, result)
	}
}
