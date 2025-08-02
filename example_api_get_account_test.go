package flon_test

import (
	"context"
	"encoding/json"
	"fmt"
)

func ExampleAPI_GetAccount() {
	api := flon.New(getAPIURL())

	account := flon.AccountName("flon.rex")
	info, err := api.GetAccount(context.Background(), account)
	if err != nil {
		if err == flon.ErrNotFound {
			fmt.Printf("unknown account: %s", account)
			return
		}

		panic(fmt.Errorf("get account: %w", err))
	}

	bytes, err := json.Marshal(info)
	if err != nil {
		panic(fmt.Errorf("json marshal response: %w", err))
	}

	fmt.Println(string(bytes))
}
