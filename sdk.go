package sui

const (
	Devnet  string = "https://fullnode.devnet.sui.io:443"
	Testnet string = "https://fullnode.testnet.sui.io:443"
)

var (
	Endpoint string
)

func init() {
	Endpoint = Devnet
}
