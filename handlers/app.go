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

	dbUser := os.Getenv("DB_USER")
	dbPassword := os.Getenv("DB_PASSWD")
	dbAddr := os.Getenv("DB_ADDR")
	dbPort := os.Getenv("DB_PORT")
	dbName := os.Getenv("DB_NAME")
	serverPort := os.Getenv("SERVER_PORT")

	dbURL := fmt.Sprintf("postgres://%s:%s@%s:%s/%s", dbUser, dbPassword, dbAddr, dbPort, dbName)
	db, err := gorm.Open(postgres.Open(dbURL), &gorm.Config{})
	fmt.Println(db, err)

	//Repository
	vehiclesRepositoryDB := domain.NewVehiclesRepostoryDB(db)

	//service
	vehiclesServiceDB := service.NewVehiclesService(&vehiclesRepositoryDB)

	//handler
	vehicleHandler := NewVehiclesHandler(vehiclesServiceDB)
	router := gin.Default()
	router.GET("/vehicles", vehicleHandler.GetAllVehicles)

	routerRun := fmt.Sprintf(":%s", serverPort)
	router.Run(routerRun)

}
