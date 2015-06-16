package soap

import (
	"encoding/xml"
	"github.com/xtraclabs/encoding/payloads"
)

type CreateEnvelope struct {
	CreateBody *CreateBody `xml:"Body"`
}

type CreateBody struct {
	Create *Create `xml:"create"`
}

type Create struct {
	WorkItem *payloads.WorkItem `xml:"workItem"`
}

func parseCreateWorkItem(payload string) (*CreateEnvelope, error) {
	var createEnv CreateEnvelope
	err := xml.Unmarshal([]byte(payload), &createEnv)
	if err != nil {
		return nil, err
	}

	return &createEnv, nil
}
