package internal

import (
	"crypto/sha512"
	"encoding/hex"
)

func HashData(jsonData string) string {
	hasher := sha512.New384()
	hasher.Write([]byte(jsonData))
	sign := hex.EncodeToString(hasher.Sum(nil))
	return sign
}
