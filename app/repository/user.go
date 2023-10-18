package repository

import (
	"chat/app/database"
	"chat/app/model/po"
	"chat/app/utils/errs"
	"context"

	"golang.org/x/xerrors"
)

type IUserRepo interface {
	GetUserList(ctx context.Context) ([]*po.User, error)
	GetUser(ctx context.Context, cond *po.GetUserCond) (*po.User, error)
	UserRegister(ctx context.Context, cond *po.UserRegCond) error
	UpdateUserInfo(ctx context.Context, cond *po.UpdateUserInfoIdCond, data *po.UpdateUserInfoCond) error
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
		return nil, xerrors.Errorf("userRepo GetUserList error ! : %w", errs.ParseDBError(err))
	}

	return userList, nil
}

func (repo *userRepo) GetUser(ctx context.Context, cond *po.GetUserCond) (*po.User, error) {
	user := &po.User{}

	db := repo.db.Session()
	if err := db.Where("account = ?", cond.Account).First(user).Error; err != nil {
		return nil, xerrors.Errorf("userRepo GetUser error ! : %w", errs.ParseDBError(err))
	}

	return user, nil
}

func (repo *userRepo) UserRegister(ctx context.Context, cond *po.UserRegCond) error {
	db := repo.db.Session()
	if err := db.Create(&po.User{
		Account:  cond.Account,
		Password: cond.Password,
		Nickname: cond.Nickname,
	}).Error; err != nil {
		return xerrors.Errorf("userRepo GetUser error ! : %w", errs.ParseDBError(err))
	}

	return nil
}

func (repo *userRepo) UpdateUserInfo(ctx context.Context, cond *po.UpdateUserInfoIdCond, data *po.UpdateUserInfoCond) error {
	db := repo.db.Session()
	if err := db.
		Model(&po.User{}).
		Where("id = ?", cond.Id).
		Updates(data).Error; err != nil {
		return xerrors.Errorf("userRepo UpdateUserInfo error ! : %w", errs.ParseDBError(err))
	}

	return nil
}
