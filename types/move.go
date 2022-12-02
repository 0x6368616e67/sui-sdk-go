package types

import (
	"encoding/json"
)

type IMoveFunctionArgType interface {
	IMoveFunctionArgType()
}

type MoveFunctionArgTypeString string

func (t MoveFunctionArgTypeString) IMoveFunctionArgType() {
}

type MoveFunctionArgTypeObject struct {
	Object string `json:"Object"`
}

func (t MoveFunctionArgTypeObject) IMoveFunctionArgType() {

}

type MoveFunctionArgTypes struct {
	Types []IMoveFunctionArgType
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
