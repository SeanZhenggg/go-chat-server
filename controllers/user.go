package controllers

import "github.com/gin-gonic/gin"

type IUserCtrl interface {
	GetUserList(ctx *gin.Context)
	GetUser(ctx *gin.Context)
	PostUserLogin(ctx *gin.Context)
	PostUserRegister(ctx *gin.Context)
}

type UserCtrl struct {
}

func (ctrl *UserCtrl) GetUserList(ctx *gin.Context) {

}

func (ctrl *UserCtrl) GetUser(ctx *gin.Context) {

}

func (ctrl *UserCtrl) PostUserLogin(ctx *gin.Context) {

}

func (ctrl *UserCtrl) PostUserRegister(ctx *gin.Context) {

}
