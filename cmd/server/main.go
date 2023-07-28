package main

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/mohammadgh1370/url-shortner/internal/config"
	"github.com/mohammadgh1370/url-shortner/internal/database"
	"github.com/mohammadgh1370/url-shortner/internal/route"
)

func main() {
	app := fiber.New()

	db := database.ConnectDB()

	route.InitRouts(app, db)

	address := fmt.Sprintf(":%s", config.APP_PORT)

	app.Listen(address)
}
