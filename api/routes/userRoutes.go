package routes

import (
	"encoding/json"
	"log"
	"movie-api/api/controller"
	"net/http"

	"gorm.io/gorm"
)

type UserRouter struct {
	Db gorm.DB
}

func UserRoutes(db gorm.DB) {
	ur := &UserRouter{Db: db}
	http.HandleFunc("GET /users", ur.GetAllUsers)
	http.HandleFunc("POST /user", ur.CreateUser)
}

func (ur *UserRouter) GetAllUsers(w http.ResponseWriter, r *http.Request) {
	ctl := controller.NewUserController(ur.Db)

	users, error := ctl.GetUsers()
	if error != nil {
		w.WriteHeader(400)
		w.Write([]byte("Error on getting users"))
		return
	}

	jsonResponse, error := json.Marshal(users)

	w.WriteHeader(200)
	w.Write(jsonResponse)
}

func (ur *UserRouter) CreateUser(w http.ResponseWriter, r *http.Request) {
	ctl := controller.NewUserController(ur.Db)

	log.Print(r.Method)
	log.Print(r.Body)
	var dto controller.UserDto

	error := json.NewDecoder(r.Body).Decode(&dto)

	if error != nil {
		w.WriteHeader(400)
		w.Write([]byte(error.Error()))
		return
	}

	user, error := ctl.CreateNewUser(dto)

	if error != nil {
		w.WriteHeader(400)
		w.Write([]byte("Fail to create new user"))
		return
	}

	jsonResponse, _ := json.Marshal(user)
	w.WriteHeader(201)
	w.Write(jsonResponse)
}
