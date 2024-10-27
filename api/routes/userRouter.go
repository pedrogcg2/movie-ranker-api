package routes

import (
	"encoding/json"
	"movie-api/api/contracts"
	"movie-api/api/controller"
	"net/http"

	"gorm.io/gorm"
)

type userRouter struct {
	Db gorm.DB
}

func UserRoutes(db gorm.DB) {
	ur := &userRouter{Db: db}
	http.HandleFunc("GET /users", ur.GetAllUsers)
	http.HandleFunc("POST /user", ur.CreateUser)
}

func (ur *userRouter) GetAllUsers(w http.ResponseWriter, r *http.Request) {
	ctl := controller.NewUserController(ur.Db)

	users, error := ctl.GetUsers()
	if error != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(error.Error()))
		return
	}

	jsonResponse, error := json.Marshal(users)

	w.WriteHeader(http.StatusOK)
	w.Write(jsonResponse)
}

func (ur *userRouter) CreateUser(w http.ResponseWriter, r *http.Request) {
	ctl := controller.NewUserController(ur.Db)

	var dto contracts.UserDto

	error := json.NewDecoder(r.Body).Decode(&dto)

	if error != nil {
		w.WriteHeader(http.StatusUnprocessableEntity)
		w.Write([]byte(error.Error()))
		return
	}

	user, error := ctl.CreateNewUser(dto)

	if error != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(error.Error()))
		return
	}

	jsonResponse, _ := json.Marshal(user)
	w.WriteHeader(http.StatusCreated)
	w.Write(jsonResponse)
}
