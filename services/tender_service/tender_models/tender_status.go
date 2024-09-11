package tender_models

type TenderStatus string

const (
	TenderStatusCreated   TenderStatus = "CREATED"
	TenderStatusPublished TenderStatus = "PUBLISHED"
	TenderStatusClosed    TenderStatus = "CLOSED"
)
