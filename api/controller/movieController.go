package controller

import (
	"movie-api/api/contracts"
	"movie-api/domain/entities"
	"movie-api/repository"

	"gorm.io/gorm"
)

type MovieController struct {
	Rp *repository.Repository[entities.Movie]
}

func NewMovieController(db gorm.DB) *MovieController {
	rp := repository.NewRepository[entities.Movie](&db)

	return &MovieController{Rp: rp}
}

func (ctl *MovieController) GetMovies(page int, pageSize int) (*[]entities.Movie, error) {
	movies, error := ctl.Rp.GetMany(pageSize, (page-1)*pageSize)

	if error != nil {
		return nil, error
	}

	return movies, nil
}

func (ctl *MovieController) GetMovie(id int) (*entities.Movie, error) {
	movie, error := ctl.Rp.GetById(id)

	if error != nil {
		return nil, error
	}

	return movie, nil
}

func (ctl *MovieController) InsertMovies(moviesDto *[]contracts.MovieDto) (*[]entities.Movie, error) {
	var movies []entities.Movie

	for _, movie := range *moviesDto {
		newMovie := entities.Movie{
			Name:        movie.Name,
			Externalid:  movie.Externalid,
			Description: movie.Description,
			Image:       movie.Image,
		}

		movies = append(movies, newMovie)
	}

	result, error := ctl.Rp.CreateBatch(movies)

	if error != nil {
		return nil, error
	}
	return result, nil
}
