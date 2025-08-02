package system

import flon "github.com/fullon-labs/flon-go"

// NewDelegateBW returns a `delegatebw` action that lives on the
// `flon.system` contract.
func NewDelegateBW(from, receiver flon.AccountName, stakeCPU, stakeNet flon.Asset, transfer bool) *flon.Action {
	return &flon.Action{
		Account: AN("flon"),
		Name:    ActN("delegatebw"),
		Authorization: []flon.PermissionLevel{
			{Actor: from, Permission: PN("active")},
		},
		ActionData: flon.NewActionData(DelegateBW{
			From:     from,
			Receiver: receiver,
			StakeNet: stakeNet,
			StakeCPU: stakeCPU,
			Transfer: flon.Bool(transfer),
		}),
	}
}

// DelegateBW represents the `flon.system::delegatebw` action.
type DelegateBW struct {
	From     flon.AccountName `json:"from"`
	Receiver flon.AccountName `json:"receiver"`
	StakeNet flon.Asset       `json:"stake_net_quantity"`
	StakeCPU flon.Asset       `json:"stake_cpu_quantity"`
	Transfer flon.Bool        `json:"transfer"`
}
