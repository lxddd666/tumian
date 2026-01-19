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
	Id         any         // 主键ID
	Symbol     any         // 股票代码 (如: 000001.SZ)
	T          *gtime.Time // 交易时间 (通常为HHMMSS格式的整数)
	Zmbzds     any         // 主买单总单数
	Zmszds     any         // 主卖单总单数
	Dddx       any         // 大单动向
	Zddy       any         // 涨跌动因
	Ddcf       any         // 大单差分
	Zmbzdszl   any         // 主买单总单数增量
	Zmszdszl   any         // 主卖单总单数增量
	Cjbszl     any         // 成交笔数增量
	Zmbtdcje   any         // 主买特大单成交额
	Zmbddcje   any         // 主买大单成交额
	Zmbzdcje   any         // 主买中单成交额
	Zmbxdcje   any         // 主买小单成交额
	Zmstdcje   any         // 主卖特大单成交额
	Zmsddcje   any         // 主卖大单成交额
	Zmszdcje   any         // 主卖中单成交额
	Zmsxdcje   any         // 主卖小单成交额
	Bdmbtdcje  any         // 被动买特大单成交额
	Bdmbddcje  any         // 被动买大单成交额
	Bdmbzdcje  any         // 被动买中单成交额
	Bdmbxdcje  any         // 被动买小单成交额
	Bdmstdcje  any         // 被动卖特大单成交额
	Bdmsddcje  any         // 被动卖大单成交额
	Bdmszdcje  any         // 被动卖中单成交额
	Bdmsxdcje  any         // 被动卖小单成交额
	Zmbtdcjl   any         // 主买特大单成交量
	Zmbddcjl   any         // 主买大单成交量
	Zmbzdcjl   any         // 主买中单成交量
	Zmbxdcjl   any         // 主买小单成交量
	Zmstdcjl   any         // 主卖特大单成交量
	Zmsddcjl   any         // 主卖大单成交量
	Zmszdcjl   any         // 主卖中单成交量
	Zmsxdcjl   any         // 主卖小单成交量
	Bdmbtdcjl  any         // 被动买特大单成交量
	Bdmbddcjl  any         // 被动买大单成交量
	Bdmbzdcjl  any         // 被动买中单成交量
	Bdmbxdcjl  any         // 被动买小单成交量
	Bdmstdcjl  any         // 被动卖特大单成交量
	Bdmsddcjl  any         // 被动卖大单成交量
	Bdmszdcjl  any         // 被动卖中单成交量
	Bdmsxdcjl  any         // 被动卖小单成交量
	Zmbtdcjzl  any         // 主买特大单成交额增量
	Zmbddcjzl  any         // 主买大单成交额增量
	Zmbtdcjzlv any         // 主买特大单成交量增量
	Zmbddcjzlv any         // 主买大单成交量增量
	CreatedAt  *gtime.Time // 创建时间
	UpdatedAt  *gtime.Time // 更新时间
}
