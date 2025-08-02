package system

import (
	"encoding/hex"
	"encoding/json"
	"fmt"
	"testing"

	"github.com/fullon-labs/flon-go/ecc"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// TODO: Move this test to the `system` contract.. and take out
// `NewAccount` from this package.
func TestActionNewAccount(t *testing.T) {
	pubKey, err := ecc.NewPublicKey(ecc.PublicKeyPrefixCompat + "6MRyAjQq8ud7hVNYcfnVPJqcVpscN5So8BhtHuGYqET5GDW5CV")
	require.NoError(t, err)
	a := &flon.Action{
		Account: flon.AccountName("eosio"),
		Name:    flon.ActionName("newaccount"),
		Authorization: []flon.PermissionLevel{
			{
				Actor:      flon.AccountName("eosio"),
				Permission: flon.PermissionName("active"),
			},
		},
		ActionData: flon.NewActionData(NewAccount{
			Creator: flon.AccountName("eosio"),
			Name:    flon.AccountName("abourget"),
			Owner: flon.Authority{
				Threshold: 1,
				Keys: []flon.KeyWeight{
					{
						PublicKey: pubKey,
						Weight:    1,
					},
				},
			},
			Active: flon.Authority{
				Threshold: 1,
				Keys: []flon.KeyWeight{
					{
						PublicKey: pubKey,
						Weight:    1,
					},
				},
			},
		}),
	}
	tx := &flon.Transaction{
		Actions: []*flon.Action{a},
	}

	buf, err := flon.MarshalBinary(tx)
	// println(string(buf))
	assert.NoError(t, err)

	assert.Equal(t, `00096e8800000000000000000000010000000000ea305500409e9a2264b89a010000000000ea305500000000a8ed3232660000000000ea305500000059b1abe93101000000010002c0ded2bc1f1305fb0faac5e6c03ee3a1924234985427b6167ca569d13df435cf0100000001000000010002c0ded2bc1f1305fb0faac5e6c03ee3a1924234985427b6167ca569d13df435cf0100000000`, hex.EncodeToString(buf))

	buf, err = json.Marshal(a)
	assert.NoError(t, err)
	assert.Equal(t, `{"account":"eosio","name":"newaccount","authorization":[{"actor":"eosio","permission":"active"}],"data":"0000000000ea305500000059b1abe93101000000010002c0ded2bc1f1305fb0faac5e6c03ee3a1924234985427b6167ca569d13df435cf0100000001000000010002c0ded2bc1f1305fb0faac5e6c03ee3a1924234985427b6167ca569d13df435cf01000000"}`, string(buf))

	buf, err = json.Marshal(a.ActionData.Data)
	assert.NoError(t, err)

	assert.Equal(t, "{\"creator\":\"eosio\",\"name\":\"abourget\",\"owner\":{\"threshold\":1,\"keys\":[{\"key\":\""+ecc.PublicKeyPrefixCompat+"6MRyAjQq8ud7hVNYcfnVPJqcVpscN5So8BhtHuGYqET5GDW5CV\",\"weight\":1}]},\"active\":{\"threshold\":1,\"keys\":[{\"key\":\""+ecc.PublicKeyPrefixCompat+"6MRyAjQq8ud7hVNYcfnVPJqcVpscN5So8BhtHuGYqET5GDW5CV\",\"weight\":1}]}}", string(buf))
	// 00096e88 0000 0000 00000000 00 00 00 00 01 0000000000ea3055

	// WUTz that ?
	// var newAct *Action
	// newAct.DecodeAs(&NewAccount{})
	// require.NoError(t, UnmarshalBinary(buf, &newAct))
	// assert.Equal(t, a, newAct)
}

func TestMarshalTransactionAndSigned(t *testing.T) {
	a := &flon.Action{
		Account: flon.AccountName("eosio"),
		Name:    flon.ActionName("newaccount"),
		Authorization: []flon.PermissionLevel{
			{
				Actor:      flon.AccountName("eosio"),
				Permission: flon.PermissionName("active"),
			},
		},
		ActionData: flon.NewActionData(NewAccount{
			Creator: flon.AccountName("eosio"),
			Name:    flon.AccountName("abourget"),
		}),
	}
	tx := &flon.SignedTransaction{Transaction: &flon.Transaction{
		Actions: []*flon.Action{a},
	}}

	buf, err := flon.MarshalBinary(tx)
	assert.NoError(t, err)
	// 00096e88 0000 0000 00000000 0000 0000 00
	// actions: 01
	// 0000000000ea3055 00409e9a2264b89a 01 0000000000ea3055 00000000a8ed3232
	// len: 22
	// 0000000000ea3055 00000059b1abe931 000000000000000000000000000000000000

	assert.Equal(t, `00096e8800000000000000000000010000000000ea305500409e9a2264b89a010000000000ea305500000000a8ed32321e0000000000ea305500000059b1abe9310000000000000000000000000000000000`, hex.EncodeToString(buf))

	buf, err = json.Marshal(a)
	assert.NoError(t, err)
	assert.Equal(t, `{"account":"eosio","name":"newaccount","authorization":[{"actor":"eosio","permission":"active"}],"data":"0000000000ea305500000059b1abe9310000000000000000000000000000"}`, string(buf))
}

func TestMarshalTransactionAndPack(t *testing.T) {
	a := &flon.Action{
		Account: flon.AccountName("eosio"),
		Name:    flon.ActionName("newaccount"),
		Authorization: []flon.PermissionLevel{
			{
				Actor:      flon.AccountName("eosio"),
				Permission: flon.PermissionName("active"),
			},
		},
		ActionData: flon.NewActionData(NewAccount{
			Creator: flon.AccountName("eosio"),
			Name:    flon.AccountName("abourget"),
		}),
	}
	b := &flon.Action{
		Account: flon.AccountName("eosio"),
		Name:    flon.ActionName("transfer"),
		Authorization: []flon.PermissionLevel{
			{
				Actor:      flon.AccountName("eosio"),
				Permission: flon.PermissionName("active"),
			},
		},
		ActionData: flon.NewActionData(NewAccount{
			Creator: flon.AccountName("eosio"),
			Name:    flon.AccountName("cbillett"),
		}),
	}

	tx := &flon.Transaction{
		Actions: []*flon.Action{a, b},
	}

	buf, err := json.Marshal(tx)
	fmt.Println("Transaction: ", string(buf))

	signedTx := &flon.SignedTransaction{Transaction: tx}
	buf, err = json.Marshal(signedTx)
	fmt.Println("Signed Transaction: ", string(buf))

	packedTx, err := signedTx.Pack(flon.CompressionNone)
	assert.NoError(t, err)

	buf, err = json.Marshal(packedTx)
	assert.NoError(t, err)
	fmt.Println("Pack tx: ", string(buf))
}
