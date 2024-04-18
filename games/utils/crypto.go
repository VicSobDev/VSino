package utils

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
)

func Hash256(seed string) string {
	h := sha256.New()
	h.Write([]byte(seed))
	return hex.EncodeToString(h.Sum(nil))
}

func SaltHash(hash, salt string) string {
	h := hmac.New(sha256.New, []byte(hash))
	h.Write([]byte(salt))
	return hex.EncodeToString(h.Sum(nil))
}
