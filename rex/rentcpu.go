package rex

func NewRentCPU(
	from flon.AccountName,
	receiver flon.AccountName,
	loanPayment flon.Asset,
	loanFund flon.Asset,
) *flon.Action {
	return &flon.Action{
		Account: REXAN,
		Name:    ActN("rentcpu"),
		Authorization: []flon.PermissionLevel{
			{Actor: from, Permission: flon.PermissionName("active")},
		},
		ActionData: flon.NewActionData(RentCPU{
			From:        from,
			Receiver:    receiver,
			LoanPayment: loanPayment,
			LoanFund:    loanFund,
		}),
	}
}

type RentCPU struct {
	From        flon.AccountName
	Receiver    flon.AccountName
	LoanPayment flon.Asset
	LoanFund    flon.Asset
}
