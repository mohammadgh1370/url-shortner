package controller

import (
	"github.com/gofiber/fiber/v2"
	"github.com/mohammadgh1370/url-shortner/internal/model"
	"github.com/mohammadgh1370/url-shortner/internal/repository"
	"github.com/mohammadgh1370/url-shortner/internal/request"
	"github.com/mohammadgh1370/url-shortner/internal/util"
)

type linkController struct {
	linkRepo repository.ILinkRepo
}

func NewLinkController(linkRepo repository.ILinkRepo) linkController {
	return linkController{linkRepo: linkRepo}
}

func (c linkController) Store(ctx *fiber.Ctx) error {
	request := new(request.LinkStoreRequest)
	ctx.BodyParser(&request)

	if errors := util.Validate(request); errors != nil {
		response := util.ErrorResponse{Message: "wrong data", Errors: errors}
		return ctx.Status(fiber.StatusUnprocessableEntity).JSON(response)
	}

	userId := ctx.Locals("user").(model.User).Id
	var countLink int64
	c.linkRepo.Count(model.Link{}, model.Link{UserId: userId}, &countLink)

	if countLink >= 10 {
		response := util.Response{Message: "you store 10 link already"}
		return ctx.Status(fiber.StatusBadRequest).JSON(response)
	}

	linkExist := model.Link{}
	c.linkRepo.Find(&linkExist, model.Link{Url: request.Url, UserId: userId})

	if linkExist.Url == request.Url {
		response := util.Response{Message: "the url exist"}
		return ctx.Status(fiber.StatusBadRequest).JSON(response)
	}

	link := model.Link{
		UserId: userId,
		Url:    request.Url,
	}
	err := c.linkRepo.Create(&link)

	if err != nil {
		response := util.Response{Message: "the url already store"}
		return ctx.Status(fiber.StatusBadRequest).JSON(response)
	}

	response := util.Response{Message: "successful", Data: link}

	return ctx.JSON(response)
}
