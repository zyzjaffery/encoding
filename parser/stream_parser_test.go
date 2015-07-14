package parser

import (
	"github.com/stretchr/testify/assert"
	"github.com/yaozong/encoding/sample"
	"github.com/yaozong/encoding/types"
	"testing"
)

func TestStreamParsing(t *testing.T) {
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
		soapEnv, err := streamParseB2bCreateWorkItem(testCase.samplePayload)
		assert.Nil(t, err)

		req, err := soapEnv.Body.GetPayload()
		assert.Nil(t, err)
		assert.IsType(t, new(types.CreateWorkItem), req)

		assert.Equal(t, testCase.partyAmount, len(*soapEnv.Body.CreateWorkItem.Parties.Partys))
		assert.Equal(t, 4, len(*soapEnv.Body.CreateWorkItem.TSFs.TSFs))
	}
}

func TestStreamParsingElementsWithSameName(t *testing.T) {
	soapEnv, err := streamParseB2bCreateWorkItem(sample.B2bCreateWorkItemWith2Parties)
	assert.Nil(t, err)

	req, err := soapEnv.Body.GetPayload()
	assert.Nil(t, err)
	assert.IsType(t, new(types.CreateWorkItem), req)

	assert.Equal(t, "some memo", *soapEnv.Body.CreateWorkItem.Memo)
	assert.Equal(t, "some memo 2", *soapEnv.Body.CreateWorkItem.ResolutionNote.Memo)
}

func TestStreamParsingElementAttributes(t *testing.T) {
	soapEnv, err := streamParseB2bCreateWorkItem(sample.B2bCreateWorkItemWith2Parties)
	assert.Nil(t, err)

	req, err := soapEnv.Body.GetPayload()
	assert.Nil(t, err)
	assert.IsType(t, new(types.CreateWorkItem), req)

	assert.Equal(t, "resolution note name", *soapEnv.Body.CreateWorkItem.ResolutionNote.Name)
	assert.Equal(t, "AB#", *(*soapEnv.Body.CreateWorkItem.TSFs.TSFs)[0].Name)
	assert.Equal(t, "RRQT", *(*soapEnv.Body.CreateWorkItem.TSFs.TSFs)[1].Name)
}
