package system

import flon "github.com/fullon-labs/flon-go"

// NewNonce returns a `nonce` action that lives on the
// `flon.bios` contract. It should exist only when booting a new
// network, as it is replaced using the `eos-bios` boot process by the
// `flon.system` contract.
func NewNonce(nonce string) *flon.Action {
	a := &flon.Action{
		Account:       AN("flon"),
		Name:          ActN("nonce"),
		Authorization: []flon.PermissionLevel{
			//{Actor: AN("flon"), Permission: PN("active")},
		},
		ActionData: flon.NewActionData(Nonce{
			Value: nonce,
		}),
	}
	return a
}
