package handlers

import (
	"fmt"
	"mini-project/domain"
	"mini-project/service"
	"net/http"
	"strconv"

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

func (vh *VehiclesHandlers) GetVehiclesByID(g *gin.Context) {
	vehicleId := g.Param("vehicle_id")
	vehicle, err := vh.service.GetVehiclesByID(vehicleId)
	if err != nil {
		errMessage := fmt.Sprintf("Your ID %v error!", vehicleId)
		g.JSON(http.StatusBadRequest, errMessage)
		return
	} else {
		g.JSON(http.StatusOK, vehicle)
	}
}

func (vh *VehiclesHandlers) DeleteVehiclesByID(g *gin.Context) {
	vehicleId := g.Param("vehicle_id")
	vehicle, err := vh.service.DeleteVehiclesByID(vehicleId)
	if err != nil {
		errMessage := fmt.Sprintf("Delete unsuccessfull %v error!", vehicleId)
		g.JSON(http.StatusBadRequest, errMessage)
		return
	} else {
		g.JSON(http.StatusOK, vehicle)
	}
}

func (vh *VehiclesHandlers) CreateVehiclesByID(g *gin.Context) {

	var input domain.InputVehicles
	err := g.ShouldBindJSON(&input)
	if err != nil {
		errMessage := "Create Vehicles error!"
		g.JSON(http.StatusBadRequest, errMessage)
		return
	} else {
		vehicle, err := vh.service.CreateVehiclesByID(input)
		if err != nil {
			errMessage := "Create Vehicle error!"
			g.JSON(http.StatusBadRequest, errMessage)
			return
		}
		g.JSON(http.StatusCreated, vehicle)
	}
	// vehicleId := g.Param("vehicle_id")

}

func (vh *VehiclesHandlers) UpdateVehiclesByID(g *gin.Context) {
	vehicleId := g.Param("vehicle_id")
	id, _ := strconv.Atoi(vehicleId)
	var input domain.UpdateVehicles
	err := g.ShouldBindJSON(&input)
	if err != nil {
		errMessage := "Update Vehicles error!"
		g.JSON(http.StatusBadRequest, errMessage)
		return
	} else {
		vehicle, err := vh.service.UpdateVehiclesByID(input, id)
		if err != nil {
			errMessage := "Update Vehicle error!"
			g.JSON(http.StatusBadRequest, errMessage)
			return
		}
		g.JSON(http.StatusCreated, vehicle)
	}
	// vehicleId := g.Param("vehicle_id")

}

func NewVehiclesHandler(s service.VehiclesService) *VehiclesHandlers {
	return &VehiclesHandlers{s}

}
