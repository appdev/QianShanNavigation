package main

import (
	"encoding/json"
	"fmt"
	"github.com/gin-contrib/gzip"
	"github.com/gin-gonic/gin"
	_ "goNav/api/database"
	"goNav/middleware"
	"goNav/model"
	"goNav/router"
	"goNav/util"
	"net/http"
)

func main() {
	engine := gin.Default()
	engine.Use(middleware.Cors())
	engine.Use(gzip.Gzip(gzip.DefaultCompression))

	router.InitRouter(engine)
	//加载静态资源，例如网页的css、js
	//engine.Use(middleware.Serve("/", middleware.LocalFile("static", false)))
	//engine.LoadHTMLGlob(getRootPath() + "static/views/*")
	//engine.GET("/", GetIndex)
	engine.Run(":9080")

}

func loadJsonFile() {
	webSite := make([]model.WebSite, 0)
	if json := util.LoadJson(); json != nil {
		for _, k := range json {
			k.Favicon = util.Fetch(k.Url)
			webSite = append(webSite, k)
		}
	}
	bytes, _ := json.Marshal(webSite)
	fmt.Println(string(bytes))
}

func GetIndex(c *gin.Context) {
	c.HTML(http.StatusOK, "index.html", gin.H{})
}
