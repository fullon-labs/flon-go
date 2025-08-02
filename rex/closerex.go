package rex

func NewCloseREX(owner flon.AccountName) *flon.Action {
	return &flon.Action{
		Account: REXAN,
		Name:    ActN("closerex"),
		Authorization: []flon.PermissionLevel{
			{Actor: owner, Permission: flon.PermissionName("active")},
		},
		ActionData: flon.NewActionData(CloseREX{
			Ownwer: owner,
		}),
	}
}

type CloseREX struct {
	Ownwer flon.AccountName
}
