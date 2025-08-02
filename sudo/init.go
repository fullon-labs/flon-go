package sudo

import "github.com/fullon-labs/flon-go"

func init() {
	flon.RegisterAction(AN("eosio.wrap"), ActN("exec"), Exec{})
}

var AN = flon.AN
var ActN = flon.ActN
