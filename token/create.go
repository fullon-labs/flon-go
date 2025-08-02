package token

import "github.com/fullon-labs/flon-go"

func NewCreate(issuer flon.AccountName, maxSupply flon.Asset) *flon.Action {
	return &flon.Action{
		Account: AN("flon.token"),
		Name:    ActN("create"),
		Authorization: []flon.PermissionLevel{
			{Actor: AN("flon.token"), Permission: PN("active")},
		},
		ActionData: flon.NewActionData(Create{
			Issuer:        issuer,
			MaximumSupply: maxSupply,
		}),
	}
}

// Create represents the `create` struct on the `flon.token` contract.
type Create struct {
	Issuer        flon.AccountName `json:"issuer"`
	MaximumSupply flon.Asset       `json:"maximum_supply"`
}
