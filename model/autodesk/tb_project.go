// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package autodesk

const TableNameTbProject = "tb_project"

// WebshellToken mapped from table <webshell_token>
type TbProject struct {
	TbProjectID string  `gorm:"column:tb_project_id;primaryKey" json:"tb_project_id"` // ID
	Code        *string `gorm:"column:code" json:"code"`                              // 描述
	Name        *string `gorm:"column:name" json:"name"`                              // 描述
	EtbpPath    *string `gorm:"column:etbp_path" json:"etbp_path"`                    // 描述
	CreateTime  *string `gorm:"column:create_time" json:"create_time"`                // datetime
	CreateUser  *string `gorm:"column:create_user" json:"create_user"`                // 添加人

}

// TableName WebshellToken's table name
func (*TbProject) TableName() string {
	return TableNameTbProject
}

func (*TbProject) PrimaryKey() []string {
	return []string{"tb_project_id"}
}
