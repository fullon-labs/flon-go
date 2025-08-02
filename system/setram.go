package system

import "github.com/fullon-labs/flon-go"

func NewSetRAM(maxRAMSize uint64) *flon.Action {
	a := &flon.Action{
		Account: AN("flon"),
		Name:    ActN("setram"),
		Authorization: []flon.PermissionLevel{
			{
				Actor:      AN("flon"),
				Permission: flon.PermissionName("active"),
			},
		},
		ActionData: flon.NewActionData(SetRAM{
			MaxRAMSize: flon.Uint64(maxRAMSize),
		}),
	}
	return a
}

// SetRAM represents the hard-coded `setram` action.
type SetRAM struct {
	MaxRAMSize flon.Uint64 `json:"max_ram_size"`
}
