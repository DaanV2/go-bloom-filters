package xrand

import "crypto/rand"

// MustBytes generates n random bytes using cryptographically secure random number generation.
// It panics if an error occurs during random byte generation.
func MustBytes(n int) []byte {
	b, err := Bytes(n)
	if err != nil {
		panic(err)
	}

	return b
}

// Bytes generates n random bytes using cryptographically secure random number generation.
// It returns the random bytes and an error if the random byte generation fails.
func Bytes(n int) ([]byte, error) {
	b := make([]byte, n)
	_, err := rand.Read(b)
	if err != nil {
		return b, err
	}

	return b, nil
}
