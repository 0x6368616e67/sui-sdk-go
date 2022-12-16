# sui-sdk-go

[![Go Build status](https://github.com/0x6368616e67/sui-sdk-go/actions/workflows/build.yml/badge.svg?branch=main)](https://github.com/0x6368616e67/sui-sdk-go/actions/workflows/build.yml)[![Test status](https://github.com/0x6368616e67/sui-sdk-go/actions/workflows/ci.yml/badge.svg?branch=main)](https://github.com/0x6368616e67/sui-sdk-go/actions/workflows/ci.yml) [![SDK Documentation](https://img.shields.io/badge/SDK-Documentation-blue)](https://pkg.go.dev/github.com/0x6368616e67/sui-sdk-go) [![MIT License](https://img.shields.io/badge/license-MIT-blue.svg)](https://github.com/0x6368616e67/sui-sdk-go/blob/main/LICENSE)

`sui-sdk-go` is a golang sdk for [Sui](https://sui.io/). Which contains 
all [RPC API](https://docs.sui.io/sui-jsonrpc) with `Client` and some operation 
for `Account` object such as `Transfer`, `Balance` etc...


## Getting started

Add SDK Dependencies

    $ go get -u  github.com/0x6368616e67/sui-sdk-go

A "Alice and Bob" example  demonstrate two user Alice and Bob.
each of them faucet 0.05 coin first. and then Alice transfer 
0.01 to Bob.

Here is the code

    package main

    import (
        "fmt"
        "time"

        "github.com/0x6368616e67/sui-sdk-go"
    )

    func main() {
        alice := sui.NewAccount()
        bob := sui.NewAccount()
        fmt.Printf("faucet first \n")
        err := sui.Faucet(alice.Address().String())
        if err != nil {
            panic(err.Error())
        }
        err = sui.Faucet(bob.Address().String())
        if err != nil {
            panic(err.Error())
        }

        cli, err := sui.Dial(sui.Devnet)
        if err != nil {
            panic(err.Error())
        }

        fmt.Printf("wait 3 second ...\n")
        time.Sleep(10 * time.Second) // wait for stat
        aliceBalance, err := alice.Balance(cli)
        if err != nil {
            panic(err.Error())
        }
        fmt.Printf("Alice balance:%f\n", aliceBalance)

        bobBalance, err := bob.Balance(cli)
        if err != nil {
            panic(err.Error())
        }
        fmt.Printf("Bob balance:%f\n", bobBalance)

        hash, err := alice.Transfer(cli, bob.Address(), 0.01)
        if err != nil {
            panic(err.Error())
        }
        fmt.Printf("Alice transfer %f SUI to bob with hash:%s\n", 0.01, hash)
        fmt.Printf("====================================================\n")
        fmt.Printf("wait 10 second ...\n")
        time.Sleep(10 * time.Second) // wait for stat
        aliceBalance, err = alice.Balance(cli)
        if err != nil {
            panic(err.Error())
        }
        fmt.Printf("Alice balance:%f\n", aliceBalance)

        bobBalance, err = bob.Balance(cli)
        if err != nil {
            panic(err.Error())
        }
        fmt.Printf("Bob balance:%f\n", bobBalance)

    }

Then in the explorer , we can see what happend .

## Support RPC Status -> v0.19.0

- [ ] sui_batchTransaction
- [ ] sui_dryRunTransaction
- [x] sui_executeTransaction
- [x] sui_executeTransactionSerializedSig
- [x] sui_getCoinMetadata
- [x] sui_getCommitteeInfo
- [ ] sui_getEvents
- [x] sui_getMoveFunctionArgTypes
- [x] sui_getNormalizedMoveFunction
- [x] sui_getNormalizedMoveModule
- [x] sui_getNormalizedMoveModulesByPackage
- [x] sui_getNormalizedMoveStruct
- [x] sui_getObject
- [x] sui_getObjectsOwnedByAddress
- [x] sui_getObjectsOwnedByObject
- [x] sui_getRawObject
- [x] sui_getTotalTransactionNumber
- [x] sui_getTransaction
- [x] sui_getTransactionAuthSigners
- [x] sui_getTransactions
- [x] sui_getTransactionsInRange
- [x] sui_mergeCoins
- [x] sui_moveCall
- [x] sui_pay
- [x] sui_payAllSui
- [x] sui_paySui
- [x] sui_publish
- [x] sui_splitCoin
- [x] sui_splitCoinEqual
- [ ] sui_subscribeEvent
- [x] sui_transferObject
- [x] sui_transferSui
- [ ] sui_tryGetPastObject