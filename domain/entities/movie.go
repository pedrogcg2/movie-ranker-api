package entities

import (
	"gorm.io/gorm"
)

type Movie struct {
	gorm.Model
	name        string
	description string
	image       string
	externalId  int
}
