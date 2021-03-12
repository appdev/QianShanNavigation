package api

import (
	"errors"
	"github.com/gin-gonic/gin"
	"goNav/middleware"
	"goNav/model"
	"strconv"
)

//列表数据
func Users(c *gin.Context) {
	var user model.User
	user.Username = c.Request.FormValue("username")
	user.Password = c.Request.FormValue("password")
	result, err := user.Users()

	if err != nil {
		responseFail(c,"抱歉未找到相关信息")
		return
	}
	middleware.ResponseSuccess(c, result)
}

func responseFail(c *gin.Context, msg string) {
	middleware.ResponseError(c, -1, errors.New(msg))
}

//添加数据
func Store(c *gin.Context) {
	var user model.User
	user.Username = c.Request.FormValue("username")
	user.Password = c.Request.FormValue("password")
	id, err := user.Insert()

	if err != nil {
		responseFail(c,"添加失败")
		return
	}
	middleware.ResponseSuccessMsg(c, id, "添加成功")
}

//修改数据
func Update(c *gin.Context) {
	var user model.User
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	user.Password = c.Request.FormValue("password")
	result, err := user.Update(id)
	if err != nil || result.ID == 0 {
		responseFail(c,"修改失败")
		return
	}
	middleware.ResponseSuccessMsg(c, "", "修改成功")
}

//删除数据
func Destroy(c *gin.Context) {
	var user model.User
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	result, err := user.Destroy(id)
	if err != nil || result.ID == 0 {
		responseFail(c,"删除失败")
		return
	}
	middleware.ResponseSuccessMsg(c, "", "删除成功")

}
