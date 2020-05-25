package model

import "time"

type Contact struct {
	ID        uint `gorm:"primary_key" json:"id"`
	UId 	  uint `gorm:"column:uId" json:"_"`
	CUID 	  uint `gorm:"column:cuId" json:"cuId"`
	Remark    string `gorm:"column:remark" json:"remark"`
	CreatedAt time.Time `gorm:"column:createAt" json:"_"`
	DeletedAt time.Time `gorm:"column:deleteAt" json:"_"`
	User 	  User `gorm:"FOREIGNKEY:ID;ASSOCIATION_FOREIGNKEY:CUID"`
}

func (Contact) TableName() string {
	return "contact"
}

func GetContactAll(userId string) (contactList []Contact) {
	DB.Preload("User").Where("uId = ?", userId).Find(&contactList)
	return
}