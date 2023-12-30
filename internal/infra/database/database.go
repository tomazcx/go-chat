package database

import (
	"log"

	"github.com/tomazcx/go-chat-app/internal/entity"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var db *gorm.DB

func GetDBInstance() *gorm.DB {
	return db
}

func InitializaDB() error {
	var err error
	db, err = gorm.Open(sqlite.Open("chat.db"), &gorm.Config{})
	if err != nil {
		return err
	}

	log.Println("Connected to SQLite database.")
	db.AutoMigrate(&entity.User{})
	log.Println("Ran migrations.")

	return nil
}
