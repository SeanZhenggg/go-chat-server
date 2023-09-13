package main

import (
	"chat/app/server"
)

func main() {
	app := server.NewAppServer()

	app.Init()
	app.Run()
}
