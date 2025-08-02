package system

import "github.com/fullon-labs/flon-go"

// NewDeleteAuth creates an action from the `flon.system` contract
// called `deleteauth`.
//
// You cannot delete the `owner` or `active` permissions.  Also, if a
// permission is still linked through a previous `updatelink` action,
// you will need to `unlinkauth` first.
func NewDeleteAuth(account flon.AccountName, permission flon.PermissionName) *flon.Action {
	a := &flon.Action{
		Account: AN("flon"),
		Name:    ActN("deleteauth"),
		Authorization: []flon.PermissionLevel{
			{Actor: account, Permission: flon.PermissionName("active")},
		},
		ActionData: flon.NewActionData(DeleteAuth{
			Account:    account,
			Permission: permission,
		}),
	}

	return a
}

// DeleteAuth represents the native `deleteauth` action, reachable
// through the `flon.system` contract.
type DeleteAuth struct {
	Account    flon.AccountName    `json:"account"`
	Permission flon.PermissionName `json:"permission"`
}
