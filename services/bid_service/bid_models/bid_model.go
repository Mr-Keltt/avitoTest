package bid_models

import (
	"time"
)

type BidModel struct {
	ID             int       `json:"id"`
	Name           string    `json:"name"`
	Description    string    `json:"description"`
	Status         string    `json:"status"`
	TenderID       int       `json:"tender_id"`
	OrganizationID int       `json:"organization_id"`
	CreatorID      int       `json:"creator_id"`
	ApprovalCount  int       `json:"approval_count"`
	CreatedAt      time.Time `json:"created_at"`
	Version        int       `json:"version"`
}
