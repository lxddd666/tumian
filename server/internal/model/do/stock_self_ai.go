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
	Id        interface{} // 自增主键
	Name      interface{} // ai模型全称名称 例如deepseek-plus 千问
	Model     interface{} // ai model
	BaseUrl   interface{} // ai base url
	ApiKey    interface{} // ai api key
	CreatedAt *gtime.Time // 创建时间
	AiModel   interface{} // 语言模型 qianwen deepseek
	Status    interface{} // 0正常 -1不正常
}
