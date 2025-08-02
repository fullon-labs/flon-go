package system

import flon "github.com/fullon-labs/flon-go"

func init() {
	flon.RegisterAction(AN("flon"), ActN("setcode"), SetCode{})
	flon.RegisterAction(AN("flon"), ActN("setabi"), SetABI{})
	flon.RegisterAction(AN("flon"), ActN("newaccount"), NewAccount{})
	flon.RegisterAction(AN("flon"), ActN("delegatebw"), DelegateBW{})
	flon.RegisterAction(AN("flon"), ActN("undelegatebw"), UndelegateBW{})
	flon.RegisterAction(AN("flon"), ActN("refund"), Refund{})
	flon.RegisterAction(AN("flon"), ActN("regproducer"), RegProducer{})
	flon.RegisterAction(AN("flon"), ActN("unregprod"), UnregProducer{})
	flon.RegisterAction(AN("flon"), ActN("regproxy"), RegProxy{})
	flon.RegisterAction(AN("flon"), ActN("voteproducer"), VoteProducer{})
	flon.RegisterAction(AN("flon"), ActN("claimrewards"), ClaimRewards{})
	flon.RegisterAction(AN("flon"), ActN("buyram"), BuyRAM{})
	flon.RegisterAction(AN("flon"), ActN("buyrambytes"), BuyRAMBytes{})
	flon.RegisterAction(AN("flon"), ActN("linkauth"), LinkAuth{})
	flon.RegisterAction(AN("flon"), ActN("unlinkauth"), UnlinkAuth{})
	flon.RegisterAction(AN("flon"), ActN("deleteauth"), DeleteAuth{})
	flon.RegisterAction(AN("flon"), ActN("rmvproducer"), RemoveProducer{})
	flon.RegisterAction(AN("flon"), ActN("setprods"), SetProds{})
	flon.RegisterAction(AN("flon"), ActN("setpriv"), SetPriv{})
	flon.RegisterAction(AN("flon"), ActN("canceldelay"), CancelDelay{})
	flon.RegisterAction(AN("flon"), ActN("bidname"), Bidname{})
	// flon.RegisterAction(AN("flon"), ActN("nonce"), &Nonce{})
	flon.RegisterAction(AN("flon"), ActN("sellram"), SellRAM{})
	flon.RegisterAction(AN("flon"), ActN("updateauth"), UpdateAuth{})
	flon.RegisterAction(AN("flon"), ActN("setramrate"), SetRAMRate{})
	flon.RegisterAction(AN("flon"), ActN("setalimits"), Setalimits{})
}

var AN = flon.AN
var PN = flon.PN
var ActN = flon.ActN
