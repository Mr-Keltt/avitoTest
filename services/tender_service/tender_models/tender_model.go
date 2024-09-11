package tender_models

import "time"

type TenderModel struct {
	ID             int          `json:"id"`
	OrganizationID int          `json:"organization_id"`
	Name           string       `json:"name"`
	Description    string       `json:"description"`
	ServiceType    string       `json:"service_type"`
	Status         TenderStatus `json:"status"`
	CreatedAt      time.Time    `json:"created_at"`
	Version        int          `json:"version"`
}
