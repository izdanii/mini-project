package handlers

import (
	"fmt"
	"mini-project/domain"
	"mini-project/service"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func Start() {
	err := godotenv.Load()
	if err != nil {
		// logger.Fatal("error loading .env file")
		fmt.Println("error loading .env file")
	}

	// logger.Info("load environment variables...")

	fmt.Println("load environment variables...")

	// dbUser := os.Getenv("DB_USER")
	// dbPassword := os.Getenv("DB_PASSWD")
	// dbAddr := os.Getenv("DB_ADDR")
	// dbPort := os.Getenv("DB_PORT")
	// dbName := os.Getenv("DB_NAME")
	serverPort := os.Getenv("SERVER_PORT")

	dbURL := "postgres://postgres:f41zd4n11@localhost:5432/miniproject"
	// dbURL := fmt.Sprintf("postgres://%s:%s@%s:%s/%s", dbUser, dbPassword, dbAddr, dbPort, dbName)
	fmt.Println("env", dbURL)
	db, err := gorm.Open(postgres.Open(dbURL), &gorm.Config{})
	fmt.Println(db, err)

	//Repository
	vehiclesRepositoryDB := domain.NewVehiclesRepostoryDB(db)
	usersRepositoryDB := domain.NewUsersRepositoryDB(db)

	//service
	vehiclesServiceDB := service.NewVehiclesService(&vehiclesRepositoryDB)
	usersServiceDB := service.NewUsersService(&usersRepositoryDB)

	//handler
	vehicleHandler := NewVehiclesHandler(vehiclesServiceDB)
	userHandler := NewUsersHandler(usersServiceDB)
	router := gin.Default()
	router.GET("/vehicles", vehicleHandler.GetAllVehicles)
	router.GET("/vehicles/:vehicle_id", vehicleHandler.GetVehiclesByID)
	router.DELETE("/vehicles/:vehicle_id", vehicleHandler.DeleteVehiclesByID)
	router.POST("/vehicles", vehicleHandler.CreateVehiclesByID)
	router.PUT("/vehicles/:vehicle_id", vehicleHandler.UpdateVehiclesByID)
	router.POST("/register", userHandler.CreateUsers)
	router.POST("/login", userHandler.LoginUsers)
	fmt.Println("server", serverPort)
	// routerRun := fmt.Sprintf(":%v", serverPort)
	router.Run(":8000")

}
