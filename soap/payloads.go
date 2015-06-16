package soap

var payload2Attachments = `
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

var payload20Attachments = `
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

var payload200Attachments = `
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
