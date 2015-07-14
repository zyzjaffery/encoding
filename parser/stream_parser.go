package parser

import (
	"encoding/xml"
	"github.com/yaozong/encoding/types"
	"strconv"
	"strings"
)

func streamParseB2bCreateWorkItem(payload string) (*types.SoapEnvelope, error) {
	decoder := xml.NewDecoder(strings.NewReader(payload))

	var soapEnv types.SoapEnvelope
	var party *types.Party
	var tsf *types.TSFType
	var docAttach *types.DocAttachment
	var revolutionNote *types.ResolutionNoteType
	var currentElementName string
	parentElementNames := make([]string, 0)

	for {
		t, _ := decoder.Token()
		if t == nil {
			break
		}

		switch t.(type) {

		case xml.StartElement:
			currentElementName = t.(xml.StartElement).Name.Local
			parentElementNames = append(parentElementNames, currentElementName)
			switch currentElementName {
			case "Body":
				soapEnv.Body = *new(types.SoapBody)
			case "CreateWorkItem":
				soapEnv.Body.CreateWorkItem = new(types.CreateWorkItem)
			case "TSFs":
				soapEnv.Body.CreateWorkItem.TSFs = new(types.TSFs)
			case "TSF":
				if soapEnv.Body.CreateWorkItem.TSFs.TSFs == nil {
					newSlice := make([]types.TSFType, 0)
					soapEnv.Body.CreateWorkItem.TSFs.TSFs = &newSlice
				}
				tsf = new(types.TSFType)
				for _, attr := range (t.(xml.StartElement)).Attr {
					if attr.Name.Local == "Name" {
						tsf.Name = &attr.Value
					}
				}
			case "Parties":
				soapEnv.Body.CreateWorkItem.Parties = new(types.Parties)
			case "Party":
				if soapEnv.Body.CreateWorkItem.Parties.Partys == nil {
					newSlice := make([]types.Party, 0)
					soapEnv.Body.CreateWorkItem.Parties.Partys = &newSlice
				}
				party = new(types.Party)
				for _, attr := range (t.(xml.StartElement)).Attr {
					if attr.Name.Local == "Name" {
						party.Name = &attr.Value
					}
				}
			case "DocAttachments":
				soapEnv.Body.CreateWorkItem.DocAttachments = new(types.DocAttachmentsType)
			case "DocAttachment":
				if soapEnv.Body.CreateWorkItem.DocAttachments.DocAttachments == nil {
					newSlice := make([]types.DocAttachment, 0)
					soapEnv.Body.CreateWorkItem.DocAttachments.DocAttachments = &newSlice
				}
				docAttach = new(types.DocAttachment)
			case "ResolutionNote":
				revolutionNote = new(types.ResolutionNoteType)
			}

		case xml.EndElement:
			endElement := t.(xml.EndElement).Name.Local
			switch endElement {
			case "TSF":
				newSlice := append(*soapEnv.Body.CreateWorkItem.TSFs.TSFs, *tsf)
				soapEnv.Body.CreateWorkItem.TSFs.TSFs = &newSlice
			case "Party":
				newSlice := append(*soapEnv.Body.CreateWorkItem.Parties.Partys, *party)
				soapEnv.Body.CreateWorkItem.Parties.Partys = &newSlice
			case "DocAttachment":
				newSlice := append(*soapEnv.Body.CreateWorkItem.DocAttachments.DocAttachments, *docAttach)
				soapEnv.Body.CreateWorkItem.DocAttachments.DocAttachments = &newSlice
			case "ResolutionNote":
				soapEnv.Body.CreateWorkItem.ResolutionNote = revolutionNote
			}
			parentElementNames = parentElementNames[:len(parentElementNames)-1]
			currentElementName = ""

		case xml.CharData:
			switch currentElementName {
			case "CurrentNode":
				cn := string(t.(xml.CharData))
				soapEnv.Body.CreateWorkItem.CurrentNode = &cn
			case "ItemType":
				cn := string(t.(xml.CharData))
				soapEnv.Body.CreateWorkItem.ItemType = &cn
			case "ItemSubtype":
				cn := string(t.(xml.CharData))
				soapEnv.Body.CreateWorkItem.ItemSubtype = &cn
			case "Description":
				cn := string(t.(xml.CharData))
				soapEnv.Body.CreateWorkItem.Description = &cn
			case "Cause":
				cn := string(t.(xml.CharData))
				soapEnv.Body.CreateWorkItem.Cause = &cn
			case "SubCause":
				cn := string(t.(xml.CharData))
				soapEnv.Body.CreateWorkItem.SubCause = &cn
			case "ProblemInd":
				cn := string(t.(xml.CharData))
				soapEnv.Body.CreateWorkItem.ProblemInd = &cn
			case "Priority":
				cn := string(t.(xml.CharData))
				soapEnv.Body.CreateWorkItem.Priority = &cn
			case "Status":
				cn := string(t.(xml.CharData))
				soapEnv.Body.CreateWorkItem.Status = &cn
			case "Amount":
				am := string(t.(xml.CharData))
				soapEnv.Body.CreateWorkItem.Amount = &am
			case "CommType":
				cn := string(t.(xml.CharData))
				soapEnv.Body.CreateWorkItem.CommType = &cn
			case "Memo":
				cn := string(t.(xml.CharData))
				if parentElementNames[len(parentElementNames)-2] == "ResolutionNote" {
					revolutionNote.Memo = &cn
				} else if parentElementNames[len(parentElementNames)-2] == "CreateWorkItem" {
					soapEnv.Body.CreateWorkItem.Memo = &cn
				}
			case "TransferTo":
				cn := string(t.(xml.CharData))
				soapEnv.Body.CreateWorkItem.TransferTo = &cn
			case "EvaluateRuleInd":
				cn := string(t.(xml.CharData))
				soapEnv.Body.CreateWorkItem.EvaluateRuleInd = &cn
			case "ParentWorkItemNo":
				cn := string(t.(xml.CharData))
				soapEnv.Body.CreateWorkItem.ParentWorkItemNo = &cn
			case "SplitId":
				cn, _ := strconv.Atoi(string(t.(xml.CharData)))
				soapEnv.Body.CreateWorkItem.SplitId = &cn
			case "ParentSplitStat":
				cn := string(t.(xml.CharData))
				soapEnv.Body.CreateWorkItem.ParentSplitStat = &cn
			case "SplitLevelNumber":
				cn, _ := strconv.Atoi(string(t.(xml.CharData)))
				soapEnv.Body.CreateWorkItem.SplitLevelNumber = &cn

			case "AccountNo":
				cn := string(t.(xml.CharData))
				party.AccountNo = &cn
			case "AccountType":
				cn := string(t.(xml.CharData))
				party.AccountType = &cn
			case "Address1":
				cn := string(t.(xml.CharData))
				party.Address1 = &cn
			case "Address2":
				cn := string(t.(xml.CharData))
				party.Address2 = &cn
			case "Address3":
				cn := string(t.(xml.CharData))
				party.Address3 = &cn
			case "Address4":
				cn := string(t.(xml.CharData))
				party.Address4 = &cn
			case "Address5":
				cn := string(t.(xml.CharData))
				party.Address5 = &cn
			case "Address6":
				cn := string(t.(xml.CharData))
				party.Address6 = &cn
			case "City":
				cn := string(t.(xml.CharData))
				party.City = &cn
			case "State":
				cn := string(t.(xml.CharData))
				party.State = &cn
			case "PostalCode":
				cn := string(t.(xml.CharData))
				party.PostalCode = &cn
			case "Country":
				cn := string(t.(xml.CharData))
				party.Country = &cn
			case "Phone":
				cn := string(t.(xml.CharData))
				party.Phone = &cn
			case "Extension":
				cn := string(t.(xml.CharData))
				party.Extension = &cn
			case "PartyId":
				cn := string(t.(xml.CharData))
				party.PartyId = &cn
			case "PartyIdType":
				cn := string(t.(xml.CharData))
				party.PartyIdType = &cn
			case "OrgName":
				cn := string(t.(xml.CharData))
				party.OrgName = &cn
			case "FirstName":
				cn := string(t.(xml.CharData))
				party.FirstName = &cn
			case "MiddleName":
				cn := string(t.(xml.CharData))
				party.MiddleName = &cn
			case "LastName":
				cn := string(t.(xml.CharData))
				party.Country = &cn
			case "PlanNo":
				cn := string(t.(xml.CharData))
				party.Phone = &cn

			case "TSF":
				cn := string(t.(xml.CharData))
				tsf.Value = &cn

			case "DocumentId":
				cn := string(t.(xml.CharData))
				docAttach.DocumentId = &cn
			case "DocumentDesc":
				cn := string(t.(xml.CharData))
				docAttach.DocumentDesc = &cn
			case "SourceSysConnInfo":
				cn := string(t.(xml.CharData))
				docAttach.SourceSysConnInfo = &cn
			case "SourceSysVersion":
				cn := string(t.(xml.CharData))
				docAttach.SourceSysVersion = &cn
			case "SourceSysName":
				cn := string(t.(xml.CharData))
				docAttach.SourceSysName = &cn

			case "ControlNo":
				cn := string(t.(xml.CharData))
				revolutionNote.ControlNo = &cn
			case "Name":
				cn := string(t.(xml.CharData))
				revolutionNote.Name = &cn
			case "Note":
				cn := string(t.(xml.CharData))
				revolutionNote.Note = &cn
			}

		}
	}

	return &soapEnv, nil
}
