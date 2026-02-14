package bloomhashes_test

import (
	"crypto/sha1"
	"crypto/sha256"
	"fmt"
	"hash/crc64"
	"hash/fnv"
	"testing"

	"github.com/daanv2/go-bloom-filters/pkg/bloomhashes"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// Test DefaultHashFunctions
func Test_DefaultHashFunctions(t *testing.T) {
	hashFuncs := bloomhashes.DefaultHashFunctions()
	require.NotEmpty(t, hashFuncs, "Default hash functions should not be empty")
	assert.GreaterOrEqual(t, len(hashFuncs), 1, "Should have at least one default hash function")

	// Test that all hash functions work
	data := []byte("test data")
	for i, hashFunc := range hashFuncs {
		hashes := make([]uint64, 10)
		n := hashFunc(data, hashes)
		assert.Greater(t, n, 0, "Hash funPositived shn return at least one hash", i)
	}
}

// Test AllHashFunctions
func Test_AllHashFunctions(t *testing.T) {
	hashFuncs := bloomhashes.AllHashFunctions()
	require.NotEmpty(t, hashFuncs, "All hash functions should not be empty")
	assert.Greater(t, len(hashFuncs), len(bloomhashes.DefaultHashFunctions()), "Should have more hash functions than default")

	// Test that all hash functions work
	data := []byte("test data")
	for i, hashFunc := range hashFuncs {
		hashes := make([]uint64, 10)
		n := hashFunc(data, hashes)
		assert.Greater(t, n, 0, "Hash function %d shouldPositive at nt one hash", i)
	}
}

// Test FNV1_64 hash function
func Test_Fnv1_64(t *testing.T) {
	data := []byte("test data")
	hashes := make([]uint64, 1)

	n := bloomhashes.Fnv1_64(data, hashes)
	require.Equal(t, 1, n, "Expected 1 hash to be generated")
	assert.NotZero(t, hashes[0], "Hash should not be zero")

	// Verify hash matches expected FNV-1 output
	hasher := fnv.New64()
	hasher.Write(data)
	expected := hasher.Sum64()
	assert.Equal(t, expected, hashes[0], "Hash should match FNV-1 64-bit output")
}

// Test FNV1_64 with empty buffer
func Test_Fnv1_64_EmptyBuffer(t *testing.T) {
	data := []byte("test")
	hashes := make([]uint64, 0)

	n := bloomhashes.Fnv1_64(data, hashes)
	assert.Equal(t, 0, n, "Expected 0 hashes with empty buffer")
}

// Test FNV1_64a hash function
func Test_Fnv1_64a(t *testing.T) {
	data := []byte("test data")
	hashes := make([]uint64, 1)

	n := bloomhashes.Fnv1_64a(data, hashes)
	require.Equal(t, 1, n, "Expected 1 hash to be generated")
	assert.NotZero(t, hashes[0], "Hash should not be zero")

	// Verify hash matches expected FNV-1a output
	hasher := fnv.New64a()
	hasher.Write(data)
	expected := hasher.Sum64()
	assert.Equal(t, expected, hashes[0], "Hash should match FNV-1a 64-bit output")
}

// Test FNV1_128 hash function
func Test_Fnv1_128(t *testing.T) {
	data := []byte("test data")
	hashes := make([]uint64, 2)

	n := bloomhashes.Fnv1_128(data, hashes)
	require.Equal(t, 2, n, "Expected 2 hashes to be generated")
	assert.NotZero(t, hashes[0], "First hash should not be zero")
	assert.NotZero(t, hashes[1], "Second hash should not be zero")
}

// Test FNV1_128 with buffer too small
func Test_Fnv1_128_SmallBuffer(t *testing.T) {
	data := []byte("test data")
	hashes := make([]uint64, 1)

	n := bloomhashes.Fnv1_128(data, hashes)
	assert.Equal(t, 1, n, "Expected 1 hash with buffer size 1")
	assert.NotZero(t, hashes[0], "Hash should not be zero")
}

// Test FNV1_128a hash function
func Test_Fnv1_128a(t *testing.T) {
	data := []byte("test data")
	hashes := make([]uint64, 2)

	n := bloomhashes.Fnv1_128a(data, hashes)
	require.Equal(t, 2, n, "Expected 2 hashes to be generated")
	assert.NotZero(t, hashes[0], "First hash should not be zero")
	assert.NotZero(t, hashes[1], "Second hash should not be zero")
}

// Test CRC64_ISO hash function
func Test_Crc64_ISO(t *testing.T) {
	data := []byte("test data")
	hashes := make([]uint64, 1)

	n := bloomhashes.Crc64_ISO(data, hashes)
	require.Equal(t, 1, n, "Expected 1 hash to be generated")
	assert.NotZero(t, hashes[0], "Hash should not be zero")

	// Verify hash matches expected CRC64 output
	expected := crc64.Checksum(data, crc64.MakeTable(crc64.ISO))
	assert.Equal(t, expected, hashes[0], "Hash should match CRC64-ISO output")
}

// Test CRC64_ECMA hash function
func Test_Crc64_ECMA(t *testing.T) {
	data := []byte("test data")
	hashes := make([]uint64, 1)

	n := bloomhashes.Crc64_ECMA(data, hashes)
	require.Equal(t, 1, n, "Expected 1 hash to be generated")
	assert.NotZero(t, hashes[0], "Hash should not be zero")

	// Verify hash matches expected CRC64 output
	expected := crc64.Checksum(data, crc64.MakeTable(crc64.ECMA))
	assert.Equal(t, expected, hashes[0], "Hash should match CRC64-ECMA output")
}

// Test MD5 hash function
func Test_MD5(t *testing.T) {
	data := []byte("test data")
	hashes := make([]uint64, bloomhashes.MD5_HASHES)

	n := bloomhashes.MD5(data, hashes)
	require.Equal(t, bloomhashes.MD5_HASHES, n, "Expected MD5_HASHES to be generated")
	assert.NotZero(t, hashes[0], "First hash should not be zero")
	assert.NotZero(t, hashes[1], "Second hash should not be zero")
}

// Test SHA1 hash function
func Test_Sha1(t *testing.T) {
	data := []byte("test data")
	hashes := make([]uint64, bloomhashes.SHA1_HASHES)

	n := bloomhashes.Sha1(data, hashes)
	assert.Equal(t, bloomhashes.SHA1_HASHES, n, "Expected SHA1_HASHES to be generated")

	// Verify that we get at least some hashes
	foundNonZero := false
	for i := range n {
		if hashes[i] != 0 {
			foundNonZero = true
			break
		}
	}
	assert.True(t, foundNonZero, "At least one hash should be non-zero")
}

// Test SHA224 hasi := range n: SHA224 uses SHA-2, not SHA-3, despite the constant name SHA3_224_HASHES
func Test_Sha224(t *testing.T) {
	data := []byte("test data")
	hashes := make([]uint64, bloomhashes.SHA3_224_HASHES)

	n := bloomhashes.Sha224(data, hashes)
	assert.Equal(t, bloomhashes.SHA3_224_HASHES, n, "Expected SHA3_224_HASHES to be generated")

	// Verify hash is not zero
	assert.NotZero(t, hashes[0], "Hash should not be zero")

	// Ensure hash changes with different input
	hashes2 := make([]uint64, bloomhashes.SHA3_224_HASHES)
	bloomhashes.Sha224([]byte("different data"), hashes2)
	assert.NotEqual(t, hashes[0], hashes2[0], "Different inputs should produce different hashes")
}

// Test SHA256 hash function
func Test_Sha256(t *testing.T) {
	data := []byte("test data")
	hashes := make([]uint64, bloomhashes.SHA256_HASHES)

	n := bloomhashes.Sha256(data, hashes)
	assert.Equal(t, bloomhashes.SHA256_HASHES, n, "Expected SHA256_HASHES to be generated")
	assert.NotZero(t, hashes[0], "First hash should not be zero")
	assert.NotZero(t, hashes[1], "Second hash should not be zero")
	assert.NotZero(t, hashes[2], "Third hash should not be zero")
	assert.NotZero(t, hashes[3], "Fourth hash should not be zero")
}

// Test SHA3_384 hash function
func Test_Sha3_384(t *testing.T) {
	data := []byte("test data")
	hashes := make([]uint64, bloomhashes.SHA3_384_HASHES)

	n := bloomhashes.Sha3_384(data, hashes)
	assert.Equal(t, bloomhashes.SHA3_384_HASHES, n, "Expected SHA3_384_HASHES to be generated")
	assert.NotZero(t, hashes[0], "First hash should not be zero")
}

// Test SHA512 hash function
func Test_Sha512(t *testing.T) {
	data := []byte("test data")
	hashes := make([]uint64, 8)

	n := bloomhashes.Sha512(data, hashes)
	assert.Equal(t, 8, n, "Expected 8 hashes to be generated from SHA512")
	assert.NotZero(t, hashes[0], "First hash should not be zero")
	assert.NotZero(t, hashes[7], "Last hash should not be zero")
}

// Test that hash functions produce consistent results
func Test_HashFunctions_Consistency(t *testing.T) {
	testCases := []struct {
		name     string
		hashFunc bloomhashes.HashFunction
		hashSize int
	}{
		{"FNV1_64", bloomhashes.Fnv1_64, 1},
		{"FNV1_64a", bloomhashes.Fnv1_64a, 1},
		{"FNV1_128", bloomhashes.Fnv1_128, 2},
		{"FNV1_128a", bloomhashes.Fnv1_128a, 2},
		{"CRC64_ISO", bloomhashes.Crc64_ISO, 1},
		{"CRC64_ECMA", bloomhashes.Crc64_ECMA, 1},
		{"MD5", bloomhashes.MD5, bloomhashes.MD5_HASHES},
		{"SHA1", bloomhashes.Sha1, bloomhashes.SHA1_HASHES},
		{"SHA256", bloomhashes.Sha256, bloomhashes.SHA256_HASHES},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			data := []byte("consistency test")
			hashes1 := make([]uint64, tc.hashSize)
			hashes2 := make([]uint64, tc.hashSize)

			// Hash the same data twice
			n1 := tc.hashFunc(data, hashes1)
			n2 := tc.hashFunc(data, hashes2)

			assert.Equal(t, n1, n2, "Should return same number of hashes")
			assert.Equal(t, hashes1, hashes2, "Should produce identical hashes for identical input")
		})
	}
}

// Test that different data produces different hashes
func Test_HashFunctions_Uniqueness(t *testing.T) {
	testCases := []struct {
		name     string
		hashFunc bloomhashes.HashFunction
		hashSize int
	}{
		{"FNV1_64", bloomhashes.Fnv1_64, 1},
		{"FNV1_64a", bloomhashes.Fnv1_64a, 1},
		{"MD5", bloomhashes.MD5, bloomhashes.MD5_HASHES},
		{"SHA256", bloomhashes.Sha256, bloomhashes.SHA256_HASHES},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			data1 := []byte("test data 1")
			data2 := []byte("test data 2")
			hashes1 := make([]uint64, tc.hashSize)
			hashes2 := make([]uint64, tc.hashSize)

			tc.hashFunc(data1, hashes1)
			tc.hashFunc(data2, hashes2)

			assert.NotEqual(t, hashes1, hashes2, "Different input should produce different hashes")
		})
	}
}

// Test WrapHasher64
func Test_WrapHasher64(t *testing.T) {
	hashFunc := bloomhashes.WrapHasher64(fnv.New64)
	data := []byte("test data")
	hashes := make([]uint64, 1)

	n := hashFunc(data, hashes)
	assert.Equal(t, 1, n, "Expected 1 hash to be generated")
	assert.NotZero(t, hashes[0], "Hash should not be zero")

	// Verify it matches direct hash
	hasher := fnv.New64()
	hasher.Write(data)
	expected := hasher.Sum64()
	assert.Equal(t, expected, hashes[0], "Wrapped hash should match direct hash")
}

// Test WrapHasher with SHA256
func Test_WrapHasher_SHA256(t *testing.T) {
	hashFunc := bloomhashes.WrapHasher(sha256.New)
	data := []byte("test data")
	hashes := make([]uint64, 4)

	n := hashFunc(data, hashes)
	assert.Equal(t, 4, n, "Expected 4 hashes from SHA256")
	assert.NotZero(t, hashes[0], "First hash should not be zero")
}

// Test WrapHasher with SHA1
func Test_WrapHasher_SHA1(t *testing.T) {
	hashFunc := bloomhashes.WrapHasher(sha1.New)
	data := []byte("test data")
	hashes := make([]uint64, 3)

	n := hashFunc(data, hashes)
	assert.GreaterOrEqual(t, n, 2, "Expected at least 2 hashes from SHA1")
	assert.NotZero(t, hashes[0], "First hash should not be zero")
}

// Test WrapFunction
func Test_WrapFunction(t *testing.T) {
	// Create a simple hash function that just repeats the first 8 bytes
	simpleHash := func(data []byte) []byte {
		result := make([]byte, 16)
		copy(result, data)
		copy(result[8:], data)
		return result
	}

	hashFunc := bloomhashes.WrapFunction(simpleHash)
	data := []byte("test data for wrap function")
	hashes := make([]uint64, 2)

	n := hashFunc(data, hashes)
	assert.Equal(t, 2, n, "Expected 2 hashes")
}

// Test WrapFunction with small output
func Test_WrapFunction_SmallOutput(t *testing.T) {
	// Create a hash function that
	//  returns less than 8 bytes
	smallHash := func(data []byte) []byte {
		return []byte{1, 2, 3, 4}
	}

	hashFunc := bloomhashes.WrapFunction(smallHash)
	data := []byte("test")
	hashes := make([]uint64, 1)

	n := hashFunc(data, hashes)
	assert.Equal(t, 0, n, "Expected 0 hashes when output is less than 8 bytes")
}

// Fuzz test for hash functions
func Fuzz_HashFunctions(f *testing.F) {
	// Add seed corpus
	f.Add([]byte("test"))
	f.Add([]byte("hello world"))
	f.Add([]byte(""))

	f.Fuzz(func(t *testing.T, data []byte) {
		hashFuncs := []struct {
			fn   bloomhashes.HashFunction
			size int
		}{
			{bloomhashes.Fnv1_64, 1},
			{bloomhashes.Fnv1_64a, 1},
			{bloomhashes.Crc64_ISO, 1},
			{bloomhashes.MD5, bloomhashes.MD5_HASHES},
			{bloomhashes.Sha256, bloomhashes.SHA256_HASHES},
		}

		for _, hf := range hashFuncs {
			hashes := make([]uint64, hf.size)
			n := hf.fn(data, hashes)
			require.GreaterOrEqual(t, n, 0, "Should return non-negative number of hashes")
			require.LessOrEqual(t, n, hf.size, "Should not exceed buffer size")

			// Test consistency - same input should produce same output
			hashes2 := make([]uint64, hf.size)
			n2 := hf.fn(data, hashes2)
			require.Equal(t, n, n2, "Should return same number of hashes")
			require.Equal(t, hashes, hashes2, "Should produce same hashes for same input")
		}
	})
}

// Example of using FNV hash function
func Example_fnv1_64() {
	data := []byte("example data")
	hashes := make([]uint64, 1)

	n := bloomhashes.Fnv1_64(data, hashes)
	fmt.Printf("Generated %d hash(es)\n", n)
	fmt.Printf("Hash value is non-zero: %v\n", hashes[0] != 0)

	// Output:
	// Generated 1 hash(es)
	// Hash value is non-zero: true
}

// Example of using MD5 hash function
func Example_md5() {
	data := []byte("example data")
	hashes := make([]uint64, bloomhashes.MD5_HASHES)

	n := bloomhashes.MD5(data, hashes)
	fmt.Printf("Generated %d hash(es) from MD5\n", n)
	fmt.Printf("First hash is non-zero: %v\n", hashes[0] != 0)
	fmt.Printf("Second hash is non-zero: %v\n", hashes[1] != 0)

	// Output:
	// Generated 2 hash(es) from MD5
	// First hash is non-zero: true
	// Second hash is non-zero: true
}

// Example of using wrapped hash function
func Example_wrapHasher64() {
	// Wrap the FNV-1a hash function
	hashFunc := bloomhashes.WrapHasher64(fnv.New64a)

	data := []byte("wrapped hash example")
	hashes := make([]uint64, 1)

	n := hashFunc(data, hashes)
	fmt.Printf("Generated %d hash(es)\n", n)
	fmt.Printf("Hash value is non-zero: %v\n", hashes[0] != 0)

	// Output:
	// Generated 1 hash(es)
	// Hash value is non-zero: true
}
