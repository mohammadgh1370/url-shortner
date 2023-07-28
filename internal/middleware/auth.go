package middleware

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
	utils "github.com/mohammadgh1370/url-shortner/internal/util"
	"strings"
)

func NewAuthMiddleware(secret string) fiber.Handler {
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

			return []byte(secret), nil
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

		ctx.Locals("identifier", claims["identifier"])

		return ctx.Next()
	}
}
