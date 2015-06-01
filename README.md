## Overview

This project is for benchmarking various xml and json marshalling and
unmarshalling schemes.

## Learnings

### Pointer Fields

The XML Unmarshaller allocates pointer fields to hold potentially optional content.

For example, if the work item struct looks like

<pre>
type workItem struct {
	CorrespondenceCount          *int                          `xml:"correspondenceCount,omitempty" json:"correspondenceCount"`
	DestinationQueue             *key                          `xml:"destinationQueue,omitempty" json:"destinationQueue"`
	DocumentAttachmentCount      *int                          `xml:"documentAttachmentCount" json:"documentAttachmentCount"`
	DocumentAttachmentReferences *documentAttachmentReferences `xml:"documentAttachmentReferences" json:"documentAttachmentReferences"`
    etc...
</pre>

And we have this document:


    <ns:workItem>
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

 
 The unmarshaller will allocate doc attachment count and references fields to hold the document content.
 
 
### Omitting Fields in XML Output

Omitting fields works as simple as adding omitempty to the xml tags associated with a field - see the above
document snippet for an example.

## Next

* Rename soap parts to where we can reuse the domain specific internal items in the context of multiple soap
transactions
