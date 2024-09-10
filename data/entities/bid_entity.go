package entities

import (
	"time"
)

// Bid represents a bid made by a participant for a Tender.
type Bid struct {
	ID             int          `gorm:"primaryKey"`
	TenderID       int          `gorm:"not null"`
	Tender         Tender       `gorm:"foreignKey:TenderID"`
	OrganizationID int          `gorm:"not null"`
	Organization   Organization `gorm:"foreignKey:OrganizationID"`
	CreatorID      int          `gorm:"not null"`
	Creator        User         `gorm:"foreignKey:CreatorID;constraint:OnDelete:CASCADE;"`
	ApprovalCount  int          `gorm:"not null;default:0"`
	Versions       []BidVersion `gorm:"foreignKey:BidID;constraint:OnDelete:CASCADE;"`
	CreatedAt      time.Time    `gorm:"autoCreateTime"`
}

// BidVersion represents a version of a Bid.
type BidVersion struct {
	ID          int       `gorm:"primaryKey"`
	BidID       int       `gorm:"not null"`
	Bid         Bid       `gorm:"foreignKey:BidID"`
	Name        string    `gorm:"not null;size:100"`
	Description string    `gorm:"size:255"`
	Status      string    `gorm:"size:50"`
	Version     int       `gorm:"not null"`
	UpdatedAt   time.Time `gorm:"autoUpdateTime"`
}
