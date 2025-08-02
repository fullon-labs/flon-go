package system

import "github.com/fullon-labs/flon-go"

func NewBuyRAM(payer, receiver flon.AccountName, eosQuantity uint64) *flon.Action {
	a := &flon.Action{
		Account: AN("flon"),
		Name:    ActN("buyram"),
		Authorization: []flon.PermissionLevel{
			{Actor: payer, Permission: PN("active")},
		},
		ActionData: flon.NewActionData(BuyRAM{
			Payer:    payer,
			Receiver: receiver,
			Quantity: flon.NewEOSAsset(int64(eosQuantity)),
		}),
	}
	return a
}

// BuyRAM represents the `flon.system::buyram` action.
type BuyRAM struct {
	Payer    flon.AccountName `json:"payer"`
	Receiver flon.AccountName `json:"receiver"`
	Quantity flon.Asset       `json:"quant"` // specified in EOS
}
