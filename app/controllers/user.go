package controllers

import (
	"chat/app/model/bo"
	"chat/app/model/dto"
	"chat/app/service"
	"chat/app/utils/errortool"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

type IUserCtrl interface {
	GetUserList(ctx *gin.Context)
	GetUser(ctx *gin.Context)
	PostUserLogin(ctx *gin.Context)
	PostUserRegister(ctx *gin.Context)
	PostUpdateUserInfo(ctx *gin.Context)
}

func ProvideUserCtrl(userSrv service.IUserSrv) IUserCtrl {
	return &UserCtrl{
		userSrv: userSrv,
	}
}

type UserCtrl struct {
	userSrv     service.IUserSrv
	SetResponse *StandardResponse
}

func (ctrl *UserCtrl) GetUserList(ctx *gin.Context) {
	dtoUserList := make([]*dto.UserInfoResp, 0)
	boUserList, err := ctrl.userSrv.GetUserList(ctx)

	if err != nil {
		ctrl.SetResponse.SetStandardResponse(ctx, http.StatusBadRequest, err)
		return
	}

	for _, boUser := range boUserList {
		userInfoResp := &dto.UserInfoResp{
			Id:          boUser.Id,
			Account:     boUser.Account,
			Nickname:    boUser.Nickname,
			Gender:      boUser.Gender,
			Country:     boUser.Country,
			Address:     boUser.Address,
			RegionCode:  boUser.RegionCode,
			PhoneNumber: boUser.PhoneNumber,
			CreateAt:    boUser.CreateAt.Format(time.DateTime),
			UpdateAt:    boUser.UpdateAt.Format(time.DateTime),
		}
		if boUser.Birthdate.IsZero() {
			userInfoResp.Birthdate = ""
		} else {
			userInfoResp.Birthdate = boUser.Birthdate.Format(time.DateOnly)
		}
		dtoUserList = append(dtoUserList, userInfoResp)

	}

	ctrl.SetResponse.SetStandardResponse(ctx, http.StatusOK, dtoUserList)
}

func (ctrl *UserCtrl) GetUser(ctx *gin.Context) {
	dtoUserCond := &dto.GetUserCond{}
	if err := ctx.BindUri(dtoUserCond); err != nil {
		ctrl.SetResponse.SetStandardResponse(ctx, http.StatusBadRequest, errortool.CommonErr.RequestParamError)
		return
	}

	boUserCond := &bo.GetUserCond{
		Account: dtoUserCond.Account,
	}
	boUser, err := ctrl.userSrv.GetUser(ctx, boUserCond)
	if err != nil {
		ctrl.SetResponse.SetStandardResponse(ctx, http.StatusBadRequest, err)
		return
	}

	dtoUser := &dto.UserInfoResp{
		Id:          boUser.Id,
		Account:     boUser.Account,
		Nickname:    boUser.Nickname,
		Gender:      boUser.Gender,
		Country:     boUser.Country,
		Address:     boUser.Address,
		RegionCode:  boUser.RegionCode,
		PhoneNumber: boUser.PhoneNumber,
		CreateAt:    boUser.CreateAt.Format(time.DateTime),
		UpdateAt:    boUser.UpdateAt.Format(time.DateTime),
	}
	if boUser.Birthdate.IsZero() {
		dtoUser.Birthdate = ""
	} else {
		dtoUser.Birthdate = boUser.Birthdate.Format(time.DateOnly)
	}

	ctrl.SetResponse.SetStandardResponse(ctx, http.StatusOK, dtoUser)
}

func (ctrl *UserCtrl) PostUserLogin(ctx *gin.Context) {
	dtoUserLogin := &dto.UserLoginCond{}

	if err := ctx.BindJSON(dtoUserLogin); err != nil {
		ctrl.SetResponse.SetStandardResponse(ctx, http.StatusBadRequest, errortool.CommonErr.RequestParamError)
		return
	}

	boUserLogin := &bo.UserLoginCond{
		Account:  dtoUserLogin.Account,
		Password: dtoUserLogin.Password,
	}

	boUserLoginResp, err := ctrl.userSrv.UserLogin(ctx, boUserLogin)
	if err != nil {
		ctrl.SetResponse.SetStandardResponse(ctx, http.StatusBadRequest, err)
		return
	}

	dtoUserLoginResp := &dto.UserLoginResp{
		Token: boUserLoginResp.Token,
	}

	ctx.SetCookie("account", dtoUserLogin.Account, int(time.Hour.Seconds()), "/", "localhost", false, false)
	ctx.SetCookie("token", dtoUserLoginResp.Token, int(time.Hour.Seconds()), "/", "localhost", false, false)
	ctrl.SetResponse.SetStandardResponse(ctx, http.StatusOK, dtoUserLoginResp)
}

func (ctrl *UserCtrl) PostUserRegister(ctx *gin.Context) {
	dtoUserRegData := &dto.UserRegCond{}
	if err := ctx.ShouldBindJSON(dtoUserRegData); err != nil {
		ctrl.SetResponse.SetStandardResponse(ctx, http.StatusBadRequest, errortool.CommonErr.RequestParamError)
		return
	}

	boUserReg := &bo.UserRegCond{
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

func (ctrl *UserCtrl) PostUpdateUserInfo(ctx *gin.Context) {

	dtoUpdateUserIdCond := &dto.UpdateUserIdCond{}
	err := ctx.BindUri(dtoUpdateUserIdCond)
	if err != nil {
		ctrl.SetResponse.SetStandardResponse(ctx, http.StatusBadRequest, errortool.CommonErr.RequestParamError)
		return
	}

	intId, err := strconv.ParseInt(dtoUpdateUserIdCond.Id, 10, 64)
	if err != nil || intId == 0 {
		ctrl.SetResponse.SetStandardResponse(ctx, http.StatusBadRequest, errortool.CommonErr.RequestParamError)
		return
	}

	dtoUpdateUserInfoCond := &dto.UpdateUserInfoCond{}
	if err := ctx.ShouldBindJSON(dtoUpdateUserInfoCond); err != nil {
		ctrl.SetResponse.SetStandardResponse(ctx, http.StatusBadRequest, errortool.CommonErr.RequestParamError)
		return
	}

	boUpdateUserInfoCond := &bo.UpdateUserInfoCond{
		Id:          uint(intId),
		Password:    dtoUpdateUserInfoCond.Password,
		Nickname:    dtoUpdateUserInfoCond.Nickname,
		Birthdate:   dtoUpdateUserInfoCond.Birthdate,
		Gender:      dtoUpdateUserInfoCond.Gender,
		Country:     dtoUpdateUserInfoCond.Country,
		Address:     dtoUpdateUserInfoCond.Country,
		RegionCode:  dtoUpdateUserInfoCond.RegionCode,
		PhoneNumber: dtoUpdateUserInfoCond.PhoneNumber,
	}
	if err := ctrl.userSrv.UpdateUserInfo(ctx, boUpdateUserInfoCond); err != nil {
		ctrl.SetResponse.SetStandardResponse(ctx, http.StatusBadRequest, err)
		return
	}

	ctrl.SetResponse.SetStandardResponse(ctx, http.StatusOK, nil)
}
