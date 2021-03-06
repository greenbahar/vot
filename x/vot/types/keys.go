package types

const (
	// ModuleName defines the module name
	ModuleName = "vot"

	// StoreKey defines the primary module store key
	StoreKey = ModuleName

	// RouterKey is the message route for slashing
	RouterKey = ModuleName

	// QuerierRoute defines the module's query routing key
	QuerierRoute = ModuleName

	// MemStoreKey defines the in-memory store key
	MemStoreKey = "mem_vot"
)

func KeyPrefix(p string) []byte {
	return []byte(p)
}

const (
	NextVoteKey = "NextVote-value-"

	StoredVoteEventKey     = "NewVoteCreated" // Indicates what key to listen to
	StoredVoteEventCreator = "Creator"
	StoredVoteEventIndex   = "Index" // What vote is relevant

	SelectVotingOptionEventKey     = "VotingOptionSelected"
	SelectVotingOptionCreator      = "Creator"
	SelectVotingOptionEventIdValue = "IdValue"
	SelectVotingOptionEventOption  = "Option"
	SelectVotingOptionEventResult  = "Result"
)
