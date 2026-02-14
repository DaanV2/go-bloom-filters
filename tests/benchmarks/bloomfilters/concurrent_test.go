package bloomfilters_test

import (
	"sync"
	"testing"

	bloomfilters "github.com/daanv2/go-bloom-filters"
	"github.com/daanv2/go-bloom-filters/tests/testutil"
	"github.com/stretchr/testify/require"
)

func Benchmark_Concurrent(b *testing.B) {
	const bf_size = 100
	const arrays = 100
	const array_length = 32

	data := testutil.MoreBytes(arrays*8, array_length)
	chunks := make([][][]byte, 8)
	for i := range 8 {
		chunks[i] = data[(arrays/8)*i : (arrays/8)*(i+1)]
	}

	b.Run("Add", func(b *testing.B) {
		bg, err := bloomfilters.NewConcurrentBloomFilter(
			bloomfilters.WithDefaultHashFunctions(),
			bloomfilters.WithSize(bf_size),
		)
		require.NoError(b, err)

		for b.Loop() {
			wg := sync.WaitGroup{}

			for _, d := range chunks {
				wg.Go(func() {
					for i := range d {
						bg.Add(d[i])
					}
				})
			}

			wg.Wait()
		}
	})

	b.Run("Add_Test", func(b *testing.B) {
		bg, err := bloomfilters.NewConcurrentBloomFilter(
			bloomfilters.WithDefaultHashFunctions(),
			bloomfilters.WithSize(bf_size),
		)
		require.NoError(b, err)

		for b.Loop() {
			wg := sync.WaitGroup{}

			for _, d := range chunks {
				wg.Go(func() {
					for i := range d {
						bg.Add(d[i])
					}
				})
			}

			wg.Wait()
			wg = sync.WaitGroup{}

			for _, d := range chunks {
				wg.Go(func() {
					for i := range d {
						v := bg.Test(d[i])
						if !v {
							b.Fatalf("expected to find %v", d[i])
						}
					}
				})
			}

			wg.Wait()
		}
	})

	b.Run("Test_Nothing", func(b *testing.B) {
		bg, err := bloomfilters.NewConcurrentBloomFilter(
			bloomfilters.WithDefaultHashFunctions(),
			bloomfilters.WithSize(bf_size),
		)
		require.NoError(b, err)

		for b.Loop() {
			wg := sync.WaitGroup{}

			for _, d := range chunks {
				wg.Go(func() {
					for i := range d {
						v := bg.Test(d[i])
						if v {
							b.Fatalf("found %v", d[i])
						}
					}
				})
			}

			wg.Wait()
		}
	})
}
