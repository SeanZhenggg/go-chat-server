package web

import (
	"chat/app/controllers"

	"github.com/gin-gonic/gin"
)

type IWebApp interface {
	Init(g *gin.Engine)
}

func ProvideWebApp(
	ctrl *controllers.Controller,
	mw controllers.IResponseMiddleware,
) *webApp {
	return &webApp{
		Ctrl: ctrl,
		Mw:   mw,
	}
}

type webApp struct {
	Ctrl *controllers.Controller
	Mw   controllers.IResponseMiddleware
}

func (app *webApp) Init(g *gin.Engine) {
	app.setApiRoutes(g)
	app.setWsRoutes(g)
}
