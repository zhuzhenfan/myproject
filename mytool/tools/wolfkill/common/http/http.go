// handler
package http

import (
	"github.com/gin-gonic/gin"
)

var CodeSuccess = 200       // 返回成功
var CodeCreateSuccess = 201 // 请求成功并且服务器创建了新的资源
var CodeParamError = 400    // 参数有误
var CodeAuthError = 401     // auth认证出错
var CodeNotFound = 404      // 未找到
var CodeConflict = 409      // 请求与当前状态冲突
var CodeUnknowError = 500   // 未知错误

const SuccessStr = "success"

type Response struct {
	// 返回码
	// 200表示成功,其他表示失败
	Code int `json:"code" example:"200"`

	// 返回的错误信息
	Message string `json:"msg,omitempty"`

	// 返回数据
	Data interface{} `json:"data,omitempty"`
}

type Success struct {
	// 返回码
	// 200 表示成功,其他表示失败
	Code int `json:"code" example:"200"`

	// 返回数据
	Data interface{} `json:"data,omitempty"`
}

type Failure struct {
	// 返回码
	// 200表示成功,其他表示失败
	Code int `json:"code" example:"500"`

	// 返回错误信息
	Message string `json:"msg,omitempty"`
}

func Failed(c *gin.Context, err error) {
	//	logger.Error(result)
	response := Failure{
		Code:    CodeUnknowError,
		Message: err.Error(),
	}
	c.JSON(CodeUnknowError, response)
	c.Abort()
}

func ParamErr(c *gin.Context, err error) {
	//logger.Error(result)
	response := Failure{
		Code:    CodeParamError,
		Message: err.Error(),
	}
	c.JSON(CodeParamError, response)
	c.Abort()

}

func Ok(c *gin.Context, result interface{}) {
	response := Success{
		Code: CodeSuccess,
		Data: result,
	}
	c.JSON(CodeSuccess, response)
	c.Abort()
}
