package payloads

func NewWorkItem() *workItem {
	return new(workItem)
}

type workItem struct {
	CorrespondenceCount          *int                          `xml:"correspondenceCount,omitempty" json:"correspondenceCount,omitempty"`
	DestinationQueue             *key                          `xml:"destinationQueue,omitempty" json:"destinationQueue"`
	DocumentAttachmentCount      *int                          `xml:"documentAttachmentCount" json:"documentAttachmentCount"`
	DocumentAttachmentReferences *documentAttachmentReferences `xml:"documentAttachmentReferences" json:"documentAttachmentReferences"`
	JeopardyFields               jeopardyFieldss               `xml:"jeopardyFields" json:"jeopardyFields"`
	Milestones                   milestoness                   `xml:"milestones" json:"milestones"`
	Fields                       fields                        `xml:"fields" json:"fields"`
	ItemType                     string                        `xml:"itemType" json:"itemType"`
	LockState                    lockState                     `xml:"lockState" json:"lockState"`
	GlobalFieldLockState         lockState                     `xml:"globalFieldLockState" json:"globalFieldLockState"`
	NoteCount                    int                           `xml:"noteCount" json:"noteCount"`
	LinkCount                    int                           `xml:"linkCount" json:"linkCount"`
	Parties                      partys                        `xml:"parties" json:"parties"`
	Status                       string                        `xml:"status" json:"status"`
	StatusType                   string                        `xml:"statusType" json:"statusType"`
	Subtype                      string                        `xml:"subtype" json:"subtype"`
	WorkItemNumber               string                        `xml:"workItemNumber" json:"workItemNumber"`
	CopiedWorkItemNumber         string                        `xml:"copiedWorkItemNumber" json:"copiedWorkItemNumber"`
	WorkItemState                workItemState                 `xml:"workItemState" json:"workItemState"`
	SmartStoreId                 string                        `xml:"smartStoreId" json:"smartStoreId"`
	SmartStoreURL                string                        `xml:"smartStoreURL" json:"smartStoreURL"`
	HyperlinkSpecifications      hyperlinkSpecifications       `xml:"hyperlinkSpecifications" json:"hyperlinkSpecifications"`
	FamilyId                     string                        `xml:"familyId" json:"familyId"`
	ArchiveState                 string                        `xml:"archiveState" json:"archiveState"`
	RepositoryDocumentDetails    string                        `xml:"repositoryDocumentDetails" json:"repositoryDocumentDetails"`
	PurgeDate                    string                        `xml:"purgeDate" json:"purgeDate"`
}
type key struct {
	Id   int64  `xml:"id" json:"id"`
	Name string `xml:"name" json:"name"`
}
type documentAttachmentReferences struct {
	DocumentAttachmentReferences []documentAttachmentReference `xml:"documentAttachmentReference" json:"documentAttachmentReference"`
}
type documentAttachmentReference struct {
	DocumentId  string `xml:"documentId" json:"documentId"`
	Description string `xml:"description" json:"description"`
	Connection  string `xml:"connection" json:"connection"`
}
type jeopardyFieldss struct {
	JeopardyFields []jeopardyField `xml:"jeopardyField" json:"jeopardyField"`
}
type jeopardyField struct {
	JeopardyComponent string `xml:"jeopardyComponent" json:"jeopardyComponent"`
	TargetDateValueC  string `xml:"targetDateValueC" json:"targetDateValueC"`
	State             string `xml:"state" json:"state"`
}
type milestoness struct {
	Milestones []milestone `xml:"milestone" json:"milestone"`
}
type milestone struct {
	MilestoneId          int64  `xml:"milestoneId" json:"milestoneId"`
	AmberThresholdValueC string `xml:"amberThresholdValueC" json:"amberThresholdValueC"`
	RedThresholdValueC   string `xml:"redThresholdValueC" json:"redThresholdValueC"`
	TargetDateValueC     string `xml:"targetDateValueC" json:"targetDateValueC"`
	MilestoneName        string `xml:"milestoneName" json:"milestoneName"`
	State                string `xml:"state" json:"state"`
	IsComplete           bool   `xml:"isComplete" json:"isComplete"`
	DateStart            string `xml:"dateStart" json:"dateStart"`
	DateComplete         string `xml:"dateComplete" json:"dateComplete"`
}
type fields struct {
	Fields []field `xml:"field" json:"field"`
}
type field struct {
	FieldName string `xml:"fieldName" json:"fieldName"`
	FieldVal  string `xml:"fieldVal" json:"fieldVal"`
}
type lockState struct {
	LockAcquired          bool   `xml:"lockAcquired" json:"lockAcquired"`
	LockTime              string `xml:"lockTime" json:"lockTime"`
	LockOwnerInfo         string `xml:"lockOwnerInfo" json:"lockOwnerInfo"`
	LockOwnerOperatorId   string `xml:"lockOwnerOperatorId" json:"lockOwnerOperatorId"`
	LockOwnerOperatorName string `xml:"lockOwnerOperatorName" json:"lockOwnerOperatorName"`
}
type partys struct {
	Partys []party `xml:"party" json:"party"`
}
type party struct {
	PartyType string `xml:"partyType" json:"partyType"`
	Fields    fields `xml:"fields" json:"fields"`
}
type workItemState struct {
	CreateOperator                       string                   `xml:"createOperator" json:"createOperator"`
	LastUpdateOperator                   string                   `xml:"lastUpdateOperator" json:"lastUpdateOperator"`
	ArchiveStatus                        string                   `xml:"archiveStatus" json:"archiveStatus"`
	CreateDate                           string                   `xml:"createDate" json:"createDate"`
	CurrentNode                          string                   `xml:"currentNode" json:"currentNode"`
	CurrentQueue                         string                   `xml:"currentQueue" json:"currentQueue"`
	CurrentStatusDate                    string                   `xml:"currentStatusDate" json:"currentStatusDate"`
	IbrStatus                            string                   `xml:"ibrStatus" json:"ibrStatus"`
	LastEventDate                        string                   `xml:"lastEventDate" json:"lastEventDate"`
	Locked                               bool                     `xml:"locked" json:"locked"`
	RemoteWorkItemNumber                 string                   `xml:"remoteWorkItemNumber" json:"remoteWorkItemNumber"`
	Rule351                              bool                     `xml:"rule351" json:"rule351"`
	SourceLinkedItem                     string                   `xml:"sourceLinkedItem" json:"sourceLinkedItem"`
	SplitStatus                          string                   `xml:"splitStatus" json:"splitStatus"`
	SuspendStatus                        string                   `xml:"suspendStatus" json:"suspendStatus"`
	TotalMinutesContributingSuspended    int                      `xml:"totalMinutesContributingSuspended" json:"totalMinutesContributingSuspended"`
	TotalMinutesNonContributingSuspended int                      `xml:"totalMinutesNonContributingSuspended" json:"totalMinutesNonContributingSuspended"`
	TotalTimesContributingSuspended      suspensionTimeLimitValue `xml:"totalTimesContributingSuspended" json:"totalTimesContributingSuspended"`
	TotalTimesNonContributingSuspended   suspensionTimeLimitValue `xml:"totalTimesNonContributingSuspended" json:"totalTimesNonContributingSuspended"`
}
type suspensionTimeLimitValue struct {
	SuspensionTimeLimit int `xml:"suspensionTimeLimit" json:"suspensionTimeLimit"`
}
type hyperlinkSpecifications struct {
	HyperlinkSpecifications []hyperlinkSpecification `xml:"hyperlinkSpecification" json:"hyperlinkSpecification"`
}
type hyperlinkSpecification struct {
	FieldName       string          `xml:"fieldName" json:"fieldName"`
	Name            string          `xml:"name" json:"name"`
	ResolvedURLData string          `xml:"resolvedURLData" json:"resolvedURLData"`
	WindowName      string          `xml:"windowName" json:"windowName"`
	HyperlinkFields hyperlinkFields `xml:"hyperlinkFields" json:"hyperlinkFields"`
}
type hyperlinkFields struct {
	HyperlinkFields []string `xml:"hyperlinkField" json:"hyperlinkField"`
}
