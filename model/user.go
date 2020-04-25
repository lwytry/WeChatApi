package model

import (
	"github.com/jinzhu/gorm"
	"time"
)

type User struct {
	gorm.Model
	Name string `gorm:"column:username" json:"username"`
	Identifier string `gorm:"column:identifier" json:"identifier"`
	Phone string `gorm:"column:phone" json:"phone"`
}

func (User) TableName() string {
	return "user"
}

func AddUser(phone string, createdBy int) bool{
	var user User
	user.Name = phone
	user.Identifier = "wx_2kd8Dflco"
	user.Phone = phone
	DB.Table("wechat_user").Create(&user)
	return true
}

func GetUserList(pageNum int, pageSize int, maps interface{}) (userList []User) {

	offset := (pageNum -1) * pageSize
	if (maps != nil) {
		DB.Where(maps).Offset(offset).Limit(pageSize).Find(&userList)
	} else {
		DB.Offset(offset).Limit(pageSize).Find(&userList)
	}

	return
}

func GetUserTotal(maps interface {}) (count int){
	DB.Model(&User{}).Where(maps).Count(&count)

	return
}

func ExistUserByPhone(phone string) bool {
	var user User
	DB.Select("id").Where("phone = ?", phone).First(&user)
	if user.ID > 0 {
		return true
	}

	return false
}


func EditTag(id int, data interface{}) bool {
	DB.Model(&User{}).Where("id = ?", id).Updates(data)

	return true
}

func DeleteTag(id int) bool {
	DB.Where("id = ?", id).Delete(&User{})

	return true
}

func (tag *User) BeforeCreate(scope *gorm.Scope) error {
	scope.SetColumn("CreatedOn", time.Now().Unix())
	return nil
}

func (tag *User) BeforeUpdate(scope *gorm.Scope) error {
	scope.SetColumn("MofifiedOn", time.Now().Unix())
	return nil
}
