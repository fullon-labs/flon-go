package system

import flon "github.com/fullon-labs/flon-go"

// NewUnlinkAuth creates an action from the `flon.system` contract
// called `unlinkauth`.
//
// `unlinkauth` detaches a previously set permission from a
// `code::actionName`. See `linkauth`.
func NewUnlinkAuth(account, code flon.AccountName, actionName flon.ActionName) *flon.Action {
	a := &flon.Action{
		Account: AN("flon"),
		Name:    ActN("unlinkauth"),
		Authorization: []flon.PermissionLevel{
			{
				Actor:      account,
				Permission: flon.PermissionName("active"),
			},
		},
		ActionData: flon.NewActionData(UnlinkAuth{
			Account: account,
			Code:    code,
			Type:    actionName,
		}),
	}

	return a
}

// UnlinkAuth represents the native `unlinkauth` action, through the
// system contract.
type UnlinkAuth struct {
	Account flon.AccountName `json:"account"`
	Code    flon.AccountName `json:"code"`
	Type    flon.ActionName  `json:"type"`
}
