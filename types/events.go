package types

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
	MoveEvent              = EventType("MoveEvent")
	PublishEvent           = EventType("Publish")
	TransferObjectEvent    = EventType("TransferObject")
	MutateObjectEvent      = EventType("MutateObject")
	CoinBalanceChangeEvent = EventType("CoinBalanceChange")
	DeleteObjectEvent      = EventType("DeleteObject")
	NewObjectEvent         = EventType("NewObject")
	EpochChangeEvent       = EventType("EpochChange")
	CheckpointEvent        = EventType("Checkpoint")
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

type SuiEventEnvelope struct {
	Timestamp uint64            `json:"timestamp"`
	TxDigest  TransactionDigest `json:"txDigest"`
	ID        EventID           `json:"id"`
	//Event     SuiEvent          `json:"event"`
}

type SuiEvents []SuiEventEnvelope

type PaginatedEvents struct {
	Data       SuiEvents `json:"data"`
	NextCursor *EventID  `json:"nextCursor"`
}
