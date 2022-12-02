package types

type SuiObjectRef struct {
	/** Base64 string representing the object digest */
	Digest TransactionDigest `json:"digest"`
	/** Hex code as string representing the object id */
	ObjectID ObjectID `json:"objectId"`
	/** Object version */
	Version uint64 `json:"version"`
}

type SuiObject struct {
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
