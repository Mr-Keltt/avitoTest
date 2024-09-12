package constants

type BidStatus string

const (
	BidStatusCreated   TenderStatus = "CREATED"
	BidStatusPublished TenderStatus = "PUBLISHED"
	BidStatusClosed    TenderStatus = "CANCELED"
	BidStatusRejected  TenderStatus = "REJECTED"
)
