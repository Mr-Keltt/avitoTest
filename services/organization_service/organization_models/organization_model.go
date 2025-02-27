package organization_models

import (
	"time"
)

// OrganizationModel is a structure that represents the organization in the service layer.
type OrganizationModel struct {
	ID          int       `json:"id"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	Type        string    `json:"type"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}
