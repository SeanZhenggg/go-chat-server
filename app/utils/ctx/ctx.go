package ctx

import (
	"chat/app/model/bo"
	"chat/app/utils/errs"
	"github.com/gin-gonic/gin"
)

const (
	UserInfo = "UserInfo"
)

func SetUserInfo(ctx *gin.Context, info *bo.UserInfo) {
	if info == nil {
		return
	}
	ctx.Set(UserInfo, info)
}

func GetUserInfo(ctx *gin.Context) (*bo.UserInfo, error) {
	val := ctx.Value(UserInfo)

	if val != nil {
		if userInfo, ok := val.(*bo.UserInfo); ok {
			return userInfo, nil
		}
	}

	return nil, errs.ReqErr.AuthFailedError
}
