package service

import (
	"chat/app/model/bo"
	"chat/app/model/po"
	"chat/app/repository"
	"chat/app/utils/auth"
	encryptUtil "chat/app/utils/encrypt"
	"chat/app/utils/errortool"
	"context"
	"errors"
	"golang.org/x/xerrors"
)

type IUserSrv interface {
	GetUserList(ctx context.Context) ([]*bo.UserInfo, error)
	GetUser(ctx context.Context, cond *bo.GetUserCond) (*bo.UserInfo, error)
	UserLogin(ctx context.Context, data *bo.UserLoginCond) (*bo.UserLoginResp, error)
	UserRegister(ctx context.Context, data *bo.UserRegCond) error
	UpdateUserInfo(ctx context.Context, cond *bo.UpdateUserInfoIdCond, data *bo.UpdateUserInfoCond) error
}

func ProvideUserSrv(userRepo repository.IUserRepo) IUserSrv {
	return &userService{userRepo}
}

type userService struct {
	userRepo repository.IUserRepo
}

func (srv *userService) GetUserList(ctx context.Context) ([]*bo.UserInfo, error) {
	poUserList, err := srv.userRepo.GetUserList(ctx)

	if err != nil {
		return nil, xerrors.Errorf("userService GetUserList error! : %w", err)
	}

	users := make([]*bo.UserInfo, len(poUserList))
	for i, poUser := range poUserList {
		users[i] = &bo.UserInfo{
			Id:          poUser.Id,
			Account:     poUser.Account,
			Password:    poUser.Password,
			Birthdate:   poUser.Birthdate,
			Nickname:    poUser.Nickname,
			Gender:      poUser.Gender,
			CountryCode: poUser.CountryCode,
			Address:     poUser.Address,
			PhoneNumber: poUser.PhoneNumber,
			CreateAt:    poUser.CreateAt,
			UpdateAt:    poUser.UpdateAt,
		}
	}

	return users, nil
}

func (srv *userService) GetUser(ctx context.Context, cond *bo.GetUserCond) (*bo.UserInfo, error) {
	poUserCond := &po.GetUserCond{
		Account: cond.Account,
	}

	poUser, err := srv.userRepo.GetUser(ctx, poUserCond)

	if err != nil {
		return nil, xerrors.Errorf("userService GetUser error! : %w", err)
	}

	user := &bo.UserInfo{
		Id:          poUser.Id,
		Account:     poUser.Account,
		Password:    poUser.Password,
		Birthdate:   poUser.Birthdate,
		Nickname:    poUser.Nickname,
		Gender:      poUser.Gender,
		CountryCode: poUser.CountryCode,
		Address:     poUser.Address,
		PhoneNumber: poUser.PhoneNumber,
		CreateAt:    poUser.CreateAt,
		UpdateAt:    poUser.UpdateAt,
	}

	return user, nil
}

func (srv *userService) UserLogin(ctx context.Context, cond *bo.UserLoginCond) (*bo.UserLoginResp, error) {
	poUserCond := &po.GetUserCond{
		Account: cond.Account,
	}

	poUser, err := srv.userRepo.GetUser(ctx, poUserCond)
	if err != nil {
		customErr, ok := errortool.ParseError(err)
		if ok && errors.Is(customErr, errortool.DbErr.NoRow) {
			return nil, xerrors.Errorf("userService UserLogin srv.userRepo.GetUser error! : %w", errortool.UserSrvErr.AccountOrPasswordError)
		}
		return nil, xerrors.Errorf("userService UserLogin srv.userRepo.GetUser error! : %w", err)
	}

	isCorrect, err := encryptUtil.PasswordCompare([]byte(poUser.Password), cond.Password)
	if err != nil || !isCorrect {
		return nil, xerrors.Errorf("userService UserLogin encrypt.EncryptPassword error! : %w", errortool.UserSrvErr.AccountOrPasswordError)
	}

	token, err := auth.TokenGenerate(cond.Account)
	if err != nil {
		return nil, xerrors.Errorf("userService UserLogin TokenGenerate error! : %w", err)
	}

	userLoginResp := &bo.UserLoginResp{
		Token: token,
	}

	return userLoginResp, nil
}

func (srv *userService) UserRegister(ctx context.Context, data *bo.UserRegCond) error {
	encrypted, err := encryptUtil.PasswordEncrypt(data.Password)
	if err != nil {
		return xerrors.Errorf("userService UserRegister encrypt.EncryptPassword error! : %w", err)
	}

	poUserReg := &po.UserRegCond{
		Account:  data.Account,
		Password: encrypted,
		Nickname: data.Nickname,
	}

	if err := srv.userRepo.UserRegister(ctx, poUserReg); err != nil {
		if errors.Is(err, errortool.DbErr.UniqueViolation) {
			return xerrors.Errorf("userService UserRegister error! : %w", errortool.UserSrvErr.AccountOrNicknameDuplicateError)
		}
		return xerrors.Errorf("userService UserRegister error! : %w", err)
	}

	return nil
}

func (srv *userService) UpdateUserInfo(ctx context.Context, cond *bo.UpdateUserInfoIdCond, data *bo.UpdateUserInfoCond) error {
	if data.Password != nil && *data.Password == "" {
		return xerrors.Errorf("userService UpdateUserInfo error! : %w", errortool.UserSrvErr.PasswordRequiredError)
	}

	if data.Gender != nil && (*data.Gender < 1 || *data.Gender > 3) {
		return xerrors.Errorf("userService UpdateUserInfo error! : %w", errortool.UserSrvErr.GenderMismatchError)
	}

	if data.CountryCode != nil && len(*data.CountryCode) > 10 {
		return xerrors.Errorf("userService UpdateUserInfo error! : %w", errortool.UserSrvErr.CountryCodeError)
	}

	poUpdateUserInfoIdCond := &po.UpdateUserInfoIdCond{}
	poUpdateUserInfoIdCond.Id = cond.Id
	poUser := &po.UpdateUserInfoCond{
		Nickname:    data.Nickname,
		Birthdate:   data.Birthdate,
		Gender:      data.Gender,
		CountryCode: data.CountryCode,
		Address:     data.Address,
		PhoneNumber: data.PhoneNumber,
	}

	if data.Password != nil {
		encrypted, err := encryptUtil.PasswordEncrypt(*data.Password)
		if err != nil {
			return xerrors.Errorf("userService UpdateUserInfo encrypt.EncryptPassword error! : %w", err)
		}
		poUser.Password = &encrypted
	}

	if err := srv.userRepo.UpdateUserInfo(ctx, poUpdateUserInfoIdCond, poUser); err != nil {
		return xerrors.Errorf("userService UpdateUserInfo error! : %w", err)
	}

	return nil
}
