package entities

import (
	valueobjects "movie-api/domain/value-objects"

	"gorm.io/gorm"
)

type MovieAvaliation struct {
	User User `gorm:"foreignKey:UserId"`
	gorm.Model
	Movie   Movie `gorm:"foreignKey:MovieId"`
	MovieId int
	UserId  int
	Note    valueobjects.Note
}
