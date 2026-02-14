package xrand

func MustString(n int) string {
	s, err := String(n)
	if err != nil {
		panic(err)
	}

	return s
}

func String(n int) (string, error) {
	b, err := Bytes(n)
	if err != nil {
		return "", err
	}

	return string(b), nil
}
