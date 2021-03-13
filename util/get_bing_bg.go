package util

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
)

type Response struct {
	Images []Images `json:"images"`
}
type Images struct {
	Startdate     string `json:"startdate"`
	Enddate       string `json:"enddate"`
	URL           string `json:"url"`
	Copyright     string `json:"copyright"`
	Copyrightlink string `json:"copyrightlink"`
	Title         string `json:"title"`
}

type RandomWallpaper struct {
	Data struct {
		Enddate   string `json:"enddate"`
		URL       string `json:"url"`
		Copyright string `json:"copyright"`
	} `json:"data"`
}

const (
	host            = `https://cn.bing.com`
	bingAPI         = `https://cn.bing.com/HPImageArchive.aspx?format=js&idx=0&n=1`
	randomWallpaper = `https://bing.ioliu.cn/v1/rand?type=json`
)

func RootHandler() (Images, error) {

	resp, err := http.Get(bingAPI)
	if err != nil {
		return Images{}, err
	}
	defer resp.Body.Close()
	if body, err := ioutil.ReadAll(resp.Body); err == nil {
		var res Response
		if err := json.Unmarshal(body, &res); err != nil {
			return Images{}, err
		}
		image := res.Images[0]
		image.URL = host + image.URL
		return image, nil
	}

	return Images{}, err
}

func RandomHandler() (Images, error) {

	resp, err := http.Get(randomWallpaper)
	if err != nil {
		return Images{}, err
	}
	defer resp.Body.Close()
	if body, err := ioutil.ReadAll(resp.Body); err == nil {
		var res RandomWallpaper
		if err := json.Unmarshal(body, &res); err != nil {
			return Images{}, err
		}
		return Images{URL: res.Data.URL, Copyright: res.Data.Copyright}, nil
	}

	return Images{}, err
}
