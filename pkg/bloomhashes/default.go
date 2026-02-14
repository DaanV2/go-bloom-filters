package bloomhashes

// DefaultHashFunctions returns a slice of recommended hash functions for bloom filters.
// These hash functions are sorted by computational cost, from lowest to highest.
// This provides a good balance between performance and hash distribution.
func DefaultHashFunctions() []HashFunction {
	return []HashFunction{
		// Sorted on cost, from lowest to highest
		Fnv1_64,
		Fnv1_64a,
		Crc64_ECMA,
		Crc64_ISO,
		Fnv1_128a,
		Fnv1_128,
	}
}

// AllHashFunctions returns a slice of all available hash functions for bloom filters.
// These hash functions are sorted by computational cost, from lowest to highest.
// Use this when you need more hash functions or better distribution at the cost of performance.
func AllHashFunctions() []HashFunction {
	return []HashFunction{
		// Sorted on cost, from lowest to highest
		Fnv1_64,
		Fnv1_64a,
		Crc64_ECMA,
		Crc64_ISO,
		Fnv1_128a,
		Fnv1_128,
		Sha256,
		Sha224,
		Sha1,
		MD5,
		Sha512,
		Sha3_384,
	}
}
