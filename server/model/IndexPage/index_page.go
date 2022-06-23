// 自动生成模板IndexPage
package IndexPage

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
)

// IndexPage 结构体
// 如果含有time.Time 请自行import time包
type IndexPage struct {
	global.GVA_MODEL
	Name string `json:"name" form:"name" gorm:"column:name;comment:;"`
}

// TableName IndexPage 表名
func (IndexPage) TableName() string {
	return "Index_page"
}
