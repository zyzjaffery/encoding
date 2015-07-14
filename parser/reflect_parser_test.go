package parser

import (
	"github.com/stretchr/testify/assert"
	"github.com/yaozong/encoding/sample"
	"github.com/yaozong/encoding/types"
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
		soapEnv, err := ReflectParseB2bCreateWorkItem(testCase.samplePayload)
		assert.Nil(t, err)

		req, err := soapEnv.Body.GetPayload()
		assert.Nil(t, err)
		assert.IsType(t, new(types.CreateWorkItem), req)

		assert.Equal(t, testCase.partyAmount, len(*soapEnv.Body.CreateWorkItem.Parties.Partys))
	}

}
