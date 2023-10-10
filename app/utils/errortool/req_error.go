package errortool

const (
	ReqGroupCode int = 2
)

func ProvideReqError(groups IGroupRepo, codes ICodeRepo) interface{} {
	group := Define.GenGroup(ReqGroupCode)

	return &reqError{
		AuthFailedError: group.GenError(1, "登入驗證失敗"),
	}
}

type reqError struct {
	AuthFailedError error
}
