package middleware

import (
	"chat/app/model/bo"
	"chat/app/service"
	"chat/app/utils/auth"
	ctxUtil "chat/app/utils/ctx"
	"chat/app/utils/errs"
	"chat/app/utils/logger"
	"github.com/gin-gonic/gin"
	"golang.org/x/xerrors"
	"net/http"
	"strings"
)

const (
	HEADER_AUTHORIZATION = "Authorization"
	ATHORIZATION_PREFIX  = "Bearer "
)

type IAuthMiddleware interface {
	AuthValidationHandler(ctx *gin.Context)
}

func ProvideAuthMiddleware(logger logger.ILogger, userSrv service.IUserSrv) IAuthMiddleware {
	return &AuthMiddleware{
		logger:  logger,
		userSrv: userSrv,
	}
}

type AuthMiddleware struct {
	logger  logger.ILogger
	userSrv service.IUserSrv
}

func (respMw *AuthMiddleware) AuthValidationHandler(ctx *gin.Context) {
	// before request
	tokenStr := ctx.GetHeader(HEADER_AUTHORIZATION)

	token, _ := strings.CutPrefix(tokenStr, ATHORIZATION_PREFIX)

	userAccount, err := auth.TokenValidation(token)
	if err != nil {
		respMw.logger.Error(xerrors.Errorf("authMiddleware AuthValidationHandler TokenValidation error : %w", err))
		SetResp(ctx, http.StatusUnauthorized, errs.ReqErr.AuthFailedError)
		ctx.Abort()
		return
	}

	boGetUserCond := &bo.GetUserCond{
		Account: userAccount,
	}

	userInfo, err := respMw.userSrv.GetUser(ctx, boGetUserCond)
	if err != nil {
		respMw.logger.Error(xerrors.Errorf("authMiddleware AuthValidationHandler TokenValidation error : %w", err))
		SetResp(ctx, http.StatusUnauthorized, errs.ReqErr.AuthFailedError)
		ctx.Abort()
		return
	}
	ctxUtil.SetUserInfo(ctx, userInfo)

	ctx.Next()

	// after request
}
