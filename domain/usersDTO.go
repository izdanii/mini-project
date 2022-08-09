package domain

import (
	"time"
)

type UsersDTO struct {
	Username  string    `json:"username"`
	Password  string    `json:"password"`
	Role      string    `json:"role"`
	Token     string    `json:"token"`
	CreatedOn time.Time `json:"created_on"`
}

func FormatMemberDTO(user Users, token string) UsersDTO {
	usersDTO := UsersDTO{
		Username:  user.Username,
		Password:  user.Password,
		Role:      user.Role,
		Token:     token,
		CreatedOn: user.CreatedOn,
	}
	return usersDTO
}
