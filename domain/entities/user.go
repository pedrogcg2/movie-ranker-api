package entities

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Name     string
	Image    string
	Password string
}
