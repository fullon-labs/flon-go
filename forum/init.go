package forum

func init() {
	flon.RegisterAction(ForumAN, ActN("clnproposal"), CleanProposal{})
	flon.RegisterAction(ForumAN, ActN("expire"), Expire{})
	flon.RegisterAction(ForumAN, ActN("post"), Post{})
	flon.RegisterAction(ForumAN, ActN("propose"), Propose{})
	flon.RegisterAction(ForumAN, ActN("status"), Status{})
	flon.RegisterAction(ForumAN, ActN("unpost"), UnPost{})
	flon.RegisterAction(ForumAN, ActN("unvote"), UnVote{})
	flon.RegisterAction(ForumAN, ActN("vote"), Vote{})
}

var AN = flon.AN
var PN = flon.PN
var ActN = flon.ActN

var ForumAN = AN("eosio.forum")
