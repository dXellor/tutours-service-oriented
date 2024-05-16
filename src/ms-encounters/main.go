package main

import (
	"fmt"
	"log"
	"net"
	"os"

	encounterrepository "tutours/soa/ms-encounters/dataservice/encounterRepository"
	"tutours/soa/ms-encounters/handler"
	encounterservice "tutours/soa/ms-encounters/usecase/encounterService"

	ms_encounters "tutours/soa/ms-encounters/proto"

	"github.com/joho/godotenv"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

func main() {

	loadConfig();
	database := initDB()
	populateDB(database)

	encounterRepository := encounterrepository.EncounterRepository{}
	encounterRepository.Init(database)
	encounterCompletionRepository := encounterrepository.EncounterCompletionRepository{}
	encounterCompletionRepository.Init(database)

	encounterService := encounterservice.EncounterService{}
	encounterService.Init(&encounterRepository)
	encounterStatsService := encounterservice.EncounterStatsService{}
	encounterStatsService.Init(&encounterCompletionRepository)
	encounterCompletionService := encounterservice.EncounterComlpetionService{}
	encounterCompletionService.Init(&encounterCompletionRepository)
	encounterHandler := handler.EncounterHandler{}

	// router := encounterHandler.InitRouter(&encounterService, &encounterStatsService, &encounterCompletionService)
	// fmt.Println("Encounters micro-service running")
	// http.ListenAndServe(":7007", router)

	lis, err := net.Listen("tcp", "localhost:8000")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	var opts []grpc.ServerOption
	grpcServer := grpc.NewServer(opts...)

	ms_encounters.RegisterEncountersServer(grpcServer, &encounterHandler)
	reflection.Register(grpcServer)
	grpcServer.Serve(lis)
}

func loadConfig() {
	envErr := godotenv.Load("config/.env")

	if envErr != nil {
		log.Fatalf(envErr.Error())
	}
}

/* not sure where to put this*/
func initDB() *gorm.DB {
	/* TODO: lazy to think of something easier: */
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
	c, ioErr := os.ReadFile("script/encounters.sql")
	if ioErr != nil {
		log.Fatalf(ioErr.Error())
	}
	sql := string(c)

	database.Exec(sql)
}
