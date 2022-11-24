package sui

import (
	"context"

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

func (c *Client) GetTotalTransactionNumber(ctx context.Context) (number uint64, err error) {
	err = c.c.CallContext(ctx, &number, "sui_getTotalTransactionNumber")
	if err != nil {
		err = ErrNumber
	}
	return
}
