package errortool

const (
	CommonGroupCode int = 1
)

func ProvideCommonError(groups IGroupRepo, codes ICodeRepo) interface{} {
	group := Define.GenGroup(CommonGroupCode)

	return &commonError{
		UnknownError:      group.GenError(1, "未知錯誤"),
		RequestParamError: group.GenError(2, "請求參數錯誤"),
	}
}

type commonError struct {
	UnknownError      error
	RequestParamError error
}
