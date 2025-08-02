package msig

func init() {
	flon.RegisterAction(AN("flon.msig"), ActN("propose"), &Propose{})
	flon.RegisterAction(AN("flon.msig"), ActN("approve"), &Approve{})
	flon.RegisterAction(AN("flon.msig"), ActN("unapprove"), &Unapprove{})
	flon.RegisterAction(AN("flon.msig"), ActN("cancel"), &Cancel{})
	flon.RegisterAction(AN("flon.msig"), ActN("exec"), &Exec{})
}

var AN = flon.AN
var PN = flon.PN
var ActN = flon.ActN
