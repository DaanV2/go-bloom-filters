package bloomhashes

func bytesToUint64(buf []byte) uint64 {
	if len(buf) < 8 {
		return 0
	}

	return uint64(buf[0]) | uint64(buf[1])<<8 | uint64(buf[2])<<16 | uint64(buf[3])<<24 |
		uint64(buf[4])<<32 | uint64(buf[5])<<40 | uint64(buf[6])<<48 | uint64(buf[7])<<56
}
