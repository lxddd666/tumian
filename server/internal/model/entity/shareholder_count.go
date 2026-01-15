// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// ShareholderCount is the golang structure for table shareholder_count.
type ShareholderCount struct {
	Id            uint64      `json:"id"            orm:"id"             description:"自增主键"`
	Symbol        string      `json:"symbol"        orm:"symbol"         description:"公司代码/股票代码 (例如: 000001.SZ, 600000.SS)"`
	Jzrq          *gtime.Time `json:"jzrq"          orm:"jzrq"           description:"截止日期 (报告期结束日, 如2023-09-30)"`
	ReportYear    int         `json:"reportYear"    orm:"report_year"    description:"报告年度"`
	ReportQuarter int         `json:"reportQuarter" orm:"report_quarter" description:"报告季度 (1-4)"`
	ReportType    string      `json:"reportType"    orm:"report_type"    description:"报告类型: annual-年报, half_year-中报, quarter-季报"`
	Gdzs          int         `json:"gdzs"          orm:"gdzs"           description:"股东总数 (户)"`
	Agdhs         int         `json:"agdhs"         orm:"agdhs"          description:"A股东户数 (户)"`
	Bgdhs         int         `json:"bgdhs"         orm:"bgdhs"          description:"B股东户数 (户)"`
	Hgdhs         int         `json:"hgdhs"         orm:"hgdhs"          description:"H股东户数 (户)"`
	Yltgdhs       int         `json:"yltgdhs"       orm:"yltgdhs"        description:"已流通股东户数 (户)"`
	Wltgdhs       int         `json:"wltgdhs"       orm:"wltgdhs"        description:"未流通股东户数 (户)"`
	AgRatio       float64     `json:"agRatio"       orm:"ag_ratio"       description:"A股股东占比(%)"`
	YltRatio      float64     `json:"yltRatio"      orm:"ylt_ratio"      description:"已流通股东占比(%)"`
	DataSource    string      `json:"dataSource"    orm:"data_source"    description:"数据来源"`
	IsLatest      int         `json:"isLatest"      orm:"is_latest"      description:"是否为该报告期最新数据"`
	CreatedAt     *gtime.Time `json:"createdAt"     orm:"created_at"     description:"创建时间"`
	UpdatedAt     *gtime.Time `json:"updatedAt"     orm:"updated_at"     description:"更新时间"`
}
