package common

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type BaseResponse struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

type ScesResponse struct {
	BaseResponse
	Data interface{} `json:"data"`
}

type ErrResponse struct {
	BaseResponse
	ErrDetails error `json:"errDetails"`
}

func NewBaseResp(code int, message string) BaseResponse {
	return BaseResponse{
		Code:    code,
		Message: message,
	}
}

func NewResp() BaseResponse {
	return NewBaseResp(0, "success")
}

func NewScesResp(data interface{}) ScesResponse {
	return ScesResponse{
		BaseResponse: NewResp(),
		Data:         data,
	}
}

func NewErrResp(code int, message string, err error) ErrResponse {
	resp := ErrResponse{
		BaseResponse: NewBaseResp(code, message),
		ErrDetails:   err,
	}
	return resp
}

func SetScesResp(ctx *gin.Context, data interface{}) {
	ctx.JSON(http.StatusOK, NewScesResp(data))
}

func SetErrResp(ctx *gin.Context, code int, message string, err error) {
	ctx.JSON(code, NewErrResp(code, message, err))
}
