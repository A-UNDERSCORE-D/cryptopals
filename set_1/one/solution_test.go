package one

import (
	"testing"
)

const (
	hexTest = "49276d206b696c6c696e6720796f757220627261696e206c696b65206120706f69736f6e6f7573206d757368726f6f6d"
	b64     = "SSdtIGtpbGxpbmcgeW91ciBicmFpbiBsaWtlIGEgcG9pc29ub3VzIG11c2hyb29t"
)

func TestHex2Base64(t *testing.T) {
	res, err := Hex2Base64(hexTest)
	if err != nil {
		t.Fatal(err)
	}

	if res != b64 {
		t.Fatalf("%q != %q", res, b64)
	}
}

func TestBase642Hex(t *testing.T) {
	res, err := Base642Hex(b64)
	if err != nil {
		t.Fatal(err)
	}

	if res != hexTest {
		t.Fatalf("%q != %q", res, hexTest)
	}
}
