package bloomfilters_test

import (
	"math/rand/v2"
	"slices"
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
	blobs := make([][][]byte, 8)
	for i := range 8 {
		blobs[i] = data[(arrays/8)*i : (arrays/8)*(i+1)]
	}

	b.Run("Add", func(b *testing.B) {
		bg, err := bloomfilters.NewConcurrentBloomFilter(
			bloomfilters.WithDefaultHashFunctions(),
			bloomfilters.WithSize(bf_size),
		)
		require.NoError(b, err)

		for b.Loop() {
			wg := sync.WaitGroup{}

			for _, todoblob := range blobs {
				wg.Go(func() {
					for _, blob := range todoblob {
						bg.Add(blob)
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

			for _, todoblob := range blobs {
				wg.Go(func() {
					for _, blob := range todoblob {
						bg.Add(blob)
					}
				})
			}

			wg.Wait()
			wg = sync.WaitGroup{}

			for _, todoblob := range blobs {
				wg.Go(func() {
					for _, blob := range todoblob {
						v := bg.Test(blob)
						if !v {
							b.Fatalf("expected to find %v", blob)
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

			for _, d := range blobs {
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
func Benchmark_Concurrent_Shuffled(b *testing.B) {
	const bf_size = 100
	const arrays = 100
	const array_length = 32

	data := testutil.MoreBytes(arrays*8, array_length)
	blobs := make([][][]byte, 8)
	for i := range 8 {
		blobs[i] = data[(arrays/8)*i : (arrays/8)*(i+1)]
	}

	b.Run("Add", func(b *testing.B) {
		bg, err := bloomfilters.NewConcurrentBloomFilter(
			bloomfilters.WithDefaultHashFunctions(),
			bloomfilters.WithSize(bf_size),
		)
		require.NoError(b, err)

		for b.Loop() {
			wg := sync.WaitGroup{}

			for _, todoblob := range blobs {
				shuffledblob := CloneShuffle(todoblob)

				wg.Go(func() {
					for _, blob := range shuffledblob {
						bg.Add(blob)
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

			for _, d := range blobs {
				shuffledblob := CloneShuffle(d)

				wg.Go(func() {
					for _, blob := range shuffledblob {
						bg.Add(blob)
					}
				})
			}

			wg.Wait()
			wg = sync.WaitGroup{}

			for _, todoblob := range blobs {
				shuffledblob := CloneShuffle(todoblob)

				wg.Go(func() {
					for _, blob := range shuffledblob {
						v := bg.Test(blob)
						if !v {
							b.Fatalf("expected to find %v", blob)
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

			for _, todoblob := range blobs {
				shuffledblob := CloneShuffle(todoblob)

				wg.Go(func() {
					for _, blob := range shuffledblob {
						v := bg.Test(blob)
						if v {
							b.Fatalf("found %v", blob)
						}
					}
				})
			}

			wg.Wait()
		}
	})
}

func CloneShuffle[T any](data []T) []T {
	d := slices.Clone(data)
	rand.Shuffle(len(d), func(i, j int) { // nolint:gosec // Not needed here
		d[i], d[j] = d[j], d[i]
	})

	return d
}
