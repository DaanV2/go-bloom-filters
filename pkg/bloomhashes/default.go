package bloomhashes

func DefaultHashFunctions() []HashFunction {
	return []HashFunction{
		Sha1,
		Sha224,
		Sha3_384,
		Sha256,
		Sha512,
	}
}

func AllHashFunctions() []HashFunction {
	return []HashFunction{
		Sha1,
		Sha224,
		Sha3_384,
		Sha256,
		Sha512,
		Crc64_ISO,
		Crc64_ECMA,
		Fnv1_64,
		Fnv1_64a,
		Fnv1_128,
		Fnv1_128a,
	}
}
