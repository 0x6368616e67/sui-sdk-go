package types

import (
	"encoding/json"
)

type SuiObjectRef struct {
	Digest   TransactionDigest `json:"digest"`
	ObjectID ObjectID          `json:"objectId"`
	Version  uint64            `json:"version"`
}

type SuiObject struct {
	Data                SuiData           `json:"data"`
	Owner               ObjectOwner       `json:"owner"`
	PreviousTransaction TransactionDigest `json:"previousTransaction"`

	StorageRebate int64        `json:"storageRebate"`
	Reference     SuiObjectRef `json:"reference"`
}

type SuiPackage struct {
	Digest   string `json:"digest"`
	ObjectId string `json:"objectId"`
	Version  int64  `json:"version"`
}

type SuiObjectInfo struct {
	SuiObjectRef
	Type                string            `json:"type"`
	Owner               SuiObjectOwner    `json:"owner"`
	PreviousTransaction TransactionDigest `json:"previousTransaction"`
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
	DataType SuiObjectType `json:"dataType"`

	Disassembled map[string]string `json:"disassembled"`

	Fields            map[string]json.RawMessage `json:"fields"`
	Type              string                     `json:"type"`
	HasPublicTransfer bool                       `json:"has_public_transfer"`
	BCSBytes          string                     `json:"bcs_bytes"`
}

type ObjectOwnerType interface {
	ObjectOwnerType()
}
type ImmutableOwner string

func (o ImmutableOwner) ObjectOwnerType() {

}

type AddressOwner struct {
	AddressOwner string `json:"AddressOwner"`
}

func (o AddressOwner) ObjectOwnerType() {

}

type ObjectOwner struct {
	ObjectOwner string `json:"ObjectOwner"`
}

func (o ObjectOwner) ObjectOwnerType() {

}

type SharedOwnerVersion struct {
	Version uint64 `json:"initial_shared_version"`
}

type SharedOwner struct {
	Shared SharedOwnerVersion `json:"Shared"`
}

func (o SharedOwner) ObjectOwnerType() {

}

type SuiObjectOwner struct {
	Type ObjectOwnerType
}

func (soo *SuiObjectOwner) UnmarshalJSON(b []byte) (err error) {
	var ior ImmutableOwner
	err = json.Unmarshal(b, &ior)
	if err == nil && len(ior) > 0 {
		soo.Type = ior
		return
	}

	var oo ObjectOwner
	err = json.Unmarshal(b, &oo)
	if err == nil && len(oo.ObjectOwner) > 0 {
		soo.Type = oo
		return
	}

	var ao AddressOwner
	err = json.Unmarshal(b, &ao)
	if err == nil && len(ao.AddressOwner) > 0 {
		soo.Type = ao
		return
	}

	var so SharedOwner
	err = json.Unmarshal(b, &so)
	if err == nil && so.Shared.Version > 0 {
		soo.Type = so
		return
	}
	return
}

type SuiObejct struct {
	Data                SuiData           `josn:"data"`
	Owner               SuiObjectOwner    `json:"owner"`
	PreviousTransaction TransactionDigest `json:"previousTransaction"`
	StorageRebate       int64             `josn:"storageRebate"`
	Reference           SuiObjectRef      `json:"reference"`
}

func (so SuiObejct) SuiObjectDetailer() {

}

func (o ObjectID) SuiObjectDetailer() {

}

func (sor SuiObjectRef) SuiObjectDetailer() {

}

type SuiObjectData struct {
	Status  SuiObjectStatus   `json:"status"`
	Details SuiObjectDetailer `json:"details"`
}

func (smnt *SuiObjectData) UnmarshalJSON(b []byte) (err error) {
	type suiObjectData struct {
		Status  SuiObjectStatus `json:"status"`
		Details json.RawMessage `json:"details"`
	}
	var obj suiObjectData
	err = json.Unmarshal(b, &obj)
	if err != nil {
		return
	}
	smnt.Status = obj.Status

	var oid ObjectID
	err = json.Unmarshal(obj.Details, &oid)
	if err == nil && len(oid) > 0 {
		smnt.Details = oid
		return
	}

	var sor SuiObjectRef
	err = json.Unmarshal(obj.Details, &sor)
	if err == nil && len(sor.Digest) > 0 {
		smnt.Details = sor
		return
	}

	var so SuiObejct
	err = json.Unmarshal(obj.Details, &so)
	if err == nil && len(so.PreviousTransaction) > 0 {
		smnt.Details = so
		return
	}

	return
}

type OwnedObjectRef struct {
	Owner     ObjectOwner  `json:"owner"`
	Reference SuiObjectRef `json:"reference"`
}

type GasCostSummary struct {
	ComputationCost int64 `json:"computationCost"`
	StorageCost     int64 `json:"storageCost"`
	StorageRebate   int64 `json:"storageRebate"`
}
