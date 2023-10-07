package server

import (
	"chat/app/server/web"
	"github.com/gin-gonic/gin"
	"time"
)

type appServer struct {
	gin     *gin.Engine
	iWebApp web.IWebApp
}

func (app *appServer) Init() {
	app.gin = gin.New()
	app.gin.Use(gin.Recovery())

	app.iWebApp.Init(app.gin)

	time.Local = time.UTC
	//_, err := time.LoadLocation()
	//if err != nil {
	//	log.Fatalf("時區設置異常：%v\n", err)
	//	return
	//}
}

func (app *appServer) Run() {
	app.gin.Run(":8080")
}
