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
	Responsibles []User           `gorm:"many2many:organization_responsibles;"` // связь многие-ко-многим через промежуточную таблицу
	Tenders      []Tender         `gorm:"foreignKey:OrganizationID"`
}

type OrganizationResponsible struct {
	ID             int          `gorm:"primaryKey"`
	OrganizationID int          `gorm:"not null;index"`                                         // Внешний ключ на организацию
	Organization   Organization `gorm:"foreignKey:OrganizationID;constraint:OnDelete:CASCADE;"` // Связь с организацией
	UserID         int          `gorm:"not null;index"`                                         // Внешний ключ на пользователя
	User           User         `gorm:"foreignKey:UserID;constraint:OnDelete:CASCADE;"`         // Связь с пользователем
	CreatedAt      time.Time    `gorm:"autoCreateTime"`
	UpdatedAt      time.Time    `gorm:"autoUpdateTime"`
}

type OrganizationType string

const (
	IE  OrganizationType = "IE"
	LLC OrganizationType = "LLC"
	JSC OrganizationType = "JSC"
)
