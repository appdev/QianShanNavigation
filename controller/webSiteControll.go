package controller

import (
	"errors"
	"github.com/gin-gonic/gin"
	"goNav/middleware"
	"goNav/model"
	"goNav/util"
	"net/http"
	"strconv"
)

type AddWebSite struct {
	Name     string `json:"name" form:"name" binding:"required"`
	Url      string `json:"url" form:"url" binding:"required"`
	Icon     string `json:"icon" form:"Icon"`
	Category string `json:"category" form:"category" binding:"required"`
	Weight   int    `json:"weight" form:"weight"`
	ShowIcon bool   `json:"showIcon" form:"weight"`
	Id       int64  `json:"id" form:"id"`
}
type Category struct {
	NewCategory string `json:"newCategory" form:"category" binding:"required"`
	OldCategory string `json:"oldCategory" form:"category" binding:"required"`
}

const ErrorImage = "https://static.apkdv.com/image/web.png!/format/webp/lossless/true"

func AddWeb(c *gin.Context) {
	addWeb, userId, err := checkData(c)
	if err != nil {
		return
	}
	webSite := model.WebSite{Url: addWeb.Url, UserID: userId, Name: addWeb.Name, Category: addWeb.Category}
	id, err := webSite.Insert()
	if err != nil {
		middleware.ResponseError(c, -1, errors.New("添加失败"))
		return
	}
	middleware.ResponseSuccessMsg(c, id, "添加成功")
}

func AddCategory(c *gin.Context) {
	addWeb, userId, err := checkData(c)
	if err != nil {
		return
	}
	webSite := model.WebSite{Url: addWeb.Url, UserID: userId, Name: addWeb.Name, Category: addWeb.Category}
	id, err := webSite.Insert()
	if err != nil {
		middleware.ResponseError(c, -1, errors.New("添加失败"))
		return
	}
	middleware.ResponseSuccessMsg(c, id, "添加成功")
}

func checkData(c *gin.Context) (AddWebSite, int64, error) {
	var addWeb AddWebSite
	var userId int64

	if id, ok := c.Get(util.USERID); ok {
		userId = id.(int64)
	} else {
		parameterError(c)
		return AddWebSite{}, 0, errors.New("参数错误")
	}

	err := c.BindJSON(&addWeb)
	if err != nil {
		middleware.ResponseError(c, -1, err)
		return AddWebSite{}, 0, err
	}
	return addWeb, userId, nil
}

func parameterError(c *gin.Context) {
	middleware.ResponseError(c, -1, errors.New("参数错误"))
}

func Update(c *gin.Context) {
	addWeb, _, err := checkData(c)
	if err != nil {
		return
	}

	webSite := model.WebSite{Url: addWeb.Url, Name: addWeb.Name, Category: addWeb.Category}
	webSite.Favicon = util.Fetch(addWeb.Url)
	_, err = webSite.Update(addWeb.Id)
	if err != nil {
		middleware.ResponseError(c, -1, errors.New("修改失败"))
		return
	}
	middleware.ResponseSuccess(c, addWeb.Id)
}

func UpdateCategory(c *gin.Context) {
	var category Category
	var userId int64

	if id, ok := c.Get(util.USERID); ok {
		userId = id.(int64)
	} else {
		parameterError(c)
		middleware.ResponseError(c, -1, errors.New("参数错误"))
		return
	}
	err := c.BindJSON(&category)
	if err != nil {
		middleware.ResponseError(c, -1, err)
		return
	}
	webSite := model.WebSite{UserID: userId, Category: category.OldCategory}
	if err := webSite.ChangeCategory(category.NewCategory); err != nil {
		middleware.ResponseError(c, -1, errors.New("修改失败"))
		return
	}
	middleware.ResponseSuccess(c, "修改成功")
}

func Delete(c *gin.Context) {
	ids := c.Query("id")
	id, _ := strconv.ParseInt(ids, 10, 64)
	webSite := model.WebSite{ID: id}
	resId, err := webSite.Destroy(id)
	if err != nil {
		middleware.ResponseError(c, -1, errors.New("删除失败"))
		return
	}
	middleware.ResponseSuccess(c, resId.ID)
}

func Finds(c *gin.Context) {
	var userId int64
	if id, ok := c.Get(util.USERID); ok {
		userId = id.(int64)
	} else {
		parameterError(c)
	}
	if userId != 0 {
		user := model.User{ID: userId}
		webs, _ := user.User()
		middleware.ResponseSuccess(c, webs)
	} else {
		parameterError(c)
	}
}

func GetNewBg(c *gin.Context) {
	if image, err := util.RootHandler(); err == nil {
		middleware.ResponseSuccess(c, image)
	} else {
		middleware.ResponseError(c, -1, errors.New("获取最新图片失败"))
	}
}
func GetBg(c *gin.Context) {
	if image, err := util.RandomHandler(); err == nil {
		middleware.ResponseSuccess(c, image)
	} else {
		middleware.ResponseError(c, -1, errors.New("获取最新图片失败"))
	}
}
func Favicon(c *gin.Context) {
	if requestUrl, ok := c.GetQuery("url"); ok {

		//if urlObj, parseErr := url.Parse(requestUrl); parseErr == nil {
		getFaviconByNet(c, requestUrl)
		//if data, err := model.FindOne(urlObj.Host); err == nil {
		//	if data.Favicon == ErrorImage {
		//		getFaviconByNet(c,parseErr,requestUrl,urlObj)
		//	} else {
		//		c.String(http.StatusOK, data.Favicon)
		//	}
		//} else {
		//	getFaviconByNet(c, err, requestUrl, urlObj)
		//}
		//} else {
		//	c.String(http.StatusOK, ErrorImage)
		//}
	} else {
		c.String(http.StatusOK, ErrorImage)
	}

}

func getFaviconByNet(c *gin.Context, requestUrl string) {
	iconUrl := util.Fetch(requestUrl)
	c.String(http.StatusOK, iconUrl)
}
