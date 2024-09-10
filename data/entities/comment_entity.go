package entities

import "time"

// Comment represents a comment made by a user regarding their work on a tender.
type Comment struct {
	ID          int       `gorm:"primaryKey"`
	UserID      int       `gorm:"not null"`
	User        User      `gorm:"foreignKey:UserID"`
	CompanyName string    `gorm:"not null;size:100"`
	TenderName  string    `gorm:"not null;size:100"`
	Description string    `gorm:"size:255"`
	ServiceType string    `gorm:"size:100"`
	Content     string    `gorm:"not null;size:500"`
	CreatedAt   time.Time `gorm:"autoCreateTime"`
	UpdatedAt   time.Time `gorm:"autoUpdateTime"`
}
