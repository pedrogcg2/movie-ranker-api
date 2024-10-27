package routes

import (
	"encoding/json"
	"movie-api/api/contracts"
	"movie-api/api/controller"
	"net/http"
	"strconv"

	"gorm.io/gorm"
)

type movieRouter struct {
	Db gorm.DB
}

func MovieRoutes(db gorm.DB) {
	mr := movieRouter{Db: db}
	http.HandleFunc("GET /movies", mr.GetMovies)
	http.HandleFunc("POST /movies", mr.InsertMovies)
}

func (mr *movieRouter) GetMovies(w http.ResponseWriter, r *http.Request) {
	ctl := controller.NewMovieController(mr.Db)
	pageQuery := r.URL.Query().Get("page")
	pageSizeQuery := r.URL.Query().Get("pageSize")

	page := -1
	pageSize := -1

	if pageQuery != "" && pageSizeQuery != "" {
		page, _ = strconv.Atoi(pageQuery)
		pageSize, _ = strconv.Atoi(pageSizeQuery)
	}

	movies, error := ctl.GetMovies(page, pageSize)

	if error != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(error.Error()))
		return
	}

	jsonResponse, _ := json.Marshal(movies)

	w.WriteHeader(http.StatusOK)
	w.Write(jsonResponse)
}

func (mr *movieRouter) InsertMovies(w http.ResponseWriter, r *http.Request) {
	ctl := controller.NewMovieController(mr.Db)

	var moviesDto *[]contracts.MovieDto

	error := json.NewDecoder(r.Body).Decode(moviesDto)

	if error != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(error.Error()))
		return
	}

	movies, error := ctl.InsertMovies(moviesDto)

	if error != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(error.Error()))
		return
	}

	jsonResponse, error := json.Marshal(movies)

	if error != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(error.Error()))
		return
	}

	w.WriteHeader(http.StatusCreated)
	w.Write(jsonResponse)
}
