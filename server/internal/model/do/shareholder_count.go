// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// ShareholderCount is the golang structure of table hg_shareholder_count for DAO operations like Where/Data.
type ShareholderCount struct {
	g.Meta        `orm:"table:hg_shareholder_count, do:true"`
	Id            any         // 自增主键
	Symbol        any         // 公司代码/股票代码 (例如: 000001.SZ, 600000.SS)
	Jzrq          *gtime.Time // 截止日期 (报告期结束日, 如2023-09-30)
	ReportYear    any         // 报告年度
	ReportQuarter any         // 报告季度 (1-4)
	ReportType    any         // 报告类型: annual-年报, half_year-中报, quarter-季报
	Gdzs          any         // 股东总数 (户)
	Agdhs         any         // A股东户数 (户)
	Bgdhs         any         // B股东户数 (户)
	Hgdhs         any         // H股东户数 (户)
	Yltgdhs       any         // 已流通股东户数 (户)
	Wltgdhs       any         // 未流通股东户数 (户)
	AgRatio       any         // A股股东占比(%)
	YltRatio      any         // 已流通股东占比(%)
	DataSource    any         // 数据来源
	IsLatest      any         // 是否为该报告期最新数据
	CreatedAt     *gtime.Time // 创建时间
	UpdatedAt     *gtime.Time // 更新时间
}
