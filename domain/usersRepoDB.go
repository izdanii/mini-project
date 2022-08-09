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
	err := u.db.Create(&users).Error

	if err != nil {
		fmt.Println("Error to registration user")
		return users, nil
	}
	return users, nil
}

func (u UsersRepositoryDB) LoginUsersInput(username string) (Users, error) {
	var users Users
	err := u.db.Where("username = ? ", username).Find(&users).Error
	if err != nil {
		fmt.Println("Error to login")
		return users, err
	}
	fmt.Println("users!!!!!", users)
	return users, nil
}

func (u UsersRepositoryDB) FindByID(id int) (Users, error) {
	var users Users
	fmt.Println("id", id)
	err := u.db.First(&users, "user_id = ?", id).Error

	if err != nil {
		fmt.Println("Error to login")
		return users, err
	}
	return users, nil
}
