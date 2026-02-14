package bloomhashes

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
