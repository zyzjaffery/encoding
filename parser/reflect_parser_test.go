package parser

import (
	"github.com/stretchr/testify/assert"
	"github.com/xtraclabs/xtrac-api-poc/types/b2b/v1/createWorkItem"
	"github.com/yaozong/encoding/sample"
	"testing"
)

func TestReflectParsing(t *testing.T) {
	testCases := []struct {
		samplePayload string
		partyAmount   int
	}{
		{sample.B2bCreateWorkItemWith2Parties, 2},
		{sample.B2bCreateWorkItemWith20Parties, 20},
		{sample.B2bCreateWorkItemWith200Parties, 200},
		{sample.B2bCreateWorkItemWith2000Parties, 2000},
	}

	for _, testCase := range testCases {
		soapEnv, err := reflectParseB2bCreateWorkItem(testCase.samplePayload)
		assert.Nil(t, err)

		req, err := soapEnv.Body.GetPayload()
		assert.Nil(t, err)
		assert.IsType(t, new(createWorkItem.CreateWorkItem), req)

		assert.Equal(t, testCase.partyAmount, len(*soapEnv.Body.CreateWorkItem.Parties.Partys))
	}

}
