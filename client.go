package sui

import (
	"context"
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

func (c *Client) GetObjectsOwnedByAddress(ctx context.Context, address types.Address) (object []*types.SuiObject, err error) {
	object = make([]*types.SuiObject, 0)
	err = c.c.CallContext(ctx, &object, "sui_getObjectsOwnedByAddress", address.String())
	if err != nil {
		fmt.Printf("err:%s\n", err.Error())
		err = ErrNumber
	}
	return
}

func (c *Client) PayAllSui(ctx context.Context, signer types.Address, coins []types.ObjectID, recipient types.Address, gasBudget uint64) (txBytes *types.TransactionBytes, err error) {
	txBytes = &types.TransactionBytes{}
	err = c.c.CallContext(ctx, txBytes, "sui_payAllSui", signer.String(), coins, recipient.String(), gasBudget)
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
func (c *Client) GetNormalizedMoveFunction(ctx context.Context, packagee types.ObjectID, module string, function string) (argTypes types.SuiMoveNormalizedFunction, err error) {
	err = c.c.CallContext(ctx, &argTypes, "sui_getNormalizedMoveFunction", packagee, module, function)
	if err != nil {
		fmt.Printf("err:%s", err.Error())
	}
	return
}
