package model

import "time"

type GroupUser struct {
	ID         uint64    `gorm:"primary_key;auto_increment;comment:'自增主键'" json:"id"`
	GroupID    uint64    `gorm:"not null;comment:'组id'" json:"group_id"`
	UserID     uint64    `gorm:"not null;comment:'用户id'" json:"user_id"`
	CreateTime time.Time `gorm:"not null;default:CURRENT_TIMESTAMP;comment:'创建时间'" json:"create_time"`
	UpdateTime time.Time `gorm:"not null;default:CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP;comment:'更新时间'" json:"update_time"`
}

func (*GroupUser) TableName() string {
	return "group_user"
}
