package middleware

import (
	"chat/app/model/bo"
	"chat/app/service"
	"chat/app/utils/auth"
	"chat/app/utils/errortool"
	"chat/app/utils/logger"
	"github.com/gin-gonic/gin"
	"golang.org/x/xerrors"
	"net/http"
	"strings"
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
	tokenStr := ctx.GetHeader("Authorization")

	token, _ := strings.CutPrefix(tokenStr, "Bearer ")

	userAccount, err := auth.TokenValidation(token)
	if err != nil {
		respMw.logger.Error(xerrors.Errorf("authMiddleware AuthValidationHandler TokenValidation error : %w", err))
		SetResp(ctx, http.StatusUnauthorized, errortool.ReqErr.RequestTokenError)
		ctx.Abort()
		return
	}

	boGetUserCond := &bo.GetUserCond{
		Account: userAccount,
	}

	_, err = respMw.userSrv.GetUser(ctx, boGetUserCond)
	if err != nil {
		respMw.logger.Error(xerrors.Errorf("authMiddleware AuthValidationHandler TokenValidation error : %w", err))
		SetResp(ctx, http.StatusUnauthorized, errortool.ReqErr.RequestTokenError)
		ctx.Abort()
		return
	}

	ctx.Next()

	// after request
}
