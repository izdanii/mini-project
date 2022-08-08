package handlers

import (
	"mini-project/domain"
	"mini-project/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

type UsersHandlers struct {
	service service.UsersService
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

	g.JSON(http.StatusOK, users)
}

func NewUsersHandler(k service.UsersService) *UsersHandlers {
	return &UsersHandlers{k}

}
