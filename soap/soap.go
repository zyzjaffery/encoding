package main

import (
	"encoding/xml"
	"fmt"
	"github.com/xtraclabs/encoding/payloads"
)

var payload = `
<soapenv:Envelope xmlns:soapenv="http://schemas.xmlsoap.org/soap/envelope/" xmlns:ns="http://xmlns.fmr.com/systems/dev/xtrac/2004/06/" xmlns:ser="http://xmlns.fmr.com/common/headers/2005/12/ServiceProcessingDirectives" xmlns:ser1="http://xmlns.fmr.com/common/headers/2005/12/ServiceCallContext" xmlns:typ="http://xmlns.fmr.com/systems/dev/xtrac/2004/06/types">
<soapenv:Body>
<ns:create>
<ns:workItem>
	<typ:correspondenceCount>10</typ:correspondenceCount>
	<typ:documentAttachmentCount>2</typ:documentAttachmentCount>
	<typ:documentAttachmentReferences>
      <typ:documentAttachmentReference>
         <typ:documentId>1</typ:documentId>
         <typ:description>doc 1 desc</typ:description>
         <typ:connection>doc 1 connection</typ:connection>
      </typ:documentAttachmentReference>
      <typ:documentAttachmentReference>
         <typ:documentId>2</typ:documentId>
         <typ:description>doc 2 desc</typ:description>
         <typ:connection>doc 2 connection</typ:connection>
      </typ:documentAttachmentReference>
   </typ:documentAttachmentReferences>
</ns:workItem>
</ns:create>
</soapenv:Body>
</soapenv:Envelope>
`

type CreateEnvelope struct {
	CreateBody *CreateBody `xml:"Body"`
}

type CreateBody struct {
	Create *Create `xml:"create"`
}

type Create struct {
	WorkItem *payloads.WorkItem `xml:"workItem"`
}

func main() {
	var createEnv CreateEnvelope

	err := xml.Unmarshal([]byte(payload), &createEnv)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	witem := createEnv.CreateBody.Create.WorkItem
	fmt.Printf("%v\n", *witem.CorrespondenceCount)

	docs := witem.DocumentAttachmentReferences.DocumentAttachmentReferences
	for _, v := range docs {
		fmt.Printf("doc %s, '%s' via %s\n", v.DocumentId, v.Description, v.Connection)
	}

}
