package controller

import (
	"encoding/json"
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/golang/mock/gomock"
	"github.com/mohammadgh1370/url-shortner/internal/model"
	"github.com/mohammadgh1370/url-shortner/internal/repository/mocks"
	"github.com/stretchr/testify/assert"
	"github.com/valyala/fasthttp"
	"testing"
)

func TestShow(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	mockLinkRepo := mocks.NewMockILinkRepo(mockCtrl)
	mockViewRepo := mocks.NewMockIViewRepo(mockCtrl)

	viewCtrl := viewController{
		linkRepo: mockLinkRepo,
		viewRepo: mockViewRepo,
	}

	viewCount := int64(5)
	address := "https://example.com/dgdfgd"
	mockLink := model.Link{
		Id:     uint(1),
		Url:    "https://example.com",
		UserId: uint(1),
		Hash:   "dgdfgd",
	}

	mockLinkRepo.EXPECT().First(gomock.Any(), model.Link{Hash: mockLink.Hash, UserId: mockLink.UserId}).SetArg(0, mockLink).Return(nil)

	mockViewRepo.EXPECT().Count(gomock.Any(), model.View{LinkId: mockLink.Id}, gomock.Any()).Do(func(_ interface{}, _ model.View, count *int64) {
		*count = viewCount
	}).Return(nil)

	app := fiber.New()
	ctx := app.AcquireCtx(&fasthttp.RequestCtx{})
	defer app.ReleaseCtx(ctx)
	ctx.Request().Header.SetContentType("application/json")

	ctx.Locals("user", model.User{Id: mockLink.UserId})
	ctx.Request().URI().SetQueryStringBytes([]byte(fmt.Sprintf(`url=%s`, address)))

	err := viewCtrl.Show(ctx)

	assert.Nil(t, err)

	assert.Equal(t, fiber.StatusOK, ctx.Response().StatusCode())

	var response map[string]int
	err = json.Unmarshal(ctx.Response().Body(), &response)
	assert.Nil(t, err)
	assert.Equal(t, map[string]int{"count": int(viewCount)}, response)
}
