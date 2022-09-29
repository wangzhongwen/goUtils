package goUtils

import (
	"crypto/sha1"
	"encoding/hex"
	"io"
	"os"
)

func Sha1File(path string) (string, error) {
	file, err := os.Open(path)
	if err != nil {
		return "", err
	}
	defer file.Close()
	h := sha1.New()
	_, erro := io.Copy(h, file)
	if erro != nil {
		return "", err
	}
	return hex.EncodeToString(h.Sum(nil)), nil
}
