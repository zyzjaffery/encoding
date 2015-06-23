package soap

import (
	"encoding/xml"
	"fmt"
	"github.com/xtraclabs/encoding/payloads"
	"strconv"
	"strings"
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

func streamParseCreateWorkItem(payload string) (*CreateEnvelope, error) {
	reader := strings.NewReader(payload)
	decoder := xml.NewDecoder(reader)

	var createEnv CreateEnvelope
	var currentElement string
	for {
		t, _ := decoder.Token()
		if t == nil {
			break
		}

		switch t.(type) {

		case xml.StartElement:
			name := t.(xml.StartElement).Name
			currentElement = name.Local
			switch currentElement {
			case "Body":
				createEnv.CreateBody = new(CreateBody)
			case "create":
				createEnv.CreateBody.Create = new(Create)
			case "workItem":
				createEnv.CreateBody.Create.WorkItem = new(payloads.WorkItem)

			}

			//fmt.Println("StartElement")
			//fmt.Printf("\tElement: %s\n", currentElement)

		case xml.EndElement:
			//fmt.Println("end element")
			currentElement = ""

		case xml.CharData:
			//fmt.Println("char data, current element is ", currentElement)
			switch currentElement {
			case "correspondenceCount":
				fmt.Println("handling corr count")
				cc, _ := strconv.Atoi(string(t.(xml.CharData)))
				createEnv.CreateBody.Create.WorkItem.CorrespondenceCount = &cc
			}

		}
	}

	return &createEnv, nil
}
