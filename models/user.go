package models

import (
	"github.com/jinzhu/gorm"
)

// User is a Gorm model
type User struct {
	gorm.Model
	Name  string
	Todos *[]Todo
}
