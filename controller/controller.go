package controller

import (
	"errors"
	"fmt"
	_ "fmt"
	"github.com/gin-gonic/gin"
	"goNav/middleware"
	"goNav/model"
	"goNav/util"
)

// 用户注册接口
func RegisterUser(c *gin.Context, registerInfo LoginReq) {
	// 用户注册
	var user = model.User{Username: registerInfo.Name, Password: registerInfo.Pwd}
	id, err := user.Insert()
	if err != nil {
		middleware.ResponseError(c, -1, errors.New("用户注册失败"))
		return
	}
	// 注册成功后 进行数据库迁移
	if json := util.LoadJson(); json != nil {
		for _, k := range json {
			k.UserID = id
			k.ShowIcon = true
			k.Insert()
		}
	}
	MakeToken(c, id, user.Username, "注册成功")
}

// 登陆结果
type LoginResult struct {
	Token string `json:"token"`
	// 用户模型
	Name string `json:"name"`
	//model.User
}

// LoginReq请求参数
type LoginReq struct {
	Name string `json:"username" form:"username" binding:"required"`
	Pwd  string `json:"password" form:"password" binding:"required"`
}

// 登陆接口 用户名和密码登陆
// name,password
func Login(c *gin.Context) {
	var loginReq LoginReq

	err := c.BindJSON(&loginReq)
	if err != nil {
		fmt.Println("register failed")
		middleware.ResponseError(c, -1, err)
		return
	}
	hasUser, user, err := model.LoginCheck(loginReq.Name, loginReq.Pwd)
	if err != nil {
		middleware.ResponseError(c, -1, err)
	} else {
		if hasUser {
			MakeToken(c, user.ID, user.Username, "登陆成功")
		} else {
			RegisterUser(c, loginReq)
		}
	}
}

func MakeToken(c *gin.Context, id int64, username string, msg string) {
	token, tokenErr := util.GenerateToken(id, username)
	if tokenErr != nil {
		middleware.ResponseError(c, -1, tokenErr)
	}
	// 获取用户相关数据
	data := LoginResult{
		Name:  username,
		Token: token,
	}
	middleware.ResponseSuccessMsg(c, data, msg)
}
