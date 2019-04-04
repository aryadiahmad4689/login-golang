package models

import (
	"github.com/jinzhu/gorm"
)

//Struct User
type User struct {
	gorm.Model
	Username string `gorm:"type:varchar(50)"`
	Password string `gorm:"type:text"`
	Nama     string `gorm:"type:varchar(100)"`
}
