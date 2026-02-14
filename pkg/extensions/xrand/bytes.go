package xrand

import "crypto/rand"

func MustBytes(n int) []byte {
	b, err := Bytes(n)
	if err != nil {
		panic(err)
	}

	return b
}

func Bytes(n int) ([]byte, error) {
	b := make([]byte, n)
	_, err := rand.Read(b)
	if err != nil {
		return b, err
	}

	return b, nil
}
