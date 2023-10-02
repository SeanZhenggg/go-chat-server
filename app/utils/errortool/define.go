package errortool

var (
	Define    = ProvideDefine()
	CommonErr = Define.Plugin(ProvideCommonError).(*commonError)
	ReqErr    = Define.Plugin(ProvideReqError).(*reqError)
	DbErr     = Define.Plugin(ProvideDBError).(*dbError)
)
