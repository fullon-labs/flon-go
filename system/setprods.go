package system

import (
	flon "github.com/fullon-labs/flon-go"
	"github.com/fullon-labs/flon-go/ecc"
)

// NewSetPriv returns a `setpriv` action that lives on the
// `flon.bios` contract. It should exist only when booting a new
// network, as it is replaced using the `eos-bios` boot process by the
// `flon.system` contract.
func NewSetProds(producers []ProducerKey) *flon.Action {
	a := &flon.Action{
		Account: AN("flon"),
		Name:    ActN("setprods"),
		Authorization: []flon.PermissionLevel{
			{Actor: AN("flon"), Permission: PN("active")},
		},
		ActionData: flon.NewActionData(SetProds{
			Schedule: producers,
		}),
	}
	return a
}

// SetProds is present in `flon.bios` contract. Used only at boot time.
type SetProds struct {
	Schedule []ProducerKey `json:"schedule"`
}

type ProducerKey struct {
	ProducerName    flon.AccountName `json:"producer_name"`
	BlockSigningKey ecc.PublicKey    `json:"block_signing_key"`
}
