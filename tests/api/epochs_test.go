package tests

import (
	"fmt"
	"testing"

	web1337 "github.com/KlyntarNetwork/Web1337Golang"
)

func TestGetCurrentEpochOnThreads(t *testing.T) {
	myOptions := web1337.Options{
		NodeURL:         "http://localhost:7332",
		ChainID:         "aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa",
		WorkflowVersion: 1,
	}

	sdkHandler, _ := web1337.NewWeb1337(myOptions)

	threadID := "vt"

	epochData, err := sdkHandler.GetCurrentEpochOnThreads(threadID)
	if err != nil {
		t.Fatalf("Error fetching current epoch on threads: %v", err)
	}

	fmt.Println("Result: ", string(epochData))
}

func TestGetCurrentLeader(t *testing.T) {
	myOptions := web1337.Options{
		NodeURL:         "http://localhost:7332",
		ChainID:         "aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa",
		WorkflowVersion: 1,
	}

	sdkHandler, _ := web1337.NewWeb1337(myOptions)

	leader, err := sdkHandler.GetCurrentLeader()
	if err != nil {
		t.Fatalf("Error fetching current leader: %v", err)
	}

	fmt.Println("Result: ", string(leader))
}

func TestGetEpochByIndex(t *testing.T) {
	myOptions := web1337.Options{
		NodeURL:         "http://localhost:7332",
		ChainID:         "aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa",
		WorkflowVersion: 1,
	}

	sdkHandler, _ := web1337.NewWeb1337(myOptions)

	epochIndex := uint(0)

	epochData, err := sdkHandler.GetEpochByIndex(epochIndex)
	if err != nil {
		t.Fatalf("Error fetching epoch by index: %v", err)
	}

	fmt.Println("Result: ", string(epochData))
}

func TestGetTotalBlocksAndTxsByEpochIndex(t *testing.T) {
	myOptions := web1337.Options{
		NodeURL:         "http://localhost:7332",
		ChainID:         "aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa",
		WorkflowVersion: 1,
	}

	sdkHandler, _ := web1337.NewWeb1337(myOptions)

	epochIndex := uint(0)

	blockAndTxStats, err := sdkHandler.GetVerificationThreadStatsPerEpoch(epochIndex)
	if err != nil {
		t.Fatalf("Error fetching total blocks and txs by epoch index: %v", err)
	}

	fmt.Println("Result: ", string(blockAndTxStats))
}

func TestGetHistoricalStatsPerEpoch(t *testing.T) {
	myOptions := web1337.Options{
		NodeURL:         "http://localhost:7332",
		ChainID:         "aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa",
		WorkflowVersion: 1,
	}

	sdkHandler, _ := web1337.NewWeb1337(myOptions)

	startIndex := uint(0)
	limit := uint(10)

	historicalStats, err := sdkHandler.GetHistoricalStatsPerEpoch(startIndex, limit)
	if err != nil {
		t.Fatalf("Error fetching historical stats per epoch: %v", err)
	}

	fmt.Println("Result: ", string(historicalStats))
}
