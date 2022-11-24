package sui

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

const (
	devNetFaucetURL  = "https://faucet.devnet.sui.io/gas"
	testNetFaucetURL = "https://faucet.testnet.sui.io/gas"
)

func Faucet(addr string) (err error) {
	type faucetReq struct {
		FixedAmountRequest struct {
			Recipient string `json:"recipient"`
		} `json:"FixedAmountRequest"`
	}
	req := &faucetReq{}
	req.FixedAmountRequest.Recipient = addr
	body, err := json.Marshal(req)
	if err != nil {
		return
	}
	faucetURL := devNetFaucetURL
	if Endpoint == Testnet {
		faucetURL = testNetFaucetURL
	}
	rsp, err := http.Post(faucetURL, "application/json", bytes.NewReader(body))
	rspBody, _ := ioutil.ReadAll(rsp.Body)
	fmt.Printf("response:%s\n", rspBody)
	return err

}
