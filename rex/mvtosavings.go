package rex

func NewMoveToSavings(owner flon.AccountName, rex flon.Asset) *flon.Action {
	return &flon.Action{
		Account: REXAN,
		Name:    ActN("mvtosavings"),
		Authorization: []flon.PermissionLevel{
			{Actor: owner, Permission: flon.PermissionName("active")},
		},
		ActionData: flon.NewActionData(MoveToSavings{
			Owner: owner,
			REX:   rex,
		}),
	}
}

type MoveToSavings struct {
	Owner flon.AccountName
	REX   flon.Asset
}
