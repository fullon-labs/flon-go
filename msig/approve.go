package msig

// NewApprove returns a `approve` action that lives on the
// `flon.msig` contract.
func NewApprove(proposer flon.AccountName, proposalName flon.Name, level flon.PermissionLevel) *flon.Action {
	return &flon.Action{
		Account:       flon.AccountName("flon.msig"),
		Name:          flon.ActionName("approve"),
		Authorization: []flon.PermissionLevel{level},
		ActionData:    flon.NewActionData(Approve{proposer, proposalName, level}),
	}
}

type Approve struct {
	Proposer     flon.AccountName     `json:"proposer"`
	ProposalName flon.Name            `json:"proposal_name"`
	Level        flon.PermissionLevel `json:"level"`
}
