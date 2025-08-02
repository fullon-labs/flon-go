package msig

// NewCancel returns a `cancel` action that lives on the
// `flon.msig` contract.
func NewCancel(proposer flon.AccountName, proposalName flon.Name, canceler flon.AccountName) *flon.Action {
	return &flon.Action{
		Account: flon.AccountName("flon.msig"),
		Name:    flon.ActionName("cancel"),
		// TODO: double check in this package that the `Actor` is always the `proposer`..
		Authorization: []flon.PermissionLevel{
			{Actor: canceler, Permission: flon.PermissionName("active")},
		},
		ActionData: flon.NewActionData(Cancel{proposer, proposalName, canceler}),
	}
}

type Cancel struct {
	Proposer     flon.AccountName `json:"proposer"`
	ProposalName flon.Name        `json:"proposal_name"`
	Canceler     flon.AccountName `json:"canceler"`
}
