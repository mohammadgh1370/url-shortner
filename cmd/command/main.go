package main

import (
	"flag"
	"fmt"
	"github.com/go-co-op/gocron"
	"github.com/mohammadgh1370/url-shortner/internal/database"
	"github.com/mohammadgh1370/url-shortner/internal/model"
	"github.com/mohammadgh1370/url-shortner/internal/repository/mysql"
	"math/rand"
	"os"
	"strconv"
	"time"
)

func main() {
	if len(os.Args) > 1 {
		commands()
		return
	}
}

func commands() {
	flag.NewFlagSet("migrate", flag.ExitOnError)

	switch os.Args[1] {
	case "migrate":
		fmt.Println("migrate")
		database.Migrate(os.Args[1])
	case "migrate:rollback":
		fmt.Println("migrate:rollback")
		database.Migrate(os.Args[1])
	case "schedule:run":
		fmt.Println("schedule run")
		schedule()
	default:
		fmt.Println("command invalid")
	}
}

func schedule() {
	scheduler := gocron.NewScheduler(time.UTC)

	_, err := scheduler.Every(10).Seconds().Do(task)
	if err != nil {
		fmt.Println("Error scheduling task:", err)
		return
	}

	scheduler.StartBlocking()
}

func task() {
	db := database.ConnectDB()

	userRepo := mysql.NewMysqlUserRepo(db)

	id := strconv.Itoa(rand.Int())
	now := time.Now()
	newUser := model.User{
		FirstName: "mohammad",
		LastName:  "ghorbani",
		Username:  "mohammad" + id,
		Password:  "123456",
		CreatedAt: now,
		UpdatedAt: now,
	}

	userRepo.Create(&newUser)
}
