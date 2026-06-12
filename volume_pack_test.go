package chalk

import (
	"bytes"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestDataPackSealParsesAndAddressesMembers(t *testing.T) {
	files := [][]byte{
		[]byte("alpha payload"),
		[]byte("bravo payload longer"),
		[]byte("charlie"),
	}
	builder := newDataPackBuilder()
	hashes := make([][32]byte, 0, len(files))
	for _, file := range files {
		hash := blake3Sum(file)
		hashes = append(hashes, hash)
		builder.append(hash, file)
	}

	sealed, err := builder.seal()
	require.NoError(t, err)
	require.Equal(t, []byte("CDP1"), sealed.bytes[:4])
	require.Equal(t, blake3Sum(sealed.bytes), sealed.chunkID)

	entries, bodyStart, err := parseDataPackHeader(sealed.bytes)
	require.NoError(t, err)
	require.Equal(t, sealed.entries, entries)

	for i, hash := range hashes {
		entry, ok := sealed.entryFor(hash)
		require.True(t, ok)
		require.GreaterOrEqual(t, int(entry.offset), bodyStart)
		got := sealed.bytes[entry.offset : entry.offset+uint64(entry.length)]
		require.Equal(t, files[i], got)
	}
}

func TestDataPackDedupesIdenticalContent(t *testing.T) {
	data := []byte("same bytes")
	hash := blake3Sum(data)
	builder := newDataPackBuilder()
	builder.append(hash, data)
	builder.append(hash, data)

	sealed, err := builder.seal()
	require.NoError(t, err)
	require.Len(t, sealed.entries, 1)
}

func TestDataPackSealIsDeterministic(t *testing.T) {
	build := func() []byte {
		builder := newDataPackBuilder()
		for _, data := range [][]byte{[]byte("alpha"), []byte("charlie")} {
			builder.append(blake3Sum(data), data)
		}
		sealed, err := builder.seal()
		require.NoError(t, err)
		return sealed.bytes
	}

	require.True(t, bytes.Equal(build(), build()))
}

func TestDataPackFitsMatchesSealedLength(t *testing.T) {
	builder := newDataPackBuilder()
	for i := 0; i < 30; i++ {
		data := bytes.Repeat([]byte{byte('a' + i%26)}, 24+i)
		require.True(t, builder.fits(uint64(len(data)), 10000))
		builder.append(blake3Sum(data), data)
	}

	sealed, err := builder.seal()
	require.NoError(t, err)
	require.Equal(t, uint64(len(sealed.bytes)), builder.objectLen())
}
