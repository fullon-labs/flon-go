package rex

func NewFundNetLoan(from flon.AccountName, loanNumber uint64, payment flon.Asset) *flon.Action {
	return &flon.Action{
		Account: REXAN,
		Name:    ActN("fundnetloan"),
		Authorization: []flon.PermissionLevel{
			{Actor: from, Permission: flon.PermissionName("active")},
		},
		ActionData: flon.NewActionData(FundNetLoan{
			From:       from,
			LoanNumber: loanNumber,
			Payment:    payment,
		}),
	}
}

type FundNetLoan struct {
	From       flon.AccountName
	LoanNumber uint64
	Payment    flon.Asset
}
