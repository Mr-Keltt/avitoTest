package entities

import (
	"time"
)

// Tender represents a procurement or bidding process.
type Tender struct {
	ID             int             `gorm:"primaryKey"`
	OrganizationID int             `gorm:"not null"`
	Organization   Organization    `gorm:"foreignKey:OrganizationID"`
	CreatorID      int             `gorm:"not null"`
	Creator        User            `gorm:"foreignKey:CreatorID;constraint:OnDelete:CASCADE;"`
	Bids           []Bid           `gorm:"foreignKey:TenderID;constraint:OnDelete:CASCADE;"`
	Versions       []TenderVersion `gorm:"foreignKey:TenderID;constraint:OnDelete:CASCADE;"`
	Status         string          `gorm:"type:varchar(50);not null"`
	ServiceType    string          `gorm:"size:100"`
	CreatedAt      time.Time       `gorm:"autoCreateTime"`
}

// TenderVersion represents a version of a Tender.
type TenderVersion struct {
	ID          int       `gorm:"primaryKey"`
	TenderID    int       `gorm:"not null"`
	Tender      Tender    `gorm:"foreignKey:TenderID"`
	Name        string    `gorm:"not null;size:100"`
	Description string    `gorm:"size:255"`
	Version     int       `gorm:"not null"`
	UpdatedAt   time.Time `gorm:"autoUpdateTime"`
}
