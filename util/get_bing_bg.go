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
	URL  string `json:"url"`
	Text string `json:"text"`
}

const (
	host            = `https://cn.bing.com`
	bingAPI         = `https://cn.bing.com/HPImageArchive.aspx?format=js&idx=0&n=1`
	randomWallpaper = `http://api.cucldk.com/bing.php?f=json&key=%E9%A3%8E%E6%99%AF`
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
		return Images{URL: res.URL, Copyright: res.Text}, nil
	}

	return Images{}, err
}
