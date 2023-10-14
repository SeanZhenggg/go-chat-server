package errortool

import "github.com/SeanZhenggg/go-utils/errortool"

const (
	CommonGroupCode int = 1
)

func ProvideCommonError(groups errortool.IGroupRepo, codes errortool.ICodeRepo) interface{} {
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
