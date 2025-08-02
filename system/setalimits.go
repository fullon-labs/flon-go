package system

import flon "github.com/fullon-labs/flon-go"

// NewSetalimits sets the account limits. Requires signature from `flon@active` account.
func NewSetalimits(account flon.AccountName, ramBytes, netWeight, cpuWeight int64) *flon.Action {
	a := &flon.Action{
		Account: AN("flon"),
		Name:    ActN("setalimit"),
		Authorization: []flon.PermissionLevel{
			{Actor: flon.AccountName("flon"), Permission: PN("active")},
		},
		ActionData: flon.NewActionData(Setalimits{
			Account:   account,
			RAMBytes:  ramBytes,
			NetWeight: netWeight,
			CPUWeight: cpuWeight,
		}),
	}
	return a
}

// Setalimits represents the `flon.system::setalimit` action.
type Setalimits struct {
	Account   flon.AccountName `json:"account"`
	RAMBytes  int64            `json:"ram_bytes"`
	NetWeight int64            `json:"net_weight"`
	CPUWeight int64            `json:"cpu_weight"`
}
