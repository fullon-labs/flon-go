package rex

func NewFundCPULoan(from flon.AccountName, loanNumber uint64, payment flon.Asset) *flon.Action {
	return &flon.Action{
		Account: REXAN,
		Name:    ActN("fundcpuloan"),
		Authorization: []flon.PermissionLevel{
			{Actor: from, Permission: flon.PermissionName("active")},
		},
		ActionData: flon.NewActionData(FundCPULoan{
			From:       from,
			LoanNumber: loanNumber,
			Payment:    payment,
		}),
	}
}

type FundCPULoan struct {
	From       flon.AccountName
	LoanNumber uint64
	Payment    flon.Asset
}
