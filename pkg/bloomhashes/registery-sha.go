package bloomhashes

import (
	"crypto/sha1" // nolint:gosec // #nosec G401 -- This is safe because we are only using SHA-1 for hashing and not for cryptographic purposes.
	"crypto/sha256"
	"crypto/sha3"
	"crypto/sha512"
)

// SHA hashes

// Sha1 computes a SHA-1 hash and converts it to a uint64.
// It XORs portions of the hash to create a 64-bit value.
// Note: SHA-1 is used for hashing purposes only, not for cryptographic security.
func Sha1(data []byte) uint64 {
	b := sha1.Sum(data) // nolint:gosec // #nosec G401 -- This is safe because we are only using SHA-1 for hashing and not for cryptographic purposes.

	return bytesToUint64(b[:]) ^ bytesToUint64(b[8:])
}

// Sha224 computes a SHA-224 hash and converts it to a uint64.
// It XORs portions of the hash to create a 64-bit value.
func Sha224(data []byte) uint64 {
	b := sha256.Sum224(data)

	return bytesToUint64(b[:]) ^ bytesToUint64(b[8:]) ^ bytesToUint64(b[16:])
}

// Sha3_384 computes a SHA3-384 hash and converts it to a uint64.
// It XORs portions of the hash to create a 64-bit value.
func Sha3_384(data []byte) uint64 {
	b := sha3.Sum384(data)

	return bytesToUint64(b[:]) ^ bytesToUint64(b[8:]) ^ bytesToUint64(b[16:]) ^ bytesToUint64(b[24:]) ^ bytesToUint64(b[32:]) ^ bytesToUint64(b[40:])
}

// Sha256 computes a SHA-256 hash and converts it to a uint64.
// It XORs portions of the hash to create a 64-bit value.
func Sha256(data []byte) uint64 {
	b := sha256.Sum256(data)

	return bytesToUint64(b[:]) ^ bytesToUint64(b[8:]) ^ bytesToUint64(b[16:]) ^ bytesToUint64(b[24:])
}

// Sha512 computes a SHA-512 hash and converts it to a uint64.
// It XORs portions of the hash to create a 64-bit value.
func Sha512(data []byte) uint64 {
	b := sha512.Sum512(data)

	return bytesToUint64(b[:]) ^ bytesToUint64(b[8:]) ^ bytesToUint64(b[16:]) ^ bytesToUint64(b[24:]) ^ bytesToUint64(b[32:]) ^ bytesToUint64(b[40:]) ^ bytesToUint64(b[48:]) ^ bytesToUint64(b[56:])
}
