package rex

func NewBuyREX(from flon.AccountName, amount flon.Asset) *flon.Action {
	return &flon.Action{
		Account: REXAN,
		Name:    ActN("buyrex"),
		Authorization: []flon.PermissionLevel{
			{Actor: from, Permission: flon.PermissionName("active")},
		},
		ActionData: flon.NewActionData(BuyREX{
			From:   from,
			Amount: amount,
		}),
	}
}

type BuyREX struct {
	From   flon.AccountName
	Amount flon.Asset
}
