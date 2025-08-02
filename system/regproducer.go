package system

import (
	flon "github.com/fullon-labs/flon-go"
	"github.com/fullon-labs/flon-go/ecc"
)

// NewRegProducer returns a `regproducer` action that lives on the
// `flon.system` contract.
func NewRegProducer(producer flon.AccountName, producerKey ecc.PublicKey, url string, location uint16) *flon.Action {
	return &flon.Action{
		Account: AN("flon"),
		Name:    ActN("regproducer"),
		Authorization: []flon.PermissionLevel{
			{Actor: producer, Permission: PN("active")},
		},
		ActionData: flon.NewActionData(RegProducer{
			Producer:    producer,
			ProducerKey: producerKey,
			URL:         url,
			Location:    location,
		}),
	}
}

// RegProducer represents the `flon.system::regproducer` action
type RegProducer struct {
	Producer    flon.AccountName `json:"producer"`
	ProducerKey ecc.PublicKey    `json:"producer_key"`
	URL         string           `json:"url"`
	Location    uint16           `json:"location"` // what,s the meaning of that anyway ?
}
