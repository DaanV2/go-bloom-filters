package bloomfilters_test

import (
	"testing"

	bloomfilters "github.com/daanv2/go-bloom-filters"
	"github.com/daanv2/go-bloom-filters/pkg/bloomhashes"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func Test_Generic_String(t *testing.T) {
	bf, err := bloomfilters.NewBloomFilter(
		bloomfilters.WithSize(1024),
		bloomfilters.WithHashFunctions([]bloomhashes.HashFunction{bloomhashes.Fnv1_64}),
	)
	require.NoError(t, err)
	gbf := bloomfilters.NewGenericBloomFilter(bf, func(s string) []byte {
		return []byte(s)
	})

	// Add some data
	data1 := "hello"
	data2 := "world"
	data3 := "test"

	gbf.Add(data1)
	gbf.Add(data2)

	// Test that added data is found
	assert.True(t, gbf.Test(data1), "Expected 'hello' to be in the filter")
	assert.True(t, gbf.Test(data2), "Expected 'world' to be in the filter")

	// Test that non-added data is (probably) not found
	// Note: there's a small chance of false positive, but unlikely with this small dataset
	assert.False(t, gbf.Test(data3), "Expected 'test' to not be in the filter")
}

func Fuzz_Generic_String(f *testing.F) {
	f.Add("hello")
	f.Add("world")
	f.Add("test")

	f.Fuzz(func(t *testing.T, item string) {
		bf, err := bloomfilters.NewBloomFilter(
			bloomfilters.WithSize(1024),
			bloomfilters.WithHashFunctions([]bloomhashes.HashFunction{bloomhashes.Fnv1_64}),
		)
		require.NoError(t, err)
		gbf := bloomfilters.NewGenericBloomFilter(bf, func(s string) []byte {
			return []byte(s)
		})

		gbf.Add(item)

		// Test that added data is found
		assert.True(t, gbf.Test(item))
	})
}
