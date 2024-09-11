package constants

type TenderStatus string

const (
	TenderStatusCreated   TenderStatus = "CREATED"
	TenderStatusPublished TenderStatus = "PUBLISHED"
	TenderStatusClosed    TenderStatus = "CLOSED"
)
