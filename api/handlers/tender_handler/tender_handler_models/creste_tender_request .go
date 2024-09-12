package tender_handler_models

// CreateTenderRequest represents the request payload for creating a tender.
type CreateTenderRequest struct {
	Name            string `json:"name"`
	Description     string `json:"description"`
	ServiceType     string `json:"service_type"`
	Status          string `json:"status"`
	OrganizationID  int    `json:"organization_id"`
	CreatorUsername string `json:"creator_username"`
}
