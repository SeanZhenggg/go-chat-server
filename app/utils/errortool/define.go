package errortool

var (
	Define     = ProvideDefine()
	CommonErr  = Define.Plugin(ProvideCommonError).(*commonError)
	ReqErr     = Define.Plugin(ProvideReqError).(*reqError)
	UserSrvErr = Define.Plugin(ProvideUserSrvError).(*userSrvError)
	DbErr      = Define.Plugin(ProvideDBError).(*dbError)
)
