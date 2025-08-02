package forum

// Status is an action to set a status update for a given account on the forum contract.
func NewStatus(account flon.AccountName, content string) *flon.Action {
	a := &flon.Action{
		Account: ForumAN,
		Name:    ActN("status"),
		Authorization: []flon.PermissionLevel{
			{Actor: account, Permission: flon.PermissionName("active")},
		},
		ActionData: flon.NewActionData(Status{
			Account: account,
			Content: content,
		}),
	}
	return a
}

// Status represents the `eosio.forum::status` action.
type Status struct {
	Account flon.AccountName `json:"account_name"`
	Content string           `json:"content"`
}
