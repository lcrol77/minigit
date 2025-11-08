package internal

import "crypto/sha1"

func computeHash(data *[]byte) []byte {
	h := sha1.New()
	h.Write(*data)
	sum := h.Sum(nil)
	return sum
}
