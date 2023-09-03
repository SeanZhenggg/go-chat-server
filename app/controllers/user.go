package controllers

import (
	"chat/app/model/bo"
	"chat/app/model/dto"
	"chat/app/service"
	"chat/app/utils/errortool"
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
		ctrl.SetResponse.SetStandardResponse(ctx, http.StatusBadRequest, errortool.ReqErr.RequestParamError)
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
	dtoUserLogin := &dto.UserLoginDto{}

	if err := ctx.BindJSON(dtoUserLogin); err != nil {
		ctrl.SetResponse.SetStandardResponse(ctx, http.StatusBadRequest, errortool.ReqErr.RequestParamError)
		return
	}

	boUserLogin := &bo.UserLoginData{
		Account:  dtoUserLogin.Account,
		Password: dtoUserLogin.Password,
	}

	boUserLoginResp, err := ctrl.userSrv.UserLogin(ctx, boUserLogin)
	if err != nil {
		ctrl.SetResponse.SetStandardResponse(ctx, http.StatusBadRequest, err)
		return
	}

	dtoUserLoginResp := &dto.UserLoginRespDto{
		Token: boUserLoginResp.Token,
	}

	ctrl.SetResponse.SetStandardResponse(ctx, http.StatusOK, dtoUserLoginResp)
}

func (ctrl *UserCtrl) PostUserRegister(ctx *gin.Context) {
	dtoUserRegData := &dto.UserRegDto{}
	if err := ctx.BindJSON(dtoUserRegData); err != nil {
		ctrl.SetResponse.SetStandardResponse(ctx, http.StatusBadRequest, errortool.ReqErr.RequestParamError)
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
