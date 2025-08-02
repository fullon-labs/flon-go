package forum

// NewExpire is an action to expire a proposal ahead of its natural death.
func NewExpire(proposer flon.AccountName, proposalName flon.Name) *flon.Action {
	a := &flon.Action{
		Account: ForumAN,
		Name:    ActN("expire"),
		Authorization: []flon.PermissionLevel{
			{Actor: proposer, Permission: flon.PermissionName("active")},
		},
		ActionData: flon.NewActionData(Expire{
			ProposalName: proposalName,
		}),
	}
	return a
}

// Expire represents the `eosio.forum::propose` action.
type Expire struct {
	ProposalName flon.Name `json:"proposal_name"`
}
