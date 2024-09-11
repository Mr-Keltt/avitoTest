// File: data/repositories/tender_repository/tender_repository.go

package tender_repository

import (
	"avitoTest/data/entities"
	"context"
)

// TenderRepository defines the interface for interacting with tenders in the database.
type TenderRepository interface {
	Create(ctx context.Context, tender *entities.Tender) error
	Update(ctx context.Context, tender *entities.Tender) error
	FindByID(ctx context.Context, id int) (*entities.Tender, error)
	GetAll(ctx context.Context) ([]*entities.Tender, error)
	GetAllByServiceType(ctx context.Context, serviceType string) ([]*entities.Tender, error)

	// Tender Version Management
	CreateVersion(ctx context.Context, version *entities.TenderVersion) error
	FindVersionByNumber(ctx context.Context, tenderID int, versionNumber int) (*entities.TenderVersion, error)

	// New method to find the latest version
	FindLatestVersion(ctx context.Context, tenderID int) (*entities.TenderVersion, error)

	// User Responsibility Check
	FindUserOrganizationResponsibility(ctx context.Context, userID, orgID int) (*entities.OrganizationResponsible, error)
}
