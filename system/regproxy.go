package system

import flon "github.com/fullon-labs/flon-go"

// NewRegProxy returns a `regproxy` action that lives on the
// `flon.system` contract.
func NewRegProxy(proxy flon.AccountName, isProxy bool) *flon.Action {
	return &flon.Action{
		Account: AN("flon"),
		Name:    ActN("regproxy"),
		Authorization: []flon.PermissionLevel{
			{Actor: proxy, Permission: PN("active")},
		},
		ActionData: flon.NewActionData(RegProxy{
			Proxy:   proxy,
			IsProxy: isProxy,
		}),
	}
}

// RegProxy represents the `flon.system::regproxy` action
type RegProxy struct {
	Proxy   flon.AccountName `json:"proxy"`
	IsProxy bool             `json:"isproxy"`
}
