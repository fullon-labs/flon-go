package token

import flon "github.com/fullon-labs/flon-go"

func NewTransfer(from, to flon.AccountName, quantity flon.Asset, memo string) *flon.Action {
	return &flon.Action{
		Account: AN("flon.token"),
		Name:    ActN("transfer"),
		Authorization: []flon.PermissionLevel{
			{Actor: from, Permission: PN("active")},
		},
		ActionData: flon.NewActionData(Transfer{
			From:     from,
			To:       to,
			Quantity: quantity,
			Memo:     memo,
		}),
	}
}

// Transfer represents the `transfer` struct on `flon.token` contract.
type Transfer struct {
	From     flon.AccountName `json:"from"`
	To       flon.AccountName `json:"to"`
	Quantity flon.Asset       `json:"quantity"`
	Memo     string           `json:"memo"`
}
