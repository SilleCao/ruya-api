// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package model

import (
	"time"
)

const TableNameSysUser = "sys_user"

// SysUser mapped from table <sys_user>
type SysUser struct {
	ID         int64     `gorm:"column:id;primaryKey;autoIncrement:true" json:"id"` // id
	Username   string    `gorm:"column:username;not null" json:"username"`          // 用户名
	Password   string    `gorm:"column:password" json:"password"`                   // 密码
	RealName   string    `gorm:"column:real_name" json:"realName"`                  // 姓名
	HeadURL    string    `gorm:"column:head_url" json:"headUrl"`                    // 头像
	Gender     int32     `gorm:"column:gender" json:"gender"`                       // 性别   0：男   1：女    2：保密
	Email      string    `gorm:"column:email" json:"email"`                         // 邮箱
	Mobile     string    `gorm:"column:mobile" json:"mobile"`                       // 手机号
	DeptID     int64     `gorm:"column:dept_id" json:"deptId"`                      // 部门ID
	SuperAdmin int32     `gorm:"column:super_admin" json:"superAdmin"`              // 超级管理员   0：否   1：是
	Status     int32     `gorm:"column:status" json:"status"`                       // 状态  0：停用   1：正常
	Creator    int64     `gorm:"column:creator" json:"creator"`                     // 创建者
	CreateDate time.Time `gorm:"column:create_date" json:"createDate"`              // 创建时间
	Updater    int64     `gorm:"column:updater" json:"updater"`                     // 更新者
	UpdateDate time.Time `gorm:"column:update_date" json:"updateDate"`              // 更新时间
}

// TableName SysUser's table name
func (*SysUser) TableName() string {
	return TableNameSysUser
}
