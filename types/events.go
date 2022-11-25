package types

import "math/big"

type EventQuery interface {
	EventQuery() string
}

type AllEventQuery string

func (eq *AllEventQuery) EventQuery() string {
	return "All"
}

const AllEvent = AllEventQuery("All")

type TransactionEventQuery struct {
	Transaction TransactionDigest `json:"Transaction"`
}

func (eq *TransactionEventQuery) EventQuery() string {
	return "Transaction"
}

type MoveModule struct {
	Package ObjectID `json:"package"`
	Module  string   `json:"module"`
}

type MoveModuleEventQuery struct {
	MoveModule MoveModule `json:"MoveModule"`
}

func (eq *MoveModuleEventQuery) EventQuery() string {
	return "MoveModule"
}

type MoveEventEventQuery struct {
	MoveEvent string `json:"MoveEvent"`
}

func (eq *MoveEventEventQuery) EventQuery() string {
	return "MoveEvent"
}

type EventType string

const (
	MoveEventType              = EventType("MoveEvent")
	PublishEventType           = EventType("Publish")
	TransferObjectEventType    = EventType("TransferObject")
	MutateObjectEventType      = EventType("MutateObject")
	CoinBalanceChangeEventType = EventType("CoinBalanceChange")
	DeleteObjectEventType      = EventType("DeleteObject")
	NewObjectEventType         = EventType("NewObject")
	EpochChangeEventType       = EventType("EpochChange")
	CheckpointEventType        = EventType("Checkpoint")
)

type EventTypeEventQuery struct {
	EventType EventType `json:"EventType"`
}

func (eq *EventTypeEventQuery) EventQuery() string {
	return "EventType"
}

type SenderEventQuery struct {
	Sender string `json:"Sender"`
}

func (eq *SenderEventQuery) EventQuery() string {
	return "Sender"
}

type ObjectEventQuery struct {
	Object ObjectID `json:"Object"`
}

func (eq *ObjectEventQuery) EventQuery() string {
	return "Object"
}

type TimeRange struct {
	Start uint64 `json:"start_time"`
	End   uint64 `json:"end_time"`
}
type TimeRangeEventQuery struct {
	TimeRange TimeRange `json:"TimeRange"`
}

func (eq *TimeRangeEventQuery) EventQuery() string {
	return "TimeRange"
}

type Recipienter interface {
	Recipienter() string
}

type AddressOwner struct {
	AddressOwner string `json:"AddressOwner"`
}

func (oo *AddressOwner) Recipienter() string {
	return "AddressOwner"
}

type ObjectOwner struct {
	ObjectOwner string `json:"ObjectOwner"`
}

func (oo *ObjectOwner) Recipienter() string {
	return "ObjectOwner"
}

type Immutable string

func (oo *Immutable) Recipienter() string {
	return "Immutable"
}

type Shared struct {
	Version uint64 `json:"initial_shared_version"`
}

func (oo *Shared) Recipienter() string {
	return "Shared"
}

type RecipientEventQuery struct {
	Recipient Recipienter `json:"Recipient"`
}

func (eq *RecipientEventQuery) EventQuery() string {
	return "Recipient"
}

type EventID struct {
	TxSeq    uint64 `json:"txSeq"`
	EventSeq uint64 `json:"eventSeq"`
}

type SuiEvent interface {
	SuiEvent() string
}

type MoveEvent struct {
	PackageID         ObjectID               `json:"packageId"`
	TransactionModule string                 `json:"transactionModule"`
	Sender            string                 `json:"sender"`
	Type              string                 `json:"type"`
	Fields            map[string]interface{} `json:"fields"`
	BCS               string                 `json:"bcs"`
}

type MoveSuiEvent struct {
	MoveEvent MoveEvent `json:"moveEvent"`
}

func (se *MoveSuiEvent) SuiEvent() string {
	return "MoveSuiEvent"
}

type PublishEvent struct {
	Sender    string   `json:"sender"`
	PackageID ObjectID `json:"packageId"`
}
type PublishSuiEvent struct {
	Publish PublishEvent `json:"publish"`
}

func (se *PublishSuiEvent) SuiEvent() string {
	return "PublishSuiEvent"
}

type BalanceChangeType string

const (
	GasBalanceChangeType     = BalanceChangeType("Gas")
	PayBalanceChangeType     = BalanceChangeType("Pay")
	ReceiveBalanceChangeType = BalanceChangeType("Receive")
)

type CoinBalanceChangeEvent struct {
	PackageID         ObjectID          `json:"packageId"`
	TransactionModule string            `json:"transactionModule"`
	Sender            string            `json:"sender"`
	Owner             Recipienter       `json:"owner"`
	ChangeType        BalanceChangeType `json:"changeType"`
	CoinType          string            `json:"coinType"`
	CoinObjectID      ObjectID          `json:"coinObjectId"`
	Version           uint64            `json:"version"`
	Amount            uint64            `json:"amount"`
}

type CoinBalanceChangeSuiEvent struct {
	CoinBalanceChange CoinBalanceChangeEvent `json:"coinBalanceChange"`
}

func (se *CoinBalanceChangeSuiEvent) SuiEvent() string {
	return "CoinBalanceChangeSuiEvent"
}

type TransferObjectEvent struct {
	PackageID         ObjectID    `json:"packageId"`
	TransactionModule string      `json:"transactionModule"`
	Sender            string      `json:"sender"`
	Recipient         Recipienter `json:"recipient"`
	ObjectType        string      `json:"objectType"`
	ObjectID          ObjectID    `json:"objectId"`
	Version           uint64      `json:"version"`
}

type TransferObjectSuiEvent struct {
	TransferObject TransferObjectEvent `json:"transferObject"`
}

func (se *TransferObjectSuiEvent) SuiEvent() string {
	return "TransferObjectSuiEvent"
}

type MutateObjectEvent struct {
	PackageID         ObjectID `json:"packageId"`
	TransactionModule string   `json:"transactionModule"`
	Sender            string   `json:"sender"`
	ObjectType        string   `json:"objectType"`
	ObjectID          ObjectID `json:"objectId"`
	Version           uint64   `json:"version"`
}

type MutateObjectSuiEvent struct {
	MutateObject MutateObjectEvent `json:"mutateObject"`
}

func (se *MutateObjectSuiEvent) SuiEvent() string {
	return "MutateObjectSuiEvent"
}

type DeleteObjectEvent struct {
	PackageID         ObjectID `json:"packageId"`
	TransactionModule string   `json:"transactionModule"`
	Sender            string   `json:"sender"`
	ObjectID          ObjectID `json:"objectId"`
	Version           uint64   `json:"version"`
}
type DeleteObjectSuiEvent struct {
	DeleteObject DeleteObjectEvent `json:"deleteObject"`
}

func (se *DeleteObjectSuiEvent) SuiEvent() string {
	return "DeleteObjectSuiEvent"
}

type NewObjectEvent struct {
	PackageID         ObjectID    `json:"packageId"`
	TransactionModule string      `json:"transactionModule"`
	Sender            string      `json:"sender"`
	Recipient         Recipienter `json:"recipient"`
	ObjectType        string      `json:"objectType"`
	ObjectID          ObjectID    `json:"objectId"`
	Version           uint64      `json:"version"`
}
type NewObjectSuiEvent struct {
	NewObject NewObjectEvent `json:"newObject"`
}

func (se *NewObjectSuiEvent) SuiEvent() string {
	return "NewObjectSuiEvent"
}

type EpochChangeSuiEvent struct {
	EpochChange big.Int `json:"epochChange"`
}

func (se *EpochChangeSuiEvent) SuiEvent() string {
	return "EpochChangeSuiEvent"
}

type CheckpointSuiEvent struct {
	Checkpoint big.Int `json:"checkpoint"`
}

func (se *CheckpointSuiEvent) SuiEvent() string {
	return "CheckpointSuiEvent"
}

type SuiEventEnvelope struct {
	Timestamp uint64            `json:"timestamp"`
	TxDigest  TransactionDigest `json:"txDigest"`
	ID        EventID           `json:"id"`
	Event     SuiEvent          `json:"event"`
}

type SuiEvents []SuiEventEnvelope

type PaginatedEvents struct {
	Data       SuiEvents `json:"data"`
	NextCursor *EventID  `json:"nextCursor"`
}
