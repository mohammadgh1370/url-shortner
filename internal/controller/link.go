package controller

import (
	"github.com/gofiber/fiber/v2"
	"github.com/mohammadgh1370/url-shortner/internal/config"
	"github.com/mohammadgh1370/url-shortner/internal/model"
	"github.com/mohammadgh1370/url-shortner/internal/repository"
	"github.com/mohammadgh1370/url-shortner/internal/request"
	"github.com/mohammadgh1370/url-shortner/internal/util"
	"net/url"
	"path"
)

type linkController struct {
	linkRepo repository.ILinkRepo
}

func NewLinkController(linkRepo repository.ILinkRepo) linkController {
	return linkController{linkRepo: linkRepo}
}

func (c linkController) Store(ctx *fiber.Ctx) error {
	request := new(request.LinkRequest)
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
	c.linkRepo.First(&linkExist, model.Link{Url: request.Url, UserId: userId})

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

	response := util.Response{Message: "successful", Data: map[string]string{"url": config.APP_URL + link.Hash}}

	return ctx.JSON(response)
}

func (c linkController) Index(ctx *fiber.Ctx) error {
	userId := ctx.Locals("user").(model.User).Id

	var links []model.Link
	c.linkRepo.Find(&links, model.Link{UserId: userId})

	var data []string
	for _, link := range links {
		data = append(data, link.Url)
	}

	response := util.Response{Message: "successful", Data: data}

	return ctx.JSON(response)
}

func (c linkController) Destroy(ctx *fiber.Ctx) error {
	userId := ctx.Locals("user").(model.User).Id

	request := new(request.LinkRequest)
	ctx.BodyParser(&request)

	if errors := util.Validate(request); errors != nil {
		response := util.ErrorResponse{Message: "wrong data", Errors: errors}
		return ctx.Status(fiber.StatusUnprocessableEntity).JSON(response)
	}

	u, _ := url.Parse(request.Url)

	linkExist := model.Link{}
	c.linkRepo.First(&linkExist, model.Link{Hash: path.Base(u.Path), UserId: userId})
	if linkExist.Hash != path.Base(u.Path) {
		response := util.Response{Message: "the link does not exist"}
		return ctx.Status(fiber.StatusBadRequest).JSON(response)
	}

	c.linkRepo.Delete(model.Link{}, linkExist)

	response := util.Response{Message: "successful"}

	return ctx.JSON(response)
}
