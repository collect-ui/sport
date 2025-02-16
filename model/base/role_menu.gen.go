// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package base

const TableNameRoleMenu = "role_menu"

// Role mapped from table <role>
type RoleMenu struct {
	RoleMenuID     string  `gorm:"column:role_menu_id;primaryKey" json:"role_menu_id"`
	RoleID   *string `gorm:"column:role_id" json:"role_id"`
	SysMenuID   *string `gorm:"column:sys_menu_id" json:"sys_menu_id"`
}

// TableName Role's table name
func (*RoleMenu) TableName() string {
	return TableNameRoleMenu
}

func (*RoleMenu) PrimaryKey() []string {
	return []string{"role_menu_id"}
}