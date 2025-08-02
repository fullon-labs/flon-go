package rex

func NewDefundNetLoan(from flon.AccountName, loanNumber uint64, amount flon.Asset) *flon.Action {
	return &flon.Action{
		Account: REXAN,
		Name:    ActN("defnetloan"),
		Authorization: []flon.PermissionLevel{
			{Actor: from, Permission: flon.PermissionName("active")},
		},
		ActionData: flon.NewActionData(DefundNetLoan{
			From:       from,
			LoanNumber: loanNumber,
			Amount:     amount,
		}),
	}
}

type DefundNetLoan struct {
	From       flon.AccountName
	LoanNumber uint64
	Amount     flon.Asset
}
