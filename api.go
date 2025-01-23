package web1337

import (
	"strconv"
)

/*
Blocks
*/
func (sdk *Web1337) GetBlockByBlockID(blockID string) ([]byte, error) {
	return sdk.getRequest("/block/" + blockID)
}

func (sdk *Web1337) GetBlockBySID(index uint) ([]byte, error) {
	indexInShardStr := strconv.FormatUint(uint64(index), 10)

	return sdk.getRequest("/block_by_sid/" + indexInShardStr)
}

func (sdk *Web1337) GetLatestNBlocks(startIndex uint, limit uint) ([]byte, error) {
	startIndexStr := strconv.FormatUint(uint64(startIndex), 10)
	limitStr := strconv.FormatUint(uint64(limit), 10)

	return sdk.getRequest("/latest_n_blocks/" + startIndexStr + "/" + limitStr)
}

func (sdk *Web1337) GetVerificationThreadStats() ([]byte, error) {
	return sdk.getRequest("/verification_thread_stats")
}

/*
Epochs
*/
func (sdk *Web1337) GetCurrentEpochOnThreads(threadID string) ([]byte, error) {
	return sdk.getRequest("/current_epoch/" + threadID)
}

func (sdk *Web1337) GetCurrentLeader() ([]byte, error) {
	return sdk.getRequest("/current_leader")
}

func (sdk *Web1337) GetEpochByIndex(epochIndex uint) ([]byte, error) {
	epochIndexStr := strconv.FormatUint(uint64(epochIndex), 10)

	return sdk.getRequest("/epoch_by_index/" + epochIndexStr)
}

func (sdk *Web1337) GetVerificationThreadStatsPerEpoch(epochIndex uint) ([]byte, error) {
	epochIndexStr := strconv.FormatUint(uint64(epochIndex), 10)

	return sdk.getRequest("/verification_thread_stats_per_epoch/" + epochIndexStr)
}

func (sdk *Web1337) GetHistoricalStatsPerEpoch(startIndex uint, limit uint) ([]byte, error) {
	startIndexStr := strconv.FormatUint(uint64(startIndex), 10)
	limitStr := strconv.FormatUint(uint64(limit), 10)

	return sdk.getRequest("/historical_stats_per_epoch/" + startIndexStr + "/" + limitStr)
}

/*
Transactions
*/
func (sdk *Web1337) GetTransactionReceipt(txID string) ([]byte, error) {
	return sdk.getRequest("/tx_receipt/" + txID)
}

func (sdk *Web1337) GetTransactionsWithAccount(accountID string) ([]byte, error) {
	return sdk.getRequest("/txs_list/" + accountID)
}

/*
State
*/
func (sdk *Web1337) GetDataFromState(cellID string) ([]byte, error) {
	return sdk.getRequest("/state/" + cellID)
}

func (sdk *Web1337) GetPoolStats(poolID string) ([]byte, error) {
	return sdk.getRequest("/pool_stats/" + poolID)
}

func (sdk *Web1337) GetAccountFromState(accountID string) ([]byte, error) {
	return sdk.getRequest("/account/" + accountID)
}

/*
Consensus related
*/
func (sdk *Web1337) GetAggregatedFinalizationProof(blockID string) ([]byte, error) {
	return sdk.getRequest("/aggregated_finalization_proof/" + blockID)
}

func (sdk *Web1337) GetAggregatedEpochFinalizationProof(epochIndex uint) ([]byte, error) {
	epochIndexStr := strconv.FormatUint(uint64(epochIndex), 10)
	return sdk.getRequest("/aggregated_epoch_finalization_proof/" + epochIndexStr)
}

/*
Misc
*/
func (sdk *Web1337) GetInfrastructureInfo() ([]byte, error) {
	return sdk.getRequest("/infrastructure_info")
}

func (sdk *Web1337) GetChainInfo() ([]byte, error) {
	return sdk.getRequest("/chain_info")
}

func (sdk *Web1337) GetKlyEVMMetadata() ([]byte, error) {
	return sdk.getRequest("/kly_evm_metadata")
}

func (sdk *Web1337) GetSynchronizationStats() ([]byte, error) {
	return sdk.getRequest("/synchronization_stats")
}

func (sdk *Web1337) GetCheckpointByEpochIndex(epochIndex uint) ([]byte, error) {
	epochIndexStr := strconv.FormatUint(uint64(epochIndex), 10)
	return sdk.getRequest("/checkpoints/" + epochIndexStr)
}

func (sdk *Web1337) GetQuorumUrlsAndPubkeys() ([]byte, error) {
	return sdk.getRequest("/quorum_urls_and_pubkeys")
}
