package ship

import (
	"fmt"
)

func NewGetBlocksAck(num uint32) []byte {
	myReq := &Request{
		BaseVariant: flon.BaseVariant{
			TypeID: RequestVariant.TypeID("get_blocks_ack_request_v0"),
			Impl: &GetBlocksAckRequestV0{
				NumMessages: num,
			},
		},
	}
	bytes, err := flon.MarshalBinary(myReq)
	if err != nil {
		panic(err)
	}

	return bytes
}

func NewRequest(req *GetBlocksRequestV0) []byte {
	myReq := &Request{
		BaseVariant: flon.BaseVariant{
			TypeID: RequestVariant.TypeID("get_blocks_request_v0"),
			Impl:   req,
		},
	}
	bytes, err := flon.MarshalBinary(myReq)
	if err != nil {
		panic(err)
	}

	return bytes
}

func ParseGetBlockResultV0(in []byte) (*GetBlocksResultV0, error) {
	variant := &Result{}
	if err := flon.UnmarshalBinary(in, &variant); err != nil {
		return nil, err
	}

	v, ok := variant.Impl.(*GetBlocksResultV0)
	if !ok {
		return nil, fmt.Errorf("invalid response type: %d", variant.TypeID)
	}
	return v, nil
}
