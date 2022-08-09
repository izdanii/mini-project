package handlers

import (
	"fmt"
	"mini-project/auth"
	"mini-project/domain"
	"mini-project/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

type UsersHandlers struct {
	service     service.UsersService
	authService auth.Service
}

func (uh *UsersHandlers) CreateUsers(g *gin.Context) {
	var input domain.Register
	err := g.ShouldBindJSON(&input)
	users, _ := uh.service.CreateUsers(input)

	if err != nil {
		g.JSON(http.StatusBadRequest, nil)
		return
	}
	g.JSON(http.StatusOK, users)
}

func (uh *UsersHandlers) LoginUsers(g *gin.Context) {
	var input domain.Login
	err := g.ShouldBindJSON(&input)
	users, _ := uh.service.LoginUsers(input)

	if err != nil {
		g.JSON(http.StatusBadRequest, nil)
		return
	}
	fmt.Println("USERSID", users.ID)
	token, err := uh.authService.GenerateToken(users.ID)
	if err != nil {
		g.JSON(http.StatusBadRequest, err)
	}
	response := domain.FormatMemberDTO(users, token)

	g.JSON(http.StatusOK, response)
}

func NewUsersHandler(k service.UsersService, authService auth.Service) *UsersHandlers {
	return &UsersHandlers{k, authService}

}
