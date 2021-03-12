package model

import (
	"errors"
	"fmt"
	orm "goNav/api/database"
	"time"
)

type User struct {
	ID        int64     `json:"id"`       // 列名为 `id`
	Username  string    `json:"username"` // 列名为 `username`
	Password  string    `json:"password"` // 列名为 `password`
	WebSite   []WebSite `gorm:"ForeignKey:UserID;AssociationForeignKey:ID"`
	CreatedAt time.Time
	UpdatedAt time.Time
}

func init() {
	orm.DB.AutoMigrate(&User{})
}

var Users []User

//添加
func (user User) Insert() (id int64, err error) {

	//添加数据
	result := orm.DB.Create(&user)
	id = user.ID
	if result.Error != nil {
		err = result.Error
		return
	}
	return
}

//列表
func (user *User) Users() (users []User, err error) {
	if err = orm.DB.Find(&users).Error; err != nil {
		return
	}
	return
}

//列表
func (user *User) User() (webs []WebSite, err error) {
	if err = orm.DB.Model(user).Association("WebSite").Find(&webs); err != nil {
		return
	}
	return
}

//修改
func (user *User) Update(id int64) (updateUser User, err error) {

	if err = orm.DB.Select([]string{"id", "username"}).First(&updateUser, id).Error; err != nil {
		return
	}

	//参数1:是要修改的数据
	//参数2:是修改的数据
	if err = orm.DB.Model(&updateUser).Updates(&user).Error; err != nil {
		return
	}
	return
}

//删除数据
func (user *User) Destroy(id int64) (Result User, err error) {

	if err = orm.DB.Select([]string{"id"}).First(&user, id).Error; err != nil {
		return
	}

	if err = orm.DB.Delete(&user).Error; err != nil {
		return
	}
	Result = *user
	return
}

// LoginCheck验证
func LoginCheck(name string, password string) (bool, User, error) {
	userData := User{}
	var user User
	dbErr := orm.DB.Where("username = ?", name).Find(&user).Error

	if dbErr != nil {
		return false, userData, dbErr
	}
	// 用户存在 并且验证成功了
	if name == user.Username {
		if password == user.Password {
			userData.Username = user.Username
			userData.ID = user.ID
			fmt.Println(userData)
			return true, userData, nil
		} else {
			return true, userData, errors.New("用户名或密码错误")
		}
	} else {
		// 用户不存在
		return false, userData, nil
	}

}
