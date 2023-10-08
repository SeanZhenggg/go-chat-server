package middleware

import (
	"chat/app/utils/auth"
	"chat/app/utils/errortool"
	"chat/app/utils/logger"
	"fmt"
	"github.com/gin-gonic/gin"
	"golang.org/x/xerrors"
	"net/http"
	"strings"
)

type IAuthMiddleware interface {
	AuthValidationHandler(ctx *gin.Context)
}

func ProvideAuthMiddleware(logger logger.ILogger) IAuthMiddleware {
	return &AuthMiddleware{
		logger: logger,
	}
}

type AuthMiddleware struct {
	logger logger.ILogger
}

func (respMw *AuthMiddleware) AuthValidationHandler(ctx *gin.Context) {
	// before request
	tokenStr := ctx.GetHeader("Authorization")

	token, _ := strings.CutPrefix(tokenStr, "Bearer ")

	_, err := auth.TokenValidation(token)
	if err != nil {
		fmt.Printf("err : %v\n", err)
		respMw.logger.Error(xerrors.Errorf("authMiddleware AuthValidationHandler TokenValidation error : %w", err))
		SetResp(ctx, http.StatusUnauthorized, errortool.ReqErr.RequestTokenError)
		ctx.Abort()
		return
	}

	ctx.Next()

	// after request
}
