package bid_handler_models

import "time"

type BidResponse struct {
	ID             int       `json:"id"`
	Name           string    `json:"name"`
	Description    string    `json:"description"`
	TenderID       int       `json:"tender_id"`
	OrganizationID int       `json:"organization_id"`
	CreatorID      int       `json:"creator_id"`
	Status         string    `json:"status"`
	CreatedAt      time.Time `json:"created_at"`
	Version        int       `json:"version"`
}
