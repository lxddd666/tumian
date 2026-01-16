// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// FlowOfFunds is the golang structure for table flow_of_funds.
type FlowOfFunds struct {
	Id         int64       `json:"id"         orm:"id"         description:"主键ID"`
	Symbol     string      `json:"symbol"     orm:"symbol"     description:"股票代码 (如: 000001.SZ)"`
	T          *gtime.Time `json:"t"          orm:"t"          description:"交易时间 (通常为HHMMSS格式的整数)"`
	Zmbzds     int         `json:"zmbzds"     orm:"zmbzds"     description:"主买单总单数"`
	Zmszds     int         `json:"zmszds"     orm:"zmszds"     description:"主卖单总单数"`
	Dddx       float64     `json:"dddx"       orm:"dddx"       description:"大单动向"`
	Zddy       float64     `json:"zddy"       orm:"zddy"       description:"涨跌动因"`
	Ddcf       float64     `json:"ddcf"       orm:"ddcf"       description:"大单差分"`
	Zmbzdszl   int         `json:"zmbzdszl"   orm:"zmbzdszl"   description:"主买单总单数增量"`
	Zmszdszl   int         `json:"zmszdszl"   orm:"zmszdszl"   description:"主卖单总单数增量"`
	Cjbszl     int         `json:"cjbszl"     orm:"cjbszl"     description:"成交笔数增量"`
	Zmbtdcje   float64     `json:"zmbtdcje"   orm:"zmbtdcje"   description:"主买特大单成交额"`
	Zmbddcje   float64     `json:"zmbddcje"   orm:"zmbddcje"   description:"主买大单成交额"`
	Zmbzdcje   float64     `json:"zmbzdcje"   orm:"zmbzdcje"   description:"主买中单成交额"`
	Zmbxdcje   float64     `json:"zmbxdcje"   orm:"zmbxdcje"   description:"主买小单成交额"`
	Zmstdcje   float64     `json:"zmstdcje"   orm:"zmstdcje"   description:"主卖特大单成交额"`
	Zmsddcje   float64     `json:"zmsddcje"   orm:"zmsddcje"   description:"主卖大单成交额"`
	Zmszdcje   float64     `json:"zmszdcje"   orm:"zmszdcje"   description:"主卖中单成交额"`
	Zmsxdcje   float64     `json:"zmsxdcje"   orm:"zmsxdcje"   description:"主卖小单成交额"`
	Bdmbtdcje  float64     `json:"bdmbtdcje"  orm:"bdmbtdcje"  description:"被动买特大单成交额"`
	Bdmbddcje  float64     `json:"bdmbddcje"  orm:"bdmbddcje"  description:"被动买大单成交额"`
	Bdmbzdcje  float64     `json:"bdmbzdcje"  orm:"bdmbzdcje"  description:"被动买中单成交额"`
	Bdmbxdcje  float64     `json:"bdmbxdcje"  orm:"bdmbxdcje"  description:"被动买小单成交额"`
	Bdmstdcje  float64     `json:"bdmstdcje"  orm:"bdmstdcje"  description:"被动卖特大单成交额"`
	Bdmsddcje  float64     `json:"bdmsddcje"  orm:"bdmsddcje"  description:"被动卖大单成交额"`
	Bdmszdcje  float64     `json:"bdmszdcje"  orm:"bdmszdcje"  description:"被动卖中单成交额"`
	Bdmsxdcje  float64     `json:"bdmsxdcje"  orm:"bdmsxdcje"  description:"被动卖小单成交额"`
	Zmbtdcjl   int         `json:"zmbtdcjl"   orm:"zmbtdcjl"   description:"主买特大单成交量"`
	Zmbddcjl   int         `json:"zmbddcjl"   orm:"zmbddcjl"   description:"主买大单成交量"`
	Zmbzdcjl   int         `json:"zmbzdcjl"   orm:"zmbzdcjl"   description:"主买中单成交量"`
	Zmbxdcjl   int         `json:"zmbxdcjl"   orm:"zmbxdcjl"   description:"主买小单成交量"`
	Zmstdcjl   int         `json:"zmstdcjl"   orm:"zmstdcjl"   description:"主卖特大单成交量"`
	Zmsddcjl   int         `json:"zmsddcjl"   orm:"zmsddcjl"   description:"主卖大单成交量"`
	Zmszdcjl   int         `json:"zmszdcjl"   orm:"zmszdcjl"   description:"主卖中单成交量"`
	Zmsxdcjl   int         `json:"zmsxdcjl"   orm:"zmsxdcjl"   description:"主卖小单成交量"`
	Bdmbtdcjl  int         `json:"bdmbtdcjl"  orm:"bdmbtdcjl"  description:"被动买特大单成交量"`
	Bdmbddcjl  int         `json:"bdmbddcjl"  orm:"bdmbddcjl"  description:"被动买大单成交量"`
	Bdmbzdcjl  int         `json:"bdmbzdcjl"  orm:"bdmbzdcjl"  description:"被动买中单成交量"`
	Bdmbxdcjl  int         `json:"bdmbxdcjl"  orm:"bdmbxdcjl"  description:"被动买小单成交量"`
	Bdmstdcjl  int         `json:"bdmstdcjl"  orm:"bdmstdcjl"  description:"被动卖特大单成交量"`
	Bdmsddcjl  int         `json:"bdmsddcjl"  orm:"bdmsddcjl"  description:"被动卖大单成交量"`
	Bdmszdcjl  int         `json:"bdmszdcjl"  orm:"bdmszdcjl"  description:"被动卖中单成交量"`
	Bdmsxdcjl  int         `json:"bdmsxdcjl"  orm:"bdmsxdcjl"  description:"被动卖小单成交量"`
	Zmbtdcjzl  float64     `json:"zmbtdcjzl"  orm:"zmbtdcjzl"  description:"主买特大单成交额增量"`
	Zmbddcjzl  float64     `json:"zmbddcjzl"  orm:"zmbddcjzl"  description:"主买大单成交额增量"`
	Zmbtdcjzlv int         `json:"zmbtdcjzlv" orm:"zmbtdcjzlv" description:"主买特大单成交量增量"`
	Zmbddcjzlv int         `json:"zmbddcjzlv" orm:"zmbddcjzlv" description:"主买大单成交量增量"`
	CreatedAt  *gtime.Time `json:"createdAt"  orm:"created_at" description:"创建时间"`
	UpdatedAt  *gtime.Time `json:"updatedAt"  orm:"updated_at" description:"更新时间"`
}
