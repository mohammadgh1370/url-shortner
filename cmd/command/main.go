package main

import (
	"flag"
	"fmt"
	"github.com/go-co-op/gocron"
	"github.com/mohammadgh1370/url-shortner/internal/config"
	"github.com/mohammadgh1370/url-shortner/internal/database"
	"github.com/mohammadgh1370/url-shortner/internal/model"
	"github.com/mohammadgh1370/url-shortner/internal/repository/mysql"
	"os"
	"strconv"
	"time"
)

func init() {
	tehranTimeZone, err := time.LoadLocation(config.TIME_ZONE)
	if err != nil {
		panic(err)
	}

	time.Local = tehranTimeZone
}

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
	s := gocron.NewScheduler(time.UTC)

	_, err := s.Every(60).Seconds().Do(removeLinkNotUseYearAgo)
	if err != nil {
		fmt.Println(err.Error())
	}

	s.StartBlocking()
}

func removeLinkNotUseYearAgo() {
	db := database.ConnectDB()

	linkRepo := mysql.NewMysqlLinkRepo(db)
	viewRepo := mysql.NewMysqlViewRepo(db)

	var links []model.Link

	yearAgo := time.Now().AddDate(-1, 0, 0)
	linkRepo.Find(&links, "created_at < '"+yearAgo.Format(time.RFC3339)+"'")

	for _, link := range links {
		var count int64
		viewRepo.Count(model.View{}, "created_at > '"+yearAgo.Format(time.RFC3339)+"' and link_id = "+strconv.Itoa(int(link.Id)), &count)
		if count == 0 {
			linkRepo.Delete(model.Link{}, link)
		}
	}
}
