// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package model

import (
	"time"
)

const TableNameUser = "user"

// User mapped from table <user>
type User struct {
	ID         int64     `gorm:"column:id;primaryKey;autoIncrement:true" json:"id"`
	Username   string    `gorm:"column:username" json:"username"`
	Password   string    `gorm:"column:password" json:"password"`
	Phone      string    `gorm:"column:phone" json:"phone"`
	CreateTime time.Time `gorm:"column:create_time;not null;default:CURRENT_TIMESTAMP" json:"create_time"`
	UpdateTime time.Time `gorm:"column:update_time;not null;default:CURRENT_TIMESTAMP" json:"update_time"`
	IsDeleted  int32     `gorm:"column:is_deleted" json:"is_deleted"`
}

// TableName User's table name
func (*User) TableName() string {
	return TableNameUser
}
