package repository

import (
	"chat/app/database"
	"chat/app/model/po"
	"chat/app/utils/errortool"
	"context"

	"golang.org/x/xerrors"
)

type IUserRepo interface {
	GetUserList(ctx context.Context) ([]*po.User, error)
	GetUser(ctx context.Context, cond *po.UserCond) (*po.User, error)
	UserRegister(ctx context.Context, cond *po.UserRegData) error
	UserLogin(ctx context.Context, cond *po.UserLoginData) (*po.User, error)
}

type userRepo struct {
	db database.IPostgresDB
}

func ProvideUserRepo(db database.IPostgresDB) IUserRepo {
	return &userRepo{db}
}

func (repo *userRepo) GetUserList(ctx context.Context) ([]*po.User, error) {
	userList := make([]*po.User, 0)

	db := repo.db.Session()
	if err := db.Model(&po.User{}).Find(&userList).Error; err != nil {
		return nil, xerrors.Errorf("userRepo GetUserList error ! : %w", errortool.ParseDBError(err))
	}

	return userList, nil
}

func (repo *userRepo) GetUser(ctx context.Context, cond *po.UserCond) (*po.User, error) {
	user := &po.User{}

	db := repo.db.Session()
	if err := db.Model(&po.User{}).Where("account = ?", cond.Account).First(user).Error; err != nil {
		return nil, xerrors.Errorf("userRepo GetUser error ! : %w", errortool.ParseDBError(err))
	}

	return user, nil
}

func (repo *userRepo) UserRegister(ctx context.Context, cond *po.UserRegData) error {
	db := repo.db.Session()
	if err := db.Model(&po.User{}).
		Create(&po.User{Account: cond.Account, Password: cond.Password, Nickname: cond.Nickname}).
		Error; err != nil {
		return xerrors.Errorf("userRepo GetUser error ! : %w", errortool.ParseDBError(err))
	}

	return nil
}

func (repo *userRepo) UserLogin(ctx context.Context, cond *po.UserLoginData) (*po.User, error) {
	user := &po.User{}

	db := repo.db.Session()
	if err := db.Model(&po.User{}).
		Where(&po.User{Account: cond.Account, Password: cond.Password}).
		First(user).Error; err != nil {
		return nil, xerrors.Errorf("userRepo UserLogin error ! : %w", errortool.ParseDBError(err))
	}

	return user, nil
}
