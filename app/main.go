package main

import (
	"chat/app/controllers"
	"chat/app/repository"
	"chat/app/service"
	"chat/app/utils/logger"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	"github.com/gin-gonic/gin"
)

func Init() (*gorm.DB, error) {
	db, err := gorm.Open(postgres.New(postgres.Config{
		DriverName: "pgx",
		DSN:        "host=localhost user=postgres password=postgrespw dbname=postgres port=55000",
	}))

	if err != nil {
		return nil, err
	}

	return db, nil
}

type MyError struct{}

func (*MyError) Error() string {
	return "error!!!"
}

func setApiRoutes(g *gin.Engine, ctrl *controllers.Controller, middleware controllers.IResponseMiddleware) {
	group := g.Group("/api")

	// middleware
	group.Use(middleware.ResponseHandler)
	group.Use(gin.Recovery())

	group.GET("/user/all", ctrl.UserCtrl.GetUserList)
	group.GET("/user/:account", ctrl.UserCtrl.GetUser)
	group.POST("/user/login", ctrl.UserCtrl.PostUserLogin)
	group.POST("/user", ctrl.UserCtrl.PostUserRegister)
}

func setWsRoutes(g *gin.Engine, ctrl *controllers.Controller) {
	group := g.Group("/websocket")
	group.Use(gin.Recovery())

	group.GET("", ctrl.ChatCtrl.Conn)
}

func main() {
	// dependency injection
	iLogger := logger.ProviderLogger()

	// db init
	db, err := Init()
	if err != nil {
		iLogger.Error(err)
		panic(err)
	}

	server := gin.New()

	// dependency injection
	iMiddleware := controllers.ProvideResponseMiddleware()
	stdResp := &controllers.StandardResponse{}
	iUserRepo := repository.ProvideUserRepo(db)
	iUserSrv := service.ProvideUserSrv(iUserRepo)
	iHubSrv := service.ProvideHubSrv(iLogger)
	iUserCtrl := controllers.ProvideUserCtrl(iUserSrv, stdResp)
	iChatCtrl := controllers.ProvideChatCtrl(iHubSrv, iUserSrv, iLogger)
	iCtrls := controllers.ProvideControllers(iUserCtrl, iChatCtrl)

	// routes
	setApiRoutes(server, iCtrls, iMiddleware)
	setWsRoutes(server, iCtrls)

	server.Run(":8080")
}
