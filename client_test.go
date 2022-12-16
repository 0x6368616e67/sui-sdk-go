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
// 	cli, err := Dial(Devnet)
// 	assert.Equal(t, err, nil)
// 	address := types.HexToAddress("0xc4173a804406a365e69dfb297d4eaaf002546ebd")
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

// func TestGetCommitteeInfo(t *testing.T) {
// 	cli, err := Dial(Devnet)
// 	assert.Equal(t, err, nil)
// 	info, err := cli.GetCommitteeInfo(context.Background(), nil)
// 	assert.Equal(t, err, nil)
// 	fmt.Printf("info:%+v \n", info)
// }

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

func TestGetGormalizedMoveModule(t *testing.T) {
	cli, err := Dial(Devnet)
	assert.Equal(t, err, nil)
	info, err := cli.GetGormalizedMoveModule(context.Background(), types.ObjectID("0x0000000000000000000000000000000000000002"), "devnet_nft")
	assert.Equal(t, err, nil)
	fmt.Printf("info:%+v \n", info)
	for _, s := range info.Structs {
		fmt.Printf("struct:%+v \n", s)
	}
}

func TestGetNormalizedMoveModulesByPackage(t *testing.T) {
	cli, err := Dial(Devnet)
	assert.Equal(t, err, nil)
	modules, err := cli.GetNormalizedMoveModulesByPackage(context.Background(), types.ObjectID("0x2"))
	assert.Equal(t, err, nil)
	fmt.Printf("modules:%+v \n", modules)
	for k, m := range modules {
		fmt.Printf("%s:%+v \n", k, m)
	}
}

func TestGetNormalizedMoveStruct(t *testing.T) {
	cli, err := Dial(Devnet)
	assert.Equal(t, err, nil)
	info, err := cli.GetNormalizedMoveStruct(context.Background(), types.ObjectID("0x2"), "bag", "Bag")
	assert.Equal(t, err, nil)
	fmt.Printf("info:%+v \n", info)
}

func TestGetObject(t *testing.T) {
	cli, err := Dial(Devnet)
	assert.Equal(t, err, nil)
	obj, err := cli.GetObject(context.Background(), types.ObjectID("0x35ec8a32a1d6ca884ba8ffb6dbcfda222533661e"))
	assert.Equal(t, err, nil)
	fmt.Printf("obj:%+v \n", obj)
}

func TestGetObjectsOwnedByObject(t *testing.T) {
	cli, err := Dial(Devnet)
	assert.Equal(t, err, nil)
	objects, err := cli.GetObjectsOwnedByObject(context.Background(), "0x21dc8124f8c6a79871957860d88adf16b5813713")
	assert.Equal(t, err, nil)
	fmt.Printf("objects:%+v \n", objects)
	fmt.Printf("objects count:%d \n", len(objects))
}

func TestGetRawObject(t *testing.T) {
	cli, err := Dial(Devnet)
	assert.Equal(t, err, nil)
	obj, err := cli.GetRawObject(context.Background(), types.ObjectID("0x35ec8a32a1d6ca884ba8ffb6dbcfda222533661e"))
	assert.Equal(t, err, nil)
	fmt.Printf("obj:%+v \n", obj)
}

func TestGetTransaction(t *testing.T) {
	cli, err := Dial(Devnet)
	assert.Equal(t, err, nil)
	tx, err := cli.GetTransaction(context.Background(), types.TransactionDigest("EwUHxDSiJeKKbzvhfQgdLcehhNRJSUKLYKzqaVCw9Ltf"))
	assert.Equal(t, err, nil)
	fmt.Printf("tx:%+v \n", tx)
}

func TestGetTransactionAuthSigners(t *testing.T) {
	cli, err := Dial(Devnet)
	assert.Equal(t, err, nil)
	s, err := cli.GetTransactionAuthSigners(context.Background(), types.TransactionDigest("EwUHxDSiJeKKbzvhfQgdLcehhNRJSUKLYKzqaVCw9Ltf"))
	assert.Equal(t, err, nil)
	fmt.Printf("signers:%+v \n", s)
}

func TestGetTransactions(t *testing.T) {
	cli, err := Dial(Devnet)
	assert.Equal(t, err, nil)
	tx, err := cli.GetTransactions(context.Background(), types.TransactionQuery{All: types.TransactionQueryAll}, types.TransactionDigest("EwUHxDSiJeKKbzvhfQgdLcehhNRJSUKLYKzqaVCw9Ltf"), 2, true)
	assert.Equal(t, err, nil)
	fmt.Printf("tx:%+v \n", tx)
}

func TestGetTransactionsInRange(t *testing.T) {
	cli, err := Dial(Devnet)
	assert.Equal(t, err, nil)
	tx, err := cli.GetTransactionsInRange(context.Background(), 12298242, 12298267)
	assert.Equal(t, err, nil)
	fmt.Printf("tx:%+v \n", tx)
}

func TestMergeCoins(t *testing.T) {
	cli, err := Dial(Devnet)
	assert.Equal(t, err, nil)
	bytes, err := cli.MergeCoins(context.Background(), types.HexToAddress("0xd4cce533669ffa336547a6f33a904dcdc7834770"), types.ObjectID("0xe28444588014cd9a7c060bb0e9722628213ccaf4"), types.ObjectID("0x76a815d4033f9b6f6e48626e29b9386331eacf8a"), types.ObjectID("0x87bf662022eedca4aa213fe90565f221885a6865"), 100)
	assert.Equal(t, err, nil)
	fmt.Printf("bytes:%+v \n", bytes.Bytes)
}
