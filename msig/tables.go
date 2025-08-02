package msig

type ProposalRow struct {
	ProposalName       flon.Name              `json:"proposal_name"`
	RequestedApprovals []flon.PermissionLevel `json:"requested_approvals"`
	ProvidedApprovals  []flon.PermissionLevel `json:"provided_approvals"`
	PackedTransaction  flon.HexBytes          `json:"packed_transaction"`
}
