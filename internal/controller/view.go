package controller

import (
	"github.com/gofiber/fiber/v2"
	"github.com/mohammadgh1370/url-shortner/internal/model"
	"github.com/mohammadgh1370/url-shortner/internal/repository"
	"github.com/mohammadgh1370/url-shortner/internal/util"
	"net/url"
	"path"
)

type viewController struct {
	linkRepo repository.ILinkRepo
	viewRepo repository.IViewRepo
}

func NewViewController(linkRepo repository.ILinkRepo, viewRepo repository.IViewRepo) viewController {
	return viewController{linkRepo: linkRepo, viewRepo: viewRepo}
}

func (c viewController) Show(ctx *fiber.Ctx) error {
	userId := ctx.Locals("user").(model.User).Id
	u, _ := url.Parse(ctx.Query("url"))

	link := model.Link{}
	c.linkRepo.First(&link, model.Link{Hash: path.Base(u.Path), UserId: userId})

	if link.Hash != path.Base(u.Path) {
		response := util.Response{Message: "the url not exist"}
		return ctx.Status(fiber.StatusNotFound).JSON(response)
	}

	var countView int64
	c.viewRepo.Count(model.View{}, model.View{LinkId: link.Id}, &countView)

	return ctx.JSON(map[string]int{"count": int(countView)})
}
