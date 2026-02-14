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
		result := hashFunc(data)
		assert.NotZero(t, result, "Hash function %d should return a non-zero hash", i)
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
		result := hashFunc(data)
		assert.NotZero(t, result, "Hash function %d should return a non-zero hash", i)
	}
}

// Test FNV1_64 hash function
func Test_Fnv1_64(t *testing.T) {
	data := []byte("test data")

	result := bloomhashes.Fnv1_64(data)
	assert.NotZero(t, result, "Hash should not be zero")

	// Verify hash matches expected FNV-1 output
	hasher := fnv.New64()
	_, _ = hasher.Write(data)
	expected := hasher.Sum64()
	assert.Equal(t, expected, result, "Hash should match FNV-1 64-bit output")
}

// Test FNV1_64a hash function
func Test_Fnv1_64a(t *testing.T) {
	data := []byte("test data")

	result := bloomhashes.Fnv1_64a(data)
	assert.NotZero(t, result, "Hash should not be zero")

	// Verify hash matches expected FNV-1a output
	hasher := fnv.New64a()
	_, _ = hasher.Write(data)
	expected := hasher.Sum64()
	assert.Equal(t, expected, result, "Hash should match FNV-1a 64-bit output")
}

// Test FNV1_128 hash function
func Test_Fnv1_128(t *testing.T) {
	data := []byte("test data")

	result := bloomhashes.Fnv1_128(data)
	assert.NotZero(t, result, "Hash should not be zero")
}

// Test FNV1_128a hash function
func Test_Fnv1_128a(t *testing.T) {
	data := []byte("test data")

	result := bloomhashes.Fnv1_128a(data)
	assert.NotZero(t, result, "Hash should not be zero")
}

// Test CRC64_ISO hash function
func Test_Crc64_ISO(t *testing.T) {
	data := []byte("test data")

	result := bloomhashes.Crc64_ISO(data)
	assert.NotZero(t, result, "Hash should not be zero")

	// Verify hash matches expected CRC64 output
	expected := crc64.Checksum(data, crc64.MakeTable(crc64.ISO))
	assert.Equal(t, expected, result, "Hash should match CRC64-ISO output")
}

// Test CRC64_ECMA hash function
func Test_Crc64_ECMA(t *testing.T) {
	data := []byte("test data")

	result := bloomhashes.Crc64_ECMA(data)
	assert.NotZero(t, result, "Hash should not be zero")

	// Verify hash matches expected CRC64 output
	expected := crc64.Checksum(data, crc64.MakeTable(crc64.ECMA))
	assert.Equal(t, expected, result, "Hash should match CRC64-ECMA output")
}

// Test MD5 hash function
func Test_MD5(t *testing.T) {
	data := []byte("test data")

	result := bloomhashes.MD5(data)
	assert.NotZero(t, result, "Hash should not be zero")
}

// Test SHA1 hash function
func Test_Sha1(t *testing.T) {
	data := []byte("test data")

	result := bloomhashes.Sha1(data)
	assert.NotZero(t, result, "Hash should not be zero")
}

// Test SHA224 hash function
func Test_Sha224(t *testing.T) {
	data := []byte("test data")

	result := bloomhashes.Sha224(data)
	assert.NotZero(t, result, "Hash should not be zero")

	// Ensure hash changes with different input
	result2 := bloomhashes.Sha224([]byte("different data"))
	assert.NotEqual(t, result, result2, "Different inputs should produce different hashes")
}

// Test SHA256 hash function
func Test_Sha256(t *testing.T) {
	data := []byte("test data")

	result := bloomhashes.Sha256(data)
	assert.NotZero(t, result, "Hash should not be zero")
}

// Test SHA3_384 hash function
func Test_Sha3_384(t *testing.T) {
	data := []byte("test data")

	result := bloomhashes.Sha3_384(data)
	assert.NotZero(t, result, "Hash should not be zero")
}

// Test SHA512 hash function
func Test_Sha512(t *testing.T) {
	data := []byte("test data")

	result := bloomhashes.Sha512(data)
	assert.NotZero(t, result, "Hash should not be zero")
}

// Test that hash functions produce consistent results
func Test_HashFunctions_Consistency(t *testing.T) {
	testCases := []struct {
		name     string
		hashFunc bloomhashes.HashFunction
	}{
		{"FNV1_64", bloomhashes.Fnv1_64},
		{"FNV1_64a", bloomhashes.Fnv1_64a},
		{"FNV1_128", bloomhashes.Fnv1_128},
		{"FNV1_128a", bloomhashes.Fnv1_128a},
		{"CRC64_ISO", bloomhashes.Crc64_ISO},
		{"CRC64_ECMA", bloomhashes.Crc64_ECMA},
		{"MD5", bloomhashes.MD5},
		{"SHA1", bloomhashes.Sha1},
		{"SHA256", bloomhashes.Sha256},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			data := []byte("consistency test")

			// Hash the same data twice
			result1 := tc.hashFunc(data)
			result2 := tc.hashFunc(data)

			assert.Equal(t, result1, result2, "Should produce identical hashes for identical input")
		})
	}
}

// Test that different data produces different hashes
func Test_HashFunctions_Uniqueness(t *testing.T) {
	testCases := []struct {
		name     string
		hashFunc bloomhashes.HashFunction
	}{
		{"FNV1_64", bloomhashes.Fnv1_64},
		{"FNV1_64a", bloomhashes.Fnv1_64a},
		{"MD5", bloomhashes.MD5},
		{"SHA256", bloomhashes.Sha256},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			data1 := []byte("test data 1")
			data2 := []byte("test data 2")

			result1 := tc.hashFunc(data1)
			result2 := tc.hashFunc(data2)

			assert.NotEqual(t, result1, result2, "Different input should produce different hashes")
		})
	}
}

// Test WrapHasher64
func Test_WrapHasher64(t *testing.T) {
	hashFunc := bloomhashes.WrapHasher64(fnv.New64)
	data := []byte("test data")

	result := hashFunc(data)
	assert.NotZero(t, result, "Hash should not be zero")

	// Verify it matches direct hash
	hasher := fnv.New64()
	hasher.Write(data)
	expected := hasher.Sum64()
	assert.Equal(t, expected, result, "Wrapped hash should match direct hash")
}

// Test WrapHasher with SHA256
func Test_WrapHasher_SHA256(t *testing.T) {
	hashFunc := bloomhashes.WrapHasher(sha256.New)
	data := []byte("test data")

	result := hashFunc(data)
	assert.NotZero(t, result, "Hash should not be zero")
}

// Test WrapHasher with SHA1
func Test_WrapHasher_SHA1(t *testing.T) {
	hashFunc := bloomhashes.WrapHasher(sha1.New)
	data := []byte("test data")

	result := hashFunc(data)
	assert.NotZero(t, result, "Hash should not be zero")
}

// Test WrapFunction
func Test_WrapFunction(t *testing.T) {
	// Create a simple hash function that returns 16 bytes
	simpleHash := func(data []byte) []byte {
		result := make([]byte, 16)
		copy(result, data)
		copy(result[8:], data)

		return result
	}

	hashFunc := bloomhashes.WrapFunction(simpleHash)
	data := []byte("test data for wrap function")

	result := hashFunc(data)
	assert.NotZero(t, result, "Hash should not be zero")
}

// Test WrapFunction with small output
func Test_WrapFunction_SmallOutput(t *testing.T) {
	// Create a hash function that returns less than 8 bytes
	smallHash := func(data []byte) []byte {
		return []byte{1, 2, 3, 4}
	}

	hashFunc := bloomhashes.WrapFunction(smallHash)
	data := []byte("test")

	// With less than 8 bytes, bytesToUint64 still produces a value
	result := hashFunc(data)
	_ = result
}

// Fuzz test for hash functions
func Fuzz_HashFunctions(f *testing.F) {
	// Add seed corpus
	f.Add([]byte("test"))
	f.Add([]byte("hello world"))
	f.Add([]byte(""))

	f.Fuzz(func(t *testing.T, data []byte) {
		hashFuncs := []bloomhashes.HashFunction{
			bloomhashes.Fnv1_64,
			bloomhashes.Fnv1_64a,
			bloomhashes.Crc64_ISO,
			bloomhashes.MD5,
			bloomhashes.Sha256,
		}

		for _, hf := range hashFuncs {
			// Should not panic
			result := hf(data)

			// Test consistency - same input should produce same output
			result2 := hf(data)
			require.Equal(t, result, result2, "Should produce same hash for same input")
		}
	})
}

// Example of using FNV hash function
func Example_fnv1_64() {
	data := []byte("example data")

	result := bloomhashes.Fnv1_64(data)
	fmt.Printf("Hash value is non-zero: %v\n", result != 0)

	// Output:
	// Hash value is non-zero: true
}

// Example of using MD5 hash function
func Example_md5() {
	data := []byte("example data")

	result := bloomhashes.MD5(data)
	fmt.Printf("Hash value is non-zero: %v\n", result != 0)

	// Output:
	// Hash value is non-zero: true
}

// Example of using wrapped hash function
func Example_wrapHasher64() {
	// Wrap the FNV-1a hash function
	hashFunc := bloomhashes.WrapHasher64(fnv.New64a)

	data := []byte("wrapped hash example")

	result := hashFunc(data)
	fmt.Printf("Hash value is non-zero: %v\n", result != 0)

	// Output:
	// Hash value is non-zero: true
}
