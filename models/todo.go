package models

import (
	"github.com/jinzhu/gorm"
)

// Todo is a Gorm model
type Todo struct {
	gorm.Model
	Title string
	User  *User
}
