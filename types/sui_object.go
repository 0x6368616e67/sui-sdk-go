package types

type SuiObjectRef struct {
	/** Base64 string representing the object digest */
	Digest TransactionDigest `json:"digest"`
	/** Hex code as string representing the object id */
	ObjectID ObjectID `json:"objectId"`
	/** Object version */
	Version uint64 `json:"version"`
}

type SuiObjectInfo struct {
	SuiObjectRef
	Type  string    `json:"type"`
	Owner Recipient `json:"owner"`
}

type SuiCoinMetadata struct {
	Decimals    uint8    `json:"decimals"`
	Description string   `json:"description"`
	IconUrl     string   `json:"iconUrl,omitempty"`
	ID          ObjectID `json:"id,omitempty"`
	Name        string   `json:"name"`
	Symbol      string   `json:"symbol"`
}

type SuiObjectStatus string

const (
	SuiObjectStatusExists    = SuiObjectStatus("Exists")
	SuiObjectStatusNotExists = SuiObjectStatus("NotExists")
	SuiObjectStatusDeleted   = SuiObjectStatus("Deleted")
)

type SuiObjectType string

const (
	SuiObjectTypeObject  = SuiObjectType("moveObject")
	SuiObjectTypePackage = SuiObjectType("package")
)

type SuiObjectDetailer interface {
	SuiObjectDetailer()
}

type SuiData struct {
}

type ObjectOwner struct {
}

type SuiObejct struct {
	Data                SuiData           `josn:"data"`
	Owner               ObjectOwner       `json:"owner"`
	PreviousTransaction TransactionDigest `json:"previousTransaction"`
	StorageRebate       int64             `josn:"storageRebate"`
	Reference           SuiObjectRef      `json:"reference"`
}

func (so *SuiObejct) SuiObjectDetailer() {

}

type SuiObjectData struct {
	Status  SuiObjectStatus   `json:"status"`
	Details SuiObjectDetailer `json:"details"`
}
