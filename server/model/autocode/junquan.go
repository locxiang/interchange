// 自动生成模板Junquan
package autocode

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/system"
)

// Junquan 结构体
// 如果含有time.Time 请自行import time包
type Junquan struct {
	global.GVA_MODEL
	URL      string `json:"url" form:"url" gorm:"column:url;comment:"`
	Username string `json:"username" form:"username" gorm:"column:username;comment:"`
	Password string `json:"password" form:"password" gorm:"column:password;comment:"`
	UserId uint    `json:"user_id" form:"user_id" gorm:"column:user_id;comment:"`
	User   system.SysUser `json:"user" form:"-"`
}
