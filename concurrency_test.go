package bloomfilters_test

import (
	"sync"
	"testing"

	bloomfilters "github.com/daanv2/go-bloom-filters"
	"github.com/daanv2/go-bloom-filters/tests/testutil"
	"github.com/stretchr/testify/require"
)

// Test that we can add all the items concurrently, no panics or error and all bits sets
func Test_Concurrent_Bloomfilter_Add(t *testing.T) {
	const bf_size = 1000_000
	const arrays = 1000
	const array_length = 32

	data := testutil.MoreBytes(arrays*8, array_length)
	chunks := make([][][]byte, 8)
	for i := range 8 {
		chunks[i] = data[(arrays/8)*i : (arrays/8)*(i+1)]
	}

	bg, err := bloomfilters.NewConcurrentBloomFilter(
		bloomfilters.WithDefaultHashFunctions(),
		bloomfilters.WithSize(bf_size),
	)
	require.NoError(t, err)
	wg := sync.WaitGroup{}

	for _, d := range chunks {
		wg.Go(func() {
			for i := range d {
				bg.Add(d[i])
			}
		})
	}

	wg.Wait()
	// Validate
	for _, d := range chunks {
		for i := range d {
			require.True(t, bg.Test(d[i]), "Expected item to be in the filter")
		}
	}
}

func Test_Concurrent_Bloomfilter_Test(t *testing.T) {
	const bf_size = 1000_000
	const arrays = 1000
	const array_length = 32

	type DataSet struct {
		data []byte
		set  bool
	}

	data := testutil.MoreBytes(arrays*8, array_length)
	chunks := make([][]DataSet, 8)
	for i := range 8 {
		items := data[(arrays/8)*i : (arrays/8)*(i+1)]

		chunks[i] = make([]DataSet, len(items))
		for j := range items {
			chunks[i][j] = DataSet{
				data: items[j],
				set:  j%2 == 0, // Set every other item
			}
		}
	}

	bg, err := bloomfilters.NewConcurrentBloomFilter(
		bloomfilters.WithDefaultHashFunctions(),
		bloomfilters.WithSize(bf_size),
	)
	require.NoError(t, err)
	wg := sync.WaitGroup{}

	for _, d := range chunks {
		wg.Go(func() {
			for i := range d {
				if d[i].set {
					bg.Add(d[i].data)
				}
			}
		})
	}

	wg.Wait()

	// Each chunk is now tested 10 times concurrently, we should find all the set items and some of the unset items (false positives)
	for _, d := range chunks {
		wg.Go(func() {
			for range 10 {
				for i := range d {
					v := bg.Test(d[i].data)
					if d[i].set {
						require.True(t, v, "Expected item to be in the filter")
					} else if v {
						// False positive, we can't assert on this but we can log it
						t.Logf("False positive for item: %v", d[i].data)
					}
				}
			}
		})
	}

	wg.Wait()
}