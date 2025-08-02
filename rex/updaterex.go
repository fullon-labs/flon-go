package rex

func NewUpdateREX(owner flon.AccountName) *flon.Action {
	return &flon.Action{
		Account: REXAN,
		Name:    ActN("updaterex"),
		Authorization: []flon.PermissionLevel{
			{Actor: owner, Permission: flon.PermissionName("active")},
		},
		ActionData: flon.NewActionData(UpdateREX{
			Owner: owner,
		}),
	}
}

type UpdateREX struct {
	Owner flon.AccountName
}
