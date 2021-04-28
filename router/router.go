package router

import (
	"github.com/labstack/echo/v4"
	"go-ms-session/repository"
	"go-ms-session/service"
)

type Routes struct {
	app                *echo.Echo
	accessTokenService *service.AccessTokenService
}

func NewRoutes(app *echo.Echo,mg *repository.MongoRepository) *Routes {
	return &Routes{
		app:                app,
		accessTokenService: service.NewVerifyAccessTokenService(mg),
	}
}
func (r *Routes) InitRoute() {
	route := r.app.POST
	route("/verify-accessToken", r.accessTokenService.VerifyToKen)
	//route("/getFacebookAccessToken", c.GetFacebookAccessTokenController, r.firebaseMiddleware)
	//route("/getGoogleAccessToken", c.GetGoogleAccessTokenController, r.firebaseMiddleware)
	//route("/getProfileSettings", c.GetProfileSetting, r.firebaseMiddleware)
}
