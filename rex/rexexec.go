package rex

func NewREXExec(user flon.AccountName, max uint16) *flon.Action {
	return &flon.Action{
		Account: REXAN,
		Name:    ActN("rexexec"),
		Authorization: []flon.PermissionLevel{
			{Actor: user, Permission: flon.PermissionName("active")},
		},
		ActionData: flon.NewActionData(REXExec{
			User: user,
			Max:  max,
		}),
	}
}

type REXExec struct {
	User flon.AccountName
	Max  uint16
}
