package controller

import (
	"github.com/gofiber/fiber/v2"
	"github.com/mohammadgh1370/url-shortner/internal/model"
	"github.com/mohammadgh1370/url-shortner/internal/repository"
	"github.com/mohammadgh1370/url-shortner/internal/util"
	"net/url"
)

type publicController struct {
	linkRepo repository.ILinkRepo
	viewRepo repository.IViewRepo
}

func NewPublicController(linkRepo repository.ILinkRepo, viewRepo repository.IViewRepo) publicController {
	return publicController{linkRepo: linkRepo, viewRepo: viewRepo}
}

func (c publicController) Redirect(ctx *fiber.Ctx) error {
	linkExist := model.Link{}
	c.linkRepo.First(&linkExist, model.Link{Hash: ctx.Params("hash")})

	if linkExist.Hash != ctx.Params("hash") {
		response := util.Response{Message: "the url not exist"}
		return ctx.Status(fiber.StatusNotFound).JSON(response)
	}

	url := generateUrl(linkExist.Url, ctx.Queries())

	view := model.View{
		LinkId:    linkExist.Id,
		Ip:        ctx.IP(),
		UserAgent: ctx.Get("User-Agent"),
		Referer:   ctx.Get("Referer"),
	}
	c.viewRepo.Create(&view)

	return ctx.Redirect(url)
}

func generateUrl(address string, queryParams map[string]string) string {
	u, _ := url.Parse(address)

	params := u.Query()
	u.RawQuery = ""
	for key, value := range queryParams {
		params.Add(key, value)
	}

	u.RawQuery = params.Encode()

	return u.String()
}
