package entities

import "gorm.io/gorm"

type User struct {
	gorm.Model
	name     string
	image    string
	password string
}
