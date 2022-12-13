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
	Bytes string       `json:"txBytes"`
	Gas   SuiObjectRef `json:"gas"`
}

type TransferObject struct {
	Recipient string       `json:"recipient"`
	ObjectRef SuiObjectRef `json:"objectRef"`
}

type SuiTransferSui struct {
	Recipient string `json:"recipient"`
	Amount    int64  `json:"amount"`
}

type SuiChangeEpoch struct {
	Epoch             int64 `json:"epoch"`
	StorageCharge     int64 `json:"storage_charge"`
	ComputationCharge int64 `json:"computation_charge"`
}

type PayTransaction struct {
	Coins      []SuiObjectRef `json:"coins"`
	Recipients []string       `json:"recipients"`
	Amounts    []int64        `json:"amounts"`
}

type PaySuiTransaction struct {
	Coins      []SuiObjectRef `json:"coins"`
	Recipients []string       `json:"recipients"`
	Amounts    []int64        `json:"amounts"`
}

type PayAllTransaction struct {
	Coins      []SuiObjectRef `json:"coins"`
	Recipients []string       `json:"recipients"`
}

type SuiTransactionKind struct {
	TransferObject *TransferObject    `json:"TransferObject,omitempty"`
	Publish        *SuiMovePackage    `json:"Publish,omitempty"`
	Call           *MoveCall          `json:"Call,omitempty"`
	TransferSui    *SuiTransferSui    `json:"TransferSui,omitempty"`
	ChangeEpoch    *SuiChangeEpoch    `json:"ChangeEpoch,omitempty"`
	Pay            *PayTransaction    `json:"Pay,omitempty"`
	PaySui         *PaySuiTransaction `json:"PaySui,omitempty"`
	PayAllSui      *PayAllTransaction `json:"PayAllSui,omitempty"`
}

type SuiTransactionData struct {
	Transactions []*SuiTransactionKind `json:"transactions"`
	Sender       string                `json:"sender"`
	GasPayment   SuiObjectRef          `json:"GasPayment"`
	GasBudget    uint64                `json:"GasBudget"`
}

type GenericAuthoritySignature []string

func (gas GenericAuthoritySignature) UnmarshalJSON(b []byte) (err error) {
	var sig string
	err = json.Unmarshal(b, &sig)
	if err == nil && len(sig) > 0 {
		gas = []string{sig}
		return
	}
	gas = make(GenericAuthoritySignature, 1)
	err = json.Unmarshal(b, &gas)
	return
}

type AuthorityQuorumSignInfo struct {
	Epoch     int64                     `json:"epoch"`
	Signature GenericAuthoritySignature `json:"signature"`
}

type CertifiedTransaction struct {
	TransactionDigest TransactionDigest       `json:"transactionDigest"`
	Data              SuiTransactionData      `json:"data"`
	TxSignature       string                  `json:"txSignature"`
	AuthSignInfo      AuthorityQuorumSignInfo `json:"authSignInfo"`
}

type ExecutionStatusType string

const (
	ExecutionStatusSuccess = ExecutionStatusType("success")
	ExecutionStatusFailure = ExecutionStatusType("failure")
)

type ExecutionStatus struct {
	Status ExecutionStatusType `json:"status"`
	Error  *string             `json:"error,omitempty"`
}

type TransactionEffects struct {
	Status            ExecutionStatus     `json:"status"`
	GasUsed           GasCostSummary      `json:"gasUsed"`
	SharedObjects     []SuiObjectRef      `json:"sharedObjects,omitempty"`
	TransactionDigest TransactionDigest   `json:"transactionDigest"`
	Created           []OwnedObjectRef    `json:"created,omitempty"`
	Mutated           []OwnedObjectRef    `json:"mutated,omitempty"`
	Unwrapped         []OwnedObjectRef    `json:"unwrapped,omitempty"`
	Deleted           []SuiObjectRef      `json:"deleted,omitempty"`
	Wrapped           []SuiObjectRef      `json:"wrapped,omitempty"`
	GasObject         OwnedObjectRef      `json:"gasObject"`
	Events            []SuiEvent          `json:"events,omitempty"`
	Dependencies      []TransactionDigest `json:"dependencies,omitempty"`
}

type SuiCertifiedTransactionEffects struct {
	Effects TransactionEffects `json:"effects"`
}

type SuiParsedSplitCoinResponse struct {
	UpdatedCoin SuiObject   `json:"updatedCoin"`
	NewCoins    []SuiObject `json:"newCoins"`
	UpdatedGas  SuiObject   `json:"updatedGas"`
}

type SuiParsedMergeCoinResponse struct {
	UpdatedCoin SuiObject `json:"updatedCoin"`
	UpdatedGas  SuiObject `json:"updatedGas"`
}
type SuiParsedPublishResponse struct {
	CreatedObjects []SuiObject `json:"createdObjects"`
	Package        SuiPackage  `json:"package"`
	UpdatedGas     SuiObject   `json:"updatedGas"`
}

type SuiParsedTransactionResponse struct {
	SplitCoin *SuiParsedSplitCoinResponse `json:"SplitCoin"`
	MergeCoin *SuiParsedMergeCoinResponse `json:"MergeCoin"`
	Publish   *SuiParsedPublishResponse   `json:"Publish"`
}

type SuiExecuteTransactionResponse struct {
	TxDigest    *TransactionDigest              `json:"tx_digest,omitempty"`
	Certificate *CertifiedTransaction           `json:"certificate,omitempty"`
	Effects     *SuiCertifiedTransactionEffects `json:"effects,omitempty"`
}

type SuiTransactionResponse struct {
	Certificate CertifiedTransaction          `json:"certificate"`
	Effects     TransactionEffects            `json:"effects"`
	TimestampMS *int64                        `json:"timestamp_ms,omitempty"`
	ParsedData  *SuiParsedTransactionResponse `json:"parsed_data,omitempty"`
}

type SuiTransactionAuthSignersResponse struct {
	Signers []string `json:"signers"`
}

type TransactionQueryALL string

const (
	TransactionQueryAll = TransactionQueryALL("All")
)

func (tq TransactionQueryALL) TransactionQueryer() {

}

type TransactionQueryMoveFunction struct {
	Package  ObjectID `json:"Package"`
	Module   *string  `json:"module"`
	Function *string  `json:"function"`
}

type TransactionQuery struct {
	All           TransactionQueryALL
	MoveFunction  *TransactionQueryMoveFunction `json:"MoveFunction"`
	InputObject   *ObjectID                     `json:"InputObject"`
	MutatedObject *ObjectID                     `json:"MutatedObject"`
	FromAddress   *string                       `json:"FromAddress"`
	ToAddress     *string                       `json:"ToAddress"`
}

func (tq TransactionQuery) MarshalJSON() ([]byte, error) {
	if tq.All == TransactionQueryAll {
		return []byte(`"All"`), nil
	}
	return json.Marshal(tq)
}

type PaginatedTransactionDigests struct {
	Data       []TransactionDigest `json:"data"`
	NextCursor *TransactionDigest  `json:"nextCursor,omitempty"`
}
