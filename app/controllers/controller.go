package controllers

import (
	"chat/app/controllers/middleware"
	"github.com/gin-gonic/gin"
)

type Controller struct {
	UserCtrl IUserCtrl
	ChatCtrl IChatCtrl
}

func ProvideControllers(userCtrl IUserCtrl, chatCtrl IChatCtrl) *Controller {
	return &Controller{
		UserCtrl: userCtrl,
		ChatCtrl: chatCtrl,
	}
}

type StandardResponse struct{}

func (stdResp *StandardResponse) SetStandardResponse(ctx *gin.Context, statusCode int, data interface{}) {
	middleware.SetResp(ctx, statusCode, data)
}
