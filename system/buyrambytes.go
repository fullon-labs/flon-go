package system

import flon "github.com/fullon-labs/flon-go"

// NewBuyRAMBytes will buy at current market price a given number of
// bytes of RAM, and grant them to the `receiver` account.
func NewBuyRAMBytes(payer, receiver flon.AccountName, bytes uint32) *flon.Action {
	a := &flon.Action{
		Account: AN("flon"),
		Name:    ActN("buyrambytes"),
		Authorization: []flon.PermissionLevel{
			{Actor: payer, Permission: flon.PermissionName("active")},
		},
		ActionData: flon.NewActionData(BuyRAMBytes{
			Payer:    payer,
			Receiver: receiver,
			Bytes:    bytes,
		}),
	}
	return a
}

// BuyRAMBytes represents the `flon.system::buyrambytes` action.
type BuyRAMBytes struct {
	Payer    flon.AccountName `json:"payer"`
	Receiver flon.AccountName `json:"receiver"`
	Bytes    uint32           `json:"bytes"`
}
