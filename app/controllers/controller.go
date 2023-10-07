package controllers

import (
	"github.com/gin-gonic/gin"
)

const (
	RespData   = "Resp_Data"
	RespStatus = "Resp_Status"
	RespCode   = "Resp_Code"
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
	ctx.Set(RespStatus, statusCode)
	ctx.Set(RespCode, 0)
	ctx.Set(RespData, data)
}
