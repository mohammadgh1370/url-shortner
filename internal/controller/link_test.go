package controller

import (
	"encoding/json"
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/golang/mock/gomock"
	"github.com/mohammadgh1370/url-shortner/internal/config"
	"github.com/mohammadgh1370/url-shortner/internal/model"
	"github.com/mohammadgh1370/url-shortner/internal/repository/mocks"
	"github.com/mohammadgh1370/url-shortner/internal/util"
	"github.com/stretchr/testify/assert"
	"github.com/valyala/fasthttp"
	"net/url"
	"path"
	"testing"
)

func TestStore(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	mockLinkRepo := mocks.NewMockILinkRepo(mockCtrl)
	linkCtrl := linkController{linkRepo: mockLinkRepo}

	mockLink := model.Link{
		Url:    "https://example.com",
		UserId: uint(1),
		Hash:   "dhgjhfj",
	}

	mockLinkRepo.EXPECT().Count(gomock.Any(), model.Link{UserId: mockLink.UserId}, gomock.Any()).Do(func(_ interface{}, _ model.Link, count *int64) {
		*count = 5
	}).Return(nil)

	mockLinkRepo.EXPECT().First(gomock.Any(), model.Link{Url: mockLink.Url, UserId: mockLink.UserId}).Return(nil)

	mockLinkRepo.EXPECT().Create(gomock.Any()).Do(func(link *model.Link) {
		link.Id = mockLink.Id
		link.Url = mockLink.Url
		link.Hash = mockLink.Hash
	}).Return(nil)

	app := fiber.New()
	ctx := app.AcquireCtx(&fasthttp.RequestCtx{})
	defer app.ReleaseCtx(ctx)
	ctx.Request().Header.SetContentType("application/json")

	ctx.Locals("user", model.User{Id: mockLink.UserId})

	payload := map[string]string{"url": mockLink.Url}
	jsonPayload, _ := json.Marshal(payload)
	ctx.Request().SetBody(jsonPayload)

	err := linkCtrl.Store(ctx)

	type response struct {
		Data    struct{ Url string }
		Message string
	}
	resp := response{}
	json.Unmarshal(ctx.Response().Body(), &resp)

	assert.Nil(t, err)
	assert.Equal(t, fiber.StatusOK, ctx.Response().StatusCode())
	assert.Equal(t, resp.Data.Url, config.APP_URL+mockLink.Hash)
}

func TestIndex(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	mockLinkRepo := mocks.NewMockILinkRepo(mockCtrl)
	linkCtrl := linkController{linkRepo: mockLinkRepo}

	userId := uint(1)
	mockLinks := []model.Link{
		{UserId: userId, Url: "https://example.com/1"},
		{UserId: userId, Url: "https://example.com/2"},
	}

	mockLinkRepo.EXPECT().Find(gomock.Any(), model.Link{UserId: userId}).Do(func(links *[]model.Link, _ model.Link) {
		*links = mockLinks
	}).Return(nil)

	app := fiber.New()
	ctx := app.AcquireCtx(&fasthttp.RequestCtx{})
	defer app.ReleaseCtx(ctx)
	ctx.Request().Header.SetContentType("application/json")

	ctx.Locals("user", model.User{Id: userId})

	err := linkCtrl.Index(ctx)

	assert.Nil(t, err)
	assert.Equal(t, fiber.StatusOK, ctx.Response().StatusCode())

	var response util.Response
	err = json.Unmarshal(ctx.Response().Body(), &response)

	var data []string
	for _, link := range mockLinks {
		data = append(data, link.Url)
	}

	assert.Nil(t, err)
	assert.Equal(t, "successful", response.Message)
	assert.ElementsMatch(t, data, response.Data)
}

func TestDestroy(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	mockLinkRepo := mocks.NewMockILinkRepo(mockCtrl)
	linkCtrl := linkController{linkRepo: mockLinkRepo}

	userId := uint(1)
	address := "https://example.com/abc"
	u, _ := url.Parse(address)
	hash := path.Base(u.Path)

	mockLinkRepo.EXPECT().First(gomock.Any(), model.Link{Hash: hash, UserId: userId}).Do(func(link *model.Link, _ model.Link) {
		*link = model.Link{UserId: userId, Hash: hash}
	}).Return(nil)

	mockLinkRepo.EXPECT().Delete(gomock.Any(), model.Link{UserId: userId, Hash: hash}).Return(nil)

	app := fiber.New()
	ctx := app.AcquireCtx(&fasthttp.RequestCtx{})
	defer app.ReleaseCtx(ctx)
	ctx.Request().Header.SetContentType("application/json")
	ctx.Request().SetBody([]byte(fmt.Sprintf(`{"url": "%s"}`, address)))

	ctx.Locals("user", model.User{Id: userId})

	err := linkCtrl.Destroy(ctx)

	assert.Nil(t, err)
	assert.Equal(t, fiber.StatusOK, ctx.Response().StatusCode())

	var response util.Response
	err = json.Unmarshal(ctx.Response().Body(), &response)
	assert.Nil(t, err)
	assert.Equal(t, "successful", response.Message)
}
