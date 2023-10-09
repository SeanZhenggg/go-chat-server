package server

import (
	"chat/app/controllers"
	"chat/app/controllers/middleware"
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
	iRespMiddleware := middleware.ProvideResponseMiddleware()
	iUserRepo := repository.ProvideUserRepo(iPostgresDB)
	iUserSrv := service.ProvideUserSrv(iUserRepo)
	iHubSrv := service.ProvideHubSrv(iLogger)
	iUserCtrl := controllers.ProvideUserCtrl(iUserSrv, iLogger)
	iChatCtrl := controllers.ProvideChatCtrl(iHubSrv, iUserSrv, iLogger)
	iCtrls := controllers.ProvideControllers(iUserCtrl, iChatCtrl)
	iAuthMiddleware := middleware.ProvideAuthMiddleware(iLogger, iUserSrv)
	iWebApp := web.ProvideWebApp(iCtrls, iRespMiddleware, iAuthMiddleware)

	server := &appServer{
		iWebApp: iWebApp,
	}

	return server
}
