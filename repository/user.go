package repository

import (
	"chat/model/po"

	"golang.org/x/xerrors"
	"gorm.io/gorm"
)

type IUserRepo interface {
	GetUserList() ([]*po.User, error)
	GetUser(cond *po.UserCond) (*po.User, error)
	UserRegister(data *po.UserRegData) error
}

type userRepo struct {
	db *gorm.DB
}

func ProvideUserRepo(db *gorm.DB) IUserRepo {
	return &userRepo{db}
}

func (repo *userRepo) GetUserList() ([]*po.User, error) {
	userList := make([]*po.User, 0)

	if err := repo.db.Model(&po.User{}).Find(&userList).Error; err != nil {
		return nil, xerrors.Errorf("userRepo GetUserList error ! : %w", err)
	}

	return userList, nil
}

func (repo *userRepo) GetUser(cond *po.UserCond) (*po.User, error) {
	user := &po.User{}
	if err := repo.db.Model(user).Where("account = ?", cond.Account).First(user).Error; err != nil {
		return nil, xerrors.Errorf("userRepo GetUser error ! : %w", err)
	}

	return user, nil
}

func (repo *userRepo) UserRegister(data *po.UserRegData) error {

	if err := repo.db.Model(&po.User{}).Create(data).Error; err != nil {
		return xerrors.Errorf("userRepo GetUser error ! : %w", err)
	}

	return nil
}
