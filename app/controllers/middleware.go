package controllers

import (
	"chat/app/utils/errortool"
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
			if parsed, ok := errortool.ParseError(err); ok {
				code = parsed.GetCode()
				message = parsed.GetMessage()
				data = nil
			}
		} else {
			err, _ := errortool.ParseError(errortool.ReqErr.UnknownError)
			code = err.GetCode()
			message = err.GetMessage()
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
