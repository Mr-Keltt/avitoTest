package entities

import (
	"time"
)

type Tender struct {
	ID             int          `gorm:"primaryKey"`
	Name           string       `gorm:"not null;size:100"`
	Description    string       `gorm:"size:255"`
	ServiceType    string       `gorm:"size:100"`
	Status         string       `gorm:"size:50"`
	OrganizationID int          `gorm:"not null"`
	Organization   Organization `gorm:"foreignKey:OrganizationID;constraint:OnDelete:CASCADE;"`
	CreatorID      int          `gorm:"not null"`
	Creator        User         `gorm:"foreignKey:CreatorID;constraint:OnDelete:CASCADE;"`
	Version        int          `gorm:"default:1"`
	CreatedAt      time.Time    `gorm:"autoCreateTime"`
	UpdatedAt      time.Time    `gorm:"autoUpdateTime"`
	Bids           []Bid        `gorm:"foreignKey:TenderID"`
}
