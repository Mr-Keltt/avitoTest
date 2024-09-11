package tender_models

type TenderUpdateModel struct {
	ID          int    `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	ServiceType string `json:"service_type"`
}
