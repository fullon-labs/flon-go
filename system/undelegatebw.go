package system

import flon "github.com/fullon-labs/flon-go"

// NewUndelegateBW returns a `undelegatebw` action that lives on the
// `flon.system` contract.
func NewUndelegateBW(from, receiver flon.AccountName, unstakeCPU, unstakeNet flon.Asset) *flon.Action {
	return &flon.Action{
		Account: AN("flon"),
		Name:    ActN("undelegatebw"),
		Authorization: []flon.PermissionLevel{
			{Actor: from, Permission: PN("active")},
		},
		ActionData: flon.NewActionData(UndelegateBW{
			From:       from,
			Receiver:   receiver,
			UnstakeNet: unstakeNet,
			UnstakeCPU: unstakeCPU,
		}),
	}
}

// UndelegateBW represents the `flon.system::undelegatebw` action.
type UndelegateBW struct {
	From       flon.AccountName `json:"from"`
	Receiver   flon.AccountName `json:"receiver"`
	UnstakeNet flon.Asset       `json:"unstake_net_quantity"`
	UnstakeCPU flon.Asset       `json:"unstake_cpu_quantity"`
}
