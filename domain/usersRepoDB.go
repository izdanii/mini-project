package domain

import (
	"fmt"

	"gorm.io/gorm"
)

type UsersRepositoryDB struct {
	db *gorm.DB
}

func NewUsersRepositoryDB(client *gorm.DB) UsersRepositoryDB {
	return UsersRepositoryDB{client}
}

func (u UsersRepositoryDB) RegisterUsersInput(users Users) (Users, error) {
	err := u.db.Create(&users)

	if err != nil {
		fmt.Println("Error to registration user")
		return users, err.Error
	}
	return users, nil
}

func (u UsersRepositoryDB) LoginUsersInput(username string) (Users, error) {
	var users Users
	err := u.db.First(&users, "username= ?", username)

	if err != nil {
		fmt.Println("Error to login")
		return users, err.Error
	}
	return users, nil
}
