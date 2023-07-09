package model

import (
	"chat/pkg/db"
	"time"
)

type Friend struct {
	ID         uint64    `json:"id" gorm:"primary_key;auto_increment;comment:'自增主键'"`
	UserID     uint64    `json:"user_id" gorm:"not null;comment:'用户id'"`
	FriendID   uint64    `json:"friend_id" gorm:"not null;comment:'好友id'"`
	CreateTime time.Time `json:"create_time" gorm:"not null;default:CURRENT_TIMESTAMP;COMMENT:'创建时间'"`
	UpdateTime time.Time `json:"update_time" gorm:"not null;default:CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP;comment:'更新时间'"`
}

func (*Friend) TableName() string {
	return "friend"
}

// IsFriend 查询是否为好友关系
func IsFriend(userId, friendId uint64) (bool, error) {
	var cnt int64
	err := db.DB.Model(&Friend{}).Where("user_id = ? and friend_id = ?", userId, friendId).
		Or("friend_id = ? and user_id = ?", userId, friendId). // 反查
		Count(&cnt).Error
	return cnt > 0, err
}

func CreateFriend(friend *Friend) error {
	return db.DB.Create(friend).Error
}
