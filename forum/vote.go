package forum

// NewVote is an action representing a simple vote to be broadcast
// through the chain network.
func NewVote(voter flon.AccountName, proposalName flon.Name, voteValue uint8, voteJSON string) *flon.Action {
	a := &flon.Action{
		Account: ForumAN,
		Name:    ActN("vote"),
		Authorization: []flon.PermissionLevel{
			{Actor: voter, Permission: flon.PermissionName("active")},
		},
		ActionData: flon.NewActionData(Vote{
			Voter:        voter,
			ProposalName: proposalName,
			Vote:         voteValue,
			VoteJSON:     voteJSON,
		}),
	}
	return a
}

// Vote represents the `eosio.forum::vote` action.
type Vote struct {
	Voter        flon.AccountName `json:"voter"`
	ProposalName flon.Name        `json:"proposal_name"`
	Vote         uint8            `json:"vote"`
	VoteJSON     string           `json:"vote_json"`
}
