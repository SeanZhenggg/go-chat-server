package controllers

import (
	"chat/utils/errortool"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type IResponseMiddleware interface {
	ResponseHandler(ctx *gin.Context)
}

func ProvideResponseMiddleware() IResponseMiddleware {
	return &ResponseMiddleware{}
}

type ResponseMiddleware struct{}

func (respMw *ResponseMiddleware) ResponseHandler(ctx *gin.Context) {
	// before request

	ctx.Next()

	// after request
	respMw.standardResponse(ctx)
}

func (respMw *ResponseMiddleware) generateStandardResponse(ctx *gin.Context) response {
	status := ctx.GetInt(Resp_Status)
	data := ctx.MustGet(Resp_Data)
	var code int
	var message string

	if status >= http.StatusBadRequest {
		if err, ok := data.(error); ok {
			if parsed, ok := errortool.UnwrapErrors(err); ok {
				code = 1234
				message = fmt.Sprintf("%v", parsed)
				data = nil
			}
		} else {
			code = 1005
			message = "unknown error"
			data = nil
		}
	}

	return response{
		Code:    code,
		Message: message,
		Data:    data,
	}
}

func (respMw *ResponseMiddleware) standardResponse(ctx *gin.Context) {
	response := respMw.generateStandardResponse(ctx)

	resp_status := ctx.GetInt(Resp_Status)

	ctx.JSON(
		resp_status,
		response,
	)
}

type response struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}
