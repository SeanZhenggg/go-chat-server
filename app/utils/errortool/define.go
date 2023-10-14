package errortool

import (
	"github.com/SeanZhenggg/go-utils/errortool"
)

var (
	Define     = errortool.ProvideDefine()
	CommonErr  = Define.Plugin(ProvideCommonError).(*commonError)
	ReqErr     = Define.Plugin(ProvideReqError).(*reqError)
	UserSrvErr = Define.Plugin(ProvideUserSrvError).(*userSrvError)
	DbErr      = Define.Plugin(ProvideDBError).(*dbError)
)
