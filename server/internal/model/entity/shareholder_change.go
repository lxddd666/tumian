// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// ShareholderChange is the golang structure for table shareholder_change.
type ShareholderChange struct {
	Id              uint64      `json:"id"              orm:"id"               description:"自增主键"`
	Symbol          string      `json:"symbol"          orm:"symbol"           description:"股票代码 (如: 000001.SZ)"`
	Jzrq            *gtime.Time `json:"jzrq"            orm:"jzrq"             description:"截止日期 (统计截止日，如2025-12-31)[citation:9]"`
	Gdhs            int         `json:"gdhs"            orm:"gdhs"             description:"股东户数 (统计截止日的总户数)[citation:3][citation:6]"`
	Bh              float64     `json:"bh"              orm:"bh"               description:"比上期变化情况 (通常为百分比，如 -6.23 表示减少6.23%)[citation:9]"`
	ChangeDirection string      `json:"changeDirection" orm:"change_direction" description:"变化方向 (衍生字段)"`
	DataSource      string      `json:"dataSource"      orm:"data_source"      description:"数据来源 (如: 交易所公告[citation:6]、互动平台[citation:3])"`
	AnnDate         *gtime.Time `json:"annDate"         orm:"ann_date"         description:"公告日期 (信息发布日期)[citation:1]"`
	CreatedAt       *gtime.Time `json:"createdAt"       orm:"created_at"       description:"创建时间"`
	UpdatedAt       *gtime.Time `json:"updatedAt"       orm:"updated_at"       description:"更新时间"`
}
