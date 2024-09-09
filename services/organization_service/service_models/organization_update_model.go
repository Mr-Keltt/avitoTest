package service_models

// OrganizationUpdateModel - model for updating an existing organization.
type OrganizationUpdateModel struct {
	ID          int    `json:"id" validate:"required"`
	Name        string `json:"name" validate:"required,min=3,max=100"`
	Description string `json:"description" validate:"max=255"`
	Type        string `json:"type" validate:"required,oneof=IE LLC JSC"`
}
