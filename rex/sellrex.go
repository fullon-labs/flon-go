package rex

func NewSellREX(from flon.AccountName, rex flon.Asset) *flon.Action {
	return &flon.Action{
		Account: REXAN,
		Name:    ActN("sellrex"),
		Authorization: []flon.PermissionLevel{
			{Actor: from, Permission: flon.PermissionName("active")},
		},
		ActionData: flon.NewActionData(SellREX{
			From: from,
			REX:  rex,
		}),
	}
}

type SellREX struct {
	From flon.AccountName
	REX  flon.Asset
}
