package domain

import "github.com/vnniciusg/microservice-auth-api/internal/dto"

type User struct {
	ID        int    `json:"id"`
	FirstName string `json:"firstname"`
	LastName  string `json:"lastname"`
	Email     string `json:"email"`
	Password  string `json:"password"`
}

type UserService interface {
	Create()
	Fetch()
	Delete()
	FindByEmail()
}

type UserRepository interface {
	Create(*dto.CreateUserRequest) (*dto.CreateUserResponse, error)
	Fetch() ([]*dto.FetchUser, error)
	Delete(userRequest *dto.DeleteUserRequest) error
	FindByEmail(userRequest *dto.FindByEmailRequest) (*dto.FindByEmailResponse, error)
}

type UserUseCase interface {
	Create(*dto.CreateUserRequest) (*dto.CreateUserResponse, error)
	Fetch() ([]*dto.FetchUser, error)
	Delete(userRequest *dto.DeleteUserRequest) error
	FindByEmail(userRequest *dto.FindByEmailRequest) (*dto.FindByEmailResponse, error)
}
