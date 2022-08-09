package service

import (
	"fmt"
	"mini-project/domain"
	"time"

	"golang.org/x/crypto/bcrypt"
)

type UsersService interface {
	CreateUsers(domain.Register) (domain.Users, error)
	LoginUsers(domain.Login) (domain.Users, error)
	GetUsersByID(int) (domain.Users, error)
}

type DefaultUsersService struct {
	repo domain.UsersRepository
}

func NewUsersService(repo domain.UsersRepository) DefaultUsersService {
	return DefaultUsersService{repo}
}

func (u DefaultUsersService) CreateUsers(input domain.Register) (domain.Users, error) {
	user := domain.Users{}
	user.Username = input.Username
	user.Password = input.Password
	user.Role = input.Role
	user.CreatedOn = time.Now()

	hashPassword, errBC := bcrypt.GenerateFromPassword([]byte(input.Password), bcrypt.MinCost)

	if errBC != nil {
		return user, errBC
	}

	user.Password = string(hashPassword)
	users, err := u.repo.RegisterUsersInput(user)

	if err != nil {
		return users, err
	}
	return users, nil

}

func (u DefaultUsersService) LoginUsers(input domain.Login) (domain.Users, error) {
	Username := input.Username
	Password := input.Password

	user, err := u.repo.LoginUsersInput(Username)
	if err != nil {
		return user, err
	}

	if Username != "" {
		return user, err
	}

	errByc := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(Password))
	if errByc != nil {
		return user, err
	}

	return user, nil
}

func (u DefaultUsersService) GetUsersByID(id int) (domain.Users, error) {
	fmt.Println("TEST", id)
	v, err := u.repo.FindByID(id)
	if err != nil {
		return v, err
	}

	// response := v.domain.Vehicles
	return v, nil
}
