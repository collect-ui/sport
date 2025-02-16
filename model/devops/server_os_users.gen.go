// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package devops

import (
	"time"
)

const TableNameServerOsUsers = "server_os_users"

// ServerOsUsers mapped from table <server_os_users>
type ServerOsUsers struct {
	ServerOsUsersID   string     `gorm:"column:server_os_users_id;primaryKey" json:"server_os_users_id"`
	UserName          string     `gorm:"column:user_name;not null" json:"user_name"`                                      // 用户名
	UserPwd           *string    `gorm:"column:user_pwd" json:"user_pwd"`                                                 // 服务器密码
	UserHome          *string    `gorm:"column:user_home" json:"user_home"`                                               // 用户主目录
	OsPrivileageCode  *string    `gorm:"column:os_privileage_code" json:"os_privileage_code"`                             // 用户操作系统权限等级，见码表
	CreateTime        *time.Time `gorm:"column:create_time;not null;" json:"create_time"`        // 记录创建时间（数据库自动写入）
	ModifyTime        *time.Time `gorm:"column:modify_time;not null;" json:"modify_time"`        // 记录修改时间（数据库自动写入）
	Comments          *string    `gorm:"column:comments" json:"comments"`                                                 // 备注说明
	DefaultUser       *string    `gorm:"column:default_user" json:"default_user"`
	ServerID          *string    `gorm:"column:server_id" json:"server_id"`
	CreateShell       *string    `gorm:"column:create_shell" json:"create_shell"`
	OssoftUserGroupID *string    `gorm:"column:ossoft_user_group_id" json:"ossoft_user_group_id"`
}

// TableName ServerOsUsers's table name
func (*ServerOsUsers) TableName() string {
	return TableNameServerOsUsers
}

func (*ServerOsUsers) PrimaryKey() []string {
	return []string{"server_os_user_id"}
}