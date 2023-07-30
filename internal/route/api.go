package route

import (
	"github.com/gofiber/fiber/v2"
	"github.com/mohammadgh1370/url-shortner/internal/controller"
	"github.com/mohammadgh1370/url-shortner/internal/middleware"
	"github.com/mohammadgh1370/url-shortner/internal/repository/mysql"
	"gorm.io/gorm"
)

func InitRouts(app *fiber.App, db *gorm.DB) {
	apiRoute := app.Group("/api")

	authRoute := apiRoute.Group("/auth")
	authMiddleware := middleware.NewAuthMiddleware(db)

	userRepo := mysql.NewMysqlUserRepo(db)
	authController := controller.NewAuthController(userRepo)

	authRoute.Post("/register", authController.Register)
	authRoute.Post("/login", authController.Login)
	authRoute.Get("/me", authMiddleware, authController.Me)

	linkRoute := apiRoute.Group("/link")
	linkRepo := mysql.NewMysqlLinkRepo(db)
	linkController := controller.NewLinkController(linkRepo)
	linkRoute.Post("/", authMiddleware, linkController.Store)
	linkRoute.Get("/", authMiddleware, linkController.Index)
	linkRoute.Delete("/:id", authMiddleware, linkController.Destroy)

	viewRepo := mysql.NewMysqlViewRepo(db)
	publicController := controller.NewPublicController(linkRepo, viewRepo)
	app.Get("/:hash", publicController.Redirect)
}
