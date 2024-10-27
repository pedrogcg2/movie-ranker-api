package entities

import (
	"gorm.io/gorm"
)

type Movie struct {
	gorm.Model
	Name        string
	Description string
	Image       string
	Externalid  int
}
