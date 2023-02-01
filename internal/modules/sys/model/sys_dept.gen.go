// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package model

import (
	"time"
)

const TableNameSysDept = "sys_dept"

// SysDept mapped from table <sys_dept>
type SysDept struct {
	ID         int64     `gorm:"column:id;primaryKey" json:"id"`       // id
	Pid        int64     `gorm:"column:pid" json:"pid"`                // 上级ID
	Pids       string    `gorm:"column:pids" json:"pids"`              // 所有上级ID，用逗号分开
	Name       string    `gorm:"column:name" json:"name"`              // 部门名称
	Sort       int32     `gorm:"column:sort" json:"sort"`              // 排序
	Creator    int64     `gorm:"column:creator" json:"creator"`        // 创建者
	CreateDate time.Time `gorm:"column:create_date" json:"createDate"` // 创建时间
	Updater    int64     `gorm:"column:updater" json:"updater"`        // 更新者
	UpdateDate time.Time `gorm:"column:update_date" json:"updateDate"` // 更新时间
}

// TableName SysDept's table name
func (*SysDept) TableName() string {
	return TableNameSysDept
}