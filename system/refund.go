package system

import flon "github.com/fullon-labs/flon-go"

// NewRefund returns a `refund` action that lives on the
// `flon.system` contract.
func NewRefund(owner flon.AccountName) *flon.Action {
	return &flon.Action{
		Account: AN("flon"),
		Name:    ActN("refund"),
		Authorization: []flon.PermissionLevel{
			{Actor: owner, Permission: PN("active")},
		},
		ActionData: flon.NewActionData(Refund{
			Owner: owner,
		}),
	}
}

// Refund represents the `flon.system::refund` action
type Refund struct {
	Owner flon.AccountName `json:"owner"`
}
