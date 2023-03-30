// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package model

import (
	"time"
)

const TableNameRole = "sys_role"

// Role mapped from table <sys_role>
type Role struct {
	ID         int64      `gorm:"column:id;type:bigint(20);primaryKey;autoIncrement:true" json:"id"`
	CreateBy   int64      `gorm:"column:create_by;type:bigint(20);not null" json:"create_by"`
	CreateTime *time.Time `gorm:"column:create_time;type:datetime;not null;default:CURRENT_TIMESTAMP" json:"create_time"` // 创建时间
	Remarks    string     `gorm:"column:remarks;type:varchar(255);not null" json:"remarks"`
	UpdateBy   int64      `gorm:"column:update_by;type:bigint(20);not null" json:"update_by"`
	UpdateTime *time.Time `gorm:"column:update_time;type:datetime;not null;default:CURRENT_TIMESTAMP" json:"update_time"` // 更新时间
	Code       string     `gorm:"column:code;type:varchar(255);not null" json:"code"`
	Name       string     `gorm:"column:name;type:varchar(255);not null" json:"name"`
}

// TableName Role's table name
func (*Role) TableName() string {
	return TableNameRole
}
