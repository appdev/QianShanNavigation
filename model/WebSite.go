package model

import (
	"fmt"
	orm "goNav/api/database"
	"time"
)

type WebSite struct {
	ID        int64     `json:"id"` // 列名为 `id`
	Url       string    `json:"url"`
	Category  string    `json:"category"`
	Name      string    `json:"name"`
	Weight    int       `json:"weight" gorm:"AUTO_INCREMENT"`
	UserID    int64     `json:"user_id"`
	CreatedAt time.Time `json:"-"`
	UpdatedAt time.Time `json:"-"`
}

func init() {
	orm.DB.AutoMigrate(&WebSite{})
}

//添加
func (webSite WebSite) Insert() (id int64, err error) {
	// 先查询当前分类
	var lastWeb WebSite
	if err = orm.DB.Last(&lastWeb, "category = ?", webSite.Category).Error; err != nil {
		if err.Error() == "record not found" {
			webSite.Weight = 0
			err = nil
		}
	} else {
		webSite.Weight = lastWeb.Weight + 1
	}
	//添加数据
	result := orm.DB.Create(&webSite)
	id = webSite.ID
	if result.Error != nil {
		err = result.Error
		return
	}
	return
}

//列表
func (webSite *WebSite) WebSites(id int64) (users []WebSite, err error) {
	if err = orm.DB.Find(&webSite, WebSite{UserID: id}).Error; err != nil {
		return
	}
	return
}

func (webSite *WebSite) WebSiteAll() (users []WebSite, err error) {
	if err = orm.DB.Find(&webSite).Error; err != nil {
		return
	}
	return

}

//修改
func (webSite *WebSite) Update(id int64) (updateWeb WebSite, err error) {

	if err = orm.DB.Select([]string{"id"}).First(&updateWeb, id).Error; err != nil {
		return
	}

	//参数1:是要修改的数据
	//参数2:是修改的数据
	if err = orm.DB.Model(&updateWeb).Updates(&webSite).Error; err != nil {
		return
	}
	return
}

//删除数据
func (webSite *WebSite) Destroy(id int64) (Result WebSite, err error) {

	if err = orm.DB.Select([]string{"id"}).First(&webSite, id).Error; err != nil {
		return
	}

	if err = orm.DB.Delete(&webSite).Error; err != nil {
		return
	}
	Result = *webSite
	return
}

func (webSite *WebSite) ChangeCategory(newCategory string) (err error) {
	fmt.Println(webSite)
	fmt.Println(newCategory)
	if err = orm.DB.Model(WebSite{}).Where("category = ? AND user_id = ?", webSite.Category, webSite.UserID).Updates(WebSite{Category: newCategory}).Error; err != nil {
		return
	}
	return nil
}
