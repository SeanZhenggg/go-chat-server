package errs

import (
	"github.com/SeanZhenggg/go-utils/errortool"
)

var (
	Define = errortool.ProvideDefine().
		SetErrCodeOptions(errortool.ErrCodeOptions{
			Min:   1,
			Max:   999,
			Range: 1000,
		})
	CommonErr  = ProvideCommonError()
	ReqErr     = ProvideReqError()
	UserSrvErr = ProvideUserSrvError()
	DbErr      = ProvideDBError()
)
