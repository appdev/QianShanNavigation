package main

import (
	"fmt"
	"github.com/gin-contrib/gzip"
	"github.com/gin-gonic/gin"
	_ "goNav/api/database"
	"goNav/middleware"
	"goNav/router"
	"goNav/util"
	"net/http"
	"os"
	"path/filepath"
)

func main() {

	res, _ := util.RootHandler()
	fmt.Println(res)
	engine := gin.Default()
	engine.Use(middleware.Cors())
	engine.Use(gzip.Gzip(gzip.DefaultCompression))

	router.InitRouter(engine)
	//加载静态资源，例如网页的css、js
	engine.Use(middleware.Serve("/", middleware.LocalFile("static", false)))
	engine.LoadHTMLGlob(getRootPath() + "static/views/*")
	engine.GET("/", GetIndex)

	engine.Run(":9080")
}

func GetIndex(c *gin.Context) {
	c.HTML(http.StatusOK, "index.html", gin.H{})
}

func getRootPath() string {
	ex, err := os.Executable()
	if err != nil {
		panic(err)
	}
	exPath := filepath.Dir(ex)
	return exPath + "/"
}
