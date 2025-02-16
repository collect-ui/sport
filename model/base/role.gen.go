// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package base

const TableNameRole = "role"

// Role mapped from table <role>
type Role struct {
	RoleID     string  `gorm:"column:role_id;primaryKey" json:"role_id"`
	RoleName   *string `gorm:"column:role_name" json:"role_name"`
	OrderIndex *int32  `gorm:"column:order_index" json:"order_index"`
	RoleCode   *string `gorm:"column:role_code" json:"role_code"`
}

// TableName Role's table name
func (*Role) TableName() string {
	return TableNameRole
}

func (*Role) PrimaryKey() []string {
	return []string{"role_id"}
}