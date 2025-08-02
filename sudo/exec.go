package sudo

import "github.com/fullon-labs/flon-go"

// NewExec creates an `exec` action, found in the `eosio.wrap`
// contract.
//
// Given an `flon.Transaction`, call `flon.MarshalBinary` on it first,
// pass the resulting bytes as `flon.HexBytes` here.
func NewExec(executer flon.AccountName, transaction flon.Transaction) *flon.Action {
	a := &flon.Action{
		Account: flon.AccountName("eosio.wrap"),
		Name:    flon.ActionName("exec"),
		Authorization: []flon.PermissionLevel{
			{Actor: executer, Permission: flon.PermissionName("active")},
		},
		ActionData: flon.NewActionData(Exec{
			Executer:    executer,
			Transaction: transaction,
		}),
	}
	return a
}

// Exec represents the `eosio.system::exec` action.
type Exec struct {
	Executer    flon.AccountName `json:"executer"`
	Transaction flon.Transaction `json:"trx"`
}
