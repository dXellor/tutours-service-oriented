package main

/* uwuw should go to cmd */

import (
	"fmt"
	"log"
	"ms-lil-service/adapter"
	"ms-lil-service/dataservice/repos"
	"ms-lil-service/model"
	usecases "ms-lil-service/usecase"
	"ms-lil-service/usecase/channel"
	"ms-lil-service/usecase/routine"
	service "ms-lil-service/usecase/services"
	"net/http"
	"os"
	"sync"
	"time"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {
	/* STRUCT AND INTERFACES */
	fmt.Println("helo, this is start")
	// l := model.Lil{
	// 	Name: "text",
	// 	Nickname: "aa",
	// 	Age: 1,
	// }
	var l = model.Lil{
		Name: "text",
		Nickname: "aa",
		Age: 1,
	}
	fmt.Println(l)

	l.Nickname = model.GenerateNickname(l)
	fmt.Println(l)

	l.Nickname = "aa"
	l.GenerateNicknameButBetter()
	fmt.Println(l)

	l.Nickname = "aa"
	l.GenerateNicknameButEvenBetter()
	fmt.Println(l)

	fmt.Println(model.Structure.AgePlus(l))
	fmt.Println(model.Structure.AgeMinus(l))

	/* MAPS, ARRAYS AND SLICES*/
	var lilArray [2]string
	lilArray[1] = "helo"
	fmt.Println(lilArray)

	var lilSlice = []string{}
	// lilSlice[0] = "helo"
	lilSlice = append(lilSlice, "helo")
	lilSlice = append(lilSlice, "helo2")
	fmt.Println(lilSlice)

	lilFullArray := [4]int{1,2,3,4}
	fmt.Println(lilFullArray)

	/* ROUTINES */
	var wg sync.WaitGroup
	wg.Add(1)
	rs := new(routine.RoutineService)
	go usecases.IRoutine.Init(rs, &wg)
	fmt.Println("Let's see where this goes")
	wg.Wait()
	fmt.Println("Services dead")

	/* CHANNELS */
	mainCh := make(chan string)

	receiver := new(channel.ReceiverService)
	sender := new(channel.SenderService)

	go usecases.IReceiver.Init(receiver, mainCh)
	go usecases.ISender.Init(sender, mainCh)

	time.Sleep(1 * time.Second)
	usecases.ISender.Send(sender)

	/* ADDITIONAL: ENV FILE - WILL PROBABLY LATER BE MOVED TO SOME HIGHER LEVEL SCRIPT/FILE*/
	envErr := godotenv.Load("config/.env")

	if envErr != nil {
		log.Fatalf("Error loading .env file") // log log
	}
	fmt.Println(os.Getenv("DATABASE_TYPE")) // woooo

	/* DATABASE THINGS */
	database := initDB()
	populateDB(database)

	repo := repos.RandomRepo{DatabaseConnection: database}
	service := service.RandomService{Rr: &repo}
	adapter := adapter.LilAdapter{Rs: &service}

	// /* HTTPS THINGS */
	router := adapter.InitRouter()
	fmt.Println(http.ListenAndServe(":7007", router))
}

/* not sure where to put this*/
func initDB() *gorm.DB {
	fmt.Println(os.Getenv("DATABASE_PORT")) // moree wooo
	/* TODO: lazy to think of something easier: */
	dbType := os.Getenv("DATABASE_TYPE")
	dbUser := os.Getenv("DATABASE_USER")
	dbSecret := os.Getenv("DATABASE_SECRET")
	dbHost := os.Getenv("DATABASE_HOST")
	dbPort := os.Getenv("DATABASE_PORT")
	dbName := os.Getenv("DATABASE_NAME")
	// mby can be moved to config instead of connectionURL:
	connectionUrl := fmt.Sprintf("%s://%s:%s@%s:%s/%s", dbType, dbUser, dbSecret, dbHost, dbPort, dbName)
	//connectionUrl := "postgresql://postgres:super@localhost:5432/testDB"
	database, databaseErr := gorm.Open(postgres.Open(connectionUrl), &gorm.Config{})
	if databaseErr != nil {
		fmt.Println(databaseErr)
		return nil
	}

	//database.AutoMigrate(&model.Lil{})
	return database
}

func populateDB(database *gorm.DB) {
	c, ioErr := os.ReadFile("script/lils.sql")
	if ioErr != nil {
		fmt.Println(ioErr)
	}
	sql := string(c)

	database.Exec(sql)
}

/* dependancy injection:  https://blog.matthiasbruns.com/golang-the-ultimate-guide-to-dependency-injection */
/* lazy to do this now, shoudn't be hard but aaaa */
func initRepos() {

}

func initServices() {

}

func initAdapters() {
	
}