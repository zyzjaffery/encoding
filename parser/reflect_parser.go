package parser

import (
	"encoding/xml"
	"github.com/yaozong/encoding/types"
)

func ReflectParseB2bCreateWorkItem(payload string) (*types.SoapEnvelope, error) {
	createEnv := new(types.SoapEnvelope)
	err := xml.Unmarshal([]byte(payload), createEnv)
	if err != nil {
		return nil, err
	}
	return createEnv, nil
}
