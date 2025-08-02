package rex

func init() {
	flon.RegisterAction(REXAN, ActN("buyrex"), BuyREX{})
	flon.RegisterAction(REXAN, ActN("closerex"), CloseREX{})
	flon.RegisterAction(REXAN, ActN("cnclrexorder"), CancelREXOrder{})
	flon.RegisterAction(REXAN, ActN("consolidate"), Consolidate{})
	flon.RegisterAction(REXAN, ActN("defcpuloan"), DefundCPULoan{})
	flon.RegisterAction(REXAN, ActN("defnetloan"), DefundNetLoan{})
	flon.RegisterAction(REXAN, ActN("deposit"), Deposit{})
	flon.RegisterAction(REXAN, ActN("fundcpuloan"), FundCPULoan{})
	flon.RegisterAction(REXAN, ActN("fundnetloan"), FundNetLoan{})
	flon.RegisterAction(REXAN, ActN("mvfrsavings"), MoveFromSavings{})
	flon.RegisterAction(REXAN, ActN("mvtosavings"), MoveToSavings{})
	flon.RegisterAction(REXAN, ActN("rentcpu"), RentCPU{})
	flon.RegisterAction(REXAN, ActN("rentnet"), RentNet{})
	flon.RegisterAction(REXAN, ActN("rexexec"), REXExec{})
	flon.RegisterAction(REXAN, ActN("sellrex"), SellREX{})
	flon.RegisterAction(REXAN, ActN("unstaketorex"), UnstakeToREX{})
	flon.RegisterAction(REXAN, ActN("updaterex"), UpdateREX{})
	flon.RegisterAction(REXAN, ActN("withdraw"), Withdraw{})
}

var AN = flon.AN
var PN = flon.PN
var ActN = flon.ActN

var REXAN = AN("eosio")
