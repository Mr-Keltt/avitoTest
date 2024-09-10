package entities

import (
	"time"
)

// Organization represents an organization involved in tenders and bids.
type Organization struct {
	ID           int    `gorm:"primaryKey"`
	Name         string `gorm:"not null;size:100"`
	Description  string
	Type         OrganizationType `gorm:"type:organization_type"`
	CreatedAt    time.Time        `gorm:"autoCreateTime"`
	UpdatedAt    time.Time        `gorm:"autoUpdateTime"`
	Responsibles []User           `gorm:"many2many:organization_responsibles;"`
	Tenders      []Tender         `gorm:"foreignKey:OrganizationID;constraint:OnDelete:CASCADE;"`
	Comments     []Comment        `gorm:"foreignKey:OrganizationID;constraint:OnDelete:CASCADE;"`
}

// OrganizationResponsible represents a link between an Organization and a User.
type OrganizationResponsible struct {
	ID             int          `gorm:"primaryKey"`
	OrganizationID int          `gorm:"not null;index"`
	Organization   Organization `gorm:"foreignKey:OrganizationID"`
	UserID         int          `gorm:"not null;index"`
	User           User         `gorm:"foreignKey:UserID"`
	CreatedAt      time.Time    `gorm:"autoCreateTime"`
	UpdatedAt      time.Time    `gorm:"autoUpdateTime"`
}

type OrganizationType string

const (
	IE  OrganizationType = "IE"
	LLC OrganizationType = "LLC"
	JSC OrganizationType = "JSC"
)
