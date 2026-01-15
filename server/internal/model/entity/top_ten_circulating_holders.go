// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// TopTenCirculatingHolders is the golang structure for table top_ten_circulating_holders.
type TopTenCirculatingHolders struct {
	Id            uint64      `json:"id"            orm:"id"             description:"自增主键"`
	Symbol        string      `json:"symbol"        orm:"symbol"         description:"公司代码/股票代码 (例如: 000001.SZ)"`
	Jzrq          *gtime.Time `json:"jzrq"          orm:"jzrq"           description:"截止日期 (报告期结束日, 如2023-09-30)[citation:4]"`
	Ggrq          *gtime.Time `json:"ggrq"          orm:"ggrq"           description:"公告日期 (信息发布日期)"`
	ReportYear    int         `json:"reportYear"    orm:"report_year"    description:"报告年度"`
	ReportQuarter int         `json:"reportQuarter" orm:"report_quarter" description:"报告季度 (1-4)"`
	ReportType    string      `json:"reportType"    orm:"report_type"    description:"报告类型: annual-年报, half_year-中报, quarter-季报[citation:4]"`
	Gdmc          string      `json:"gdmc"          orm:"gdmc"           description:"股东名称"`
	Gdlx          string      `json:"gdlx"          orm:"gdlx"           description:"股东类型 (如: 基金、社保、个人等)"`
	Gfxz          string      `json:"gfxz"          orm:"gfxz"           description:"股份性质 (如: 流通A股、限售A股等)"`
	Cgsl          int64       `json:"cgsl"          orm:"cgsl"           description:"持股数量 (股)"`
	Cgbl          float64     `json:"cgbl"          orm:"cgbl"           description:"持股比例 (%)"`
	Cgpm          int         `json:"cgpm"          orm:"cgpm"           description:"持股排名 (1-10)"`
	Bdyy          string      `json:"bdyy"          orm:"bdyy"           description:"变动原因"`
	BdType        string      `json:"bdType"        orm:"bd_type"        description:"变动类型: increase-增持, decrease-减持, new-新进, unchanged-不变, other-其他"`
	DataSource    string      `json:"dataSource"    orm:"data_source"    description:"数据来源 (如: 交易所公告)[citation:4]"`
	IsLatest      int         `json:"isLatest"      orm:"is_latest"      description:"是否为该报告期最新数据: 0-历史快照, 1-最新"`
	Version       int         `json:"version"       orm:"version"        description:"数据版本"`
	CreatedAt     *gtime.Time `json:"createdAt"     orm:"created_at"     description:"创建时间"`
	UpdatedAt     *gtime.Time `json:"updatedAt"     orm:"updated_at"     description:"更新时间"`
}
