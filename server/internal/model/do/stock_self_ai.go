// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// StockSelfAi is the golang structure of table hg_stock_self_ai for DAO operations like Where/Data.
type StockSelfAi struct {
	g.Meta    `orm:"table:hg_stock_self_ai, do:true"`
	Id        any         // 自增主键
	Name      any         // ai模型全称名称 例如deepseek-plus 千问
	Model     any         // ai model
	BaseUrl   any         // ai base url
	ApiKey    any         // ai api key
	CreatedAt *gtime.Time // 创建时间
	AiModel   any         // 语言模型 qianwen deepseek
	Status    any         // 0正常 -1不正常
}
