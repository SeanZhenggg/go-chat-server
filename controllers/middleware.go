package controllers

import (
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
	// status := ctx.GetInt(Resp_Status)
	// data := ctx.MustGet(Resp_Data)
	// var code int
	// var message string

	// var newError error
	// if status >= http.StatusBadRequest {
	// 	if tempErr, ok := data.(error); ok {
	// 		for {
	// 			if tmp := errors.Unwrap(tempErr); tmp != nil {
	// 				newError = tmp
	// 			} else {
	// 				break
	// 			}
	// 		}

	// 	}
	// }

	return response{
		// 	Code:    code,
		// 	Message: message,
		// 	Data:    data,
	}
}

func (respMw *ResponseMiddleware) standardResponse(ctx *gin.Context) {
	response := respMw.generateStandardResponse(ctx)

	resp_status := ctx.GetInt(Resp_Status)
	if resp_status >= http.StatusBadRequest {
		response.Data = nil
	}

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
