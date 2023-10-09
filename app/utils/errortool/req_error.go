package errortool

const (
	ReqGroupCode int = 2
)

func ProvideReqError(groups IGroupRepo, codes ICodeRepo) interface{} {
	group := Define.GenGroup(ReqGroupCode)

	return &reqError{
		RequestTokenError:               group.GenError(1, "Token驗證失敗"),
		AccountOrPasswordError:          group.GenError(2, "帳號或密碼錯誤"),
		AccountOrNicknameDuplicateError: group.GenError(3, "帳號或暱稱重複"),
		PasswordRequiredError:           group.GenError(4, "密碼不得為空"),
		GenderMismatchError:             group.GenError(5, "性別欄位只可為男性、女性或不公開"),
		CountryCodeError:                group.GenError(6, "國籍輸入有誤"),
	}
}

type reqError struct {
	RequestTokenError               error
	AccountOrPasswordError          error
	AccountOrNicknameDuplicateError error
	PasswordRequiredError           error
	GenderMismatchError             error
	CountryCodeError                error
}
