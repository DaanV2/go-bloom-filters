package bloomhashes

import (
	"crypto/md5" // nolint:gosec // #nosec G401 -- This is safe because we are only using MD5 for hashing and not for cryptographic purposes.
	"hash/crc64"
)

// NOTE: no Adler since it only produces a 32-bit hash, which is not enough for our purposes.
// func Adler32(data []byte, hashes []uint64) int {
// }

// NOTE: no CRC-32 since it only produces a 32-bit hash, which is not enough for our purposes.
// func Crc32(data []byte, hashes []uint64) int {
// }

const CRC64_HASHES = 1 // The number of uint64 values that can be extracted from a CRC-64 hash.

func Crc64_ISO(data []byte, hashes []uint64) int {
	w := crc64.Checksum(data, crc64.MakeTable(crc64.ISO))

	if len(hashes) > 0 {
		hashes[0] = w

		return 1
	}

	return 0
}

const CRC64_ECMA_HASHES = 1 // The number of uint64 values that can be extracted from a CRC-64-ECMA hash.

func Crc64_ECMA(data []byte, hashes []uint64) int {
	w := crc64.Checksum(data, crc64.MakeTable(crc64.ECMA))

	if len(hashes) > 0 {
		hashes[0] = w

		return 1
	}

	return 0
}

const MD5_HASHES = md5.Size / uint64_size // The number of uint64 values that can be extracted from an MD5 hash.

func MD5(data []byte, hashes []uint64) int {
	b := md5.Sum(data) // nolint:gosec // #nosec G401 -- This is safe because we are only using MD5 for hashing and not for cryptographic purposes.

	return PutUint64(b[:], hashes)
}
