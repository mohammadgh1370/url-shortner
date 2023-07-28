package route

import (
	"github.com/gofiber/fiber/v2"
	"github.com/mohammadgh1370/url-shortner/internal/config"
	"github.com/mohammadgh1370/url-shortner/internal/controller"
	"github.com/mohammadgh1370/url-shortner/internal/middleware"
	"github.com/mohammadgh1370/url-shortner/internal/repository/mysql"
	"gorm.io/gorm"
)

func InitRouts(app *fiber.App, db *gorm.DB) {
	api := app.Group("/api")

	auth := api.Group("/auth")

	userRepo := mysql.NewMysqlUserRepo(db)
	controller := controller.NewAuthController(userRepo)

	authMiddleware := middleware.NewAuthMiddleware(config.JWT_SECRET_KEY)

	auth.Post("/register", controller.Register)
	auth.Post("/login", controller.Login)
	auth.Get("/me", authMiddleware, controller.Me)
}
