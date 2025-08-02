package system

import flon "github.com/fullon-labs/flon-go"

func NewBidname(bidder, newname flon.AccountName, bid flon.Asset) *flon.Action {
	a := &flon.Action{
		Account: AN("flon"),
		Name:    ActN("bidname"),
		Authorization: []flon.PermissionLevel{
			{Actor: bidder, Permission: PN("active")},
		},
		ActionData: flon.NewActionData(Bidname{
			Bidder:  bidder,
			Newname: newname,
			Bid:     bid,
		}),
	}
	return a
}

// Bidname represents the `flon.system_contract::bidname` action.
type Bidname struct {
	Bidder  flon.AccountName `json:"bidder"`
	Newname flon.AccountName `json:"newname"`
	Bid     flon.Asset       `json:"bid"` // specified in EOS
}
