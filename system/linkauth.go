package system

import flon "github.com/fullon-labs/flon-go"

// NewLinkAuth creates an action from the `flon.system` contract
// called `linkauth`.
//
// `linkauth` allows you to attach certain permission to the given
// `code::actionName`. With this set on-chain, you can use the
// `requiredPermission` to sign transactions for `code::actionName`
// and not rely on your `active` (which might be more sensitive as it
// can sign anything) for the given operation.
func NewLinkAuth(account, code flon.AccountName, actionName flon.ActionName, requiredPermission flon.PermissionName) *flon.Action {
	a := &flon.Action{
		Account: AN("flon"),
		Name:    ActN("linkauth"),
		Authorization: []flon.PermissionLevel{
			{
				Actor:      account,
				Permission: flon.PermissionName("active"),
			},
		},
		ActionData: flon.NewActionData(LinkAuth{
			Account:     account,
			Code:        code,
			Type:        actionName,
			Requirement: requiredPermission,
		}),
	}

	return a
}

// LinkAuth represents the native `linkauth` action, through the
// system contract.
type LinkAuth struct {
	Account     flon.AccountName    `json:"account"`
	Code        flon.AccountName    `json:"code"`
	Type        flon.ActionName     `json:"type"`
	Requirement flon.PermissionName `json:"requirement"`
}
