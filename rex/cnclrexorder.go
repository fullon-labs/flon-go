package rex

func NewCancelREXOrder(owner flon.AccountName) *flon.Action {
	return &flon.Action{
		Account: REXAN,
		Name:    ActN("cnclrexorder"),
		Authorization: []flon.PermissionLevel{
			{Actor: owner, Permission: flon.PermissionName("active")},
		},
		ActionData: flon.NewActionData(CancelREXOrder{
			Owner: owner,
		}),
	}
}

type CancelREXOrder struct {
	Owner flon.AccountName
}
