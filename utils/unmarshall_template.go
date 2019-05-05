package cim

import (
	"encoding/json"
	"github.com/cheekybits/genny/generic"
)

type TargetType generic.Type

func UnMarshallTargetType(rawValue json.RawMessage) (*TargetType, error) { 
	var result TargetType
    err := json.Unmarshal(rawValue, &result)
    return &result, err 
}
