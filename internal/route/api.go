package route

import (
	"github.com/gofiber/fiber/v2"
	"github.com/mohammadgh1370/url-shortner/internal/controller"
	"github.com/mohammadgh1370/url-shortner/internal/middleware"
	"github.com/mohammadgh1370/url-shortner/internal/repository/mysql"
	"gorm.io/gorm"
)

func InitRouts(app *fiber.App, db *gorm.DB) {

	userRepo := mysql.NewMysqlUserRepo(db)
	linkRepo := mysql.NewMysqlLinkRepo(db)
	viewRepo := mysql.NewMysqlViewRepo(db)

	authMiddleware := middleware.NewAuthMiddleware(userRepo)

	authController := controller.NewAuthController(userRepo)
	linkController := controller.NewLinkController(linkRepo)
	viewController := controller.NewViewController(linkRepo, viewRepo)
	publicController := controller.NewPublicController(linkRepo, viewRepo)

	apiRoute := app.Group("/api")

	authRoute := apiRoute.Group("/auth")
	authRoute.Post("/register", authController.Register)
	authRoute.Post("/login", authController.Login)
	authRoute.Get("/me", authMiddleware, authController.Me)

	linkRoute := apiRoute.Group("/link")
	linkRoute.Post("/", authMiddleware, linkController.Store)
	linkRoute.Get("/", authMiddleware, linkController.Index)
	linkRoute.Delete("/", authMiddleware, linkController.Destroy)

	viewRoute := apiRoute.Group("/view")
	viewRoute.Get("/", authMiddleware, viewController.Show)

	app.Get("/:hash", publicController.Redirect)
}
