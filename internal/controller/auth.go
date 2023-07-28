package controller

import (
	"github.com/gofiber/fiber/v2"
	"github.com/mohammadgh1370/url-shortner/internal/model"
	"github.com/mohammadgh1370/url-shortner/internal/repository"
	"github.com/mohammadgh1370/url-shortner/internal/request"
	"github.com/mohammadgh1370/url-shortner/internal/util"
	"time"
)

type AuthController struct {
	userRepo repository.IUserRepo
}

func NewAuthController(userRepo repository.IUserRepo) AuthController {
	return AuthController{userRepo: userRepo}
}

func (c *AuthController) Register(ctx *fiber.Ctx) error {
	request := new(request.UserSignUpRequest)
	ctx.BodyParser(&request)

	if errors := util.Validate(request); errors != nil {
		response := util.ErrorResponse{Message: "wrong data", Errors: errors}
		return ctx.Status(fiber.StatusUnprocessableEntity).JSON(response)
	}

	var userExist model.User

	c.userRepo.Find(&userExist, model.User{Username: request.Username})

	if userExist.Username == request.Username {
		response := util.ErrorResponse{Message: "User with this username already exist."}
		return ctx.Status(fiber.StatusBadRequest).JSON(response)
	}

	hashedPassword, err := util.HashPassword(request.Password)
	if err != nil {
		response := util.ErrorResponse{Message: err.Error()}
		return ctx.Status(fiber.StatusInternalServerError).JSON(response)
	}

	now := time.Now()
	newUser := model.User{
		FirstName: request.FirstName,
		LastName:  request.LastName,
		Username:  request.Username,
		Password:  hashedPassword,
		CreatedAt: now,
		UpdatedAt: now,
	}

	if err := c.userRepo.Create(&newUser); err != nil {
		response := util.Response{Message: err.Error()}
		return ctx.Status(fiber.StatusBadRequest).JSON(response)
	}

	token, err := util.GenerateAccessToken(newUser.Username)
	if err != nil {
		response := util.Response{Message: err.Error()}
		return ctx.Status(fiber.StatusBadRequest).JSON(response)
	}

	response := util.Response{Message: "successful", Data: map[string]string{"token": token}}

	return ctx.JSON(response)
}

func (c *AuthController) Login(ctx *fiber.Ctx) error {
	request := new(request.UserSignInRequest)
	ctx.BodyParser(&request)

	var user model.User

	if err := c.userRepo.Find(&user, model.User{Username: request.Username}); err != nil {
		return ctx.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"message": "User not exist.",
		})
	}

	if err := util.VerifyPassword(user.Password, request.Password); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Invalid password",
		})
	}

	token, err := util.GenerateAccessToken(request.Username)
	if err != nil {
		response := util.Response{Message: err.Error()}
		return ctx.Status(fiber.StatusBadRequest).JSON(response)
	}

	response := util.Response{Message: "successful", Data: map[string]string{"token": token}}

	return ctx.JSON(response)
}

func (c *AuthController) Me(ctx *fiber.Ctx) error {
	var user model.User

	if err := c.userRepo.Find(&user, model.User{Username: ctx.Locals("identifier").(string)}); err != nil {
		return ctx.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"message": "User not exist.",
		})
	}

	response := util.Response{Message: "successful", Data: user}

	return ctx.JSON(response)
}
