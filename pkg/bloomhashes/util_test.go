package bloomhashes_test

import (
	"encoding/binary"
	"testing"

	"github.com/daanv2/go-bloom-filters/pkg/bloomhashes"
	"github.com/stretchr/testify/require"
)

func Fuzz_ToUint64(f *testing.F) {
	f.Add([]byte{})
	f.Add([]byte{1, 2, 3, 4, 5, 6, 7, 8})
	f.Add([]byte{1, 2, 3, 4, 5, 6, 7, 8, 9})
	f.Add([]byte{1, 2, 3, 4, 5, 6, 7, 8, 9, 10})

	f.Fuzz(func(t *testing.T, b []byte) {
		result := bloomhashes.ToUint64(b)
		if len(b) < 8 {
			require.Empty(t, result)
		}
		if len(b) > 8 {
			require.NotEmpty(t, result)
			require.Len(t, result, len(b)/8)
		}

		for i := range result {
			var buf [8]byte
			binary.NativeEndian.PutUint64(buf[:], result[i])
			require.Equal(t, buf[:], b[i*8:(i+1)*8])
		}
	})
}
