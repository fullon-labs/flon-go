package rex

func NewUnstakeToREX(
	owner flon.AccountName,
	receiver flon.AccountName,
	fromNet flon.Asset,
	fromCPU flon.Asset,
) *flon.Action {
	return &flon.Action{
		Account: REXAN,
		Name:    ActN("unstaketorex"),
		Authorization: []flon.PermissionLevel{
			{Actor: owner, Permission: flon.PermissionName("active")},
		},
		ActionData: flon.NewActionData(UnstakeToREX{
			Owner:    owner,
			Receiver: receiver,
			FromNet:  fromNet,
			FromCPU:  fromCPU,
		}),
	}
}

type UnstakeToREX struct {
	Owner    flon.AccountName
	Receiver flon.AccountName
	FromNet  flon.Asset
	FromCPU  flon.Asset
}
