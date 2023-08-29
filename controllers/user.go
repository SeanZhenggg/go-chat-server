package controllers

import (
	"chat/service"

	"github.com/gin-gonic/gin"
)

type IUserCtrl interface {
	GetUserList(ctx *gin.Context)
	GetUser(ctx *gin.Context)
	PostUserLogin(ctx *gin.Context)
	PostUserRegister(ctx *gin.Context)
}

func ProvideUserCtrl(userSrv service.IUserSrv, setResponse *StandardResponse) IUserCtrl {
	return &UserCtrl{
		userSrv:     userSrv,
		SetResponse: setResponse,
	}
}

type UserCtrl struct {
	userSrv     service.IUserSrv
	SetResponse *StandardResponse
}

func (ctrl *UserCtrl) GetUserList(ctx *gin.Context) {

}

func (ctrl *UserCtrl) GetUser(ctx *gin.Context) {

}

func (ctrl *UserCtrl) PostUserLogin(ctx *gin.Context) {

}

func (ctrl *UserCtrl) PostUserRegister(ctx *gin.Context) {

}
