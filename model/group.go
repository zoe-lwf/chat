package model

import (
	"chat/pkg/db"
	"gorm.io/gorm"
	"time"
)

type Group struct {
	ID         uint64    `json:"id" gorm:"primary_key;auto_increment;comment:'自增主键'"`
	Name       string    `json:"name" gorm:"not null;comment:'群组名称'"`
	OwnerID    uint64    `gorm:"not null;comment:'群主id'" json:"owner_id"`
	CreateTime time.Time `gorm:"not null;default:CURRENT_TIMESTAMP;comment:'创建时间'" json:"create_time"`
	UpdateTime time.Time `gorm:"not null;default:CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP;comment:'更新时间'" json:"update_time"`
}

func (*Group) TableName() string {
	return "group"
}

// CreateGroup 开启事务创建群组
func CreateGroup(group *Group, ids []uint64) error {
	return db.DB.Transaction(func(tx *gorm.DB) error {
		err := tx.Create(group).Error
		if err != nil {
			return err
		}
		groupUsers := make([]*GroupUser, 0, len(ids))
		for _, id := range ids {
			groupUsers = append(groupUsers, &GroupUser{
				GroupID: group.ID,
				UserID:  id,
			})
		}
		return tx.Create(groupUsers).Error
	})
}
