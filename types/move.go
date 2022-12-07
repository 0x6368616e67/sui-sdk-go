package types

import (
	"encoding/json"
)

type MoveFunctionArgTyper interface {
	MoveFunctionArgTyper()
}

type MoveFunctionArgTypeString string

func (t MoveFunctionArgTypeString) MoveFunctionArgTyper() {
}

type MoveFunctionArgTypeObject struct {
	Object string `json:"Object"`
}

func (t MoveFunctionArgTypeObject) MoveFunctionArgTyper() {

}

type MoveFunctionArgTypes struct {
	Types []MoveFunctionArgTyper
}

func (ats *MoveFunctionArgTypes) UnmarshalJSON(b []byte) (err error) {

	types := make([]json.RawMessage, 0)
	err = json.Unmarshal(b, &types)
	if err != nil {
		return
	}

	for _, b := range types {
		var str string
		err = json.Unmarshal(b, &str)
		if err == nil {
			ats.Types = append(ats.Types, MoveFunctionArgTypeString(str))
			continue
		}
		var obj MoveFunctionArgTypeObject
		err = json.Unmarshal(b, &obj)
		if err == nil {
			ats.Types = append(ats.Types, obj)
		}

	}

	return
}

type SuiMoveAbilitySet struct {
	Abilities []string `json:"abilities"`
}

type SuiMoveNormalizedTyper interface {
	SuiMoveNormalizedTyper()
}

type SuiMoveNormalizedTypeStringType string

func (t SuiMoveNormalizedTypeStringType) SuiMoveNormalizedTyper() {

}

type SuiMoveNormalizedTypeParameterType struct {
	TypeParameter *int64 `json:"TypeParameter"`
}

func (t SuiMoveNormalizedTypeParameterType) SuiMoveNormalizedTyper() {

}

type SuiMoveNormalizedStructTypeStruct struct {
	Address       string                   `json:"address"`
	Module        string                   `json:"module"`
	Name          string                   `json:"name"`
	TypeArguments []*SuiMoveNormalizedType `json:"type_arguments"`
}
type SuiMoveNormalizedStructType struct {
	Struct *SuiMoveNormalizedStructTypeStruct `json:"Struct"`
}

func (t SuiMoveNormalizedStructType) SuiMoveNormalizedTyper() {

}

type SuiMoveNormalizedTypeReference struct {
	Reference *SuiMoveNormalizedType `json:"Reference"`
}

func (t SuiMoveNormalizedTypeReference) SuiMoveNormalizedTyper() {

}

type SuiMoveNormalizedTypeMutableReference struct {
	MutableReference *SuiMoveNormalizedType `json:"MutableReference"`
}

func (t SuiMoveNormalizedTypeMutableReference) SuiMoveNormalizedTyper() {

}

type SuiMoveNormalizedTypeVector struct {
	Vector *SuiMoveNormalizedType `json:"Vector"`
}

func (t SuiMoveNormalizedTypeVector) SuiMoveNormalizedTyper() {

}

type SuiMoveNormalizedType struct {
	Type SuiMoveNormalizedTyper
}

func (t SuiMoveNormalizedType) SuiMoveNormalizedTyper() {

}

func (smnt *SuiMoveNormalizedType) UnmarshalJSON(b []byte) (err error) {
	var st SuiMoveNormalizedTypeStringType
	err = json.Unmarshal(b, &st)
	if err == nil {
		smnt.Type = st
		return
	}
	var pt SuiMoveNormalizedTypeParameterType
	err = json.Unmarshal(b, &pt)
	if err == nil && pt.TypeParameter != nil {
		smnt.Type = pt
		return
	}

	var rt SuiMoveNormalizedTypeReference
	err = json.Unmarshal(b, &rt)
	if err == nil && rt.Reference != nil {
		smnt.Type = rt
		return
	}

	var mrt SuiMoveNormalizedTypeMutableReference
	err = json.Unmarshal(b, &mrt)
	if err == nil && mrt.MutableReference != nil {
		smnt.Type = mrt
		return
	}
	var vt SuiMoveNormalizedTypeVector
	err = json.Unmarshal(b, &vt)
	if err == nil && vt.Vector != nil {
		smnt.Type = vt
		return
	}

	var stt SuiMoveNormalizedStructType
	err = json.Unmarshal(b, &stt)
	if err == nil && stt.Struct != nil {
		smnt.Type = stt
		return
	}

	return nil
}

type SuiMoveNormalizedFunction struct {
	Visibility     string                  `json:"visibility"`
	IsEntry        bool                    `json:"is_entry"`
	TypeParameters []SuiMoveAbilitySet     `json:"type_parameters"`
	Parameters     []SuiMoveNormalizedType `json:"parameters"`
	Return         []SuiMoveNormalizedType `json:"return_"`
}

type SuiMoveModuleID struct {
	Address string `json:"address"`
	Name    string `json:"name"`
}

type SuiMoveStructTypeParameter struct {
	Constraints SuiMoveAbilitySet `json:"constraints"`
	IsPhantom   bool              `json:"is_phantom"`
}

type SuiMoveNormalizedField struct {
	Name string                `json:"name"`
	Type SuiMoveNormalizedType `json:"type_"`
}

type SuiMoveNormalizedStruct struct {
	Abilities      SuiMoveAbilitySet            `json:"abilities"`
	TypeParameters []SuiMoveStructTypeParameter `json:"type_parameters"`
	Fields         []SuiMoveNormalizedField     `json:"fields"`
}

type SuiMoveNormalizedModule struct {
	FileFormatVersion int64                                `json:"file_format_version"`
	Address           string                               `json:"address"`
	Name              string                               `json:"name"`
	Friends           []SuiMoveModuleID                    `json:"friends"`
	Structs           map[string]SuiMoveNormalizedStruct   `json:"structs"`
	ExposedFunctions  map[string]SuiMoveNormalizedFunction `json:"exposed_functions"`
}
