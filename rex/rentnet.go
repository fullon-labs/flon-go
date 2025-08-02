package rex

func NewRentNet(
	from flon.AccountName,
	receiver flon.AccountName,
	loanPayment flon.Asset,
	loanFund flon.Asset,
) *flon.Action {
	return &flon.Action{
		Account: REXAN,
		Name:    ActN("rentnet"),
		Authorization: []flon.PermissionLevel{
			{Actor: from, Permission: flon.PermissionName("active")},
		},
		ActionData: flon.NewActionData(RentNet{
			From:        from,
			Receiver:    receiver,
			LoanPayment: loanPayment,
			LoanFund:    loanFund,
		}),
	}
}

type RentNet struct {
	From        flon.AccountName
	Receiver    flon.AccountName
	LoanPayment flon.Asset
	LoanFund    flon.Asset
}
