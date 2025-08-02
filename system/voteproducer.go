package system

import flon "github.com/fullon-labs/flon-go"

// NewNonce returns a `nonce` action that lives on the
// `flon.bios` contract. It should exist only when booting a new
// network, as it is replaced using the `eos-bios` boot process by the
// `flon.system` contract.
func NewVoteProducer(voter flon.AccountName, proxy flon.AccountName, producers ...flon.AccountName) *flon.Action {
	a := &flon.Action{
		Account: AN("flon"),
		Name:    ActN("voteproducer"),
		Authorization: []flon.PermissionLevel{
			{Actor: voter, Permission: PN("active")},
		},
		ActionData: flon.NewActionData(
			VoteProducer{
				Voter:     voter,
				Proxy:     proxy,
				Producers: producers,
			},
		),
	}
	return a
}

// VoteProducer represents the `flon.system::voteproducer` action
type VoteProducer struct {
	Voter     flon.AccountName   `json:"voter"`
	Proxy     flon.AccountName   `json:"proxy"`
	Producers []flon.AccountName `json:"producers"`
}
