package bloomfilters_test

import (
	"testing"

	bloomfilters "github.com/daanv2/go-bloom-filters"
	"github.com/daanv2/go-bloom-filters/pkg/bloomhashes"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// Test basic bloom filter operations
func TestBloomFilter_BasicOperations(t *testing.T) {
	filter, err := bloomfilters.NewBloomFilter(
		bloomfilters.WithSize(10000),
		bloomfilters.WithHashFunctions(bloomhashes.DefaultHashFunctions()),
	)
	require.NoError(t, err)

	// Add items
	filter.Add([]byte("hello"))
	filter.Add([]byte("world"))
	filter.Add([]byte("bloom"))

	// Test membership - should find all added items
	assert.True(t, filter.Test([]byte("hello")), "Should find 'hello'")
	assert.True(t, filter.Test([]byte("world")), "Should find 'world'")
	assert.True(t, filter.Test([]byte("bloom")), "Should find 'bloom'")

	// Test non-membership - should not find items not added
	// Note: This could theoretically fail due to false positives, but with
	// a reasonable filter size and hash functions, it's very unlikely
	assert.False(t, filter.Test([]byte("notadded")), "Should not find 'notadded'")
}

// Test different hash functions
func TestBloomFilter_DifferentHashFunctions(t *testing.T) {
	// Test with all hash functions
	filter1, err := bloomfilters.NewBloomFilter(
		bloomfilters.WithSize(10000),
		bloomfilters.WithHashFunctions(bloomhashes.AllHashFunctions()),
	)
	require.NoError(t, err)

	filter1.Add([]byte("test"))
	assert.True(t, filter1.Test([]byte("test")))

	// Test with specific hash functions
	filter2, err := bloomfilters.NewBloomFilter(
		bloomfilters.WithSize(10000),
		bloomfilters.WithHashFunctions([]bloomhashes.HashFunction{
			bloomhashes.Sha256,
			bloomhashes.Fnv1_64,
		}),
	)
	require.NoError(t, err)

	filter2.Add([]byte("data"))
	assert.True(t, filter2.Test([]byte("data")))
}

// Test appending hash functions
func TestBloomFilter_AppendHashFunctions(t *testing.T) {
	filter, err := bloomfilters.NewBloomFilter(
		bloomfilters.WithSize(10000),
		bloomfilters.WithHashFunctions(bloomhashes.DefaultHashFunctions()),
		bloomfilters.WithAppendHashFunctions([]bloomhashes.HashFunction{
			bloomhashes.Fnv1_64,
		}),
	)
	require.NoError(t, err)

	filter.Add([]byte("example"))
	assert.True(t, filter.Test([]byte("example")))
}

// Test URL deduplication use case
func TestBloomFilter_URLDeduplication(t *testing.T) {
	urlFilter, err := bloomfilters.NewBloomFilter(
		bloomfilters.WithSize(1000000),
		bloomfilters.WithHashFunctions(bloomhashes.DefaultHashFunctions()),
	)
	require.NoError(t, err)

	urls := []string{
		"https://example.com",
		"https://example.com/page1",
		"https://example.com", // duplicate
		"https://example.com/page2",
		"https://example.com/page1", // duplicate
	}

	uniqueUrls := 0
	for _, url := range urls {
		urlBytes := []byte(url)
		if !urlFilter.Test(urlBytes) {
			urlFilter.Add(urlBytes)
			uniqueUrls++
		}
	}

	assert.Equal(t, 3, uniqueUrls, "Should have found 3 unique URLs")
}

// Test error conditions
func TestBloomFilter_Errors(t *testing.T) {
	// Test with no size option (default zero bits in struct)
	_, err := bloomfilters.NewBloomFilter(
		bloomfilters.WithHashFunctions(bloomhashes.DefaultHashFunctions()),
	)
	require.Error(t, err, "Should error with no size option")

	// Test with no hash functions
	_, err = bloomfilters.NewBloomFilter(
		bloomfilters.WithSize(1000),
	)
	require.Error(t, err, "Should error with no hash functions")
}
