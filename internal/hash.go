package internal

import (
	"crypto/hmac"
	"crypto/sha256"
)

func HmacSha256(key, input []byte) []byte {
	h := hmac.New(sha256.New, key)
	h.Write(input)
	return h.Sum(nil)
}
