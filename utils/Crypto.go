package utils

import (
	"crypto/sha1"
	"encoding/base64"
	"encoding/hex"
	"math/rand"
	"time"
)

func Hash(text string) string {
	h := sha1.New()
	h.Write([]byte(text))
	bs := hex.EncodeToString(h.Sum(nil))
	return string(bs)
}

func EncodeBase64(b []byte) string {
	return base64.StdEncoding.EncodeToString(b)
}

func DecodeBase64(s string) ([]byte, error) {
	data, err := base64.StdEncoding.DecodeString(s)
	if err != nil {
		return nil, err
	}
	return data, nil
}

func RandomString(strlen int) string {
	var r = rand.New(rand.NewSource(time.Now().UnixNano()))
	const chars = "abcdefghijklmnopqrstuvwxyz0123456789"
	result := make([]byte, strlen)
	for i := range result {
		result[i] = chars[r.Intn(len(chars))]
	}
	return string(result)
}
