package sui

import (
	"context"
	"fmt"
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
	assert.Equal(t, len(events.Data), 1)
	//fmt.Printf("events:%+v \n", events)
	fmt.Printf("data:%+v \n", events.Data[0].Event.CoinBalanceChangeEvent)

}

func TestGetObjectsOwnedByAddress(t *testing.T) {
	cli, err := Dial(Testnet)
	assert.Equal(t, err, nil)
	address := types.HexToAddress("")
	objects, err := cli.GetObjectsOwnedByAddress(context.Background(), address)
	assert.Equal(t, err, nil)
	fmt.Printf("objects:%+v \n", objects)
	fmt.Printf("objects count:%d \n", len(objects))
}

func TestPayAllSui(t *testing.T) {
	cli, err := Dial(Testnet)
	assert.Equal(t, err, nil)

	address := types.HexToAddress("")
	objects, err := cli.GetObjectsOwnedByAddress(context.Background(), address)
	assert.Equal(t, err, nil)
	objs := []types.ObjectID{}
	for _, o := range objects {
		objs = append(objs, o.ObjectID)
	}
	signer := types.HexToAddress("")
	txBytes, err := cli.PayAllSui(context.Background(), signer, objs, signer, 2000)
	assert.Equal(t, err, nil)
	fmt.Printf("txBytes:%+v \n", txBytes)
}

func TestExecuteTransaction(t *testing.T) {

}
