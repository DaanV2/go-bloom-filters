package bloomhashes

import (
	"crypto/md5" // nolint:gosec // #nosec G401 -- This is safe because we are only using MD5 for hashing and not for cryptographic purposes.
	"hash/crc64"
)

// NOTE: no Adler since it only produces a 32-bit hash, which is not enough for our purposes.
// func Adler32(data []byte) uint64 {
// }

// NOTE: no CRC-32 since it only produces a 32-bit hash, which is not enough for our purposes.
// func Crc32(data []byte) uint64 {
// }

// Crc64_ISO computes a CRC-64 hash using the ISO polynomial.
// It returns a 64-bit hash value suitable for use in bloom filters.
func Crc64_ISO(data []byte) uint64 {
	return crc64.Checksum(data, crc64.MakeTable(crc64.ISO))
}

// Crc64_ECMA computes a CRC-64 hash using the ECMA polynomial.
// It returns a 64-bit hash value suitable for use in bloom filters.
func Crc64_ECMA(data []byte) uint64 {
	return crc64.Checksum(data, crc64.MakeTable(crc64.ECMA))
}

// MD5 computes an MD5 hash and converts it to a uint64.
// It returns the first 8 bytes of the MD5 hash as a 64-bit value.
// Note: MD5 is used for hashing purposes only, not for cryptographic security.
func MD5(data []byte) uint64 {
	b := md5.Sum(data) // nolint:gosec // #nosec G401 -- This is safe because we are only using MD5 for hashing and not for cryptographic purposes.

	return bytesToUint64(b[:])
}
