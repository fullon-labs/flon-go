package rex

func NewWithdraw(owner flon.AccountName, amount flon.Asset) *flon.Action {
	return &flon.Action{
		Account: REXAN,
		Name:    ActN("withdraw"),
		Authorization: []flon.PermissionLevel{
			{Actor: owner, Permission: flon.PermissionName("active")},
		},
		ActionData: flon.NewActionData(Withdraw{
			Owner:  owner,
			Amount: amount,
		}),
	}
}

type Withdraw struct {
	Owner  flon.AccountName
	Amount flon.Asset
}
