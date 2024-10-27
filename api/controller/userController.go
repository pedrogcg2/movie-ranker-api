package controller

import (
	"movie-api/api/contracts"
	"movie-api/domain/entities"
	"movie-api/repository"

	"gorm.io/gorm"
)

type UserController struct {
	Rp *repository.Repository[entities.User]
}

func NewUserController(db gorm.DB) *UserController {
	repository := repository.NewRepository[entities.User](&db)
	return &UserController{Rp: repository}
}

func (controller *UserController) CreateNewUser(dto contracts.UserDto) (*entities.User, error) {
	user := &entities.User{
		Name:     dto.Name,
		Password: dto.Password,
		Image:    dto.Image,
	}

	user, error := controller.Rp.Create(user)
	if error != nil {
		return nil, error
	}

	return user, nil
}

func (controller *UserController) GetUsers() (*[]entities.User, error) {
	users, error := controller.Rp.GetAll()

	if error != nil {
		return nil, error
	}

	return users, nil
}
