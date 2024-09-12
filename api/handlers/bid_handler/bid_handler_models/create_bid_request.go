package bid_handler_models

type CreateBidRequest struct {
	Name           string `json:"name"`
	Description    string `json:"description"`
	TenderID       int    `json:"tender_id"`
	OrganizationID int    `json:"organization_id"`
	CreatorID      int    `json:"creator_id"`
	Status         string `json:"status"`
}
