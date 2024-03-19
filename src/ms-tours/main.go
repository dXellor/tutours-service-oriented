package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	tourrepository "tutours/soa/ms-tours/dataservice/tourRepository"
	"tutours/soa/ms-tours/handler"
	tourservice "tutours/soa/ms-tours/usecase/tourService"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

func main() {
	loadConfig()

	database := initDB()
	populateDB(database)

	tourRepository := tourrepository.TourRepository{}
	tourRepository.Init(database)

	tourService := tourservice.TourService{}
	tourService.Init(&tourRepository)

	tourhandler := handler.TourHandler{}

	router := tourhandler.InitRouter(&tourService)

	fmt.Println("Tours micro-service running")
	http.ListenAndServe(":7007", router)

}

func loadConfig() {
	envErr := godotenv.Load("config/.env")

	if envErr != nil {
		log.Fatalf(envErr.Error())
	}
}

func initDB() *gorm.DB {

	dbType := os.Getenv("DATABASE_TYPE")
	dbUser := os.Getenv("DATABASE_USER")
	dbSecret := os.Getenv("DATABASE_SECRET")
	dbHost := os.Getenv("DATABASE_HOST")
	dbPort := os.Getenv("DATABASE_PORT")
	dbName := os.Getenv("DATABASE_NAME")
	connectionUrl := fmt.Sprintf("%s://%s:%s@%s:%s/%s", dbType, dbUser, dbSecret, dbHost, dbPort, dbName)
	database, databaseErr := gorm.Open(postgres.Open(connectionUrl), &gorm.Config{NamingStrategy: schema.NamingStrategy{
		NoLowerCase: true,
	}})
	if databaseErr != nil {
		log.Fatalf(databaseErr.Error())
		return nil
	}
	return database
}

func populateDB(database *gorm.DB) {
	c, ioErr := os.ReadFile("script/tours.sql")
	if ioErr != nil {
		log.Fatalf(ioErr.Error())
	}
	sql := string(c)

	database.Exec(sql)
}
