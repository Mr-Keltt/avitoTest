package tender_handler_models

// UpdateTenderRequest represents the request payload for updating a tender.
type UpdateTenderRequest struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	ServiceType string `json:"service_type"`
}
