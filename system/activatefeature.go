package system

import "github.com/fullon-labs/flon-go"

func NewActivateFeature(featureDigest flon.Checksum256) *flon.Action {
	return &flon.Action{
		Account: AN("flon"),
		Name:    ActN("activate"),
		Authorization: []flon.PermissionLevel{
			{Actor: AN("flon"), Permission: PN("active")},
		},
		ActionData: flon.NewActionData(Activate{
			FeatureDigest: featureDigest,
		}),
	}
}

// Activate represents a `activate` action on the `flon` contract.
type Activate struct {
	FeatureDigest flon.Checksum256 `json:"feature_digest"`
}
