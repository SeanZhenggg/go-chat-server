package controllers

import (
	"github.com/gin-gonic/gin"
)

const (
	Resp_Data   = "Resp_Data"
	Resp_Status = "Resp_Status"
	Resp_Code   = "Resp_Code"
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
	ctx.Set(Resp_Status, statusCode)
	ctx.Set(Resp_Code, 0)
	ctx.Set(Resp_Data, data)
}
