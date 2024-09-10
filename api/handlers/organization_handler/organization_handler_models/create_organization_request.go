package organization_handler_models

// CreateOrganizationRequest - API модель для создания новой организации.
type CreateOrganizationRequest struct {
	Name        string `json:"name" validate:"required,min=3,max=100"`
	Description string `json:"description" validate:"max=255"`
	Type        string `json:"type" validate:"required,oneof=IE LLC JSC"`
}
