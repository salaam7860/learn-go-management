package store

import (
	"basic_management/learn-go-management/model"
	"basic_management/learn-go-management/util"
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Connect struct {
	DB *gorm.DB
}

func (c *Connect) NewStore() error {
	dsn := "host=192.168.100.97 user=postgres password=8603 dbname=postgres port=5432 sslmode=disable"

	util.Log(model.LogLevelInfo, model.StorePackage, model.NewStore, "CONNECTION ESTABLISHING", nil)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		util.Log(model.LogLevelError, model.StorePackage, model.NewStore, "ERROR WHILE ESTABLISHING CONNECTION", err)
		return err

	} else {
		c.DB = db
	}

	err = db.AutoMigrate(
		model.User{},
	)
	if err != nil {
		util.Log(model.LogLevelError, model.StorePackage, model.NewStore, "ERROR WHILE AUTO_MIGRATION", err)
		return err
	}
	fmt.Printf("DATABASE: %v\n", db)
	return nil
}

type StoreOperations interface {
	NewStore() error
	CreateUser(user *model.User) error
}
