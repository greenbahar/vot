package types

import "encoding/binary"

var _ binary.ByteOrder

const (
	// StoredVoteKeyPrefix is the prefix to retrieve all StoredVote
	StoredVoteKeyPrefix = "StoredVote/value/"
)

// StoredVoteKey returns the store key to retrieve a StoredVote from the index fields
func StoredVoteKey(
	index string,
) []byte {
	var key []byte

	indexBytes := []byte(index)
	key = append(key, indexBytes...)
	key = append(key, []byte("/")...)

	return key
}
