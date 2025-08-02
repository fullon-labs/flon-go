package msig

// NewPropose returns a `propose` action that lives on the
// `flon.msig` contract.
func NewPropose(proposer flon.AccountName, proposalName flon.Name, requested []flon.PermissionLevel, transaction *flon.Transaction) *flon.Action {
	return &flon.Action{
		Account: flon.AccountName("flon.msig"),
		Name:    flon.ActionName("propose"),
		Authorization: []flon.PermissionLevel{
			{Actor: proposer, Permission: flon.PermissionName("active")},
		},
		ActionData: flon.NewActionData(Propose{proposer, proposalName, requested, transaction}),
	}
}

type Propose struct {
	Proposer     flon.AccountName       `json:"proposer"`
	ProposalName flon.Name              `json:"proposal_name"`
	Requested    []flon.PermissionLevel `json:"requested"`
	Transaction  *flon.Transaction      `json:"trx"`
}
