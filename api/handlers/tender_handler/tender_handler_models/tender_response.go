package tender_handler_models

import "time"

// TenderResponse represents the response payload for a tender.
type TenderResponse struct {
	ID             int       `json:"id"`
	Name           string    `json:"name"`
	Description    string    `json:"description"`
	ServiceType    string    `json:"service_type"`
	Status         string    `json:"status"`
	OrganizationID int       `json:"organization_id"`
	CreatedAt      time.Time `json:"created_at"`
	Version        int       `json:"version"`
}
