package entities

import (
	"time"
)

type User struct {
	ID            int            `gorm:"primaryKey"`
	Username      string         `gorm:"unique;not null;size:50"`
	FirstName     string         `gorm:"size:50"`
	LastName      string         `gorm:"size:50"`
	CreatedAt     time.Time      `gorm:"autoCreateTime"`
	UpdatedAt     time.Time      `gorm:"autoUpdateTime"`
	Organizations []Organization `gorm:"many2many:organization_responsibles;"` // связь многие-ко-многим через промежуточную таблицу
	Tenders       []Tender       `gorm:"foreignKey:CreatorID"`
	Bids          []Bid          `gorm:"foreignKey:CreatorID"`
}
