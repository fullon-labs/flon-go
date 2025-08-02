package rex

func NewDeposit(owner flon.AccountName, amount flon.Asset) *flon.Action {
	return &flon.Action{
		Account: REXAN,
		Name:    ActN("deposit"),
		Authorization: []flon.PermissionLevel{
			{Actor: owner, Permission: flon.PermissionName("active")},
		},
		ActionData: flon.NewActionData(Deposit{
			Owner:  owner,
			Amount: amount,
		}),
	}
}

type Deposit struct {
	Owner  flon.AccountName
	Amount flon.Asset
}
