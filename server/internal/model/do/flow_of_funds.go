// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// FlowOfFunds is the golang structure of table hg_flow_of_funds for DAO operations like Where/Data.
type FlowOfFunds struct {
	g.Meta     `orm:"table:hg_flow_of_funds, do:true"`
	Id         interface{} // 主键ID
	Symbol     interface{} // 股票代码 (如: 000001.SZ)
	T          *gtime.Time // 交易时间 (通常为HHMMSS格式的整数)
	Zmbzds     interface{} // 主买单总单数
	Zmszds     interface{} // 主卖单总单数
	Dddx       interface{} // 大单动向
	Zddy       interface{} // 涨跌动因
	Ddcf       interface{} // 大单差分
	Zmbzdszl   interface{} // 主买单总单数增量
	Zmszdszl   interface{} // 主卖单总单数增量
	Cjbszl     interface{} // 成交笔数增量
	Zmbtdcje   interface{} // 主买特大单成交额
	Zmbddcje   interface{} // 主买大单成交额
	Zmbzdcje   interface{} // 主买中单成交额
	Zmbxdcje   interface{} // 主买小单成交额
	Zmstdcje   interface{} // 主卖特大单成交额
	Zmsddcje   interface{} // 主卖大单成交额
	Zmszdcje   interface{} // 主卖中单成交额
	Zmsxdcje   interface{} // 主卖小单成交额
	Bdmbtdcje  interface{} // 被动买特大单成交额
	Bdmbddcje  interface{} // 被动买大单成交额
	Bdmbzdcje  interface{} // 被动买中单成交额
	Bdmbxdcje  interface{} // 被动买小单成交额
	Bdmstdcje  interface{} // 被动卖特大单成交额
	Bdmsddcje  interface{} // 被动卖大单成交额
	Bdmszdcje  interface{} // 被动卖中单成交额
	Bdmsxdcje  interface{} // 被动卖小单成交额
	Zmbtdcjl   interface{} // 主买特大单成交量
	Zmbddcjl   interface{} // 主买大单成交量
	Zmbzdcjl   interface{} // 主买中单成交量
	Zmbxdcjl   interface{} // 主买小单成交量
	Zmstdcjl   interface{} // 主卖特大单成交量
	Zmsddcjl   interface{} // 主卖大单成交量
	Zmszdcjl   interface{} // 主卖中单成交量
	Zmsxdcjl   interface{} // 主卖小单成交量
	Bdmbtdcjl  interface{} // 被动买特大单成交量
	Bdmbddcjl  interface{} // 被动买大单成交量
	Bdmbzdcjl  interface{} // 被动买中单成交量
	Bdmbxdcjl  interface{} // 被动买小单成交量
	Bdmstdcjl  interface{} // 被动卖特大单成交量
	Bdmsddcjl  interface{} // 被动卖大单成交量
	Bdmszdcjl  interface{} // 被动卖中单成交量
	Bdmsxdcjl  interface{} // 被动卖小单成交量
	Zmbtdcjzl  interface{} // 主买特大单成交额增量
	Zmbddcjzl  interface{} // 主买大单成交额增量
	Zmbtdcjzlv interface{} // 主买特大单成交量增量
	Zmbddcjzlv interface{} // 主买大单成交量增量
	CreatedAt  *gtime.Time // 创建时间
	UpdatedAt  *gtime.Time // 更新时间
}
