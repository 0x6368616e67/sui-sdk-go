package types

import "encoding/json"

const (
	RAW_TRANSACTION_SALT           = "APTOS::RawTransaction"
	RAW_TRANSACTION_WITH_DATA_SALT = "APTOS::RawTransactionWithData"
)

type AuthorityPublicKeyBytes struct {
}

type CommitteeInfo struct {
	Epoch         uint64            `json:"epoch"`
	CommitteeInfo []json.RawMessage `json:"committee_info"`
}

type TransactionBytes struct {
	Bytes string `json:"txBytes"`
}

type SuiTransactionData struct {
	//transactions: SuiTransactionKind[];
	Sender     string       `json:"sender"`
	GasPayment SuiObjectRef `json:"GasPayment"`
	GasBudget  uint64       `json:"GasBudget"`
}

type CertifiedTransaction struct {
	TransactionDigest TransactionDigest  `json:"transactionDigest"`
	Data              SuiTransactionData `json:"data"`
	TxSignature       string             `json:"txSignature"`
	//authSignInfo: AuthorityQuorumSignInfo;
}

type TransactionEffects struct {
	/** The transaction digest */
	TransactionDigest TransactionDigest `json:"transactionDigest"`
}

type SuiCertifiedTransactionEffects struct {
	Effects TransactionEffects `json:"effects"`
}

type SuiExecuteTransactionResponse struct {
	TxDigest    *TransactionDigest              `json:"tx_digest,omitempty"`
	Certificate *CertifiedTransaction           `json:"certificate,omitempty"`
	Effects     *SuiCertifiedTransactionEffects `json:"effects,omitempty"`
}
