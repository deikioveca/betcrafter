package config

import (
	"fmt"
	"log"

	"github.com/deikioveca/betcrafter/internal/ticket/model"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func InitDB(cfg *Config) *gorm.DB {
	dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=%s", cfg.DBHost, cfg.DBPort, cfg.DBUser, cfg.DBPassword, cfg.DBName, cfg.DBSSLMode)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("failed to connect database: %v", err)
	}

	if err := db.AutoMigrate(&model.Ticket{}, &model.TicketMatch{}); err != nil {
		log.Fatalf("automigrate failed: %v", err)
	}

	log.Println("Database connected and migrated successfully")
	return db
}