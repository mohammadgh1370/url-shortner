package main

import (
	"fmt"
	"github.com/mohammadgh1370/url-shortner/internal/database"
)

func main() {
	db := database.ConnectDB()

	database.RunAutoMigrations(db)

	fmt.Println("migrate models")
}
