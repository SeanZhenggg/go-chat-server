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
	}
}

type reqError struct {
	RequestTokenError               error
	AccountOrPasswordError          error
	AccountOrNicknameDuplicateError error
}
