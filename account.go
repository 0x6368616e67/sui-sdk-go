package sui

import (
	"encoding/hex"
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
