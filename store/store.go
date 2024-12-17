package store

import (
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Connect struct {
	DB *gorm.DB
}

func (c *Connect) NewStore() error {
	dsn := "host=192.168.100.97 user=postgres password=8603 dbname=postgres port=5432 sslmode=disable"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		return err

	} else {
		c.DB = db
	}
	fmt.Printf("DATABASE: %v\n", db)
	return nil
}

type StoreOperations interface {
	NewStore() error
}
