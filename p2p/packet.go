package p2p

type Envelope struct {
	Sender   *Peer
	Receiver *Peer
	Packet   *flon.Packet `json:"envelope"`
}

func NewEnvelope(sender *Peer, receiver *Peer, packet *flon.Packet) *Envelope {
	return &Envelope{
		Sender:   sender,
		Receiver: receiver,
		Packet:   packet,
	}
}
