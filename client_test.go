package sui

import (
	"context"
	"testing"

	"github.com/0x6368616e67/sui-sdk-go/types"
	"github.com/stretchr/testify/assert"
)

func TestGetTotalTransactionNumber(t *testing.T) {
	cli, err := Dial(Devnet)
	assert.Equal(t, err, nil)
	num, err := cli.GetTotalTransactionNumber(context.Background())
	assert.Equal(t, err, nil)
	t.Log(num)
	assert.Greater(t, num, uint64(0))
}

func TestGetEvents(t *testing.T) {
	cli, err := Dial(Devnet)
	assert.Equal(t, err, nil)
	query := &types.TransactionEventQuery{
		Transaction: "8vaJIJ49U/Ngohpp0ARqGHhfoBOaA/FC6wGJHsjdiLI="}
	events, err := cli.GetEvents(context.Background(), query, nil, 1, false)
	assert.Equal(t, err, nil)
	t.Log(events)
}
