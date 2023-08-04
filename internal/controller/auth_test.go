package controller

import (
	"encoding/json"
	"github.com/gofiber/fiber/v2"
	"github.com/golang/mock/gomock"
	"github.com/mohammadgh1370/url-shortner/internal/model"
	"github.com/mohammadgh1370/url-shortner/internal/repository/mocks"
	"github.com/mohammadgh1370/url-shortner/internal/request"
	"github.com/mohammadgh1370/url-shortner/internal/util"
	"github.com/stretchr/testify/assert"
	"github.com/valyala/fasthttp"
	"testing"
)

func TestRegisterWithValidRequest(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	mockUserRepo := mocks.NewMockIUserRepo(mockCtrl)

	authCtrl := authController{userRepo: mockUserRepo}

	payload := request.UserRegisterRequest{
		FirstName: "mohammad",
		LastName:  "ghorbani",
		Username:  "mohammadgh",
		Password:  "password",
	}
	jsonPayload, _ := json.Marshal(payload)

	app := fiber.New()
	ctx := app.AcquireCtx(&fasthttp.RequestCtx{})
	defer app.ReleaseCtx(ctx)
	ctx.Request().Header.SetContentType("application/json")
	ctx.Request().SetBody(jsonPayload)

	mockUserRepo.EXPECT().First(&model.User{}, model.User{Username: payload.Username}).Return(nil)

	hashedPassword, _ := util.HashPassword(payload.Password)
	var newUser = model.User{
		FirstName: payload.FirstName,
		LastName:  payload.LastName,
		Username:  payload.Username,
		Password:  hashedPassword,
	}

	mockUserRepo.EXPECT().Create(gomock.Any()).Do(func(user *model.User) {
		assert.Equal(t, newUser.Username, user.Username)
	}).Return(nil)

	err := authCtrl.Register(ctx)

	type response struct {
		Data    struct{ Token string }
		Message string
	}
	resp := response{}
	json.Unmarshal(ctx.Response().Body(), &resp)

	assert.Nil(t, err)
	assert.Equal(t, 200, ctx.Response().StatusCode())
	assert.NotNil(t, resp.Data.Token)
}

func TestLoginWithValidRequest(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	mockUserRepo := mocks.NewMockIUserRepo(mockCtrl)

	authCtrl := authController{userRepo: mockUserRepo}

	payload := &request.UserRegisterRequest{
		Username: "mohammadgh",
		Password: "password",
	}
	jsonPayload, _ := json.Marshal(payload)

	app := fiber.New()
	ctx := app.AcquireCtx(&fasthttp.RequestCtx{})
	defer app.ReleaseCtx(ctx)
	ctx.Request().Header.SetContentType("application/json")
	ctx.Request().SetBody(jsonPayload)

	hashedPassword, _ := util.HashPassword(payload.Password)
	mockUser := &model.User{
		Username: payload.Username,
		Password: hashedPassword,
	}

	mockUserRepo.EXPECT().First(gomock.Any(), model.User{Username: payload.Username}).SetArg(0, *mockUser).Return(nil)

	err := authCtrl.Login(ctx)

	type response struct {
		Data    struct{ Token string }
		Message string
	}
	resp := response{}
	json.Unmarshal(ctx.Response().Body(), &resp)

	assert.Nil(t, err)
	assert.Equal(t, 200, ctx.Response().StatusCode())
	assert.NotNil(t, resp.Data.Token)
}

func TestMe(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	mockUserRepo := mocks.NewMockIUserRepo(mockCtrl)
	authCtrl := authController{userRepo: mockUserRepo}

	mockUser := model.User{
		FirstName: "mohammad",
		LastName:  "ghorbani",
		Username:  "mohammadgh",
		Password:  "password",
	}

	app := fiber.New()
	ctx := app.AcquireCtx(&fasthttp.RequestCtx{})
	defer app.ReleaseCtx(ctx)
	ctx.Request().Header.SetContentType("application/json")

	ctx.Locals("user", mockUser)

	err := authCtrl.Me(ctx)

	type response struct {
		Data    model.User
		Message string
	}
	resp := response{}
	json.Unmarshal(ctx.Response().Body(), &resp)

	assert.Nil(t, err)
	assert.Equal(t, 200, ctx.Response().StatusCode())
	assert.Equal(t, resp.Data.Username, mockUser.Username)
	assert.Equal(t, resp.Data.FirstName, mockUser.FirstName)
	assert.Equal(t, resp.Data.LastName, mockUser.LastName)
}
