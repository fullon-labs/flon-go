package token

import "github.com/fullon-labs/flon-go"

func NewIssue(to flon.AccountName, quantity flon.Asset, memo string) *flon.Action {
	return &flon.Action{
		Account: AN("flon.token"),
		Name:    ActN("issue"),
		Authorization: []flon.PermissionLevel{
			{Actor: to, Permission: PN("active")},
		},
		ActionData: flon.NewActionData(Issue{
			To:       to,
			Quantity: quantity,
			Memo:     memo,
		}),
	}
}

// Issue represents the `issue` struct on the `flon.token` contract.
type Issue struct {
	To       flon.AccountName `json:"to"`
	Quantity flon.Asset       `json:"quantity"`
	Memo     string           `json:"memo"`
}
