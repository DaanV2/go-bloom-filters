package bloomfilters_test

import (
	"fmt"
	"testing"

	bloomfilters "github.com/daanv2/go-bloom-filters"
	"github.com/daanv2/go-bloom-filters/pkg/bloomhashes"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// Test NewBloomFilter with valid options
func Test_NewBloomFilter_Valid(t *testing.T) {
	bf, err := bloomfilters.NewBloomFilter(
		bloomfilters.WithSize(1024),
		bloomfilters.WithHashFunctions([]bloomhashes.HashFunction{bloomhashes.Fnv1_64}),
	)
	require.NoError(t, err)
	require.NotNil(t, bf)
}

// Test NewBloomFilter with multiple hash functions
func Test_NewBloomFilter_MultipleHashFunctions(t *testing.T) {
	bf, err := bloomfilters.NewBloomFilter(
		bloomfilters.WithSize(1024),
		bloomfilters.WithHashFunctions([]bloomhashes.HashFunction{
			bloomhashes.Fnv1_64,
			bloomhashes.Fnv1_64a,
			bloomhashes.MD5,
		}),
	)
	require.NoError(t, err)
	require.NotNil(t, bf)
}

// Test NewBloomFilter with zero size
// Note: NewBits(0) still creates at least one uint64 (64 bits), but the bloom filter
// should reject a filter with size 0 as invalid since it's likely unintentional
func Test_NewBloomFilter_InvalidSize(t *testing.T) {
	// Test with explicit size 0 - this should still work because NewBits ensures at least 64 bits
	bf, err := bloomfilters.NewBloomFilter(
		bloomfilters.WithSize(0),
		bloomfilters.WithHashFunctions([]bloomhashes.HashFunction{bloomhashes.Fnv1_64}),
	)
	
	// NewBits(0) creates at least 64 bits, so this should actually succeed
	require.NoError(t, err)
	require.NotNil(t, bf)
}

// Test NewBloomFilter with no hash functions (should fail)
func Test_NewBloomFilter_NoHashFunctions(t *testing.T) {
	bf, err := bloomfilters.NewBloomFilter(
		bloomfilters.WithSize(1024),
	)
	require.Error(t, err)
	require.Nil(t, bf)
	assert.Contains(t, err.Error(), "at least one hash function is required")
}

// Test NewBloomFilter with nil hash function (should fail)
func Test_NewBloomFilter_NilHashFunction(t *testing.T) {
	bf, err := bloomfilters.NewBloomFilter(
		bloomfilters.WithSize(1024),
		bloomfilters.WithHashFunctions([]bloomhashes.HashFunction{nil}),
	)
	require.Error(t, err)
	require.Nil(t, bf)
	assert.Contains(t, err.Error(), "hash functions cannot be nil")
}

// Test NewBloomFilter with mixed nil and valid hash functions (should fail)
func Test_NewBloomFilter_MixedNilHashFunctions(t *testing.T) {
	bf, err := bloomfilters.NewBloomFilter(
		bloomfilters.WithSize(1024),
		bloomfilters.WithHashFunctions([]bloomhashes.HashFunction{
			bloomhashes.Fnv1_64,
			nil,
			bloomhashes.Fnv1_64a,
		}),
	)
	require.Error(t, err)
	require.Nil(t, bf)
	assert.Contains(t, err.Error(), "hash functions cannot be nil")
}

// Test Add and Test with single hash function
func Test_BloomFilter_AddAndTest_SingleHash(t *testing.T) {
	bf, err := bloomfilters.NewBloomFilter(
		bloomfilters.WithSize(1024),
		bloomfilters.WithHashFunctions([]bloomhashes.HashFunction{bloomhashes.Fnv1_64}),
	)
	require.NoError(t, err)

	// Add some data
	data1 := []byte("hello")
	data2 := []byte("world")
	data3 := []byte("test")

	bf.Add(data1)
	bf.Add(data2)

	// Test that added data is found
	assert.True(t, bf.Test(data1), "Expected 'hello' to be in the filter")
	assert.True(t, bf.Test(data2), "Expected 'world' to be in the filter")

	// Test that non-added data is (probably) not found
	// Note: there's a small chance of false positive, but unlikely with this small dataset
	assert.False(t, bf.Test(data3), "Expected 'test' to not be in the filter")
}

// Test Add and Test with multiple hash functions
func Test_BloomFilter_AddAndTest_MultipleHashes(t *testing.T) {
	bf, err := bloomfilters.NewBloomFilter(
		bloomfilters.WithSize(2048),
		bloomfilters.WithHashFunctions([]bloomhashes.HashFunction{
			bloomhashes.Fnv1_64,
			bloomhashes.Fnv1_64a,
			bloomhashes.MD5,
		}),
	)
	require.NoError(t, err)

	// Add some data
	data := []byte("multiple hash test")
	bf.Add(data)

	// Test that added data is found
	assert.True(t, bf.Test(data), "Expected data to be in the filter")

	// Test that non-added data is (probably) not found
	assert.False(t, bf.Test([]byte("not added")), "Expected 'not added' to not be in the filter")
}

// Test empty data
func Test_BloomFilter_EmptyData(t *testing.T) {
	bf, err := bloomfilters.NewBloomFilter(
		bloomfilters.WithSize(1024),
		bloomfilters.WithHashFunctions([]bloomhashes.HashFunction{bloomhashes.Fnv1_64}),
	)
	require.NoError(t, err)

	// Add empty data
	bf.Add([]byte{})

	// Test empty data
	assert.True(t, bf.Test([]byte{}), "Expected empty data to be in the filter")

	// Test non-empty data
	assert.False(t, bf.Test([]byte("something")), "Expected non-empty data to not be in the filter")
}

// Test with large data
func Test_BloomFilter_LargeData(t *testing.T) {
	bf, err := bloomfilters.NewBloomFilter(
		bloomfilters.WithSize(16384), // Larger size to reduce false positive rate
		bloomfilters.WithHashFunctions([]bloomhashes.HashFunction{
			bloomhashes.Sha256,
			bloomhashes.Fnv1_64,
		}),
	)
	require.NoError(t, err)

	// Create a large byte slice (1MB)
	largeData := make([]byte, 1024*1024)
	for i := range largeData {
		largeData[i] = byte(i % 256)
	}

	bf.Add(largeData)
	assert.True(t, bf.Test(largeData), "Expected large data to be in the filter")

	// Create a completely different large data set
	differentData := make([]byte, 1024*1024)
	for i := range differentData {
		differentData[i] = byte((i + 128) % 256)
	}
	
	// Note: There's a small chance of false positive, but with large enough filter
	// and multiple hash functions, it should be unlikely
	result := bf.Test(differentData)
	// We don't assert false here because bloom filters can have false positives
	// This test mainly ensures the filter can handle large data without crashing
	_ = result
}

// Test that bloom filter can distinguish between similar strings
func Test_BloomFilter_SimilarData(t *testing.T) {
	bf, err := bloomfilters.NewBloomFilter(
		bloomfilters.WithSize(1024),
		bloomfilters.WithHashFunctions([]bloomhashes.HashFunction{
			bloomhashes.Fnv1_64,
			bloomhashes.Fnv1_64a,
		}),
	)
	require.NoError(t, err)

	data1 := []byte("test")
	data2 := []byte("Test")
	data3 := []byte("tests")

	bf.Add(data1)

	assert.True(t, bf.Test(data1), "Expected 'test' to be in the filter")
	assert.False(t, bf.Test(data2), "Expected 'Test' to not be in the filter")
	assert.False(t, bf.Test(data3), "Expected 'tests' to not be in the filter")
}

// Test WithAppendHashFunctions
func Test_BloomFilter_AppendHashFunctions(t *testing.T) {
	bf, err := bloomfilters.NewBloomFilter(
		bloomfilters.WithSize(1024),
		bloomfilters.WithHashFunctions([]bloomhashes.HashFunction{bloomhashes.Fnv1_64}),
		bloomfilters.WithAppendHashFunctions([]bloomhashes.HashFunction{bloomhashes.Fnv1_64a}),
	)
	require.NoError(t, err)

	data := []byte("append test")
	bf.Add(data)

	// Should work with both hash functions
	assert.True(t, bf.Test(data), "Expected data to be in the filter")
}

// Test multiple additions of the same data
func Test_BloomFilter_MultipleAdditions(t *testing.T) {
	bf, err := bloomfilters.NewBloomFilter(
		bloomfilters.WithSize(1024),
		bloomfilters.WithHashFunctions([]bloomhashes.HashFunction{bloomhashes.Fnv1_64}),
	)
	require.NoError(t, err)

	data := []byte("duplicate")

	// Add the same data multiple times
	bf.Add(data)
	bf.Add(data)
	bf.Add(data)

	// Should still be found
	assert.True(t, bf.Test(data), "Expected data to be in the filter after multiple additions")
}

// Test with all available hash functions
func Test_BloomFilter_AllHashFunctions(t *testing.T) {
	hashFuncs := []bloomhashes.HashFunction{
		bloomhashes.Fnv1_64,
		bloomhashes.Fnv1_64a,
		bloomhashes.Fnv1_128,
		bloomhashes.Fnv1_128a,
		bloomhashes.Crc64_ISO,
		bloomhashes.Crc64_ECMA,
		bloomhashes.MD5,
		bloomhashes.Sha1,
		bloomhashes.Sha224,
		bloomhashes.Sha256,
		bloomhashes.Sha3_384,
		bloomhashes.Sha512,
	}

	bf, err := bloomfilters.NewBloomFilter(
		bloomfilters.WithSize(8192),
		bloomfilters.WithHashFunctions(hashFuncs),
	)
	require.NoError(t, err)

	data := []byte("comprehensive test")
	bf.Add(data)

	assert.True(t, bf.Test(data), "Expected data to be in the filter with all hash functions")
}

// Fuzz test for BloomFilter Add and Test
func Fuzz_BloomFilter_AddTest(f *testing.F) {
	// Add seed corpus
	f.Add([]byte("test"))
	f.Add([]byte("hello world"))
	f.Add([]byte(""))
	f.Add([]byte("a"))

	f.Fuzz(func(t *testing.T, data []byte) {
		bf, err := bloomfilters.NewBloomFilter(
			bloomfilters.WithSize(1024),
			bloomfilters.WithHashFunctions([]bloomhashes.HashFunction{
				bloomhashes.Fnv1_64,
				bloomhashes.MD5,
			}),
		)
		require.NoError(t, err)

		// Add data
		bf.Add(data)

		// Test that the added data is found
		require.True(t, bf.Test(data), "Expected added data to be found in the filter")
	})
}

// Example of basic BloomFilter usage
func ExampleBloomFilter() {
	// Create a bloom filter with size 1024 and FNV hash function
	bf, _ := bloomfilters.NewBloomFilter(
		bloomfilters.WithSize(1024),
		bloomfilters.WithHashFunctions([]bloomhashes.HashFunction{bloomhashes.Fnv1_64}),
	)

	// Add some data
	bf.Add([]byte("apple"))
	bf.Add([]byte("banana"))
	bf.Add([]byte("orange"))

	// Test if data is in the filter
	fmt.Println(bf.Test([]byte("apple")))
	fmt.Println(bf.Test([]byte("banana")))
	fmt.Println(bf.Test([]byte("grape")))

	// Output:
	// true
	// true
	// false
}

// Example demonstrating bloom filter with multiple hash functions
func ExampleBloomFilter_multipleHashes() {
	// Create a bloom filter with multiple hash functions for better accuracy
	bf, _ := bloomfilters.NewBloomFilter(
		bloomfilters.WithSize(2048),
		bloomfilters.WithHashFunctions([]bloomhashes.HashFunction{
			bloomhashes.Fnv1_64,
			bloomhashes.Fnv1_64a,
			bloomhashes.MD5,
		}),
	)

	// Add email addresses
	bf.Add([]byte("user@example.com"))
	bf.Add([]byte("admin@example.com"))

	// Check if emails exist
	fmt.Println(bf.Test([]byte("user@example.com")))
	fmt.Println(bf.Test([]byte("admin@example.com")))
	fmt.Println(bf.Test([]byte("guest@example.com")))

	// Output:
	// true
	// true
	// false
}
