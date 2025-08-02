package msig

// NewUnapprove returns a `unapprove` action that lives on the
// `flon.msig` contract.
func NewUnapprove(proposer flon.AccountName, proposalName flon.Name, level flon.PermissionLevel) *flon.Action {
	return &flon.Action{
		Account:       flon.AccountName("flon.msig"),
		Name:          flon.ActionName("unapprove"),
		Authorization: []flon.PermissionLevel{level},
		ActionData:    flon.NewActionData(Unapprove{proposer, proposalName, level}),
	}
}

type Unapprove struct {
	Proposer     flon.AccountName     `json:"proposer"`
	ProposalName flon.Name            `json:"proposal_name"`
	Level        flon.PermissionLevel `json:"level"`
}
