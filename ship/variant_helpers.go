package ship

// Request
var RequestVariant = flon.NewVariantDefinition([]flon.VariantType{
	{"get_status_request_v0", (*GetStatusRequestV0)(nil)},
	{"get_blocks_request_v0", (*GetBlocksRequestV0)(nil)},
	{"get_blocks_ack_request_v0", (*GetBlocksAckRequestV0)(nil)},
})

type Request struct {
	flon.BaseVariant
}

func (r *Request) UnmarshalBinary(decoder *flon.Decoder) error {
	return r.BaseVariant.UnmarshalBinaryVariant(decoder, RequestVariant)
}

// Result
var ResultVariant = flon.NewVariantDefinition([]flon.VariantType{
	{"get_status_result_v0", (*GetStatusResultV0)(nil)},
	{"get_blocks_result_v0", (*GetBlocksResultV0)(nil)},
})

type Result struct {
	flon.BaseVariant
}

func (r *Result) UnmarshalBinary(decoder *flon.Decoder) error {
	return r.BaseVariant.UnmarshalBinaryVariant(decoder, ResultVariant)
}

// TransactionTrace
var TransactionTraceVariant = flon.NewVariantDefinition([]flon.VariantType{
	{"transaction_trace_v0", (*TransactionTraceV0)(nil)},
})

type TransactionTrace struct {
	flon.BaseVariant
}

type TransactionTraceArray struct {
	Elem []*TransactionTrace
}

func (t *TransactionTraceArray) AsTransactionTracesV0() (out []*TransactionTraceV0) {
	if t == nil || t.Elem == nil {
		return nil
	}
	for _, e := range t.Elem {
		switch v := e.Impl.(type) {
		case *TransactionTraceV0:
			out = append(out, v)

		default:
			panic("wrong type for conversion")
		}
	}
	return out
}

func (r TransactionTraceArray) MarshalBinary(enc *flon.Encoder) error {
	data, err := flon.MarshalBinary(r.Elem)
	if err != nil {
		return err
	}
	return enc.Encode(data)
}

func (r *TransactionTraceArray) UnmarshalBinary(decoder *flon.Decoder) error {
	data, err := decoder.ReadByteArray()
	if err != nil {
		return err
	}
	return flon.UnmarshalBinary(data, &r.Elem)
}

func (r *TransactionTrace) UnmarshalBinary(decoder *flon.Decoder) error {
	return r.BaseVariant.UnmarshalBinaryVariant(decoder, TransactionTraceVariant)
}

// ActionTrace
var ActionTraceVariant = flon.NewVariantDefinition([]flon.VariantType{
	{"action_trace_v0", (*ActionTraceV0)(nil)},
	{"action_trace_v1", (*ActionTraceV1)(nil)},
})

type ActionTrace struct {
	flon.BaseVariant
}

func (r *ActionTrace) UnmarshalBinary(decoder *flon.Decoder) error {
	return r.BaseVariant.UnmarshalBinaryVariant(decoder, ActionTraceVariant)
}

// PartialTransaction
var PartialTransactionVariant = flon.NewVariantDefinition([]flon.VariantType{
	{"partial_transaction_v0", (*PartialTransactionV0)(nil)},
})

type PartialTransaction struct {
	flon.BaseVariant
}

func (r *PartialTransaction) UnmarshalBinary(decoder *flon.Decoder) error {
	return r.BaseVariant.UnmarshalBinaryVariant(decoder, PartialTransactionVariant)
}

// TableDelta
var TableDeltaVariant = flon.NewVariantDefinition([]flon.VariantType{
	{"table_delta_v0", (*TableDeltaV0)(nil)},
})

type TableDelta struct {
	flon.BaseVariant
}

func (d *TableDelta) UnmarshalBinary(decoder *flon.Decoder) error {
	return d.BaseVariant.UnmarshalBinaryVariant(decoder, TableDeltaVariant)
}

type TableDeltaArray struct {
	Elem []*TableDelta
}

func (d TableDeltaArray) MarshalBinary(enc *flon.Encoder) error {
	data, err := flon.MarshalBinary(d.Elem)
	if err != nil {
		return err
	}
	return enc.Encode(data)
}

func (d *TableDeltaArray) UnmarshalBinary(decoder *flon.Decoder) error {
	data, err := decoder.ReadByteArray()
	if err != nil {
		return err
	}
	return flon.UnmarshalBinary(data, &d.Elem)
}

func (t *TableDeltaArray) AsTableDeltasV0() (out []*TableDeltaV0) {
	if t == nil || t.Elem == nil {
		return nil
	}
	for _, e := range t.Elem {
		switch v := e.Impl.(type) {
		case *TableDeltaV0:
			out = append(out, v)

		default:
			panic("wrong type for conversion")
		}
	}
	return out
}

// Transaction
var TransactionVariant = flon.NewVariantDefinition([]flon.VariantType{
	{"transaction_id", (*flon.Checksum256)(nil)},
	{"packed_transaction", (*flon.PackedTransaction)(nil)},
})

type Transaction struct {
	flon.BaseVariant
}

func (d *Transaction) UnmarshalBinary(decoder *flon.Decoder) error {
	return d.BaseVariant.UnmarshalBinaryVariant(decoder, TransactionVariant)
}

// ActionReceipt
var ActionReceiptVariant = flon.NewVariantDefinition([]flon.VariantType{
	{"action_receipt_v0", (*ActionReceiptV0)(nil)},
})

type ActionReceipt struct {
	flon.BaseVariant
}

func (r *ActionReceipt) UnmarshalBinary(decoder *flon.Decoder) error {
	return r.BaseVariant.UnmarshalBinaryVariant(decoder, ActionReceiptVariant)
}
