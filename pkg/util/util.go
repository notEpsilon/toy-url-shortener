package util

import (
	"crypto/rand"
	"encoding/hex"
)

func GetRandomSlug() string {
	buff := make([]byte, 4)
	rand.Read(buff)
	return hex.EncodeToString(buff)
}
