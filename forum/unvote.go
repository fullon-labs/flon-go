package forum

// NewUnVote is an action representing the action to undoing a current vote
func NewUnVote(voter flon.AccountName, proposalName flon.Name) *flon.Action {
	a := &flon.Action{
		Account: ForumAN,
		Name:    ActN("unvote"),
		Authorization: []flon.PermissionLevel{
			{Actor: voter, Permission: flon.PermissionName("active")},
		},
		ActionData: flon.NewActionData(UnVote{
			Voter:        voter,
			ProposalName: proposalName,
		}),
	}
	return a
}

// UnVote represents the `eosio.forum::unvote` action.
type UnVote struct {
	Voter        flon.AccountName `json:"voter"`
	ProposalName flon.Name        `json:"proposal_name"`
}
