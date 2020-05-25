package model

import (
	"github.com/jinzhu/gorm"
	"time"
)

type User struct {
	ID        uint `gorm:"primary_key" json:"id"`
	Name string `gorm:"column:username" json:"username"`
	Identifier string `gorm:"column:identifier" json:"identifier"`
	Phone string `gorm:"column:phone" json:"phone"`
	Password string `gorm:"column:password" json:"_"`
	AvatarPath string `gorm:"column:avatarPath" json:"avatarPath"`
	CreatedAt time.Time `json:"_"`
	UpdatedAt time.Time `json:"_"`
	DeletedAt *time.Time `sql:"index" json:"_"`
}

func (User) TableName() string {
	return "user"
}

func AddUser(phone string, identifier string, createdBy int) bool{
	var user User
	user.Name = phone
	user.Identifier = identifier
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

func GetUserByPhone(phone string) (user User) {
	DB.Table("wechat_user").Where("phone = ?", phone).First(&user)

	return
}

func (tag *User) BeforeCreate(scope *gorm.Scope) error {
	scope.SetColumn("CreatedOn", time.Now().Unix())
	return nil
}
