package sui

import (
	"context"
	"fmt"
	"testing"

	"github.com/0x6368616e67/sui-sdk-go/types"
	"github.com/stretchr/testify/assert"
)

// func TestGetEvents(t *testing.T) {
// 	cli, err := Dial(Devnet)
// 	assert.Equal(t, err, nil)
// 	query := &types.TransactionEventQuery{
// 		Transaction: "8vaJIJ49U/Ngohpp0ARqGHhfoBOaA/FC6wGJHsjdiLI="}
// 	events, err := cli.GetEvents(context.Background(), query, nil, 1, false)
// 	assert.Equal(t, err, nil)
// 	assert.Equal(t, len(events.Data), 1)
// 	//fmt.Printf("events:%+v \n", events)
// 	fmt.Printf("data:%+v \n", events.Data[0].Event.CoinBalanceChangeEvent)
// }

// func TestGetObjectsOwnedByAddress(t *testing.T) {
// 	cli, err := Dial(Testnet)
// 	assert.Equal(t, err, nil)
// 	address := types.HexToAddress("")
// 	objects, err := cli.GetObjectsOwnedByAddress(context.Background(), address)
// 	assert.Equal(t, err, nil)
// 	fmt.Printf("objects:%+v \n", objects)
// 	fmt.Printf("objects count:%d \n", len(objects))
// }

// func TestPayAllSui(t *testing.T) {
// 	cli, err := Dial(Testnet)
// 	assert.Equal(t, err, nil)

// 	address := types.HexToAddress("")
// 	objects, err := cli.GetObjectsOwnedByAddress(context.Background(), address)
// 	assert.Equal(t, err, nil)
// 	objs := []types.ObjectID{}
// 	for _, o := range objects {
// 		objs = append(objs, o.ObjectID)
// 	}
// 	signer := types.HexToAddress("")
// 	txBytes, err := cli.PayAllSui(context.Background(), signer, objs, signer, 2000)
// 	assert.Equal(t, err, nil)
// 	fmt.Printf("txBytes:%+v \n", txBytes)
// }

func TestGetCoinMetadata(t *testing.T) {
	cli, err := Dial(Devnet)
	assert.Equal(t, err, nil)
	metadata, err := cli.GetCoinMetadata(context.Background(), "0x2::sui::SUI")
	assert.Equal(t, err, nil)
	fmt.Printf("metadata:%+v \n", metadata)
}

func TestGetCommitteeInfo(t *testing.T) {
	cli, err := Dial(Devnet)
	assert.Equal(t, err, nil)
	info, err := cli.GetCommitteeInfo(context.Background(), nil)
	assert.Equal(t, err, nil)
	fmt.Printf("info:%+v \n", info)
}

func TestGetTotalTransactionNumber(t *testing.T) {
	cli, err := Dial(Devnet)
	assert.Equal(t, err, nil)
	num, err := cli.GetTotalTransactionNumber(context.Background())
	assert.Equal(t, err, nil)
	t.Log(num)
	assert.Greater(t, num, uint64(0))
}

func TestGetMoveFunctionArgTypes(t *testing.T) {
	cli, err := Dial(Devnet)
	assert.Equal(t, err, nil)
	ats, err := cli.GetMoveFunctionArgTypes(context.Background(), types.ObjectID("0x2"), "sui", "transfer")
	assert.Equal(t, err, nil)
	fmt.Printf("ats:%+v \n", ats)
}

func TestGetNormalizedMoveFunction(t *testing.T) {
	cli, err := Dial(Devnet)
	assert.Equal(t, err, nil)
	funcs, err := cli.GetNormalizedMoveFunction(context.Background(), types.ObjectID("0x2"), "sui", "transfer")
	//funcs, err := cli.GetNormalizedMoveFunction(context.Background(), types.ObjectID("0x0000000000000000000000000000000000000002"), "devnet_nft", "mint")
	assert.Equal(t, err, nil)
	fmt.Printf("funcs:%+v \n", funcs)
	for _, p := range funcs.Parameters {

		fmt.Printf("param:%+v \n", p)
	}
}
