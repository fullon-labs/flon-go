package system

import flon "github.com/fullon-labs/flon-go"

func NewSetRAMRate(bytesPerBlock uint16) *flon.Action {
	a := &flon.Action{
		Account: AN("flon"),
		Name:    ActN("setram"),
		Authorization: []flon.PermissionLevel{
			{
				Actor:      AN("flon"),
				Permission: flon.PermissionName("active"),
			},
		},
		ActionData: flon.NewActionData(SetRAMRate{
			BytesPerBlock: bytesPerBlock,
		}),
	}
	return a
}

// SetRAMRate represents the system contract's `setramrate` action.
type SetRAMRate struct {
	BytesPerBlock uint16 `json:"bytes_per_block"`
}
