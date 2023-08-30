package controllers

import (
	"chat/model/bo"
	"chat/model/dto"
	"chat/service"
	"net/http"

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
	dtoUserList := make([]*dto.UserDto, 0)
	boUserList, err := ctrl.userSrv.GetUserList(ctx)

	if err != nil {
		ctrl.SetResponse.SetStandardResponse(ctx, http.StatusBadRequest, err)
		return
	}

	for _, boUser := range boUserList {
		dtoUserList = append(dtoUserList, &dto.UserDto{
			Id:       boUser.Id,
			Account:  boUser.Account,
			Nickname: boUser.Nickname,
			CreateAt: boUser.CreateAt,
			UpdateAt: boUser.UpdateAt,
		})
	}

	ctrl.SetResponse.SetStandardResponse(ctx, http.StatusOK, dtoUserList)
}

func (ctrl *UserCtrl) GetUser(ctx *gin.Context) {
	dtoUserCond := &dto.UserCondDto{}
	if err := ctx.BindUri(dtoUserCond); err != nil {
		ctrl.SetResponse.SetStandardResponse(ctx, http.StatusBadRequest, "參數錯誤")
		return
	}

	boUserCond := &bo.UserCond{
		Account: dtoUserCond.Account,
	}
	boUser, err := ctrl.userSrv.GetUser(ctx, boUserCond)
	if err != nil {
		ctrl.SetResponse.SetStandardResponse(ctx, http.StatusBadRequest, err)
		return
	}

	dtoUser := &dto.UserDto{
		Id:       boUser.Id,
		Account:  boUser.Account,
		Nickname: boUser.Nickname,
		CreateAt: boUser.CreateAt,
		UpdateAt: boUser.UpdateAt,
	}

	ctrl.SetResponse.SetStandardResponse(ctx, http.StatusOK, dtoUser)
}

func (ctrl *UserCtrl) PostUserLogin(ctx *gin.Context) {

}

func (ctrl *UserCtrl) PostUserRegister(ctx *gin.Context) {
	dtoUserRegData := &dto.UserRegDto{}
	if err := ctx.BindJSON(dtoUserRegData); err != nil {
		ctrl.SetResponse.SetStandardResponse(ctx, http.StatusBadRequest, "參數錯誤")
		return
	}

	boUserReg := &bo.UserRegData{
		Account:  dtoUserRegData.Account,
		Password: dtoUserRegData.Password,
		Nickname: dtoUserRegData.Nickname,
	}
	if err := ctrl.userSrv.UserRegister(ctx, boUserReg); err != nil {
		ctrl.SetResponse.SetStandardResponse(ctx, http.StatusBadRequest, err)
		return
	}

	ctrl.SetResponse.SetStandardResponse(ctx, http.StatusOK, nil)
}
