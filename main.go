package main

import (
	"chat/controllers"
	"chat/repository"
	"chat/service"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	"fmt"

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

	fmt.Println("connection success!!!")

	return db, nil
}

type MyError struct{}

func (*MyError) Error() string {
	return "error!!!"
}

func main() {
	// db
	db, err := Init()
	if err != nil {
		panic(err)
	}

	server := gin.New()

	// middleware
	server.Use(gin.Recovery())
	var middle controllers.IResponseMiddleware = controllers.ProvideResponseMiddleware()
	server.Use(middle.ResponseHandler)

	// dependency injection
	stdResp := &controllers.StandardResponse{}
	iUserRepo := repository.ProvideUserRepo(db)
	iUserSrv := service.ProvideUserSrv(iUserRepo)
	iServices := service.ProvideServices(iUserSrv)
	iUserCtrl := controllers.ProvideUserCtrl(iServices.UserSrv, stdResp)
	iCtrls := controllers.ProvideControllers(iUserCtrl)

	// routes
	server.GET("/api/user/all", iCtrls.UserCtrl.GetUserList)
	server.GET("/api/user/:account", iCtrls.UserCtrl.GetUser)
	server.POST("/api/user/login", iCtrls.UserCtrl.PostUserLogin)
	server.POST("/api/user", iCtrls.UserCtrl.PostUserRegister)

	server.Run(":8080")
}
