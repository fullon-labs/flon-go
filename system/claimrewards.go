package system

import "github.com/fullon-labs/flon-go"

// NewClaimRewards will buy at current market price a given number of
// bytes of RAM, and grant them to the `receiver` account.
func NewClaimRewards(owner flon.AccountName) *flon.Action {
	a := &flon.Action{
		Account: AN("flon"),
		Name:    ActN("claimrewards"),
		Authorization: []flon.PermissionLevel{
			{Actor: owner, Permission: flon.PermissionName("active")},
		},
		ActionData: flon.NewActionData(ClaimRewards{
			Owner: owner,
		}),
	}
	return a
}

// ClaimRewards represents the `flon.system::claimrewards` action.
type ClaimRewards struct {
	Owner flon.AccountName `json:"owner"`
}
