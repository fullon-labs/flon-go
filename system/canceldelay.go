package system

import "github.com/fullon-labs/flon-go"

// NewCancelDelay creates an action from the `flon.system` contract
// called `canceldelay`.
//
// `canceldelay` allows you to cancel a deferred transaction,
// previously sent to the chain with a `delay_sec` larger than 0.  You
// need to sign with cancelingAuth, to cancel a transaction signed
// with that same authority.
func NewCancelDelay(cancelingAuth flon.PermissionLevel, transactionID flon.Checksum256) *flon.Action {
	a := &flon.Action{
		Account: AN("flon"),
		Name:    ActN("canceldelay"),
		Authorization: []flon.PermissionLevel{
			cancelingAuth,
		},
		ActionData: flon.NewActionData(CancelDelay{
			CancelingAuth: cancelingAuth,
			TransactionID: transactionID,
		}),
	}

	return a
}

// CancelDelay represents the native `canceldelay` action, through the
// system contract.
type CancelDelay struct {
	CancelingAuth flon.PermissionLevel `json:"canceling_auth"`
	TransactionID flon.Checksum256     `json:"trx_id"`
}
