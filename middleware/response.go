package middleware

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
)

const (
	SuccessCode ResponseCode = 200
)

type ResponseCode int

type Response struct {
	Code ResponseCode `json:"code"`
	Msg  string       `json:"msg"`
	Data interface{}  `json:"data"`
}

func ResponseError(c *gin.Context, code ResponseCode, err error) {
	resp := &Response{Code: code, Msg: err.Error(), Data: nil}
	c.JSON(-1, resp)
	response, _ := json.Marshal(resp)
	c.Set("response", string(response))
	//c.AbortWithError(200, err)
}

func ResponseSuccess(c *gin.Context, data interface{}) {
	resp := &Response{Code: SuccessCode, Msg: "", Data: data}
	c.JSON(200, resp)
	response, _ := json.Marshal(resp)
	c.Set("response", string(response))
}
func ResponseSuccessMsg(c *gin.Context, data interface{}, msg string) {
	resp := &Response{Code: SuccessCode, Msg: msg, Data: data}
	c.JSON(200, resp)
	response, _ := json.Marshal(resp)
	c.Set("response", string(response))
}
