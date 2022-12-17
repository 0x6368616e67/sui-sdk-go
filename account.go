package sui

import (
	"bytes"
	"context"
	"encoding/base64"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"strconv"
	"strings"

	"github.com/0x6368616e67/sui-sdk-go/types"
)

type Account struct {
	privateKey types.PrivKey
}

func NewAccount() *Account {
	acc := &Account{}

	acc.privateKey = types.GenPrivKey()
	return acc
}

func NewAccountWithPrivateKey(key string) *Account {
	key = strings.TrimPrefix(key, "0x")
	acc := &Account{}
	var err error
	if acc.privateKey, err = types.GenPrivKeyFromHex(key); err != nil {
		return nil
	}
	return acc
}

func NewAccountWithHexSeed(seed string) *Account {
	acc := &Account{}
	seedBuf, err := hex.DecodeString(seed)
	if err != nil {
		return nil
	}
	acc.privateKey = types.GenPrivKeyFromSeed(seedBuf)

	return acc
}

func (acc *Account) PrivateKey() string {
	return acc.privateKey.String()
}

func (acc *Account) Sign(msg []byte) ([]byte, error) {
	return acc.privateKey.Sign(msg)
}

func (acc *Account) Address() types.Address {
	return acc.privateKey.PubKey().Address()
}

func (acc *Account) Balance(client *Client) (balance float64, err error) {
	objs, err := client.GetObjectsOwnedByAddress(context.TODO(), acc.Address())
	if err != nil {
		return
	}
	for _, o := range objs {
		obj, err := client.GetObject(context.TODO(), o.ObjectID)
		if err != nil {
			return 0, err
		}
		if sobj, ok := obj.Details.(types.SuiObejct); ok {
			if sobj.Data.Type == "0x2::coin::Coin<0x2::sui::SUI>" {
				if rj, ok := sobj.Data.Fields["balance"]; ok {
					var b string
					err = json.Unmarshal(rj, &b)
					if err != nil {
						return 0, err
					}
					bb, _ := strconv.ParseUint(b, 10, 64)
					balance += float64(bb) / (1000000000)
				}
			}
		}
	}
	err = nil
	return
}

func (acc *Account) Transfer(client *Client, recipient types.Address, amount float64) (hash string, err error) {
	objs, err := client.GetObjectsOwnedByAddress(context.TODO(), acc.Address())
	if err != nil {
		return
	}
	aamount := 0.0
	var coins []types.ObjectID
	for _, o := range objs {
		obj, err := client.GetObject(context.TODO(), o.ObjectID)
		if err != nil {
			return "", err
		}
		if sobj, ok := obj.Details.(types.SuiObejct); ok {
			if sobj.Data.Type == "0x2::coin::Coin<0x2::sui::SUI>" {
				if rj, ok := sobj.Data.Fields["balance"]; ok {
					var b string
					err = json.Unmarshal(rj, &b)
					if err != nil {
						return "", err
					}
					bb, _ := strconv.ParseUint(b, 10, 64)
					aamount += float64(bb) / (1000000000)

					type FieldID struct {
						ID string `json:"id"`
					}
					var fid FieldID
					err = json.Unmarshal(sobj.Data.Fields["id"], &fid)
					if err != nil {
						return "", err
					}

					coins = append(coins, types.ObjectID(fid.ID))
				}
			}
		}
		if aamount+float64(20000)/(1000000000) > amount {
			break
		}
	}
	ramount := int64(amount * 1000000000)
	fmt.Printf("ramount:%d\n", ramount)
	txBytes, err := client.PaySui(context.TODO(), acc.Address(), coins, []string{recipient.String()}, []int64{ramount}, 2000)
	//txBytes, err := client.PayAllSui(context.TODO(), acc.Address(), coins, recipient, 2000)
	if err != nil {
		return
	}
	tx, err := base64.StdEncoding.DecodeString(txBytes.Bytes)
	if err != nil {
		fmt.Printf("DecodeString error:%s", err.Error())
		return
	}
	var sigData bytes.Buffer
	sigData.Write([]byte{0, 0, 0})
	sigData.Write(tx)
	sig, err := acc.Sign(sigData.Bytes())
	if err != nil {
		fmt.Printf("Sign error:%s", err.Error())
		return
	}
	var buf bytes.Buffer
	buf.Write([]byte{0x00})
	buf.Write(sig)
	buf.Write(acc.privateKey.PubKey().Bytes())
	sigStr := base64.StdEncoding.EncodeToString(buf.Bytes())

	rsp, err := client.ExecuteTransactionSerializedSig(context.TODO(), txBytes.Bytes, sigStr, WaitForEffectsCert)
	if err != nil {
		return
	}
	fmt.Printf("rsp:%+v \n", rsp)
	//hash = string(*rsp.TxDigest)
	return
}
