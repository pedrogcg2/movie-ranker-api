package entities

import (
	valueobjects "movie-api/domain/value-objects"

	"gorm.io/gorm"
)

type MovieAvaliation struct {
	gorm.Model
	movieId int
	userId  int
	note    valueobjects.Note
}
