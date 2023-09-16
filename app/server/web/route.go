package web

import (
	"github.com/gin-gonic/gin"
)

func (app *webApp) setApiRoutes(g *gin.Engine) {
	group := g.Group("/api")

	// middleware
	group.Use(app.Mw.ResponseHandler)

	group.GET("/user/all", app.Ctrl.UserCtrl.GetUserList)
	group.GET("/user/:account", app.Ctrl.UserCtrl.GetUser)
	group.POST("/user/login", app.Ctrl.UserCtrl.PostUserLogin)
	group.POST("/user", app.Ctrl.UserCtrl.PostUserRegister)
}

func (app *webApp) setWsRoutes(g *gin.Engine) {
	group := g.Group("/websocket")

	group.GET("", app.Ctrl.ChatCtrl.Conn)
}
