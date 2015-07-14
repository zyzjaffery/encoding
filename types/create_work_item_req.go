package types

type CreateWorkItem struct {
	CurrentNode      *string             `xml:"CurrentNode" json:"CurrentNode"`
	ItemType         *string             `xml:"ItemType" json:"ItemType"`
	ItemSubtype      *string             `xml:"ItemSubtype" json:"ItemSubtype"`
	Description      *string             `xml:"Description" json:"Description"`
	Cause            *string             `xml:"Cause" json:"Cause"`
	SubCause         *string             `xml:"SubCause" json:"SubCause"`
	ProblemInd       *string             `xml:"ProblemInd" json:"ProblemInd"`
	Priority         *string             `xml:"Priority" json:"Priority"`
	Status           *string             `xml:"Status" json:"Status"`
	Amount           *string             `xml:"Amount" json:"Amount"`
	CommType         *string             `xml:"CommType" json:"CommType"`
	Memo             *string             `xml:"Memo" json:"Memo"`
	TransferTo       *string             `xml:"TransferTo" json:"TransferTo"`
	EvaluateRuleInd  *string             `xml:"EvaluateRuleInd" json:"EvaluateRuleInd"`
	ParentWorkItemNo *string             `xml:"ParentWorkItemNo" json:"ParentWorkItemNo"`
	SplitId          *int                `xml:"SplitId" json:"SplitId"`
	ParentSplitStat  *string             `xml:"ParentSplitStat" json:"ParentSplitStat"`
	SplitLevelNumber *int                `xml:"SplitLevelNumber" json:"SplitLevelNumber"`
	DocAttachments   *DocAttachmentsType `xml:"DocAttachments" json:"DocAttachments"`
	ResolutionNote   *ResolutionNoteType `xml:"ResolutionNote" json:"ResolutionNote"`
	Parties          *Parties            `xml:"Parties" json:"Parties"`
	TSFs             *TSFs               `xml:"TSFs" json:"TSFs"`
}
type TSFs struct {
	TSFs *[]TSFType `xml:"TSF" json:"TSF"`
}
type TSFType struct {
	Name  *string `xml:"Name,attr" json:"Name"`
	Value *string `xml:",chardata" json:"Value"`
}
type DocAttachmentsType struct {
	DocAttachments *[]DocAttachment
}
type DocAttachment struct {
	DocumentId        *string `xml:"DocumentId" json:"DocumentId"`
	DocumentDesc      *string `xml:"DocumentDesc" json:"DocumentDesc"`
	SourceSysConnInfo *string `xml:"SourceSysConnInfo" json:"SourceSysConnInfo"`
	SourceSysVersion  *string `xml:"SourceSysVersion" json:"SourceSysVersion"`
	SourceSysName     *string `xml:"SourceSysName" json:"SourceSysName"`
}
type ResolutionNoteType struct {
	ControlNo *string `xml:"ControlNo" json:"ControlNo"`
	Name      *string `xml:"Name" json:"Name"`
	Memo      *string `xml:"Memo" json:"Memo"`
	Note      *string `xml:"Note" json:"Note"`
}
type Parties struct {
	Partys *[]Party `xml:"Party" json:"Party"`
}
type Party struct {
	Name        *string `xml:"Name,attr" json:"Name"`
	AccountNo   *string `xml:"AccountNo" json:"AccountNo"`
	AccountType *string `xml:"AccountType" json:"AccountType"`
	Address1    *string `xml:"Address1" json:"Address1"`
	Address2    *string `xml:"Address2" json:"Address2"`
	Address3    *string `xml:"Address3" json:"Address3"`
	Address4    *string `xml:"Address4" json:"Address4"`
	Address5    *string `xml:"Address5" json:"Address5"`
	Address6    *string `xml:"Address6" json:"Address6"`
	City        *string `xml:"City" json:"City"`
	State       *string `xml:"State" json:"State"`
	PostalCode  *string `xml:"PostalCode" json:"PostalCode"`
	Country     *string `xml:"Country" json:"Country"`
	Phone       *string `xml:"Phone" json:"Phone"`
	Extension   *string `xml:"Extension" json:"Extension"`
	PartyId     *string `xml:"PartyId" json:"PartyId"`
	PartyIdType *string `xml:"PartyIdType" json:"PartyIdType"`
	OrgName     *string `xml:"OrgName" json:"OrgName"`
	FirstName   *string `xml:"FirstName" json:"FirstName"`
	MiddleName  *string `xml:"MiddleName" json:"MiddleName"`
	LastName    *string `xml:"LastName" json:"LastName"`
	PlanNo      *string `xml:"PlanNo" json:"PlanNo"`
}
