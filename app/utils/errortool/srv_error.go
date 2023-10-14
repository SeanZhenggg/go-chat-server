package errortool

import (
	"github.com/SeanZhenggg/go-utils/errortool"
)

const (
	UserSrvGroupCode int = 3
)

func ProvideUserSrvError(groups errortool.IGroupRepo, codes errortool.ICodeRepo) interface{} {
	group := Define.GenGroup(UserSrvGroupCode)

	return &userSrvError{
		AccountOrPasswordError:          group.GenError(1, "帳號或密碼錯誤"),
		AccountOrNicknameDuplicateError: group.GenError(2, "帳號或暱稱重複"),
		PasswordRequiredError:           group.GenError(3, "密碼不得為空"),
		GenderMismatchError:             group.GenError(4, "性別欄位只可為男性、女性或不公開"),
		CountryCodeError:                group.GenError(5, "國籍輸入有誤"),
	}
}

type userSrvError struct {
	AccountOrPasswordError          error
	AccountOrNicknameDuplicateError error
	PasswordRequiredError           error
	GenderMismatchError             error
	CountryCodeError                error
}
