package server

import (
	"chat/app/controllers"
	"chat/app/database"
	"chat/app/repository"
	"chat/app/server/web"
	"chat/app/service"
	"chat/app/utils/logger"
)

func NewAppServer() *appServer {
	// dependency injection
	iPostgresDB := database.ProvidePostgresDB()
	iLogger := logger.ProviderLogger()
	iMiddleware := controllers.ProvideResponseMiddleware()
	iUserRepo := repository.ProvideUserRepo(iPostgresDB)
	iUserSrv := service.ProvideUserSrv(iUserRepo)
	iHubSrv := service.ProvideHubSrv(iLogger)
	iUserCtrl := controllers.ProvideUserCtrl(iUserSrv)
	iChatCtrl := controllers.ProvideChatCtrl(iHubSrv, iUserSrv, iLogger)
	iCtrls := controllers.ProvideControllers(iUserCtrl, iChatCtrl)
	iWebApp := web.ProvideWebApp(iCtrls, iMiddleware)

	server := &appServer{
		iWebApp: iWebApp,
	}

	return server
}
