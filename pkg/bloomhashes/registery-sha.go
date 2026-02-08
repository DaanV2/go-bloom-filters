package bloomhashes

import (
	"crypto/sha1" // nolint:gosec // #nosec G401 -- This is safe because we are only using SHA-1 for hashing and not for cryptographic purposes.
	"crypto/sha256"
	"crypto/sha3"
	"crypto/sha512"
)

// SHA hashes

const SHA1_HASHES = sha1.Size / uint64_size // The number of uint64 values that can be extracted from a SHA-1 hash.

func Sha1(data []byte, hashes []uint64) int {
	b := sha1.Sum(data) // nolint:gosec // #nosec G401 -- This is safe because we are only using SHA-1 for hashing and not for cryptographic purposes.

	return PutUint64(b[:], hashes)
}

const SHA3_224_HASHES = 28 / uint64_size // The number of uint64 values that can be extracted from a SHA3-224 hash.

func Sha224(data []byte, hashes []uint64) int {
	b := sha256.Sum224(data)

	return PutUint64(b[:], hashes)
}

const SHA3_384_HASHES = 48 / uint64_size // The number of uint64 values that can be extracted from a SHA3-384 hash.

func Sha3_384(data []byte, hashes []uint64) int {
	b := sha3.Sum384(data)

	return PutUint64(b[:], hashes)
}

const SHA256_HASHES = sha256.Size / uint64_size // The number of uint64 values that can be extracted from a SHA-256 hash.

func Sha256(data []byte, hashes []uint64) int {
	b := sha256.Sum256(data)

	return PutUint64(b[:], hashes)
}

const SHA3_512_HASHES = sha512.Size / uint64_size // The number of uint64 values that can be extracted from a SHA3-512 hash.

func Sha512(data []byte, hashes []uint64) int {
	b := sha512.Sum512(data)

	return PutUint64(b[:], hashes)
}

