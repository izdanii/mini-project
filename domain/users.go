package domain

import (
	"time"
)

type Users struct {
	Username  string    `json:"username" gorm:"username"`
	Password  string    `json:"password" gorm:"column:password"`
	Role      string    `json:"role" gorm:"column:role"`
	CreatedOn time.Time `json:"created_on" gorm:"column:created_on"`
}

type Login struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type Register struct {
	Username  string    `json:"username"`
	Password  string    `json:"password"`
	Role      string    `json:"role"`
	CreatedOn time.Time `json:"created_on"`
}

type UsersRepository interface {
	RegisterUsersInput(Users) (Users, error)
	LoginUsersInput(string) (Users, error)
}
