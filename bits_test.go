package bloomfilters_test

import (
	"fmt"
	"testing"

	bloomfilters "github.com/daanv2/go-bloom-filters"
	"github.com/stretchr/testify/require"
)

func Fuzz_Bits_SetGet(f *testing.F) {
	size := uint64(1024) // Example size for the Bits structure
	f.Add(uint64(0))     // Add edge case: index 0
	f.Add(size - 1)      // Add edge case: last index
	f.Add(uint64(1090))

	f.Fuzz(func(t *testing.T, index uint64) {
		bits := bloomfilters.NewBits(size)
		bits.Setbit(index)

		// If index > size, it should not set any bit

		// Check that other bits are not set
		for i := range size {
			expect := (i == index) // Only the bit at 'index' should be set
			require.Equal(t, expect, bits.Getbit(i), "Expected bit at index %d to be %v", i, expect)
		}
	})
}


func ExampleBits() {
	bits := bloomfilters.NewBits(128) // Create a Bits structure with 128 bits
	bits.Setbit(5)                    // Set the bit at index 5
	bits.Setbit(10)                   // Set the bit at index 10

	// Check if specific bits are set
	fmt.Println(bits.Getbit(5))
	fmt.Println(bits.Getbit(10))
	fmt.Println(bits.Getbit(0))
	fmt.Println(bits.Getbit(127))

	// Output:
	// true
	// true
	// false
	// false
}
