package model

import (
	"chat/pkg/db"
	"time"
)

type User struct {
	ID          uint64    `gorm:"primary_key;auto_increment;comment:'自增主键'" json:"id"`
	PhoneNumber string    `gorm:"not null;unique;comment:'手机号'" json:"phone_number"`
	Nickname    string    `gorm:"not null;comment:'昵称'" json:"nickname"`
	Password    string    `gorm:"not null;comment:'密码'" json:"-"`
	CreateTime  time.Time `gorm:"not null;default:CURRENT_TIMESTAMP;comment:'创建时间'" json:"create_time"`
	UpdateTime  time.Time `gorm:"not null;default:CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP;comment:'更新时间'" json:"update_time"`
}

func (*User) TableName() string {
	return "user"
}

func GetUserCountByPhone(phoneNumber string) (int64, error) {
	var cnt int64
	err := db.DB.Model(&User{}).Where("phone_number = ?", phoneNumber).Count(&cnt).Error
	if err != nil {
		panic(err)
	}
	return cnt, err
}

func CreateUser(user *User) error {
	return db.DB.Create(user).Error
}
