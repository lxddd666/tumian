// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// StockList is the golang structure for table stock_list.
type StockList struct {
	Id           uint64      `json:"id"           orm:"id"            description:"自增主键"`
	Dm           string      `json:"dm"           orm:"dm"            description:"股票代码 (唯一业务标识，如: 000001)"`
	Mc           string      `json:"mc"           orm:"mc"            description:"股票名称 (如: 平安银行)"`
	Jys          string      `json:"jys"          orm:"jys"           description:"交易所代码 (如: sh, sz, bj)"`
	ExchangeName string      `json:"exchangeName" orm:"exchange_name" description:"交易所全称"`
	Symbol       string      `json:"symbol"       orm:"symbol"        description:"标准股票代码 (如: 000001.SZ)"`
	Status       int         `json:"status"       orm:"status"        description:"状态: 1-正常, 0-退市"`
	ListDate     *gtime.Time `json:"listDate"     orm:"list_date"     description:"上市日期"`
	DataSource   string      `json:"dataSource"   orm:"data_source"   description:"数据来源"`
	CreatedAt    *gtime.Time `json:"createdAt"    orm:"created_at"    description:"创建时间"`
	UpdatedAt    *gtime.Time `json:"updatedAt"    orm:"updated_at"    description:"更新时间"`
}
