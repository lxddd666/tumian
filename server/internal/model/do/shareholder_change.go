// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// ShareholderChange is the golang structure of table hg_shareholder_change for DAO operations like Where/Data.
type ShareholderChange struct {
	g.Meta          `orm:"table:hg_shareholder_change, do:true"`
	Id              interface{} // 自增主键
	Symbol          interface{} // 股票代码 (如: 000001.SZ)
	Jzrq            *gtime.Time // 截止日期 (统计截止日，如2025-12-31)[citation:9]
	Gdhs            interface{} // 股东户数 (统计截止日的总户数)[citation:3][citation:6]
	Bh              interface{} // 比上期变化情况 (通常为百分比，如 -6.23 表示减少6.23%)[citation:9]
	ChangeDirection interface{} // 变化方向 (衍生字段)
	DataSource      interface{} // 数据来源 (如: 交易所公告[citation:6]、互动平台[citation:3])
	AnnDate         *gtime.Time // 公告日期 (信息发布日期)[citation:1]
	CreatedAt       *gtime.Time // 创建时间
	UpdatedAt       *gtime.Time // 更新时间
}
