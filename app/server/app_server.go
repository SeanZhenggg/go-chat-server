package server

import (
	"chat/app/server/web"

	"github.com/gin-gonic/gin"
)

type appServer struct {
	gin     *gin.Engine
	iWebApp web.IWebApp
}

func (app *appServer) Init() {
	app.gin = gin.New()
	app.gin.Use(gin.Recovery())

	app.iWebApp.Init(app.gin)
}

func (app *appServer) Run() {
	app.gin.Run(":8080")
}
