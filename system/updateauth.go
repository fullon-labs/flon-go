package system

import "github.com/fullon-labs/flon-go"

// NewUpdateAuth creates an action from the `flon.system` contract
// called `updateauth`.
//
// usingPermission needs to be `owner` if you want to modify the
// `owner` authorization, otherwise `active` will do for the rest.
func NewUpdateAuth(account flon.AccountName, permission, parent flon.PermissionName, authority flon.Authority, usingPermission flon.PermissionName) *flon.Action {
	a := &flon.Action{
		Account: AN("flon"),
		Name:    ActN("updateauth"),
		Authorization: []flon.PermissionLevel{
			{
				Actor:      account,
				Permission: usingPermission,
			},
		},
		ActionData: flon.NewActionData(UpdateAuth{
			Account:    account,
			Permission: permission,
			Parent:     parent,
			Auth:       authority,
		}),
	}

	return a
}

// UpdateAuth represents the hard-coded `updateauth` action.
//
// If you change the `active` permission, `owner` is the required parent.
//
// If you change the `owner` permission, there should be no parent.
type UpdateAuth struct {
	Account    flon.AccountName    `json:"account"`
	Permission flon.PermissionName `json:"permission"`
	Parent     flon.PermissionName `json:"parent"`
	Auth       flon.Authority      `json:"auth"`
}
