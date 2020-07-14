package database

import (
	"go-starter-kit/models"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

var db *gorm.DB

// DB is getting gorm database instance
func DB() *gorm.DB {
	return db
}

// Connect is starting to mongodb connection
func Connect() *gorm.DB {
	connection, err := gorm.Open("sqlite3", "test.db")
	if err != nil {
		panic("failed to connect database")
	}
	db = connection

	// Migrate the schema
	db.AutoMigrate(&models.Todo{})

	var todo *models.Todo = &models.Todo{
		Title: "Bugün alışveriş yapacağım.",
	}

	db.Create(&todo)

	return connection
}

// Close is closing db connection
func Close() (err error) {
	return DB().Close()
}
