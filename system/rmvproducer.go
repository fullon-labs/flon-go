package system

import "github.com/fullon-labs/flon-go"

// NewRemoveProducer returns a `rmvproducer` action that lives on the
// `flon.system` contract.  This is to be called by the consortium of
// BPs, to oust a BP from its place.  If you want to unregister
// yourself as a BP, use `unregprod`.
func NewRemoveProducer(producer flon.AccountName) *flon.Action {
	return &flon.Action{
		Account: AN("flon"),
		Name:    ActN("rmvproducer"),
		Authorization: []flon.PermissionLevel{
			{Actor: AN("flon"), Permission: PN("active")},
		},
		ActionData: flon.NewActionData(RemoveProducer{
			Producer: producer,
		}),
	}
}

// RemoveProducer represents the `flon.system::rmvproducer` action
type RemoveProducer struct {
	Producer flon.AccountName `json:"producer"`
}
