package sui

import (
	"context"
	"encoding/base64"
	"encoding/hex"
	"encoding/json"
	"fmt"
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
					var b int64
					err = json.Unmarshal(rj, &b)
					if err != nil {
						return 0, err
					}
					balance += float64(b) / (1000000000)
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
	var coins []string
	for _, o := range objs {
		obj, err := client.GetObject(context.TODO(), o.ObjectID)
		if err != nil {
			return "", err
		}
		if sobj, ok := obj.Details.(types.SuiObejct); ok {
			if sobj.Data.Type == "0x2::coin::Coin<0x2::sui::SUI>" {
				if rj, ok := sobj.Data.Fields["balance"]; ok {
					type Balance struct {
						Balance int `json:"balance"`
						ID      struct {
							ID string `json:"id"`
						} `json:"id"`
					}

					var b Balance
					err = json.Unmarshal(rj, &b)
					if err != nil {
						return "", err
					}
					aamount += float64(b.Balance) / (1000000000)
					coins = append(coins, b.ID.ID)
				}
			}
		}
		if aamount+float64(20000)/(1000000000) > amount {
			break
		}
	}
	txBytes, err := client.PaySui(context.TODO(), acc.Address(), coins, []string{recipient.String()}, []int64{int64(amount)}, 2000)
	if err != nil {
		return
	}
	tx, err := base64.StdEncoding.DecodeString(txBytes.Bytes)
	if err != nil {
		fmt.Printf("DecodeString error:%s", err.Error())
		return
	}
	sig, err := acc.Sign(tx)
	if err != nil {
		fmt.Printf("Sign error:%s", err.Error())
		return
	}
	sigStr := base64.StdEncoding.EncodeToString(sig)
	rsp, err := client.ExecuteTransaction(context.TODO(), txBytes.Bytes, ED25519, sigStr, acc.privateKey.PubKey(), WaitForTxCert)
	if err != nil {
		return
	}
	hash = string(*rsp.TxDigest)
	return
}
