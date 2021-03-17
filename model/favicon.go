package model

import orm "goNav/api/database"

type Favicon struct {
	ID      int64  `json:"id"` // 列名为 `id`
	Favicon string `json:"favicon"`
	Host    string `json:"host"`
}

func init() {
	orm.DB.AutoMigrate(&Favicon{})
}

//添加
func (favicon Favicon) Insert() (id int64, err error) {

	//添加数据
	result := orm.DB.Create(&favicon)
	id = favicon.ID
	if result.Error != nil {
		err = result.Error
		return
	}
	return
}

//查询某一项
func FindOne(host string) (faviconIcon Favicon, err error) {
	if err = orm.DB.Where("Host = ?", host).First(&faviconIcon).Error; err != nil {
		return
	}
	return
}
