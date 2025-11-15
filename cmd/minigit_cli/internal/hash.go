package internal

import "crypto/sha1"

func ComputeHash(data []byte) []byte {
	h := sha1.New()
	h.Write(data)
	sum := h.Sum(nil)
	return sum
}
