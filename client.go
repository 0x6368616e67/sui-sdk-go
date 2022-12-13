package sui

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/0x6368616e67/sui-sdk-go/types"
	"github.com/ethereum/go-ethereum/rpc"
)

// Client defines typed wrappers for the Sui RPC API.
type Client struct {
	c *rpc.Client
}

// Dial connects a client to the given URL.
func Dial(rawurl string) (*Client, error) {
	return DialContext(context.Background(), rawurl)
}

func DialContext(ctx context.Context, rawurl string) (*Client, error) {
	c, err := rpc.DialContext(ctx, rawurl)
	if err != nil {
		return nil, err
	}
	return NewClient(c), nil
}

// NewClient creates a client that uses the given RPC client.
func NewClient(c *rpc.Client) *Client {
	return &Client{c}
}

// Close close a client's connection
func (c *Client) Close() {
	c.c.Close()
}

func (c *Client) GetEvents(ctx context.Context, query types.EventQuery, cursor *types.EventID, limit uint64, descending bool) (event *types.PaginatedEvents, err error) {
	event = &types.PaginatedEvents{}
	err = c.c.CallContext(ctx, event, "sui_getEvents", query, cursor, limit, descending)
	if err != nil {
		fmt.Printf("err:%s\n", err.Error())
		err = ErrNumber
	}
	return
}

func (c *Client) ExecuteTransaction(ctx context.Context, txBytes string, signatureScheme SignatureScheme, signature string, pubkey types.PubKey, requestType ExecuteTransactionRequestType) (response *types.SuiExecuteTransactionResponse, err error) {
	response = &types.SuiExecuteTransactionResponse{}
	err = c.c.CallContext(ctx, response, "sui_executeTransaction", txBytes, signatureScheme, signature, pubkey.Base64(), requestType)
	if err != nil {
		fmt.Printf("err:%s\n", err.Error())
		err = ErrNumber
	}
	return
}

///
/// here
///

// GetCoinMetadata return metadata(e.g., symbol, decimals) for a coin
func (c *Client) GetCoinMetadata(ctx context.Context, coinType string) (metadata types.SuiCoinMetadata, err error) {
	err = c.c.CallContext(ctx, &metadata, "sui_getCoinMetadata", coinType)
	return
}

// GetCommitteeInfo return the committee information for the asked epoch
func (c *Client) GetCommitteeInfo(ctx context.Context, epoch *uint64) (info types.CommitteeInfo, err error) {
	err = c.c.CallContext(ctx, &info, "sui_getCommitteeInfo", epoch)
	return
}

// GetTotalTransactionNumber return the total number of transactions known to the server
func (c *Client) GetTotalTransactionNumber(ctx context.Context) (number uint64, err error) {
	err = c.c.CallContext(ctx, &number, "sui_getTotalTransactionNumber")
	if err != nil {
		err = ErrNumber
	}
	return
}

// GetMoveFunctionArgTypes return the argument types of a Move function, based on normalized Type
func (c *Client) GetMoveFunctionArgTypes(ctx context.Context, packagee types.ObjectID, module string, function string) (argTypes types.MoveFunctionArgTypes, err error) {
	err = c.c.CallContext(ctx, &argTypes, "sui_getMoveFunctionArgTypes", packagee, module, function)
	return
}

// GetNormalizedMoveFunction return a structured representation of Move function
func (c *Client) GetNormalizedMoveFunction(ctx context.Context, packagee types.ObjectID, module string, function string) (functionInfo types.SuiMoveNormalizedFunction, err error) {
	err = c.c.CallContext(ctx, &functionInfo, "sui_getNormalizedMoveFunction", packagee, module, function)
	return
}

// GetGormalizedMoveModule return a structured representation of Move module
func (c *Client) GetGormalizedMoveModule(ctx context.Context, packagee types.ObjectID, module string) (moduleInfo types.SuiMoveNormalizedModule, err error) {
	err = c.c.CallContext(ctx, &moduleInfo, "sui_getNormalizedMoveModule", packagee, module)
	return
}

// GetNormalizedMoveModulesByPackage return structured representations of all modules in the given package
func (c *Client) GetNormalizedMoveModulesByPackage(ctx context.Context, packagee types.ObjectID) (modules map[string]*types.SuiMoveNormalizedModule, err error) {
	modules = make(map[string]*types.SuiMoveNormalizedModule)
	err = c.c.CallContext(ctx, &modules, "sui_getNormalizedMoveModulesByPackage", packagee)
	return
}

// GetNormalizedMoveStruct return a structured representation of Move struct
func (c *Client) GetNormalizedMoveStruct(ctx context.Context, packagee types.ObjectID, module string, structt string) (structInfo types.SuiMoveNormalizedStruct, err error) {
	err = c.c.CallContext(ctx, &structInfo, "sui_getNormalizedMoveStruct", packagee, module, structt)
	return
}

// GetObject the object information for a specified object
func (c *Client) GetObject(ctx context.Context, objectID types.ObjectID) (objectData types.SuiObjectData, err error) {
	err = c.c.CallContext(ctx, &objectData, "sui_getObject", objectID)
	return
}

// GetObjectsOwnedByAddress return the list of objects owned by an address
func (c *Client) GetObjectsOwnedByAddress(ctx context.Context, address types.Address) (object []*types.SuiObjectInfo, err error) {
	object = make([]*types.SuiObjectInfo, 0)
	err = c.c.CallContext(ctx, &object, "sui_getObjectsOwnedByAddress", address.String())
	return
}

// GetObjectsOwnedByObject return  the list of objects owned by an object
func (c *Client) GetObjectsOwnedByObject(ctx context.Context, objectID types.ObjectID) (object []*types.SuiObjectInfo, err error) {
	object = make([]*types.SuiObjectInfo, 0)
	err = c.c.CallContext(ctx, &object, "sui_getObjectsOwnedByObject", objectID)
	return
}

// GetRawObject return the raw BCS serialized move object bytes for a specified object
func (c *Client) GetRawObject(ctx context.Context, objectID types.ObjectID) (object *types.SuiObjectData, err error) {
	err = c.c.CallContext(ctx, &object, "sui_getRawObject", objectID)
	return
}

// GetTransaction return  the transaction response object
func (c *Client) GetTransaction(ctx context.Context, transaction types.TransactionDigest) (response *types.SuiTransactionResponse, err error) {
	err = c.c.CallContext(ctx, &response, "sui_getTransaction", transaction)
	return
}

// GetTransactionAuthSigners return the authority public keys that commits to the authority signature of the transaction
func (c *Client) GetTransactionAuthSigners(ctx context.Context, transaction types.TransactionDigest) (response *types.SuiTransactionAuthSignersResponse, err error) {
	err = c.c.CallContext(ctx, &response, "sui_getTransactionAuthSigners", transaction)
	return
}

// GetTransactions return list of transactions for a specified query criteria
func (c *Client) GetTransactions(ctx context.Context, query types.TransactionQuery, cursor types.TransactionDigest, limit int64, descending bool) (digests *types.PaginatedTransactionDigests, err error) {
	err = c.c.CallContext(ctx, &digests, "sui_getTransactions", query, cursor, limit, descending)
	return
}

// GetTransactionsInRange return  list of transaction digests within the queried range
func (c *Client) GetTransactionsInRange(ctx context.Context, start int64, end int64) (digests []types.TransactionDigest, err error) {
	digests = make([]types.TransactionDigest, 1)
	err = c.c.CallContext(ctx, &digests, "sui_getTransactionsInRange", start, end)
	return
}

// MergeCoins create an unsigned transaction to merge multiple coins into one coin
func (c *Client) MergeCoins(ctx context.Context, signer types.Address, primaryCoin types.ObjectID, coinToMerge types.ObjectID, gas types.ObjectID, gasBudget uint64) (bytes types.TransactionBytes, err error) {
	err = c.c.CallContext(ctx, &bytes, "sui_mergeCoins", signer.String(), primaryCoin, coinToMerge, gas, gasBudget)
	return
}

// MoveCall create an unsigned transaction to execute a Move call on the network, by calling the specified function in the module of a given package
func (c *Client) MoveCall(ctx context.Context, signer types.Address, packageObjectID types.ObjectID, module string, function string, typeArgs []string, args []json.RawMessage, gas types.ObjectID, gasBudget uint64) (bytes types.TransactionBytes, err error) {
	err = c.c.CallContext(ctx, &bytes, "sui_moveCall", signer.String(), packageObjectID, module, function, typeArgs, args, gas, gasBudget)
	return
}

// Pay send Coin<T> to a list of addresses, where `T` can be any coin type, following a list of amounts, The object specified in the `gas` field will be used to pay the gas fee for the transaction. The gas object can not appear in `input_coins`. If the gas object is not specified, the RPC server will auto-select one
func (c *Client) Pay(ctx context.Context, signer types.Address, inputCoins []string, recipients []string, amounts []int64, gas types.ObjectID, gasBudget uint64) (bytes types.TransactionBytes, err error) {
	err = c.c.CallContext(ctx, &bytes, "sui_pay", signer.String(), inputCoins, recipients, amounts, gas, gasBudget)
	return
}

// PayAllSui send all SUI coins to one recipient. This is for SUI coin only and does not require a separate gas coin object. Specifically,
// what pay_all_sui does are:
// 1. accumulate all SUI from input coins and deposit all SUI to the first input coin
// 2. transfer the updated first coin to the recipient and also use this first coin as gas coin object.
// 3. the balance of the first input coin after tx is sum(input_coins) - actual_gas_cost.
// 4. all other input coins other than the first are deleted
func (c *Client) PayAllSui(ctx context.Context, signer types.Address, coins []types.ObjectID, recipient types.Address, gasBudget uint64) (bytes types.TransactionBytes, err error) {
	err = c.c.CallContext(ctx, &bytes, "sui_payAllSui", signer.String(), coins, recipient.String(), gasBudget)
	return
}

// PaySui send SUI coins to a list of addresses, following a list of amounts.
// This is for SUI coin only and does not require a separate gas coin object. Specifically, what pay_sui does are:
// 1. debit each input_coin to create new coin following the order of amounts and assign it to the corresponding recipient.
// 2. accumulate all residual SUI from input coins left and deposit all SUI to the first input coin, then use the first input coin as the gas coin object.
// 3. the balance of the first input coin after tx is sum(input_coins) - sum(amounts) - actual_gas_cost
// 4. all other input coints other than the first one are deleted.
func (c *Client) PaySui(ctx context.Context, signer types.Address, inputCoins []string, recipients []string, amounts []int64, gasBudget uint64) (bytes types.TransactionBytes, err error) {
	err = c.c.CallContext(ctx, &bytes, "sui_paySui", signer.String(), inputCoins, recipients, amounts, gasBudget)
	return
}

// Publish create an unsigned transaction to publish Move module
func (c *Client) Publish(ctx context.Context, signer types.Address, compiledModules []string, gas types.ObjectID, gasBudget uint64) (bytes types.TransactionBytes, err error) {
	err = c.c.CallContext(ctx, &bytes, "sui_publish", signer.String(), compiledModules, gas, gasBudget)
	return
}

// SplitCoin create an unsigned transaction to split a coin object into multiple coins
func (c *Client) SplitCoin(ctx context.Context, signer types.Address, coin types.ObjectID, splitAmounts []int64, gas types.ObjectID, gasBudget uint64) (bytes types.TransactionBytes, err error) {
	err = c.c.CallContext(ctx, &bytes, "sui_splitCoin", signer.String(), coin, splitAmounts, gas, gasBudget)
	return
}

// SplitCoinEqual create an unsigned transaction to split a coin object into multiple equal-size coins
func (c *Client) SplitCoinEqual(ctx context.Context, signer types.Address, coin types.ObjectID, splitCount int64, gas types.ObjectID, gasBudget uint64) (bytes types.TransactionBytes, err error) {
	err = c.c.CallContext(ctx, &bytes, "sui_splitCoinEqual", signer.String(), coin, splitCount, gas, gasBudget)
	return
}

// TransferObject create an unsigned transaction to transfer an object from one address to another. The object's type must allow public transfers
func (c *Client) TransferObject(ctx context.Context, signer types.Address, objectID types.ObjectID, recipient types.Address, gas types.ObjectID, gasBudget uint64) (bytes types.TransactionBytes, err error) {
	err = c.c.CallContext(ctx, &bytes, "sui_transferObject", signer.String(), objectID, gas, gasBudget, recipient.String())
	return
}

// TransferSui create an unsigned transaction to send SUI coin object to a Sui address. The SUI object is also used as the gas object
func (c *Client) TransferSui(ctx context.Context, signer types.Address, objectID types.ObjectID, recipient types.Address, amount int64, gasBudget uint64) (bytes types.TransactionBytes, err error) {
	err = c.c.CallContext(ctx, &bytes, "sui_transferSui", signer.String(), objectID, gasBudget, recipient.String(), amount)
	return
}
