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
}

func ProvideController(userCtrl IUserCtrl) *Controller {
	return &Controller{
		UserCtrl: userCtrl,
	}
}

type StandardResponse struct{}

func (stdResp *StandardResponse) SetStandardResponse(ctx *gin.Context, statusCode int, code int, data interface{}) {
	ctx.Set(Resp_Status, statusCode)
	ctx.Set(Resp_Code, code)
	ctx.Set(Resp_Data, data)
}
