package system

import (
	"encoding/json"
	"fmt"
	flon "github.com/fullon-labs/flon-go"
	"io/ioutil"
)

func NewSetContract(account flon.AccountName, wasmPath, abiPath string) (out []*flon.Action, err error) {
	codeAction, err := NewSetCode(account, wasmPath)
	if err != nil {
		return nil, err
	}

	abiAction, err := NewSetABI(account, abiPath)
	if err != nil {
		return nil, err
	}

	return []*flon.Action{codeAction, abiAction}, nil
}

func NewSetContractContent(account flon.AccountName, wasmContent, abiContent []byte) (out []*flon.Action, err error) {
	codeAction := NewSetCodeContent(account, wasmContent)

	abiAction, err := NewSetAbiContent(account, abiContent)
	if err != nil {
		return nil, err
	}

	return []*flon.Action{codeAction, abiAction}, nil
}

func NewSetCode(account flon.AccountName, wasmPath string) (out *flon.Action, err error) {
	codeContent, err := ioutil.ReadFile(wasmPath)
	if err != nil {
		return nil, err
	}
	return NewSetCodeContent(account, codeContent), nil
}

func NewSetCodeContent(account flon.AccountName, codeContent []byte) *flon.Action {
	return &flon.Action{
		Account: AN("flon"),
		Name:    ActN("setcode"),
		Authorization: []flon.PermissionLevel{
			{
				Actor:      account,
				Permission: flon.PermissionName("active"),
			},
		},
		ActionData: flon.NewActionData(SetCode{
			Account:   account,
			VMType:    0,
			VMVersion: 0,
			Code:      flon.HexBytes(codeContent),
		}),
	}
}

func NewSetABI(account flon.AccountName, abiPath string) (out *flon.Action, err error) {
	abiContent, err := ioutil.ReadFile(abiPath)
	if err != nil {
		return nil, err
	}

	return NewSetAbiContent(account, abiContent)
}

func NewSetAbiContent(account flon.AccountName, abiContent []byte) (out *flon.Action, err error) {
	var abiPacked []byte
	if len(abiContent) > 0 {
		var abiDef flon.ABI
		if err := json.Unmarshal(abiContent, &abiDef); err != nil {
			return nil, fmt.Errorf("unmarshal ABI file: %w", err)
		}

		abiPacked, err = flon.MarshalBinary(abiDef)
		if err != nil {
			return nil, fmt.Errorf("packing ABI: %w", err)
		}
	}

	return &flon.Action{
		Account: AN("flon"),
		Name:    ActN("setabi"),
		Authorization: []flon.PermissionLevel{
			{
				Actor:      account,
				Permission: flon.PermissionName("active"),
			},
		},
		ActionData: flon.NewActionData(SetABI{
			Account: account,
			ABI:     flon.HexBytes(abiPacked),
		}),
	}, nil
}

func NewSetAbiFromAbi(account flon.AccountName, abi flon.ABI) (out *flon.Action, err error) {
	var abiPacked []byte
	abiPacked, err = flon.MarshalBinary(abi)
	if err != nil {
		return nil, fmt.Errorf("packing ABI: %w", err)
	}

	return &flon.Action{
		Account: AN("flon"),
		Name:    ActN("setabi"),
		Authorization: []flon.PermissionLevel{
			{
				Actor:      account,
				Permission: flon.PermissionName("active"),
			},
		},
		ActionData: flon.NewActionData(SetABI{
			Account: account,
			ABI:     flon.HexBytes(abiPacked),
		}),
	}, nil
}

// NewSetCodeTx is _deprecated_. Use NewSetContract instead, and build
// your transaction yourself.
func NewSetCodeTx(account flon.AccountName, wasmPath, abiPath string) (out *flon.Transaction, err error) {
	actions, err := NewSetContract(account, wasmPath, abiPath)
	if err != nil {
		return nil, err
	}
	return &flon.Transaction{Actions: actions}, nil
}

// SetCode represents the hard-coded `setcode` action.
type SetCode struct {
	Account   flon.AccountName `json:"account"`
	VMType    byte             `json:"vmtype"`
	VMVersion byte             `json:"vmversion"`
	Code      flon.HexBytes    `json:"code"`
}

// SetABI represents the hard-coded `setabi` action.
type SetABI struct {
	Account flon.AccountName `json:"account"`
	ABI     flon.HexBytes    `json:"abi"`
}
