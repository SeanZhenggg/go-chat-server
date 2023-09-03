package errortool

func ProvideReqError(groups IGroupRepo, codes ICodeRepo) interface{} {
	group := Define.GenGroup(1)

	return &reqError{
		UnknownError:      group.GenError(1, "未知錯誤"),
		RequestParamError: group.GenError(2, "請求參數錯誤"),
		RequestTokenError: group.GenError(3, "Token驗證失敗"),
	}
}

type reqError struct {
	UnknownError      error
	RequestParamError error
	RequestTokenError error
}
