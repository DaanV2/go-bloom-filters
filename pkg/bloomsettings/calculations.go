package bloomsettings

import "math"

// OptimalHashFunctions calculates the optimal number of hash functions (k) for a Bloom filter
// - m is the numbers of bits in the array
// - n is the amount of elements expected to be stored in the filter
func OptimalHashFunctions(m, n uint64) uint64 {
	//https://en.wikipedia.org/wiki/Bloom_filter#Optimal_number_of_hash_functions
	return uint64(math.Ceil(float64(m)/float64(n)) * math.Ln2)
}

// FalsePositiveRate calculates the false positive rate of a Bloom filter
// - m is the numbers of bits in the array
// - n is the amount of elements expected to be stored in the filter
// - k is the number of hash functions used
func FalsePositiveRate(m, n, k uint64) float64 {
	//https://en.wikipedia.org/wiki/Bloom_filter
	return math.Pow(1-math.Exp((-float64(k)*float64(n))/float64(m)), float64(k))
}
