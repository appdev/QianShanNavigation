package util

import (
	"encoding/json"
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"goNav/model"
	"log"
	"net/url"
	"regexp"
	"strings"
)

// Main access point. Given a uri, return the favicon
func Fetch(uri string) string {
	iconUrl := Detect(uri)
	if iconUrl == "" {
		iconUrl = "https://static.apkdv.com/image/web.png!/format/webp/lossless/true"
	}
	return iconUrl
}

// Attempt to get the url's HTML
func Detect(uri string) string {
	// remove trailing .ico
	re := regexp.MustCompile(".ico")
	uri = re.ReplaceAllString(uri, "")
	// add http scheme if needed so go url doesn't throw error
	hasHttp, _ := regexp.MatchString("^http", uri)
	if hasHttp != true {
		uri = "http://" + uri
	}
	urlObj, parseErr := url.Parse(uri)
	doc, queryErr := goquery.NewDocument(urlObj.String())
	if parseErr == nil && queryErr == nil {
		return FindFaviconUriInHTML(urlObj, doc)
	} else {
		if parseErr != nil {
			scheme := urlObj.Scheme
			if scheme == "" {
				scheme = "http"
			}
			host := urlObj.Host
			newUri := scheme + "://" + host
			urlObj, parseErr := url.Parse(newUri)
			if parseErr == nil {
				doc, queryErr := goquery.NewDocument(urlObj.String())
				if queryErr == nil {
					return FindFaviconUriInHTML(urlObj, doc)
				} else {
					return ""
				}
			} else {
				return ""
			}
		} else {
			return ""
		}
	}
}

// Look for <link rel="icon" and any base url
func FindFaviconUriInHTML(uri *url.URL, doc *goquery.Document) string {
	base, iconUrl := HTMLParserHandler(doc)
	// replace urls that start with  // path since go http cannot retrieve them
	re := regexp.MustCompile("^(//)")
	iconUrl = re.ReplaceAllString(iconUrl, uri.Scheme+"://")
	base = re.ReplaceAllString(base, uri.Scheme+"://")
	relIconUrl, _ := regexp.MatchString("^/", iconUrl)
	if base == "" {
		base = uri.String()
	}
	// remove trailing base slash
	trailingSlash := regexp.MustCompile("/$")
	base = trailingSlash.ReplaceAllString(base, "")

	if iconUrl == "" {
		return "https://besticon-demo.herokuapp.com/icon?url=" + base + "&size=24..32..64"

	} else {
		base = re.ReplaceAllString(base, uri.Scheme+"://")
		parseIconUrl, err := url.Parse(iconUrl)
		baseUrl, err := url.Parse(base)

		if err == nil {
			if parseIconUrl.Scheme == "" {
				if relIconUrl {
					iconUrl = uri.Scheme + "://" + baseUrl.Host + iconUrl
				} else {
					iconUrl = uri.Scheme + "://" + baseUrl.Host + "/" + iconUrl
				}
			}
			return iconUrl
		}

		return ""

	}

}

// parse the HTML to get the favicon url
func HTMLParserHandler(doc *goquery.Document) (string, string) {
	base := ""
	uri := ""

	// goroutine and channel to look for favicon
	uriChannel := make(chan string)
	go func() {
		doc.Find("link").Each(func(i int, s *goquery.Selection) {
			rel, relExists := s.Attr("rel")
			if relExists == true {
				shortcutIcon, _ := regexp.MatchString("(?i)^(shortcut )?icon$", rel)
				if shortcutIcon == true {
					tagUri, uriExists := s.Attr("href")
					if uriExists == true {
						uriChannel <- tagUri
					}
				}
			}
		})
		if len(uriChannel) == 0 {
			uriChannel <- ""
		}

	}()
	baseChannel := make(chan string)
	go func() {
		doc.Find("base").Each(func(i int, s *goquery.Selection) {
			baseUri, hrefExists := s.Attr("href")
			if hrefExists == true {
				baseChannel <- baseUri
			}
		})
		if len(baseChannel) == 0 {
			baseChannel <- ""
		}
	}()
	base = <-baseChannel
	uri = <-uriChannel
	return base, uri
}
func ExampleScrape() {
	html := `<div class="list closed">
        <ul>
          <!------>
            <a rel="nofollow" size="视频媒体" href="https://www.youtube.com/" target="_blank"><svg class="icon" aria-hidden="true"><use xlink:href="#icon-youtube"></use></svg> Youtube</a>
            <a rel="nofollow" size="视频媒体" href="https://v.qq.com/" target="_blank"><svg class="icon" aria-hidden="true"><use xlink:href="#icon-tengxunshipin"></use></svg>腾讯视频</a>
            <a rel="nofollow" size="视频媒体" href="https://www.youku.com/" target="_blank"><svg class="icon" aria-hidden="true"><use xlink:href="#icon-youku"></use></svg>优酷</a>
            <a rel="nofollow" size="视频媒体" href="https://www.iqiyi.com/" target="_blank"><svg class="icon" aria-hidden="true"><use xlink:href="#icon-aiqiyi"></use></svg>爱奇艺</a>
            <a rel="nofollow" size="视频媒体" href="http://www.acfun.cn/index.html" target="_blank"><svg class="icon" aria-hidden="true"><use xlink:href="#icon-acfun"></use></svg>ACFUN</a>
            <a rel="nofollow" size="视频媒体" href="https://www.bilibili.com/" target="_blank"><svg class="icon" aria-hidden="true"><use xlink:href="#icon-bili"></use></svg>哔哩哔哩</a>
            <a rel="nofollow" size="视频媒体" href="https://www.nfmovies.com/" target="_blank"><svg class="icon" aria-hidden="true"><use xlink:href="#icon-netflix"></use></svg>奈非影视</a>
            <a rel="nofollow" size="视频媒体" href="https://k1080.net/" target="_blank"><svg class="icon" aria-hidden="true"><use xlink:href="#icon-yingshi"></use></svg>K1080</a>
            <a rel="nofollow" size="视频媒体" href="https://www.yunbtv.com/" target="_blank"><svg class="icon" aria-hidden="true"><use xlink:href="#icon-yunbo"></use></svg>云播TV</a>
          <!------> 
            <a rel="nofollow" size="邮箱" href="https://mail.google.com/mail/u/0/#inbox" target="_blank"><svg class="icon" aria-hidden="true"><use xlink:href="#icon-gmail"></use></svg>Gmail</a>
            <a rel="nofollow" size="邮箱" href="https://outlook.live.com/mail/" target="_blank"><svg class="icon" aria-hidden="true"><use xlink:href="#icon-windows"></use></svg>Hotmail</a>
            <a rel="nofollow" size="邮箱" href="https://mail.163.com/" target="_blank"><svg class="icon" aria-hidden="true"><use xlink:href="#icon-wangyi"></use></svg>网易邮箱</a>
            <a rel="nofollow" size="邮箱" href="https://mail.sina.com.cn/" target="_blank"><svg class="icon" aria-hidden="true"><use xlink:href="#icon-xinlangwang"></use></svg>新浪邮箱</a>
            <a rel="nofollow" size="邮箱" href="https://mail.qq.com/" target="_blank"><svg class="icon" aria-hidden="true"><use xlink:href="#icon-QQ"></use></svg>QQ邮箱</a>
            <a rel="nofollow" size="邮箱" href="https://qiye.aliyun.com/" target="_blank"><svg class="icon" aria-hidden="true"><use xlink:href="#icon-yunyouxiang"></use></svg>阿里邮箱</a>
          <!------>
            <a rel="nofollow" size="社交资讯" href="https://www.weibo.com" target="_blank"><svg class="icon" aria-hidden="true"><use xlink:href="#icon-weibo"></use></svg>微博</a>
            <a rel="nofollow" size="社交资讯" href="https://www.zhihu.com/" target="_blank"><svg class="icon" aria-hidden="true"><use xlink:href="#icon-zhihu"></use></svg>知乎</a>
            <a rel="nofollow" size="社交资讯" href="https://www.douban.com/" target="_blank"><svg class="icon" aria-hidden="true"><use xlink:href="#icon-douban"></use></svg>豆瓣</a>
            <a rel="nofollow" size="社交资讯" href="https://www.jianshu.com/" target="_blank"><svg class="icon" aria-hidden="true"><use xlink:href="#icon-jianshu"></use></svg>简书</a>
            <a rel="nofollow" size="社交资讯" href="https://www.v2ex.com/" target="_blank"><svg class="icon" aria-hidden="true"><use xlink:href="#icon-v2ex"></use></svg>V2EX</a>
            <a rel="nofollow" size="社交资讯" href="https://www.instagram.com" target="_blank"><svg class="icon" aria-hidden="true"><use xlink:href="#icon-instagram"></use></svg>Instagram</a>
            <a rel="nofollow" size="社交资讯" href="https://www.twitter.com" target="_blank"><svg class="icon" aria-hidden="true"><use xlink:href="#icon-twitter"></use></svg>Twitter</a>
            <a rel="nofollow" size="社交资讯" href="https://www.facebook.com" target="_blank"><svg class="icon" aria-hidden="true"><use xlink:href="#icon-facebook"></use></svg>Facebook</a>
          <!------>
            <a rel="nofollow" size="购物" href="https://www.taobao.com" target="_blank"><svg class="icon" aria-hidden="true"><use xlink:href="#icon-taobao"></use></svg>淘宝网</a>
            <a rel="nofollow" size="购物" href="https://dyartstyle.com/juhuasuan/" target="_blank"><svg class="icon" aria-hidden="true"><use xlink:href="#icon-juhuasuan"></use></svg>聚划算</a>
            <a rel="nofollow" size="购物" href="https://dyartstyle.com/temai/" target="_blank"><svg class="icon" aria-hidden="true"><use xlink:href="#icon-temai"></use></svg>淘宝特卖</a>
            <a rel="nofollow" size="购物" href="https://mobile.yangkeduo.com/duo_cms_mall.html?pid=8742514_71211367&cpsSign=CM_190605_8742514_71211367_92fdd9f7fb637fec599bf556f263ed1f&duoduo_type=2&launch_wx=1" title="拼多多手机端专属商城，超低价商品火热抢购中，更有超多大额优惠券持续发放，一键立抢>>" target="_blank"><svg class="icon" aria-hidden="true"><use xlink:href="#icon-99"></use></svg>拼多多</a>
            <a rel="nofollow" size="购物" href="https://www.jd.com" target="_blank"><svg class="icon" aria-hidden="true"><use xlink:href="#icon-jingdong"></use></svg>京东</a>
            <a rel="nofollow" size="购物" href="https://www.suning.com/" target="_blank"><svg class="icon" aria-hidden="true"><use xlink:href="#icon-suning"></use></svg>苏宁易购</a>
            <a rel="nofollow" size="购物" href="http://you.163.com/" target="_blank"><svg class="icon" aria-hidden="true"><use xlink:href="#icon-yanxuan"></use></svg>网易严选</a>
            <a rel="nofollow" size="购物" href="https://www.amazon.cn/" target="_blank"><svg class="icon" aria-hidden="true"><use xlink:href="#icon-amazon"></use></svg>亚马逊</a>
            <a rel="nofollow" size="购物" href="http://www.dangdang.com/" target="_blank"><svg class="icon" aria-hidden="true"><use xlink:href="#icon-dangdang"></use></svg>当当</a>
            <a rel="nofollow" size="购物" href="https://wat.dyartstyle.com/" target="_blank"><svg class="icon" aria-hidden="true"><use xlink:href="#icon-wat"></use></svg>吾爱淘</a>
          <!------>
            <a rel="nofollow" size="设计视觉" href="https://hao.5iux.cn/" target="_blank"><svg class="icon" aria-hidden="true"><use xlink:href="#icon-daohang1"></use></svg>设计导航</a>
            <a rel="nofollow" size="设计视觉" href="https://www.pinterest.com/" target="_blank"><svg class="icon" aria-hidden="true"><use xlink:href="#icon-pinterest"></use></svg>Pinterest</a>
            <a rel="nofollow" size="设计视觉" href="https://www.behance.net/" target="_blank"><svg class="icon" aria-hidden="true"><use xlink:href="#icon-behance"></use></svg>Behance</a>
            <a rel="nofollow" size="设计视觉" href="https://www.dribbble.com/" target="_blank"><svg class="icon" aria-hidden="true"><use xlink:href="#icon-dribbble"></use></svg>Dribbble</a>
            <a rel="nofollow" size="设计视觉" href="https://huaban.com/" target="_blank"><svg class="icon" aria-hidden="true"><use xlink:href="#icon-huaban"></use></svg>花瓣</a>
            <a rel="nofollow" size="设计视觉" href="https://www.zcool.com.cn/" target="_blank"><svg class="icon" aria-hidden="true"><use xlink:href="#icon-zhanku"></use></svg>站酷</a>
            <a rel="nofollow" size="设计视觉" href="https://www.iconfont.cn/" target="_blank"><svg class="icon" aria-hidden="true"><use xlink:href="#icon-iconfont"></use></svg>阿里图标</a>
            <a rel="nofollow" size="设计视觉" href="https://www.iconfinder.com/" target="_blank"><svg class="icon" aria-hidden="true"><use xlink:href="#icon-eye"></use></svg>IconFinder</a>
            <a rel="nofollow" size="设计视觉" href="https://uiiiuiii.com/" target="_blank"><svg class="icon" aria-hidden="true"><use xlink:href="#icon-jiaocheng"></use></svg>优设教程</a>
          <!------>
            <a rel="nofollow" size="搜索引擎" href="https://www.google.com/" target="_blank"><svg class="icon" aria-hidden="true"><use xlink:href="#icon-google"></use></svg>Google</a>
            <a rel="nofollow" size="搜索引擎" href="https://duckduckgo.com/" target="_blank"><svg class="icon" aria-hidden="true"><use xlink:href="#icon-duck"></use></svg>DuckGo</a>
            <a rel="nofollow" size="搜索引擎" href="https://www.bing.com/" target="_blank"><svg class="icon" aria-hidden="true"><use xlink:href="#icon-bing"></use></svg>Bing</a>
            <a rel="nofollow" size="搜索引擎" href="https:/m.baidu.com/?pu=sz%401321_480&wpo=btmfast" target="_blank"><svg class="icon" aria-hidden="true"><use xlink:href="#icon-icon_baidulogo"></use></svg>百度</a>
            <a rel="nofollow" size="搜索引擎" href="https://hk.yahoo.com/" target="_blank"><svg class="icon" aria-hidden="true"><use xlink:href="#icon-yahoo"></use></svg>雅虎</a>
            <a rel="nofollow" size="搜索引擎" href="https://www.sogou.com/" target="_blank"><svg class="icon" aria-hidden="true"><use xlink:href="#icon-sougou"></use></svg>搜狗</a>
            <a rel="nofollow" size="搜索引擎" href="https://www.naver.com/" target="_blank"><svg class="icon" aria-hidden="true"><use xlink:href="#icon-icon-test"></use></svg>NAVER</a>
            <a rel="nofollow" size="搜索引擎" href="https://mijisou.com/" target="_blank"><svg class="icon" aria-hidden="true"><use xlink:href="#icon-miji"></use></svg>秘迹</a>
            <a rel="nofollow" size="搜索引擎" href="https://www.dogedoge.com/" target="_blank"><svg class="icon" aria-hidden="true"><use xlink:href="#icon-gougou"></use></svg>多吉</a>
            <a rel="nofollow" size="搜索引擎" href="https://seeres.com/" target="_blank"><svg class="icon" aria-hidden="true"><use xlink:href="#icon-sousuo"></use></svg>seeres</a>
            <!------>
            <a rel="nofollow" size="工具" href="https://tools.miku.ac/" target="_blank"><svg class="icon" aria-hidden="true"><use xlink:href="#icon-ai-tool"></use></svg>Miku工具</a>
            <a rel="nofollow" size="工具" href="https://g.5iux.cn/ip/" target="_blank"><svg class="icon" aria-hidden="true"><use xlink:href="#icon-wangluo"></use></svg>IP查询</a>
            <a rel="nofollow" size="工具" href="https://translate.google.cn/?hl=zh-CN" target="_blank"><svg class="icon" aria-hidden="true"><use xlink:href="#icon-fanyi"></use></svg>谷歌翻译</a>
            <a rel="nofollow" size="工具" href="http://www.slimego.cn/" target="_blank"><svg class="icon" aria-hidden="true"><use xlink:href="#icon-shilaimu"></use></svg>史莱姆</a>
            <a rel="nofollow" size="工具" href="https://feedly.com" target="_blank"><svg class="icon" aria-hidden="true"><use xlink:href="#icon-feedly"></use></svg>Feedly</a>         
            <a rel="nofollow" size="工具" href="https://pan.baidu.com" target="_blank"><svg class="icon" aria-hidden="true"><use xlink:href="#icon-baiduyun"></use></svg>百度网盘</a>
            <a rel="nofollow" size="工具" href="http://www.mdeditor.com/" target="_blank"><svg class="icon" aria-hidden="true"><use xlink:href="#icon-md"></use></svg>MD编辑器</a>
            <a rel="nofollow" size="工具" href="http://cubic-bezier.com" target="_blank"><svg class="icon" aria-hidden="true"><use xlink:href="#icon-quxian"></use></svg>贝赛尔曲线</a>
            <a rel="nofollow" size="工具" href="/base64/" target="_blank"><svg class="icon" aria-hidden="true"><use xlink:href="#icon-base64"></use></svg>Base64</a>
            <a rel="nofollow" size="工具" href="https://javascriptobfuscator.com/Javascript-Obfuscator.aspx" target="_blank"><svg class="icon" aria-hidden="true"><use xlink:href="#icon-jshunxiao"></use></svg>JS混淆器</a>
            <a rel="nofollow" size="工具" href="https://ping.pe" target="_blank"><svg class="icon" aria-hidden="true"><use xlink:href="#icon-wangluo1"></use></svg>Ping.pe</a>
            <a rel="nofollow" size="工具" href="https://ping.chinaz.com/" target="_blank"><svg class="icon" aria-hidden="true"><use xlink:href="#icon-pingup"></use></svg>站长Ping</a>
            <a rel="nofollow" size="工具" href="https://apkdl.in/" target="_blank"><svg class="icon" aria-hidden="true"><use xlink:href="#icon-anzhuo"></use></svg>APK下载</a>
            <!------>
            <a rel="nofollow" size="开发" href="http://www.w3school.com.cn/" target="_blank"><svg class="icon" aria-hidden="true"><use xlink:href="#icon-h5"></use></svg>W3school</a>
            <a rel="nofollow" size="开发" href="https://github.com/" target="_blank"><svg class="icon" aria-hidden="true"><use xlink:href="#icon-github"></use></svg>Github</a>
            <a rel="nofollow" size="开发" href="https://codepen.io/" target="_blank"><svg class="icon" aria-hidden="true"><use xlink:href="#icon-codepen"></use></svg>Codepen</a>
            <a rel="nofollow" size="开发" href="https://www.52pojie.cn/" target="_blank"><svg class="icon" aria-hidden="true"><use xlink:href="#icon-theater-masks"></use></svg>吾爱破解</a>
            <a rel="nofollow" size="开发" href="https://segmentfault.com/" target="_blank"><svg class="icon" aria-hidden="true"><use xlink:href="#icon-msg"></use></svg>SF思否</a>
            <a rel="nofollow" size="开发" href="https://cdnjs.com/" target="_blank"><svg class="icon" aria-hidden="true"><use xlink:href="#icon-cdnjs"></use></svg>CdnJs</a>
            <a rel="nofollow" size="开发" href="https://fontawesome.com/icons?d=gallery&m=free" target="_blank"><svg class="icon" aria-hidden="true"><use xlink:href="#icon-font-awesome"></use></svg>Font A.</a>
            <a rel="nofollow" size="开发" href="https://msdn.itellyou.cn/" target="_blank"><svg class="icon" aria-hidden="true"><use xlink:href="#icon-windows"></use></svg>MSDN下载</a>
            <a rel="nofollow" size="开发" href="https://dash.cloudflare.com/" target="_blank"><svg class="icon" aria-hidden="true"><use xlink:href="#icon-cloudflare"></use></svg>C. flare</a>
            <a rel="nofollow" size="开发" href="https://www.swiper.com.cn/" target="_blank"><svg class="icon" aria-hidden="true"><use xlink:href="#icon-S"></use></svg>Swiper</a>
        </ul>
    </div>`

	// Load the HTML document
	doc, err := goquery.NewDocumentFromReader(strings.NewReader(html))
	if err != nil {
		log.Fatal(err)
	}
	webSite := make([]model.WebSite, 0)
	// Find the review items
	//var category string
	doc.Find("ul").Find("a").Each(func(i int, s *goquery.Selection) {
		//fmt.Println(s.Find("li").Text())
		href, _ := s.Attr("href")
		size, _ := s.Attr("size")
		webSiteSingle := model.WebSite{Name: s.Text(), Category: size, Url: href}
		webSiteSingle.Favicon = Fetch(href)
		webSite = append(webSite, webSiteSingle)

	})
	b, err := json.Marshal(webSite)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(string(b))
	//
	//doc.Find("ul").Each(func(i int, s *goquery.Selection) {
	//	// For each item found, get the band and title
	//	name := s.Find(".nofollow")
	//
	//	href,_:=name.Attr("href")
	//	webSiteSingle:=model.WebSite{Name: name.Text(),Url: href}
	//	webSiteSingle.Favicon = Fetch(href)
	//	webSite=append(webSite,webSiteSingle)
	//	//title := s.Find("i").Text()
	//	//fmt.Printf("Review %d: %s - %s\n", i, band, title)
	//	fmt.Println(name.Text())
	//
	//})
	//fmt.Println(webSite)
}
