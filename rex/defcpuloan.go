package rex

func NewDefundCPULoan(from flon.AccountName, loanNumber uint64, amount flon.Asset) *flon.Action {
	return &flon.Action{
		Account: REXAN,
		Name:    ActN("defcpuloan"),
		Authorization: []flon.PermissionLevel{
			{Actor: from, Permission: flon.PermissionName("active")},
		},
		ActionData: flon.NewActionData(DefundCPULoan{
			From:       from,
			LoanNumber: loanNumber,
			Amount:     amount,
		}),
	}
}

type DefundCPULoan struct {
	From       flon.AccountName
	LoanNumber uint64
	Amount     flon.Asset
}
