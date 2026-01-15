// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// EnterpriseHistoricalData is the golang structure for table enterprise_historical_data.
type EnterpriseHistoricalData struct {
	Id            uint64      `json:"id"            orm:"id"             description:"自增主键"`
	Symbol        string      `json:"symbol"        orm:"symbol"         description:"证券代码 (如: 000001.SZ, AAPL)"`
	T             *gtime.Time `json:"t"             orm:"t"              description:"交易时间 (精确到分钟或日)"`
	Date          *gtime.Time `json:"date"          orm:"date"           description:"交易日期 (衍生字段)"`
	Year          int         `json:"year"          orm:"year"           description:"交易年份"`
	Month         int         `json:"month"         orm:"month"          description:"交易月份"`
	Weekday       int         `json:"weekday"       orm:"weekday"        description:"星期几 (1=周日,7=周六)"`
	O             float64     `json:"o"             orm:"o"              description:"开盘价"`
	H             float64     `json:"h"             orm:"h"              description:"最高价"`
	L             float64     `json:"l"             orm:"l"              description:"最低价"`
	C             float64     `json:"c"             orm:"c"              description:"收盘价"`
	Pc            float64     `json:"pc"            orm:"pc"             description:"前收盘价"`
	V             int64       `json:"v"             orm:"v"              description:"成交量 (股/手)"`
	A             float64     `json:"a"             orm:"a"              description:"成交额 (元)"`
	Change        float64     `json:"change"        orm:"change"         description:"涨跌额"`
	ChangePct     float64     `json:"changePct"     orm:"change_pct"     description:"涨跌幅 (%)"`
	Amplitude     float64     `json:"amplitude"     orm:"amplitude"      description:"振幅 (%)"`
	Sf            int         `json:"sf"            orm:"sf"             description:"停牌标志: 0-正常, 1-停牌"`
	TradingStatus string      `json:"tradingStatus" orm:"trading_status" description:"交易状态描述"`
	Period        string      `json:"period"        orm:"period"         description:"数据周期: 1min, 5min, 15min, 30min, 60min, day, week, month"`
	IsAdjusted    int         `json:"isAdjusted"    orm:"is_adjusted"    description:"是否复权: 0-不复权, 1-前复权, 2-后复权"`
	DataQuality   int         `json:"dataQuality"   orm:"data_quality"   description:"数据质量: 0-异常, 1-正常, 2-补全"`
	IsVerified    int         `json:"isVerified"    orm:"is_verified"    description:"是否已验证: 0-未验证, 1-已验证"`
	DataSource    string      `json:"dataSource"    orm:"data_source"    description:"数据来源"`
	Version       int         `json:"version"       orm:"version"        description:"数据版本"`
	CreatedAt     *gtime.Time `json:"createdAt"     orm:"created_at"     description:"创建时间"`
	UpdatedAt     *gtime.Time `json:"updatedAt"     orm:"updated_at"     description:"更新时间"`
}
