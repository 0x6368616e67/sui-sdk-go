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

type SharedRecipient struct {
	Version uint64 `json:"initial_shared_version"`
}

type Recipient struct {
	AddressOwner *string          `json:"AddressOwner,omitempty"`
	ObjectOwner  *string          `json:"ObjectOwner,omitempty"`
	Shared       *SharedRecipient `json:"Shared,omitempty"`
}

type Immutable string

type EventID struct {
	TxSeq    uint64 `json:"txSeq"`
	EventSeq uint64 `json:"eventSeq"`
}

type MoveEvent struct {
	PackageID         ObjectID               `json:"packageId"`
	TransactionModule string                 `json:"transactionModule"`
	Sender            string                 `json:"sender"`
	Type              string                 `json:"type"`
	Fields            map[string]interface{} `json:"fields"`
	BCS               string                 `json:"bcs"`
}

type PublishEvent struct {
	Sender    string   `json:"sender"`
	PackageID ObjectID `json:"packageId"`
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
	Owner             Recipient         `json:"owner"`
	ChangeType        BalanceChangeType `json:"changeType"`
	CoinType          string            `json:"coinType"`
	CoinObjectID      ObjectID          `json:"coinObjectId"`
	Version           uint64            `json:"version"`
	Amount            int64             `json:"amount"`
}

type TransferObjectEvent struct {
	PackageID         ObjectID  `json:"packageId"`
	TransactionModule string    `json:"transactionModule"`
	Sender            string    `json:"sender"`
	Recipient         Recipient `json:"recipient"`
	ObjectType        string    `json:"objectType"`
	ObjectID          ObjectID  `json:"objectId"`
	Version           uint64    `json:"version"`
}

type MutateObjectEvent struct {
	PackageID         ObjectID `json:"packageId"`
	TransactionModule string   `json:"transactionModule"`
	Sender            string   `json:"sender"`
	ObjectType        string   `json:"objectType"`
	ObjectID          ObjectID `json:"objectId"`
	Version           uint64   `json:"version"`
}

type DeleteObjectEvent struct {
	PackageID         ObjectID `json:"packageId"`
	TransactionModule string   `json:"transactionModule"`
	Sender            string   `json:"sender"`
	ObjectID          ObjectID `json:"objectId"`
	Version           uint64   `json:"version"`
}

type NewObjectEvent struct {
	PackageID         ObjectID  `json:"packageId"`
	TransactionModule string    `json:"transactionModule"`
	Sender            string    `json:"sender"`
	Recipient         Recipient `json:"recipient"`
	ObjectType        string    `json:"objectType"`
	ObjectID          ObjectID  `json:"objectId"`
	Version           uint64    `json:"version"`
}

type SuiEvent struct {
	MoveEvent              *MoveEvent              `json:"moveEvent,omitempty"`
	Publish                *PublishEvent           `json:"publish,omitempty"`
	CoinBalanceChangeEvent *CoinBalanceChangeEvent `json:"coinBalanceChange,omitempty"`
	TransferObject         *TransferObjectEvent    `json:"transferObject,omitempty"`
	MutateObject           *MutateObjectEvent      `json:"mutateObject,omitempty"`
	DeleteObject           *DeleteObjectEvent      `json:"deleteObject,omitempty"`
	NewObject              *NewObjectEvent         `json:"newObject,omitempty"`
	EpochChange            *big.Int                `json:"epochChange,omitempty"`
	Checkpoint             *big.Int                `json:"checkpoint,omitempty"`
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
