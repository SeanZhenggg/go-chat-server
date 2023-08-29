package main

import (
	// "gorm.io/driver/postgres"
	// "gorm.io/gorm"
	"chat/controllers"

	"github.com/gin-gonic/gin"
)

// func Init() {
// 	_, err := gorm.Open(postgres.New(postgres.Config{
// 		DriverName: "pgx",
// 		DSN:        "host=localhost user=postgres password=postgrespw dbname=postgres port=55000",
// 	}))

// 	if err != nil {
// 		fmt.Println(err)
// 	}

// 	fmt.Println("connection success!!!")
// }

func main() {
	server := gin.New()

	// middleware
	server.Use(gin.Recovery())
	var middle controllers.IResponseMiddleware = controllers.ProvideResponseMiddleware()
	server.Use(middle.ResponseHandler)

	// routes
	server.GET("/user/all")
	server.POST("/user/login")
	server.POST("/user/:account")

	server.Run(":8080")
}
