package system

import (
	flon "github.com/fullon-labs/flon-go"
	"github.com/fullon-labs/flon-go/ecc"
)

// NewNewAccount returns a `newaccount` action that lives on the
// `flon.system` contract.
func NewNewAccount(creator, newAccount flon.AccountName, publicKey ecc.PublicKey) *flon.Action {
	return &flon.Action{
		Account: AN("flon"),
		Name:    ActN("newaccount"),
		Authorization: []flon.PermissionLevel{
			{Actor: creator, Permission: PN("active")},
		},
		ActionData: flon.NewActionData(NewAccount{
			Creator: creator,
			Name:    newAccount,
			Owner: flon.Authority{
				Threshold: 1,
				Keys: []flon.KeyWeight{
					{
						PublicKey: publicKey,
						Weight:    1,
					},
				},
				Accounts: []flon.PermissionLevelWeight{},
			},
			Active: flon.Authority{
				Threshold: 1,
				Keys: []flon.KeyWeight{
					{
						PublicKey: publicKey,
						Weight:    1,
					},
				},
				Accounts: []flon.PermissionLevelWeight{},
			},
		}),
	}
}

// NewDelegatedNewAccount returns a `newaccount` action that lives on the
// `flon.system` contract. It is filled with an authority structure that
// delegates full control of the new account to an already existing account.
func NewDelegatedNewAccount(creator, newAccount flon.AccountName, delegatedTo flon.AccountName) *flon.Action {
	return &flon.Action{
		Account: AN("flon"),
		Name:    ActN("newaccount"),
		Authorization: []flon.PermissionLevel{
			{Actor: creator, Permission: PN("active")},
		},
		ActionData: flon.NewActionData(NewAccount{
			Creator: creator,
			Name:    newAccount,
			Owner: flon.Authority{
				Threshold: 1,
				Keys:      []flon.KeyWeight{},
				Accounts: []flon.PermissionLevelWeight{
					flon.PermissionLevelWeight{
						Permission: flon.PermissionLevel{
							Actor:      delegatedTo,
							Permission: PN("active"),
						},
						Weight: 1,
					},
				},
			},
			Active: flon.Authority{
				Threshold: 1,
				Keys:      []flon.KeyWeight{},
				Accounts: []flon.PermissionLevelWeight{
					flon.PermissionLevelWeight{
						Permission: flon.PermissionLevel{
							Actor:      delegatedTo,
							Permission: PN("active"),
						},
						Weight: 1,
					},
				},
			},
		}),
	}
}

// NewCustomNewAccount returns a `newaccount` action that lives on the
// `flon.system` contract. You can specify your own `owner` and
// `active` permissions.
func NewCustomNewAccount(creator, newAccount flon.AccountName, owner, active flon.Authority) *flon.Action {
	return &flon.Action{
		Account: AN("flon"),
		Name:    ActN("newaccount"),
		Authorization: []flon.PermissionLevel{
			{Actor: creator, Permission: PN("active")},
		},
		ActionData: flon.NewActionData(NewAccount{
			Creator: creator,
			Name:    newAccount,
			Owner:   owner,
			Active:  active,
		}),
	}
}

// NewAccount represents a `newaccount` action on the `flon.system`
// contract. It is one of the rare ones to be hard-coded into the
// blockchain.
type NewAccount struct {
	Creator flon.AccountName `json:"creator"`
	Name    flon.AccountName `json:"name"`
	Owner   flon.Authority   `json:"owner"`
	Active  flon.Authority   `json:"active"`
}
