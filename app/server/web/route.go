package web

import (
	"github.com/gin-gonic/gin"
)

func (app *webApp) setApiRoutes(g *gin.Engine) {
	group := g.Group("/api")

	// middleware
	group.Use(app.RespMw.ResponseHandler)

	group.GET("/user/all", app.Ctrl.UserCtrl.GetUserList)
	group.GET("/user/:account", app.AuthMw.AuthValidationHandler, app.Ctrl.UserCtrl.GetUser)
	group.POST("/user/login", app.Ctrl.UserCtrl.PostUserLogin)
	group.POST("/user", app.Ctrl.UserCtrl.PostUserRegister)
	group.POST("/user/:id", app.AuthMw.AuthValidationHandler, app.Ctrl.UserCtrl.PostUpdateUserInfo)
}

func (app *webApp) setWsRoutes(g *gin.Engine) {
	group := g.Group("/websocket")

	group.GET("", app.AuthMw.AuthValidationHandler, app.Ctrl.ChatCtrl.Conn)
}
