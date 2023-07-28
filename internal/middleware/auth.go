package middleware

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
	"github.com/mohammadgh1370/url-shortner/internal/config"
	"github.com/mohammadgh1370/url-shortner/internal/model"
	"github.com/mohammadgh1370/url-shortner/internal/repository/mysql"
	utils "github.com/mohammadgh1370/url-shortner/internal/util"
	"gorm.io/gorm"
	"strings"
)

func NewAuthMiddleware(db *gorm.DB) fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		var tokenString string
		authorization := ctx.Get("Authorization")

		if strings.HasPrefix(authorization, "Bearer ") {
			tokenString = strings.TrimPrefix(authorization, "Bearer ")
		}

		if tokenString == "" {
			response := utils.Response{Message: "You are not logged in"}
			return ctx.Status(fiber.StatusUnauthorized).JSON(response)
		}

		tokenByte, err := jwt.Parse(tokenString, func(jwtToken *jwt.Token) (interface{}, error) {
			if _, ok := jwtToken.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("unexpected signing method: %s", jwtToken.Header["alg"])
			}

			return []byte(config.JWT_SECRET_KEY), nil
		})

		if err != nil {
			response := utils.Response{Message: fmt.Sprintf("invalidate token: %v", err)}
			return ctx.Status(fiber.StatusUnauthorized).JSON(response)
		}

		claims, ok := tokenByte.Claims.(jwt.MapClaims)
		if !ok || !tokenByte.Valid {
			response := utils.Response{Message: "invalid token claim"}
			return ctx.Status(fiber.StatusUnauthorized).JSON(response)

		}

		var userRepo = mysql.NewMysqlUserRepo(db)
		var user model.User
		if err := userRepo.Find(&user, model.User{Username: claims["identifier"].(string)}); err != nil {
			return ctx.Status(fiber.StatusNotFound).JSON(fiber.Map{
				"message": "User not exist.",
			})
		}

		ctx.Locals("user", user)

		return ctx.Next()
	}
}
