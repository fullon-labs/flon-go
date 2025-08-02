package main

import (
	"context"
	"encoding/hex"
	"fmt"
	flon "github.com/fullon-labs/flon-go"
	"github.com/fullon-labs/flon-go/token"
)

const (
	api           = "https://test-chain.ambt.art" // 节点api
	myAccountName = ""
	myPrivateKey  = ""
	toAccountName = ""
)

func main() {
	client := flon.New(api)
	ctx := context.Background()

	keyBag := &flon.KeyBag{}
	err := keyBag.ImportAmaxPrivateKey(ctx, myPrivateKey)
	if err != nil {
		panic(fmt.Errorf("import private key: %w", err))
	}
	client.SetSigner(keyBag)
	// 是否开启 api debug
	client.Debug = true

	from := flon.AccountName(myAccountName)
	to := flon.AccountName(toAccountName)

	symbol := flon.Symbol{
		Precision: 8,
		Symbol:    "AMAX",
	}

	amounts := fmt.Sprintf("%v AMAX", 0.001)
	quantity, err := flon.NewFixedSymbolAssetFromString(symbol, amounts)
	memo := "test transfer"

	fmt.Printf("quantity: %#v\n", quantity)

	if err != nil {
		panic(fmt.Errorf("invalid quantity: %w", err))
	}

	txOpts := &flon.TxOptions{}
	if err := txOpts.FillFromChain(ctx, client); err != nil {
		panic(fmt.Errorf("filling tx opts: %w", err))
	}

	// 其他action，类似 NewTransfer
	// 也可以直接使用 SignPushActions 方法，sign & push
	tx := flon.NewTransaction([]*flon.Action{token.NewTransfer(from, to, quantity, memo)}, txOpts)

	_, packedTrx, err := client.SignTransaction(ctx, tx, txOpts.ChainID, flon.CompressionNone)
	if err != nil {
		panic(fmt.Errorf("sign transaction: %w", err))
	}

	response, err := client.PushTransaction(ctx, packedTrx)
	if err != nil {
		panic(fmt.Errorf("push transaction: %w", err))
	}

	fmt.Printf("flon-push transaction, response====: %v\n", response)

	fmt.Printf("Transaction [%s] submitted to the network succesfully.\n", hex.EncodeToString(response.Processed.ID))
}
