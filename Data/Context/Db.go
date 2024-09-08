package context

import (
	"log"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	"avitoTest/data/entities"
)

func ConnectDB() *gorm.DB {
	dsn := os.Getenv("POSTGRES_CONN")
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("DB connection error: %v", err)
	}

	// Automatic creation of tables based on entities
	err = db.AutoMigrate(&entities.User{}, &entities.Organization{}, &entities.OrganizationResponsible{}, &entities.Tender{}, &entities.Bid{})
	if err != nil {
		log.Fatalf("Migration failed: %v", err)
	}

	return db
}
