package rex

func NewMoveFromSavings(owner flon.AccountName, rex flon.Asset) *flon.Action {
	return &flon.Action{
		Account: REXAN,
		Name:    ActN("mvfrsavings"),
		Authorization: []flon.PermissionLevel{
			{Actor: owner, Permission: flon.PermissionName("active")},
		},
		ActionData: flon.NewActionData(MoveFromSavings{
			Owner: owner,
			REX:   rex,
		}),
	}
}

type MoveFromSavings struct {
	Owner flon.AccountName
	REX   flon.Asset
}
