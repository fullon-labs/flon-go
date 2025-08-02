package system

import flon "github.com/fullon-labs/flon-go"

// NewUnregProducer returns a `unregprod` action that lives on the
// `flon.system` contract.
func NewUnregProducer(producer flon.AccountName) *flon.Action {
	return &flon.Action{
		Account: AN("flon"),
		Name:    ActN("unregprod"),
		Authorization: []flon.PermissionLevel{
			{Actor: producer, Permission: PN("active")},
		},
		ActionData: flon.NewActionData(UnregProducer{
			Producer: producer,
		}),
	}
}

// UnregProducer represents the `flon.system::unregprod` action
type UnregProducer struct {
	Producer flon.AccountName `json:"producer"`
}
