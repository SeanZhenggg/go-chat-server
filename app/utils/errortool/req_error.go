package errortool

import (
	"github.com/SeanZhenggg/go-utils/errortool"
)

const (
	ReqGroupCode int = 2
)

func ProvideReqError(groups errortool.IGroupRepo, codes errortool.ICodeRepo) interface{} {
	group := Define.GenGroup(ReqGroupCode)

	return &reqError{
		AuthFailedError: group.GenError(1, "登入驗證失敗"),
	}
}

type reqError struct {
	AuthFailedError error
}
