package models

import (
	"time"
)

// OrganizationModel is a structure that represents the organization in the service layer.
type OrganizationModel struct {
	ID          int       `json:"id"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	Type        string    `json:"type"` // IE, LLC, JSC
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

// OrganizationCreateModel - model for creating a new organization.
type OrganizationCreateModel struct {
	Name        string `json:"name" validate:"required,min=3,max=100"`
	Description string `json:"description" validate:"max=255"`
	Type        string `json:"type" validate:"required,oneof=IE LLC JSC"`
}

// OrganizationUpdateModel - model for updating an existing organization.
type OrganizationUpdateModel struct {
	ID          int    `json:"id" validate:"required"`
	Name        string `json:"name" validate:"required,min=3,max=100"`
	Description string `json:"description" validate:"max=255"`
	Type        string `json:"type" validate:"required,oneof=IE LLC JSC"`
}
