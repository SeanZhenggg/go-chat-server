package errortool

var (
	Define = ProvideDefine()
	ReqErr = Define.Plugin(ProvideReqError).(*reqError)
	DbErr  = Define.Plugin(ProvideDBError).(*dbError)
)
