package handlers

import (
	"mini-project/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

type VehiclesHandlers struct {
	service service.VehiclesService
}

func (vh *VehiclesHandlers) GetAllVehicles(g *gin.Context) {
	vehicles, err := vh.service.GetAllVehicles()

	if err != nil {
		g.JSON(http.StatusBadRequest, nil)
	}

	g.JSON(http.StatusOK, vehicles)

}

func NewVehiclesHandler(s service.VehiclesService) *VehiclesHandlers {
	return &VehiclesHandlers{s}

}
