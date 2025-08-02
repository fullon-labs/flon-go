package forum

// NewPropose is an action to submit a proposal for vote.
func NewPropose(proposer flon.AccountName, proposalName flon.Name, title string, proposalJSON string, expiresAt flon.JSONTime) *flon.Action {
	a := &flon.Action{
		Account: ForumAN,
		Name:    ActN("propose"),
		Authorization: []flon.PermissionLevel{
			{Actor: proposer, Permission: flon.PermissionName("active")},
		},
		ActionData: flon.NewActionData(Propose{
			Proposer:     proposer,
			ProposalName: proposalName,
			Title:        title,
			ProposalJSON: proposalJSON,
			ExpiresAt:    expiresAt,
		}),
	}
	return a
}

// Propose represents the `eosio.forum::propose` action.
type Propose struct {
	Proposer     flon.AccountName `json:"proposer"`
	ProposalName flon.Name        `json:"proposal_name"`
	Title        string           `json:"title"`
	ProposalJSON string           `json:"proposal_json"`
	ExpiresAt    flon.JSONTime    `json:"expires_at"`
}
