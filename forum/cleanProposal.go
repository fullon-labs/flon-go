package forum

// CleanProposal is an action to flush proposal and allow RAM used by it.
func NewCleanProposal(cleaner flon.AccountName, proposalName flon.Name, maxCount uint64) *flon.Action {
	a := &flon.Action{
		Account: ForumAN,
		Name:    ActN("clnproposal"),
		Authorization: []flon.PermissionLevel{
			{Actor: cleaner, Permission: flon.PermissionName("active")},
		},
		ActionData: flon.NewActionData(CleanProposal{
			ProposalName: proposalName,
			MaxCount:     maxCount,
		}),
	}
	return a
}

// CleanProposal represents the `eosio.forum::clnproposal` action.
type CleanProposal struct {
	ProposalName flon.Name `json:"proposal_name"`
	MaxCount     uint64    `json:"max_count"`
}
