package service

import (
	"chat/model/bo"
	"chat/model/po"

	"chat/repository"

	"github.com/gin-gonic/gin"
	"golang.org/x/xerrors"
)

type IUserSrv interface {
	GetUserList(ctx *gin.Context) ([]*bo.UserInfo, error)
	GetUser(ctx *gin.Context, cond *bo.UserCond) (*bo.UserInfo, error)
	UserLogin(ctx *gin.Context, data *bo.UserLoginData) (*bo.UserLoginResp, error)
	UserRegister(ctx *gin.Context, data *bo.UserRegData) error
}

func ProvideUserSrv(userRepo repository.IUserRepo) IUserSrv {
	return &userService{userRepo}
}

type userService struct {
	userRepo repository.IUserRepo
}

func (srv *userService) GetUserList(ctx *gin.Context) ([]*bo.UserInfo, error) {
	users := make([]*bo.UserInfo, 0)

	poUserList, err := srv.userRepo.GetUserList()

	if err != nil {
		return nil, xerrors.Errorf("userService GetUserList error! : %w", err)
	}

	for _, poUser := range poUserList {
		users = append(users, &bo.UserInfo{
			Id:       poUser.Id,
			Account:  poUser.Account,
			Password: poUser.Password,
			Nickname: poUser.Nickname,
			CreateAt: poUser.CreateAt,
			UpdateAt: poUser.UpdateAt,
		})
	}

	return users, nil
}

func (srv *userService) GetUser(ctx *gin.Context, cond *bo.UserCond) (*bo.UserInfo, error) {
	poUserCond := &po.UserCond{
		Account: cond.Account,
	}

	poUser, err := srv.userRepo.GetUser(poUserCond)

	if err != nil {
		return nil, xerrors.Errorf("userService GetUser error! : %w", err)
	}

	user := &bo.UserInfo{
		Id:       poUser.Id,
		Account:  poUser.Account,
		Password: poUser.Password,
		Nickname: poUser.Nickname,
		CreateAt: poUser.CreateAt,
		UpdateAt: poUser.UpdateAt,
	}

	return user, nil
}

func (srv *userService) UserLogin(ctx *gin.Context, data *bo.UserLoginData) (*bo.UserLoginResp, error) {

	return &bo.UserLoginResp{}, nil
}

func (srv *userService) UserRegister(ctx *gin.Context, data *bo.UserRegData) error {
	poUserReg := &po.UserRegData{
		Account:  data.Account,
		Password: data.Password,
		Nickname: data.Nickname,
	}

	if err := srv.userRepo.UserRegister(poUserReg); err != nil {
		return xerrors.Errorf("userService UserRegister error! : %w", err)
	}

	return nil
}
