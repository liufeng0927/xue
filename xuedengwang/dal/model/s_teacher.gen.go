// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package model

import (
	"time"
)

const TableNameTeacher = "s_teacher"

// Teacher mapped from table <s_teacher>
type Teacher struct {
	ID         int64      `gorm:"column:id;type:bigint(20);primaryKey;autoIncrement:true" json:"id"`
	CreateBy   int64      `gorm:"column:create_by;type:bigint(20);not null" json:"create_by"`
	CreateTime *time.Time `gorm:"column:create_time;type:datetime;not null;default:CURRENT_TIMESTAMP" json:"create_time"` // 创建时间
	Remarks    string     `gorm:"column:remarks;type:varchar(255);not null" json:"remarks"`
	UpdateBy   int64      `gorm:"column:update_by;type:bigint(20);not null" json:"update_by"`
	UpdateTime *time.Time `gorm:"column:update_time;type:datetime;not null;default:CURRENT_TIMESTAMP" json:"update_time"` // 更新时间
	Name       string     `gorm:"column:name;type:varchar(255);not null" json:"name"`
	Phone      string     `gorm:"column:phone;type:varchar(255);not null" json:"phone"`
	Qq         string     `gorm:"column:qq;type:varchar(255);not null" json:"qq"`
	Sex        string     `gorm:"column:sex;type:varchar(255);not null" json:"sex"`
	TeachNo    string     `gorm:"column:teach_no;type:varchar(255);not null" json:"teach_no"`
	CourseID   int64      `gorm:"column:course_id;type:bigint(20);not null;index:FKgpacv4uc6gmdaridy3afaf5co,priority:1" json:"course_id"`
	Course     Course     `gorm:"foreignKey:CourseID" json:"course"`
}

// TableName Teacher's table name
func (*Teacher) TableName() string {
	return TableNameTeacher
}