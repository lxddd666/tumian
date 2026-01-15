// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// TopTenCirculatingHolders is the golang structure of table hg_top_ten_circulating_holders for DAO operations like Where/Data.
type TopTenCirculatingHolders struct {
	g.Meta        `orm:"table:hg_top_ten_circulating_holders, do:true"`
	Id            any         // 自增主键
	Symbol        any         // 公司代码/股票代码 (例如: 000001.SZ)
	Jzrq          *gtime.Time // 截止日期 (报告期结束日, 如2023-09-30)[citation:4]
	Ggrq          *gtime.Time // 公告日期 (信息发布日期)
	ReportYear    any         // 报告年度
	ReportQuarter any         // 报告季度 (1-4)
	ReportType    any         // 报告类型: annual-年报, half_year-中报, quarter-季报[citation:4]
	Gdmc          any         // 股东名称
	Gdlx          any         // 股东类型 (如: 基金、社保、个人等)
	Gfxz          any         // 股份性质 (如: 流通A股、限售A股等)
	Cgsl          any         // 持股数量 (股)
	Cgbl          any         // 持股比例 (%)
	Cgpm          any         // 持股排名 (1-10)
	Bdyy          any         // 变动原因
	BdType        any         // 变动类型: increase-增持, decrease-减持, new-新进, unchanged-不变, other-其他
	DataSource    any         // 数据来源 (如: 交易所公告)[citation:4]
	IsLatest      any         // 是否为该报告期最新数据: 0-历史快照, 1-最新
	Version       any         // 数据版本
	CreatedAt     *gtime.Time // 创建时间
	UpdatedAt     *gtime.Time // 更新时间
}
