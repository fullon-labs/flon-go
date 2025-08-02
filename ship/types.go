package ship

import (
	"github.com/fullon-labs/flon-go/ecc"
)

// State History Plugin Requests

type GetStatusRequestV0 struct {
}

type GetBlocksAckRequestV0 struct {
	NumMessages uint32
}

type GetBlocksRequestV0 struct {
	StartBlockNum       uint32
	EndBlockNum         uint32
	MaxMessagesInFlight uint32
	HavePositions       []*BlockPosition
	IrreversibleOnly    bool
	FetchBlock          bool
	FetchTraces         bool
	FetchDeltas         bool
}

// State History Plugin Results
type GetStatusResultV0 struct {
	Head                 *BlockPosition
	LastIrreversible     *BlockPosition
	TraceBeginBlock      uint32
	TraceEndBlock        uint32
	ChainStateBeginBlock uint32
	ChainStateEndBlock   uint32
}

type GetBlocksResultV0 struct {
	Head             *BlockPosition
	LastIrreversible *BlockPosition
	ThisBlock        *BlockPosition         `eos:"optional"`
	PrevBlock        *BlockPosition         `eos:"optional"`
	Block            *SignedBlockBytes      `eos:"optional"`
	Traces           *TransactionTraceArray `eos:"optional"`
	Deltas           *TableDeltaArray       `eos:"optional"`
}

// State History Plugin version of EOS structs
type BlockPosition struct {
	BlockNum uint32
	BlockID  flon.Checksum256
}

type Row struct {
	Present bool
	Data    []byte
}

type ActionTraceV0 struct {
	ActionOrdinal        flon.Varuint32
	CreatorActionOrdinal flon.Varuint32
	Receipt              *ActionReceipt `eos:"optional"`
	Receiver             flon.Name
	Act                  *Action
	ContextFree          bool
	Elapsed              int64
	Console              flon.SafeString
	AccountRamDeltas     []*flon.AccountDelta
	Except               string `eos:"optional"`
	ErrorCode            uint64 `eos:"optional"`
}

type ActionTraceV1 struct {
	ActionOrdinal        flon.Varuint32
	CreatorActionOrdinal flon.Varuint32
	Receipt              *ActionReceipt `eos:"optional"`
	Receiver             flon.Name
	Act                  *Action
	ContextFree          bool
	Elapsed              int64
	Console              flon.SafeString
	AccountRamDeltas     []*flon.AccountDelta
	Except               string `eos:"optional"`
	ErrorCode            uint64 `eos:"optional"`
	ReturnValue          []byte
}

type Action struct {
	Account       flon.AccountName
	Name          flon.ActionName
	Authorization []flon.PermissionLevel
	Data          []byte
}

type ActionReceiptV0 struct {
	Receiver       flon.Name
	ActDigest      flon.Checksum256
	GlobalSequence uint64
	RecvSequence   uint64
	AuthSequence   []AccountAuthSequence
	CodeSequence   flon.Varuint32
	ABISequence    flon.Varuint32
}

type AccountAuthSequence struct {
	Account  flon.Name
	Sequence uint64
}

type TableDeltaV0 struct {
	Name string
	Rows []Row
}

type PartialTransactionV0 struct {
	Expiration            uint32
	RefBlockNum           uint16
	RefBlockPrefix        uint32
	MaxNetUsageWords      flon.Varuint32
	MaxCpuUsageMs         uint8
	DelaySec              flon.Varuint32
	TransactionExtensions []*Extension
	Signatures            []ecc.Signature
	ContextFreeData       []byte
}

type TransactionTraceV0 struct {
	ID              flon.Checksum256 `json:"id"`
	Status          flon.TransactionStatus
	CPUUsageUS      uint32                `json:"cpu_usage_us"`
	NetUsageWords   flon.Varuint32        `json:"net_usage_words"`
	Elapsed         flon.Int64            `json:"elapsed"`
	NetUsage        uint64                `json:"net_usage"`
	Scheduled       bool                  `json:"scheduled"`
	ActionTraces    []*ActionTrace        `json:"action_traces"`
	AccountDelta    *flon.AccountRAMDelta `json:"account_delta" eos:"optional"`
	Except          string                `json:"except" eos:"optional"`
	ErrorCode       uint64                `json:"error_code" eos:"optional"`
	FailedDtrxTrace *TransactionTrace     `json:"failed_dtrx_trace" eos:"optional"`
	Partial         *PartialTransaction   `json:"partial" eos:"optional"`
}

type SignedBlockHeader struct {
	flon.BlockHeader
	ProducerSignature ecc.Signature // no pointer!!
}

type TransactionReceipt struct {
	flon.TransactionReceiptHeader
	Trx *Transaction
}

//type TransactionID flon.Checksum256

type SignedBlock struct {
	SignedBlockHeader
	Transactions    []*TransactionReceipt
	BlockExtensions []*Extension
}

type SignedBlockBytes SignedBlock

func (s *SignedBlockBytes) AsSignedBlock() *SignedBlock {
	if s == nil {
		return nil
	}
	ss := SignedBlock(*s)
	return &ss
}

func (s *SignedBlockBytes) UnmarshalBinary(decoder *flon.Decoder) error {
	data, err := decoder.ReadByteArray()
	if err != nil {
		return err
	}
	return flon.UnmarshalBinary(data, (*SignedBlock)(s))
}

type Extension struct {
	Type uint16
	Data []byte
}
