package entities

import (
	"time"
)

type Organization struct {
	ID           int    `gorm:"primaryKey"`
	Name         string `gorm:"not null;size:100"`
	Description  string
	Type         OrganizationType `gorm:"type:organization_type"`
	CreatedAt    time.Time        `gorm:"autoCreateTime"`
	UpdatedAt    time.Time        `gorm:"autoUpdateTime"`
	Responsibles []User           `gorm:"many2many:organization_responsibles;"`
	Tenders      []Tender         `gorm:"foreignKey:OrganizationID"`
}

type OrganizationResponsible struct {
	ID             int          `gorm:"primaryKey"`
	OrganizationID int          `gorm:"not null;index"`
	Organization   Organization `gorm:"foreignKey:OrganizationID;constraint:OnDelete:CASCADE;"`
	UserID         int          `gorm:"not null;index"`
	User           User         `gorm:"foreignKey:UserID;constraint:OnDelete:CASCADE;"`
	CreatedAt      time.Time    `gorm:"autoCreateTime"`
	UpdatedAt      time.Time    `gorm:"autoUpdateTime"`
}

type OrganizationType string

const (
	IE  OrganizationType = "IE"
	LLC OrganizationType = "LLC"
	JSC OrganizationType = "JSC"
)
