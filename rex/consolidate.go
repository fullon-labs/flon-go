package rex

func NewConsolidate(owner flon.AccountName) *flon.Action {
	return &flon.Action{
		Account: REXAN,
		Name:    ActN("consolidate"),
		Authorization: []flon.PermissionLevel{
			{Actor: owner, Permission: flon.PermissionName("active")},
		},
		ActionData: flon.NewActionData(Consolidate{
			Owner: owner,
		}),
	}
}

type Consolidate struct {
	Owner flon.AccountName
}
