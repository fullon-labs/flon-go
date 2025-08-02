package system

import flon "github.com/fullon-labs/flon-go"

// NewInitSystem returns a `init` action that lives on the
// `flon.system` contract.
func NewInitSystem(version flon.Varuint32, core flon.Symbol) *flon.Action {
	return &flon.Action{
		Account: AN("flon"),
		Name:    ActN("init"),
		Authorization: []flon.PermissionLevel{
			{
				Actor:      AN("flon"),
				Permission: flon.PermissionName("active"),
			},
		},
		ActionData: flon.NewActionData(Init{
			Version: version,
			Core:    core,
		}),
	}
}

// Init represents the `flon.system::init` action
type Init struct {
	Version flon.Varuint32 `json:"version"`
	Core    flon.Symbol    `json:"core"`
}
