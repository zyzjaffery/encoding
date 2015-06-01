package main

import (
	"encoding/xml"
	"fmt"
	"github.com/xtraclabs/encoding/payloads"
)

var wireCreate = `
<soapenv:Envelope xmlns:soapenv="http://schemas.xmlsoap.org/soap/envelope/" xmlns:ns="http://xmlns.fmr.com/systems/dev/xtrac/2004/06/" xmlns:ser="http://xmlns.fmr.com/common/headers/2005/12/ServiceProcessingDirectives" xmlns:ser1="http://xmlns.fmr.com/common/headers/2005/12/ServiceCallContext" xmlns:typ="http://xmlns.fmr.com/systems/dev/xtrac/2004/06/types">
   <soapenv:Header>
   </soapenv:Header>
   <soapenv:Body>
      <ns:create>
         <ns:workItem>
            <!--Optional:-->
            <typ:correspondenceCount>?</typ:correspondenceCount>
            <!--Optional:-->
            <typ:destinationQueue>
               <!--Optional:-->
               <typ:id>?</typ:id>
               <!--Optional:-->
               <typ:name>?</typ:name>
            </typ:destinationQueue>
            <!--Optional:-->
            <typ:documentAttachmentCount>?</typ:documentAttachmentCount>
            <!--Optional:-->
            <typ:documentAttachmentReferences>
               <!--Zero or more repetitions:-->
               <typ:documentAttachmentReference>
                  <typ:documentId>?</typ:documentId>
                  <!--Optional:-->
                  <typ:description>?</typ:description>
                  <typ:connection>?</typ:connection>
               </typ:documentAttachmentReference>
            </typ:documentAttachmentReferences>
            <!--Optional:-->
            <typ:jeopardyFields>
               <!--Zero or more repetitions:-->
               <typ:jeopardyField>
                  <typ:jeopardyComponent>?</typ:jeopardyComponent>
                  <typ:targetDateValueC>?</typ:targetDateValueC>
                  <!--Optional:-->
                  <typ:state>?</typ:state>
               </typ:jeopardyField>
            </typ:jeopardyFields>
            <!--Optional:-->
            <typ:milestones>
               <!--Zero or more repetitions:-->
               <typ:milestone>
                  <typ:milestoneId>?</typ:milestoneId>
                  <typ:amberThresholdValueC>?</typ:amberThresholdValueC>
                  <typ:redThresholdValueC>?</typ:redThresholdValueC>
                  <typ:targetDateValueC>?</typ:targetDateValueC>
                  <typ:milestoneName>?</typ:milestoneName>
                  <typ:state>?</typ:state>
                  <typ:isComplete>?</typ:isComplete>
                  <typ:dateStart>?</typ:dateStart>
                  <typ:dateComplete>?</typ:dateComplete>
               </typ:milestone>
            </typ:milestones>
            <!--Optional:-->
            <typ:fields>
               <!--Zero or more repetitions:-->
               <typ:field>
                  <typ:fieldName>?</typ:fieldName>
                  <!--Optional:-->
                  <typ:fieldVal>?</typ:fieldVal>
               </typ:field>
            </typ:fields>
            <!--Optional:-->
            <typ:itemType>?</typ:itemType>
            <!--Optional:-->
            <typ:lockState>
               <typ:lockAcquired>?</typ:lockAcquired>
               <typ:lockTime>?</typ:lockTime>
               <typ:lockOwnerInfo>?</typ:lockOwnerInfo>
               <!--Optional:-->
               <typ:lockOwnerOperatorId>?</typ:lockOwnerOperatorId>
               <typ:lockOwnerOperatorName>?</typ:lockOwnerOperatorName>
            </typ:lockState>
            <!--Optional:-->
            <typ:globalFieldLockState>
               <typ:lockAcquired>?</typ:lockAcquired>
               <typ:lockTime>?</typ:lockTime>
               <typ:lockOwnerInfo>?</typ:lockOwnerInfo>
               <!--Optional:-->
               <typ:lockOwnerOperatorId>?</typ:lockOwnerOperatorId>
               <typ:lockOwnerOperatorName>?</typ:lockOwnerOperatorName>
            </typ:globalFieldLockState>
            <!--Optional:-->
            <typ:noteCount>?</typ:noteCount>
            <!--Optional:-->
            <typ:linkCount>?</typ:linkCount>
            <typ:parties>
               <!--Zero or more repetitions:-->
               <typ:party>
                  <typ:partyType>?</typ:partyType>
                  <typ:fields>
                     <!--Zero or more repetitions:-->
                     <typ:field>
                        <typ:fieldName>?</typ:fieldName>
                        <!--Optional:-->
                        <typ:fieldVal>?</typ:fieldVal>
                     </typ:field>
                  </typ:fields>
               </typ:party>
            </typ:parties>
            <!--Optional:-->
            <typ:status>?</typ:status>
            <!--Optional:-->
            <typ:statusType>?</typ:statusType>
            <!--Optional:-->
            <typ:subtype>?</typ:subtype>
            <!--Optional:-->
            <typ:workItemNumber>?</typ:workItemNumber>
            <!--Optional:-->
            <typ:copiedWorkItemNumber>?</typ:copiedWorkItemNumber>
            <!--Optional:-->
            <typ:workItemState>
               <!--Optional:-->
               <typ:createOperator>?</typ:createOperator>
               <!--Optional:-->
               <typ:lastUpdateOperator>?</typ:lastUpdateOperator>
               <!--Optional:-->
               <typ:archiveStatus>?</typ:archiveStatus>
               <!--Optional:-->
               <typ:createDate>?</typ:createDate>
               <!--Optional:-->
               <typ:currentNode>?</typ:currentNode>
               <!--Optional:-->
               <typ:currentQueue>?</typ:currentQueue>
               <!--Optional:-->
               <typ:currentStatusDate>?</typ:currentStatusDate>
               <!--Optional:-->
               <typ:ibrStatus>?</typ:ibrStatus>
               <!--Optional:-->
               <typ:lastEventDate>?</typ:lastEventDate>
               <!--Optional:-->
               <typ:locked>?</typ:locked>
               <!--Optional:-->
               <typ:remoteWorkItemNumber>?</typ:remoteWorkItemNumber>
               <!--Optional:-->
               <typ:rule351>?</typ:rule351>
               <!--Optional:-->
               <typ:sourceLinkedItem>?</typ:sourceLinkedItem>
               <!--Optional:-->
               <typ:splitStatus>?</typ:splitStatus>
               <!--Optional:-->
               <typ:suspendStatus>?</typ:suspendStatus>
               <!--Optional:-->
               <typ:totalMinutesContributingSuspended>?</typ:totalMinutesContributingSuspended>
               <!--Optional:-->
               <typ:totalMinutesNonContributingSuspended>?</typ:totalMinutesNonContributingSuspended>
               <!--Optional:-->
               <typ:totalTimesContributingSuspended>
                  <typ:suspensionTimeLimit>?</typ:suspensionTimeLimit>
               </typ:totalTimesContributingSuspended>
               <!--Optional:-->
               <typ:totalTimesNonContributingSuspended>
                  <typ:suspensionTimeLimit>?</typ:suspensionTimeLimit>
               </typ:totalTimesNonContributingSuspended>
            </typ:workItemState>
            <!--Optional:-->
            <typ:smartStoreId>?</typ:smartStoreId>
            <!--Optional:-->
            <typ:smartStoreURL>?</typ:smartStoreURL>
            <!--Optional:-->
            <typ:hyperlinkSpecifications>
               <!--Zero or more repetitions:-->
               <typ:hyperlinkSpecification>
                  <typ:fieldName>?</typ:fieldName>
                  <typ:name>?</typ:name>
                  <!--Optional:-->
                  <typ:resolvedURLData>?</typ:resolvedURLData>
                  <!--Optional:-->
                  <typ:windowName>?</typ:windowName>
                  <!--Optional:-->
                  <typ:hyperlinkFields>
                     <!--Zero or more repetitions:-->
                     <typ:hyperlinkField>?</typ:hyperlinkField>
                  </typ:hyperlinkFields>
               </typ:hyperlinkSpecification>
            </typ:hyperlinkSpecifications>
            <!--Optional:-->
            <typ:familyId>?</typ:familyId>
            <!--Optional:-->
            <typ:archiveState>?</typ:archiveState>
            <!--Optional:-->
            <typ:repositoryDocumentDetails>?</typ:repositoryDocumentDetails>
            <!--Optional:-->
            <typ:purgeDate>?</typ:purgeDate>
         </ns:workItem>
         <ns:evaluateRule>?</ns:evaluateRule>
         <ns:note>
            <!--Optional:-->
            <typ:controlNumber>?</typ:controlNumber>
            <!--Optional:-->
            <typ:dateTime>?</typ:dateTime>
            <!--Optional:-->
            <typ:itemSubtype>?</typ:itemSubtype>
            <!--Optional:-->
            <typ:itemType>?</typ:itemType>
            <!--Optional:-->
            <typ:memo>?</typ:memo>
            <!--Optional:-->
            <typ:noteName>?</typ:noteName>
            <!--Optional:-->
            <typ:noteText>?</typ:noteText>
            <!--Optional:-->
            <typ:operatorId>?</typ:operatorId>
            <!--Optional:-->
            <typ:workItemNumber>?</typ:workItemNumber>
         </ns:note>
      </ns:create>
   </soapenv:Body>
</soapenv:Envelope>
`

var workitemPayload = `
 <ns:workItem>
   <typ:documentAttachmentCount>2</typ:documentAttachmentCount>
   <!--Optional:-->
   <typ:documentAttachmentReferences>
      <!--Zero or more repetitions:-->
      <typ:documentAttachmentReference>
         <typ:documentId>1</typ:documentId>
         <!--Optional:-->
         <typ:description>doc 1 desc</typ:description>
         <typ:connection>doc 1 connection</typ:connection>
      </typ:documentAttachmentReference>
      <typ:documentAttachmentReference>
         <typ:documentId>2</typ:documentId>
         <!--Optional:-->
         <typ:description>doc 2 desc</typ:description>
         <typ:connection>doc 2 connection</typ:connection>
      </typ:documentAttachmentReference>
   </typ:documentAttachmentReferences>
 </ns:workItem>
`

func main() {
	var workitem = payloads.NewWorkItem()
	err := xml.Unmarshal([]byte(workitemPayload), workitem)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	fmt.Printf("corr count: %v\n", workitem.CorrespondenceCount)
	fmt.Printf("doc attachment count: %v\n", *workitem.DocumentAttachmentCount)
	fmt.Printf("doc attachments: %v\n", workitem.DocumentAttachmentReferences)

	xmlbytes, err := xml.Marshal(workitem)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	fmt.Println(string(xmlbytes))

}
