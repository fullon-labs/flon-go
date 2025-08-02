package system

import "github.com/fullon-labs/flon-go"

// NewSetPriv returns a `setpriv` action that lives on the
// `flon.bios` contract. It should exist only when booting a new
// network, as it is replaced using the `eos-bios` boot process by the
// `flon.system` contract.
func NewSetPriv(account flon.AccountName) *flon.Action {
	a := &flon.Action{
		Account: AN("flon"),
		Name:    ActN("setpriv"),
		Authorization: []flon.PermissionLevel{
			{Actor: AN("flon"), Permission: PN("active")},
		},
		ActionData: flon.NewActionData(SetPriv{
			Account: account,
			IsPriv:  flon.Bool(true),
		}),
	}
	return a
}

// SetPriv sets privileged account status. Used in the bios boot mechanism.
type SetPriv struct {
	Account flon.AccountName `json:"account"`
	IsPriv  flon.Bool        `json:"is_priv"`
}
