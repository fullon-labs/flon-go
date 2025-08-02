package msig

// NewExec returns a `exec` action that lives on the
// `flon.msig` contract.
func NewExec(proposer flon.AccountName, proposalName flon.Name, executer flon.AccountName) *flon.Action {
	return &flon.Action{
		Account: flon.AccountName("flon.msig"),
		Name:    flon.ActionName("exec"),
		// TODO: double check in this package that the `Actor` is always the `proposer`..
		Authorization: []flon.PermissionLevel{
			{Actor: executer, Permission: flon.PermissionName("active")},
		},
		ActionData: flon.NewActionData(Exec{proposer, proposalName, executer}),
	}
}

type Exec struct {
	Proposer     flon.AccountName `json:"proposer"`
	ProposalName flon.Name        `json:"proposal_name"`
	Executer     flon.AccountName `json:"executer"`
}
