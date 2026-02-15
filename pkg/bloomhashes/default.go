package bloomhashes

// DefaultHashFunctions returns a slice of recommended hash functions for bloom filters.
// These hash functions are sorted by computational cost, from lowest to highest.
// This provides a good balance between performance and hash distribution.
func DefaultHashFunctions() []HashFunction {
	return []HashFunction{
		// Sorted on cost, from lowest to highest
		Fnv1_64a,
		Fnv1_64,
		Crc64_ISO,
		Crc64_ECMA,
		Fnv1_128,
		Fnv1_128a,
	}
}

// AllHashFunctions returns a slice of all available hash functions for bloom filters.
// These hash functions are sorted by computational cost, from lowest to highest.
// Use this when you need more hash functions or better distribution at the cost of performance.
func AllHashFunctions() []HashFunction {
	return []HashFunction{
		// Sorted on cost, from lowest to highest
		Fnv1_64a,
		Fnv1_64,
		Crc64_ISO,
		Crc64_ECMA,
		Fnv1_128,
		Fnv1_128a,
		Sha256,
		Sha224,
		Sha1,
		MD5,
		Sha512,
		Sha3_384,
	}
}
