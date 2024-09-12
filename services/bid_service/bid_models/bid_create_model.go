package bid_models

type BidCreateModel struct {
	Name           string `json:"name" validate:"required"`
	Description    string `json:"description" validate:"required"`
	TenderID       int    `json:"tender_id" validate:"required"`
	OrganizationID int    `json:"organization_id" validate:"required"`
	CreatorID      int    `json:"creator_id" validate:"required"`
	Status         string `json:"status" validate:"required,oneof=CREATED PUBLISHED CANCELED"`
}
