package xrand

// MustString generates a random string of n bytes using cryptographically secure random bytes.
// It panics if an error occurs during random byte generation.
func MustString(n int) string {
	s, err := String(n)
	if err != nil {
		panic(err)
	}

	return s
}

// String generates a random string of n bytes using cryptographically secure random bytes.
// It returns an error if random byte generation fails.
func String(n int) (string, error) {
	b, err := Bytes(n)
	if err != nil {
		return "", err
	}

	return string(b), nil
}
