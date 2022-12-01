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

type SignatureScheme string

const (
	ED25519   = SignatureScheme("ED25519")
	Secp256k1 = SignatureScheme("Secp256k1")
)

type ExecuteTransactionRequestType string

const (
	ImmediateReturn       = ExecuteTransactionRequestType("ImmediateReturn")
	WaitForTxCert         = ExecuteTransactionRequestType("WaitForTxCert")
	WaitForEffectsCert    = ExecuteTransactionRequestType("WaitForEffectsCert")
	WaitForLocalExecution = ExecuteTransactionRequestType("WaitForLocalExecution")
)
