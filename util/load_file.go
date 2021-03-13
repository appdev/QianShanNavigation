package util

import (
	"encoding/json"
	"goNav/model"
	"io/ioutil"
	"os"
	"path/filepath"
)

func LoadJson() []model.WebSite {
	// 获取 参数，请传入一个文件路径
	///Users/lengyue/MyCode/golang/nav/static/static/userweb.json
	// ioutil 方式读取，会一次性读取整个文件，在对大文件处理时会有内存压力
	fileData, err := ioutil.ReadFile(getRootPath() + "static/static/userweb.json")
	dropErr(err)
	//&model.WebSite{}
	res := make([]model.WebSite, 10)
	if err := json.Unmarshal(fileData, &res); err != nil {
		return nil
	}
	return res
}

// 创建一个错误处理函数，避免过多的 if err != nil{} 出现
func dropErr(e error) {
	if e != nil {
		panic(e)
	}
}

func getRootPath() string {
	ex, err := os.Executable()
	if err != nil {
		panic(err)
	}
	exPath := filepath.Dir(ex)
	return exPath + "/"
}
