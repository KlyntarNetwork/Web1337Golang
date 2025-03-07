package tests

import (
	"fmt"
	"testing"

	web1337 "github.com/KlyntarNetwork/Web1337Golang"
)

func TestGetTransactionReceipt(t *testing.T) {
	myOptions := web1337.Options{
		NodeURL:         "http://localhost:7332",
		ChainID:         "aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa",
		WorkflowVersion: 1,
	}

	sdkHandler, _ := web1337.NewWeb1337(myOptions)

	txID := "2aaa516ba5739a7db9c5544483709be51c96b905a48cbd37e59073d4d5ee403b"

	txReceipt, err := sdkHandler.GetTransactionReceipt(txID)
	if err != nil {
		t.Fatalf("Error fetching transaction receipt: %v", err)
	} else {
		fmt.Println("Result: ", string(txReceipt))
	}
}

func TestGetTransactionsList(t *testing.T) {
	myOptions := web1337.Options{
		NodeURL:         "http://localhost:7332",
		ChainID:         "aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa",
		WorkflowVersion: 1,
	}

	sdkHandler, _ := web1337.NewWeb1337(myOptions)

	accountID := "4218fb0aaace62c4bfafbdd9adb05b99a9bf1a33eeae074215a51cb644b9a85c"

	txList, err := sdkHandler.GetTransactionsWithAccount(accountID)
	if err != nil {
		t.Fatalf("Error fetching transactions list: %v", err)
	} else {
		fmt.Println("Result: ", string(txList))
	}
}
