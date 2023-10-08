package web

import (
	"chat/app/controllers"
	"chat/app/controllers/middleware"

	"github.com/gin-gonic/gin"
)

type IWebApp interface {
	Init(g *gin.Engine)
}

func ProvideWebApp(
	ctrl *controllers.Controller,
	respMw middleware.IResponseMiddleware,
	authMw middleware.IAuthMiddleware,
) *webApp {
	return &webApp{
		Ctrl:   ctrl,
		RespMw: respMw,
		AuthMw: authMw,
	}
}

type webApp struct {
	Ctrl   *controllers.Controller
	RespMw middleware.IResponseMiddleware
	AuthMw middleware.IAuthMiddleware
}

func (app *webApp) Init(g *gin.Engine) {
	app.setApiRoutes(g)
	app.setWsRoutes(g)
}
