package bloomhashes

import "unsafe"

// ToUint64 tries to convert a byte slice to a slice of uint64 values.
// It processes the input byte slice in chunks of 8 bytes, converting each chunk into a uint64 value and appending it to the resulting slice. If the length of the input byte slice is not a multiple of 8, the remaining bytes are ignored.
func ToUint64(b []byte) []uint64 {
	if len(b) < 8 {
		return nil
	}

	l := len(b) / 8
	result := make([]uint64, l)
	_ = PutUint64(b, result)

	return result
}

func PutUint64(b []byte, values []uint64) int {
	if len(b) < 8 {
		return 0
	}
	l := min(len(b)/8, len(values))
	tmp := unsafe.Slice((*uint64)(unsafe.Pointer(&b[0])), l) // nolint:gosec // #nosec G103 -- This is safe because we are only writing to the byte slice and not reading from it.
	copy(tmp, values[:l])

	return l
}
