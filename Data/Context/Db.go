package context

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	"avitoTest/data/entities"
)

// ConnectDB connects to PostgreSQL database and performs migrations
func ConnectDB(dsn string) (*gorm.DB, error) {
	// Opening a connection to the database
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	// Automatic creation of tables based on entities
	err = db.AutoMigrate(&entities.User{}, &entities.Organization{}, &entities.OrganizationResponsible{}, &entities.Tender{}, &entities.Bid{})
	if err != nil {
		return nil, err
	}

	return db, nil
}
