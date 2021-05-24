package one

import (
	"encoding/base64"
	"encoding/hex"
)

func Hex2Base64(hexStr string) (string, error) {
	res, err := hex.DecodeString(hexStr)
	if err != nil {
		return "", err
	}

	return base64.StdEncoding.EncodeToString(res), nil
}

func Base642Hex(b64 string) (string, error) {
	res, err := base64.StdEncoding.DecodeString(b64)
	if err != nil {
		return "", err
	}

	return hex.EncodeToString(res), nil
}

func MustDecodeHex(hexStr string) []byte {
	res, err := hex.DecodeString(hexStr)
	if err != nil {
		panic(err)
	}
	return res
}
