package service

import (
	"chat/model/bo"

	"github.com/gin-gonic/gin"
	"golang.org/x/xerrors"
	"gorm.io/gorm"
)

type IUserSrv interface {
	GetUserList(ctx *gin.Context) ([]*bo.UserInfo, error)
	GetUser(ctx *gin.Context, cond *bo.UserCond) (*bo.UserInfo, error)
	UserLogin(ctx *gin.Context, data *bo.UserLoginData) (*bo.UserLoginResp, error)
	UserRegister(ctx *gin.Context, data *bo.UserRegData) (*bo.UserInfo, error)
}

func ProvideUserSrv(db *gorm.DB) IUserSrv {
	return &userService{db}
}

type userService struct {
	Db *gorm.DB
}

func (srv *userService) GetUserList(ctx *gin.Context) ([]*bo.UserInfo, error) {
	users := []*bo.UserInfo{}

	if err := srv.Db.Find(&users); err != nil {
		return nil, xerrors.Errorf("userService GetUserList error!")
	}

	return users, nil
}

func (srv *userService) GetUser(ctx *gin.Context, cond *bo.UserCond) (*bo.UserInfo, error) {
	user := &bo.UserInfo{}

	if err := srv.Db.Find(&user); err != nil {
		return nil, xerrors.Errorf("userService GetUser error!")
	}

	return user, nil
}

func (srv *userService) UserLogin(ctx *gin.Context, data *bo.UserLoginData) (*bo.UserLoginResp, error) {
	return &bo.UserLoginResp{}, nil
}

func (srv *userService) UserRegister(ctx *gin.Context, data *bo.UserRegData) (*bo.UserInfo, error) {
	return &bo.UserInfo{}, nil
}
