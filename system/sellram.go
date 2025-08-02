package system

import flon "github.com/fullon-labs/flon-go"

// NewSellRAM will sell at current market price a given number of
// bytes of RAM.
func NewSellRAM(account flon.AccountName, bytes uint64) *flon.Action {
	a := &flon.Action{
		Account: AN("flon"),
		Name:    ActN("sellram"),
		Authorization: []flon.PermissionLevel{
			{Actor: account, Permission: flon.PermissionName("active")},
		},
		ActionData: flon.NewActionData(SellRAM{
			Account: account,
			Bytes:   bytes,
		}),
	}
	return a
}

// SellRAM represents the `flon.system::sellram` action.
type SellRAM struct {
	Account flon.AccountName `json:"account"`
	Bytes   uint64           `json:"bytes"`
}
