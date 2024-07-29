package utils

import (
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"messaggio/pkg/model"
	"time"
)

func NewPostgresDB(cfg *Config) (*gorm.DB, error) {
	var db *gorm.DB
	var err error

	dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
		cfg.Host, cfg.Port, cfg.Username, cfg.Password, cfg.DBName, cfg.SSLMode)

	for i := 0; i < 100; i++ {
		db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
		if err == nil {
			break
		}
		time.Sleep(2 * time.Second)
	}

	if err != nil {
		return nil, fmt.Errorf("unable to connect to the database: %v", err)
	}

	err = db.AutoMigrate(
		&model.Message{},
	)
	if err != nil {
		return nil, err
	}

	fmt.Println("Connected to the database successfully!")
	return db, nil
}
