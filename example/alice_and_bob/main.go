package main

import (
	"fmt"
	"time"

	"github.com/0x6368616e67/sui-sdk-go"
)

func main() {
	//alice := sui.NewAccount()
	alice := sui.NewAccountWithPrivateKey("0xbc25e1ca875b2e2e3adf967628fb0f4672ccd786586b1cbe964657c799a7d4c5e486221bcaeb3bcb60b974d575d930bee9937fb55ddeea4087440d38f4f57e10")
	fmt.Printf("alice: %v: %v\n", alice.Address(), alice.PrivateKey())
	//bob := sui.NewAccount()
	bob := sui.NewAccountWithPrivateKey("0x0d49cea656f2c067d30cea441df11bcb78aa62c10203dd019cf6aeefdb5bd27c493bf73535b2a3e51d83c127fdd4198b3d509abcb31b78754ee0a4cb2520e3fd")
	fmt.Printf("bob:%v\n", bob.PrivateKey())
	// fmt.Printf("faucet first \n")
	// err := sui.Faucet(alice.Address().String())
	// if err != nil {
	// 	panic(err.Error())
	// }
	// err = sui.Faucet(bob.Address().String())
	// if err != nil {
	// 	panic(err.Error())
	// }

	cli, err := sui.Dial(sui.Devnet)
	if err != nil {
		panic(err.Error())
	}

	fmt.Printf("wait 3 second ...\n")
	//time.Sleep(3 * time.Second) // wait for stat
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

	hash, err := alice.Transfer(cli, bob.Address(), 0.001)
	if err != nil {
		panic(err.Error())
	}
	fmt.Printf("Alice transfer %f SUI to bob with hash:%s\n", 0.0001, hash)
	fmt.Printf("====================================================\n")
	fmt.Printf("wait 3 second ...\n")
	time.Sleep(3 * time.Second) // wait for stat
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
