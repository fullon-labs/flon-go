package token

import flon "github.com/fullon-labs/flon-go"

func init() {
	flon.RegisterAction(AN("flon.token"), ActN("transfer"), Transfer{})
	flon.RegisterAction(AN("flon.token"), ActN("issue"), Issue{})
	flon.RegisterAction(AN("flon.token"), ActN("create"), Create{})
}
