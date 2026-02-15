package bloomfilters_test

import (
	"encoding/json"
	"testing"

	bloomfilters "github.com/daanv2/go-bloom-filters"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func Test_Bits_MarshalBinary(t *testing.T) {
	bits := bloomfilters.NewBits(128)
	for i := uint64(0); i < bits.Size(); i += 2 {
		bits.Setbit(i) // Set every even bit
	}

	data, err := bits.MarshalBinary()
	require.NoError(t, err)

	var unmarshaled bloomfilters.Bits
	err = unmarshaled.UnmarshalBinary(data)
	require.NoError(t, err)

	require.Equal(t, bits.Size(), unmarshaled.Size(), "Sizes should match after unmarshaling")
	for i := range bits.Size() {
		expect := (i%2 == 0) // Only even bits should be set
		require.Equal(t, expect, unmarshaled.Getbit(i), "Expected bit at index %d to be %v", i, expect)
	}

	assert.True(t, bits.Equals(&unmarshaled))
}

func Test_Bits_MarshalText(t *testing.T) {
	bits := bloomfilters.NewBits(128)
	for i := uint64(0); i < bits.Size(); i += 2 {
		bits.Setbit(i) // Set every even bit
	}

	data, err := bits.MarshalText()
	require.NoError(t, err)

	var unmarshaled bloomfilters.Bits
	err = unmarshaled.UnmarshalText(data)
	require.NoError(t, err)

	require.Equal(t, bits.Size(), unmarshaled.Size(), "Sizes should match after unmarshaling")
	for i := range bits.Size() {
		expect := (i%2 == 0) // Only even bits should be set
		require.Equal(t, expect, unmarshaled.Getbit(i), "Expected bit at index %d to be %v", i, expect)
	}

	assert.True(t, bits.Equals(&unmarshaled))
}

func Test_Bits_MarshalJSON(t *testing.T) {
	bits := bloomfilters.NewBits(128)
	for i := uint64(0); i < bits.Size(); i += 2 {
		bits.Setbit(i) // Set every even bit
	}

	type Data struct {
		Bits *bloomfilters.Bits `json:"bits"`
	}

	carrier := Data{Bits: &bits}
	v, err := json.Marshal(carrier)
	require.NoError(t, err)

	var unmarshaledCarrier Data
	err = json.Unmarshal(v, &unmarshaledCarrier)
	require.NoError(t, err)
	unmarshaled := unmarshaledCarrier.Bits

	require.Equal(t, bits.Size(), unmarshaled.Size(), "Sizes should match after unmarshaling")
	assert.True(t, bits.Equals(unmarshaled))
}
