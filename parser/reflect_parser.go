package parser
import (
	"github.com/xtraclabs/xtrac-api-poc/types"
	"encoding/xml"
)

func reflectParseB2bCreateWorkItem(payload string) (*types.SoapEnvelope, error) {
	createEnv := new(types.SoapEnvelope)
	err := xml.Unmarshal([]byte(payload), createEnv)
	if err != nil {
		return nil, err
	}
	return createEnv, nil
}