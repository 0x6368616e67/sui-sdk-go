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
