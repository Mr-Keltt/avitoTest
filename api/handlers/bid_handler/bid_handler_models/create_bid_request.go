package bid_handler_models

type CreateBidRequest struct {
	Name           string `json:"name"`
	Description    string `json:"description"`
	Status         string `json:"status"`
	TenderID       int    `json:"tenderId"`
	OrganizationID int    `json:"organizationId"`
	CreatorID      int    `json:"creatorId"`
}
