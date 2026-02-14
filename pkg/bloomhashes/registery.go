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

func Crc64_ISO(data []byte) uint64 {
	return crc64.Checksum(data, crc64.MakeTable(crc64.ISO))
}

func Crc64_ECMA(data []byte) uint64 {
	return crc64.Checksum(data, crc64.MakeTable(crc64.ECMA))
}

func MD5(data []byte) uint64 {
	b := md5.Sum(data) // nolint:gosec // #nosec G401 -- This is safe because we are only using MD5 for hashing and not for cryptographic purposes.

	return bytesToUint64(b[:])
}
