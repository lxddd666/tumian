// Package stockin
// @Link  https://github.com/bufanyun/hotgo
// @Copyright  Copyright (c) 2026 HotGo CLI
// @Author  Ms <133814250@qq.com>
// @License  https://github.com/bufanyun/hotgo/blob/master/LICENSE
// @AutoGenerate Version 2.18.6
package stockin

import (
	"context"
	"hotgo/internal/model/entity"
	"hotgo/internal/model/input/form"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// FlowOfFundsUpdateFields 修改资金流向明细表字段过滤
type FlowOfFundsUpdateFields struct {
	Symbol     string  `json:"symbol"     dc:"股票代码 (如: 000001.SZ)"`
	T          int     `json:"t"          dc:"交易时间 (通常为HHMMSS格式的整数)"`
	Zmbzds     int     `json:"zmbzds"     dc:"主买单总单数"`
	Zmszds     int     `json:"zmszds"     dc:"主卖单总单数"`
	Dddx       float64 `json:"dddx"       dc:"大单动向"`
	Zddy       float64 `json:"zddy"       dc:"涨跌动因"`
	Ddcf       float64 `json:"ddcf"       dc:"大单差分"`
	Zmbzdszl   int     `json:"zmbzdszl"   dc:"主买单总单数增量"`
	Zmszdszl   int     `json:"zmszdszl"   dc:"主卖单总单数增量"`
	Cjbszl     int     `json:"cjbszl"     dc:"成交笔数增量"`
	Zmbtdcje   float64 `json:"zmbtdcje"   dc:"主买特大单成交额"`
	Zmbddcje   float64 `json:"zmbddcje"   dc:"主买大单成交额"`
	Zmbzdcje   float64 `json:"zmbzdcje"   dc:"主买中单成交额"`
	Zmbxdcje   float64 `json:"zmbxdcje"   dc:"主买小单成交额"`
	Zmstdcje   float64 `json:"zmstdcje"   dc:"主卖特大单成交额"`
	Zmsddcje   float64 `json:"zmsddcje"   dc:"主卖大单成交额"`
	Zmszdcje   float64 `json:"zmszdcje"   dc:"主卖中单成交额"`
	Zmsxdcje   float64 `json:"zmsxdcje"   dc:"主卖小单成交额"`
	Bdmbtdcje  float64 `json:"bdmbtdcje"  dc:"被动买特大单成交额"`
	Bdmbddcje  float64 `json:"bdmbddcje"  dc:"被动买大单成交额"`
	Bdmbzdcje  float64 `json:"bdmbzdcje"  dc:"被动买中单成交额"`
	Bdmbxdcje  float64 `json:"bdmbxdcje"  dc:"被动买小单成交额"`
	Bdmstdcje  float64 `json:"bdmstdcje"  dc:"被动卖特大单成交额"`
	Bdmsddcje  float64 `json:"bdmsddcje"  dc:"被动卖大单成交额"`
	Bdmszdcje  float64 `json:"bdmszdcje"  dc:"被动卖中单成交额"`
	Bdmsxdcje  float64 `json:"bdmsxdcje"  dc:"被动卖小单成交额"`
	Zmbtdcjl   int     `json:"zmbtdcjl"   dc:"主买特大单成交量"`
	Zmbddcjl   int     `json:"zmbddcjl"   dc:"主买大单成交量"`
	Zmbzdcjl   int     `json:"zmbzdcjl"   dc:"主买中单成交量"`
	Zmbxdcjl   int     `json:"zmbxdcjl"   dc:"主买小单成交量"`
	Zmstdcjl   int     `json:"zmstdcjl"   dc:"主卖特大单成交量"`
	Zmsddcjl   int     `json:"zmsddcjl"   dc:"主卖大单成交量"`
	Zmszdcjl   int     `json:"zmszdcjl"   dc:"主卖中单成交量"`
	Zmsxdcjl   int     `json:"zmsxdcjl"   dc:"主卖小单成交量"`
	Bdmbtdcjl  int     `json:"bdmbtdcjl"  dc:"被动买特大单成交量"`
	Bdmbddcjl  int     `json:"bdmbddcjl"  dc:"被动买大单成交量"`
	Bdmbzdcjl  int     `json:"bdmbzdcjl"  dc:"被动买中单成交量"`
	Bdmbxdcjl  int     `json:"bdmbxdcjl"  dc:"被动买小单成交量"`
	Bdmstdcjl  int     `json:"bdmstdcjl"  dc:"被动卖特大单成交量"`
	Bdmsddcjl  int     `json:"bdmsddcjl"  dc:"被动卖大单成交量"`
	Bdmszdcjl  int     `json:"bdmszdcjl"  dc:"被动卖中单成交量"`
	Bdmsxdcjl  int     `json:"bdmsxdcjl"  dc:"被动卖小单成交量"`
	Zmbtdcjzl  float64 `json:"zmbtdcjzl"  dc:"主买特大单成交额增量"`
	Zmbddcjzl  float64 `json:"zmbddcjzl"  dc:"主买大单成交额增量"`
	Zmbtdcjzlv int     `json:"zmbtdcjzlv" dc:"主买特大单成交量增量"`
	Zmbddcjzlv int     `json:"zmbddcjzlv" dc:"主买大单成交量增量"`
}

// FlowOfFundsInsertFields 新增资金流向明细表字段过滤
type FlowOfFundsInsertFields struct {
	Symbol     string  `json:"symbol"     dc:"股票代码 (如: 000001.SZ)"`
	T          int     `json:"t"          dc:"交易时间 (通常为HHMMSS格式的整数)"`
	Zmbzds     int     `json:"zmbzds"     dc:"主买单总单数"`
	Zmszds     int     `json:"zmszds"     dc:"主卖单总单数"`
	Dddx       float64 `json:"dddx"       dc:"大单动向"`
	Zddy       float64 `json:"zddy"       dc:"涨跌动因"`
	Ddcf       float64 `json:"ddcf"       dc:"大单差分"`
	Zmbzdszl   int     `json:"zmbzdszl"   dc:"主买单总单数增量"`
	Zmszdszl   int     `json:"zmszdszl"   dc:"主卖单总单数增量"`
	Cjbszl     int     `json:"cjbszl"     dc:"成交笔数增量"`
	Zmbtdcje   float64 `json:"zmbtdcje"   dc:"主买特大单成交额"`
	Zmbddcje   float64 `json:"zmbddcje"   dc:"主买大单成交额"`
	Zmbzdcje   float64 `json:"zmbzdcje"   dc:"主买中单成交额"`
	Zmbxdcje   float64 `json:"zmbxdcje"   dc:"主买小单成交额"`
	Zmstdcje   float64 `json:"zmstdcje"   dc:"主卖特大单成交额"`
	Zmsddcje   float64 `json:"zmsddcje"   dc:"主卖大单成交额"`
	Zmszdcje   float64 `json:"zmszdcje"   dc:"主卖中单成交额"`
	Zmsxdcje   float64 `json:"zmsxdcje"   dc:"主卖小单成交额"`
	Bdmbtdcje  float64 `json:"bdmbtdcje"  dc:"被动买特大单成交额"`
	Bdmbddcje  float64 `json:"bdmbddcje"  dc:"被动买大单成交额"`
	Bdmbzdcje  float64 `json:"bdmbzdcje"  dc:"被动买中单成交额"`
	Bdmbxdcje  float64 `json:"bdmbxdcje"  dc:"被动买小单成交额"`
	Bdmstdcje  float64 `json:"bdmstdcje"  dc:"被动卖特大单成交额"`
	Bdmsddcje  float64 `json:"bdmsddcje"  dc:"被动卖大单成交额"`
	Bdmszdcje  float64 `json:"bdmszdcje"  dc:"被动卖中单成交额"`
	Bdmsxdcje  float64 `json:"bdmsxdcje"  dc:"被动卖小单成交额"`
	Zmbtdcjl   int     `json:"zmbtdcjl"   dc:"主买特大单成交量"`
	Zmbddcjl   int     `json:"zmbddcjl"   dc:"主买大单成交量"`
	Zmbzdcjl   int     `json:"zmbzdcjl"   dc:"主买中单成交量"`
	Zmbxdcjl   int     `json:"zmbxdcjl"   dc:"主买小单成交量"`
	Zmstdcjl   int     `json:"zmstdcjl"   dc:"主卖特大单成交量"`
	Zmsddcjl   int     `json:"zmsddcjl"   dc:"主卖大单成交量"`
	Zmszdcjl   int     `json:"zmszdcjl"   dc:"主卖中单成交量"`
	Zmsxdcjl   int     `json:"zmsxdcjl"   dc:"主卖小单成交量"`
	Bdmbtdcjl  int     `json:"bdmbtdcjl"  dc:"被动买特大单成交量"`
	Bdmbddcjl  int     `json:"bdmbddcjl"  dc:"被动买大单成交量"`
	Bdmbzdcjl  int     `json:"bdmbzdcjl"  dc:"被动买中单成交量"`
	Bdmbxdcjl  int     `json:"bdmbxdcjl"  dc:"被动买小单成交量"`
	Bdmstdcjl  int     `json:"bdmstdcjl"  dc:"被动卖特大单成交量"`
	Bdmsddcjl  int     `json:"bdmsddcjl"  dc:"被动卖大单成交量"`
	Bdmszdcjl  int     `json:"bdmszdcjl"  dc:"被动卖中单成交量"`
	Bdmsxdcjl  int     `json:"bdmsxdcjl"  dc:"被动卖小单成交量"`
	Zmbtdcjzl  float64 `json:"zmbtdcjzl"  dc:"主买特大单成交额增量"`
	Zmbddcjzl  float64 `json:"zmbddcjzl"  dc:"主买大单成交额增量"`
	Zmbtdcjzlv int     `json:"zmbtdcjzlv" dc:"主买特大单成交量增量"`
	Zmbddcjzlv int     `json:"zmbddcjzlv" dc:"主买大单成交量增量"`
}

// FlowOfFundsEditInp 修改/新增资金流向明细表
type FlowOfFundsEditInp struct {
	entity.FlowOfFunds
}

func (in *FlowOfFundsEditInp) Filter(ctx context.Context) (err error) {
	// 验证股票代码 (如: 000001.SZ)
	if err := g.Validator().Rules("required").Data(in.Symbol).Messages("股票代码 (如: 000001.SZ)不能为空").Run(ctx); err != nil {
		return err.Current()
	}
	// 验证交易时间 (通常为HHMMSS格式的整数)
	if err := g.Validator().Rules("required").Data(in.T).Messages("交易时间 (通常为HHMMSS格式的整数)不能为空").Run(ctx); err != nil {
		return err.Current()
	}
	return
}

type FlowOfFundsEditModel struct{}

// FlowOfFundsDeleteInp 删除资金流向明细表
type FlowOfFundsDeleteInp struct {
	Id interface{} `json:"id" v:"required#主键ID不能为空" dc:"主键ID"`
}

func (in *FlowOfFundsDeleteInp) Filter(ctx context.Context) (err error) {
	return
}

type FlowOfFundsDeleteModel struct{}

// FlowOfFundsViewInp 获取指定资金流向明细表信息
type FlowOfFundsViewInp struct {
	Id int64 `json:"id" v:"required#主键ID不能为空" dc:"主键ID"`
}

func (in *FlowOfFundsViewInp) Filter(ctx context.Context) (err error) {
	return
}

type FlowOfFundsViewModel struct {
	entity.FlowOfFunds
}

// FlowOfFundsListInp 获取资金流向明细表列表
type FlowOfFundsListInp struct {
	form.PageReq
	Id        int64         `json:"id"        dc:"主键ID"`
	CreatedAt []*gtime.Time `json:"createdAt" dc:"创建时间"`
}

func (in *FlowOfFundsListInp) Filter(ctx context.Context) (err error) {
	return
}

type FlowOfFundsListModel struct {
	Id         int64       `json:"id"         dc:"主键ID"`
	Symbol     string      `json:"symbol"     dc:"股票代码 (如: 000001.SZ)"`
	T          int         `json:"t"          dc:"交易时间 (通常为HHMMSS格式的整数)"`
	Zmbzds     int         `json:"zmbzds"     dc:"主买单总单数"`
	Zmszds     int         `json:"zmszds"     dc:"主卖单总单数"`
	Dddx       float64     `json:"dddx"       dc:"大单动向"`
	Zddy       float64     `json:"zddy"       dc:"涨跌动因"`
	Ddcf       float64     `json:"ddcf"       dc:"大单差分"`
	Zmbzdszl   int         `json:"zmbzdszl"   dc:"主买单总单数增量"`
	Zmszdszl   int         `json:"zmszdszl"   dc:"主卖单总单数增量"`
	Cjbszl     int         `json:"cjbszl"     dc:"成交笔数增量"`
	Zmbtdcje   float64     `json:"zmbtdcje"   dc:"主买特大单成交额"`
	Zmbddcje   float64     `json:"zmbddcje"   dc:"主买大单成交额"`
	Zmbzdcje   float64     `json:"zmbzdcje"   dc:"主买中单成交额"`
	Zmbxdcje   float64     `json:"zmbxdcje"   dc:"主买小单成交额"`
	Zmstdcje   float64     `json:"zmstdcje"   dc:"主卖特大单成交额"`
	Zmsddcje   float64     `json:"zmsddcje"   dc:"主卖大单成交额"`
	Zmszdcje   float64     `json:"zmszdcje"   dc:"主卖中单成交额"`
	Zmsxdcje   float64     `json:"zmsxdcje"   dc:"主卖小单成交额"`
	Bdmbtdcje  float64     `json:"bdmbtdcje"  dc:"被动买特大单成交额"`
	Bdmbddcje  float64     `json:"bdmbddcje"  dc:"被动买大单成交额"`
	Bdmbzdcje  float64     `json:"bdmbzdcje"  dc:"被动买中单成交额"`
	Bdmbxdcje  float64     `json:"bdmbxdcje"  dc:"被动买小单成交额"`
	Bdmstdcje  float64     `json:"bdmstdcje"  dc:"被动卖特大单成交额"`
	Bdmsddcje  float64     `json:"bdmsddcje"  dc:"被动卖大单成交额"`
	Bdmszdcje  float64     `json:"bdmszdcje"  dc:"被动卖中单成交额"`
	Bdmsxdcje  float64     `json:"bdmsxdcje"  dc:"被动卖小单成交额"`
	Zmbtdcjl   int         `json:"zmbtdcjl"   dc:"主买特大单成交量"`
	Zmbddcjl   int         `json:"zmbddcjl"   dc:"主买大单成交量"`
	Zmbzdcjl   int         `json:"zmbzdcjl"   dc:"主买中单成交量"`
	Zmbxdcjl   int         `json:"zmbxdcjl"   dc:"主买小单成交量"`
	Zmstdcjl   int         `json:"zmstdcjl"   dc:"主卖特大单成交量"`
	Zmsddcjl   int         `json:"zmsddcjl"   dc:"主卖大单成交量"`
	Zmszdcjl   int         `json:"zmszdcjl"   dc:"主卖中单成交量"`
	Zmsxdcjl   int         `json:"zmsxdcjl"   dc:"主卖小单成交量"`
	Bdmbtdcjl  int         `json:"bdmbtdcjl"  dc:"被动买特大单成交量"`
	Bdmbddcjl  int         `json:"bdmbddcjl"  dc:"被动买大单成交量"`
	Bdmbzdcjl  int         `json:"bdmbzdcjl"  dc:"被动买中单成交量"`
	Bdmbxdcjl  int         `json:"bdmbxdcjl"  dc:"被动买小单成交量"`
	Bdmstdcjl  int         `json:"bdmstdcjl"  dc:"被动卖特大单成交量"`
	Bdmsddcjl  int         `json:"bdmsddcjl"  dc:"被动卖大单成交量"`
	Bdmszdcjl  int         `json:"bdmszdcjl"  dc:"被动卖中单成交量"`
	Bdmsxdcjl  int         `json:"bdmsxdcjl"  dc:"被动卖小单成交量"`
	Zmbtdcjzl  float64     `json:"zmbtdcjzl"  dc:"主买特大单成交额增量"`
	Zmbddcjzl  float64     `json:"zmbddcjzl"  dc:"主买大单成交额增量"`
	Zmbtdcjzlv int         `json:"zmbtdcjzlv" dc:"主买特大单成交量增量"`
	Zmbddcjzlv int         `json:"zmbddcjzlv" dc:"主买大单成交量增量"`
	CreatedAt  *gtime.Time `json:"createdAt"  dc:"创建时间"`
	UpdatedAt  *gtime.Time `json:"updatedAt"  dc:"更新时间"`
}

// FlowOfFundsExportModel 导出资金流向明细表
type FlowOfFundsExportModel struct {
	Id         int64       `json:"id"         dc:"主键ID"`
	Symbol     string      `json:"symbol"     dc:"股票代码 (如: 000001.SZ)"`
	T          int         `json:"t"          dc:"交易时间 (通常为HHMMSS格式的整数)"`
	Zmbzds     int         `json:"zmbzds"     dc:"主买单总单数"`
	Zmszds     int         `json:"zmszds"     dc:"主卖单总单数"`
	Dddx       float64     `json:"dddx"       dc:"大单动向"`
	Zddy       float64     `json:"zddy"       dc:"涨跌动因"`
	Ddcf       float64     `json:"ddcf"       dc:"大单差分"`
	Zmbzdszl   int         `json:"zmbzdszl"   dc:"主买单总单数增量"`
	Zmszdszl   int         `json:"zmszdszl"   dc:"主卖单总单数增量"`
	Cjbszl     int         `json:"cjbszl"     dc:"成交笔数增量"`
	Zmbtdcje   float64     `json:"zmbtdcje"   dc:"主买特大单成交额"`
	Zmbddcje   float64     `json:"zmbddcje"   dc:"主买大单成交额"`
	Zmbzdcje   float64     `json:"zmbzdcje"   dc:"主买中单成交额"`
	Zmbxdcje   float64     `json:"zmbxdcje"   dc:"主买小单成交额"`
	Zmstdcje   float64     `json:"zmstdcje"   dc:"主卖特大单成交额"`
	Zmsddcje   float64     `json:"zmsddcje"   dc:"主卖大单成交额"`
	Zmszdcje   float64     `json:"zmszdcje"   dc:"主卖中单成交额"`
	Zmsxdcje   float64     `json:"zmsxdcje"   dc:"主卖小单成交额"`
	Bdmbtdcje  float64     `json:"bdmbtdcje"  dc:"被动买特大单成交额"`
	Bdmbddcje  float64     `json:"bdmbddcje"  dc:"被动买大单成交额"`
	Bdmbzdcje  float64     `json:"bdmbzdcje"  dc:"被动买中单成交额"`
	Bdmbxdcje  float64     `json:"bdmbxdcje"  dc:"被动买小单成交额"`
	Bdmstdcje  float64     `json:"bdmstdcje"  dc:"被动卖特大单成交额"`
	Bdmsddcje  float64     `json:"bdmsddcje"  dc:"被动卖大单成交额"`
	Bdmszdcje  float64     `json:"bdmszdcje"  dc:"被动卖中单成交额"`
	Bdmsxdcje  float64     `json:"bdmsxdcje"  dc:"被动卖小单成交额"`
	Zmbtdcjl   int         `json:"zmbtdcjl"   dc:"主买特大单成交量"`
	Zmbddcjl   int         `json:"zmbddcjl"   dc:"主买大单成交量"`
	Zmbzdcjl   int         `json:"zmbzdcjl"   dc:"主买中单成交量"`
	Zmbxdcjl   int         `json:"zmbxdcjl"   dc:"主买小单成交量"`
	Zmstdcjl   int         `json:"zmstdcjl"   dc:"主卖特大单成交量"`
	Zmsddcjl   int         `json:"zmsddcjl"   dc:"主卖大单成交量"`
	Zmszdcjl   int         `json:"zmszdcjl"   dc:"主卖中单成交量"`
	Zmsxdcjl   int         `json:"zmsxdcjl"   dc:"主卖小单成交量"`
	Bdmbtdcjl  int         `json:"bdmbtdcjl"  dc:"被动买特大单成交量"`
	Bdmbddcjl  int         `json:"bdmbddcjl"  dc:"被动买大单成交量"`
	Bdmbzdcjl  int         `json:"bdmbzdcjl"  dc:"被动买中单成交量"`
	Bdmbxdcjl  int         `json:"bdmbxdcjl"  dc:"被动买小单成交量"`
	Bdmstdcjl  int         `json:"bdmstdcjl"  dc:"被动卖特大单成交量"`
	Bdmsddcjl  int         `json:"bdmsddcjl"  dc:"被动卖大单成交量"`
	Bdmszdcjl  int         `json:"bdmszdcjl"  dc:"被动卖中单成交量"`
	Bdmsxdcjl  int         `json:"bdmsxdcjl"  dc:"被动卖小单成交量"`
	Zmbtdcjzl  float64     `json:"zmbtdcjzl"  dc:"主买特大单成交额增量"`
	Zmbddcjzl  float64     `json:"zmbddcjzl"  dc:"主买大单成交额增量"`
	Zmbtdcjzlv int         `json:"zmbtdcjzlv" dc:"主买特大单成交量增量"`
	Zmbddcjzlv int         `json:"zmbddcjzlv" dc:"主买大单成交量增量"`
	CreatedAt  *gtime.Time `json:"createdAt"  dc:"创建时间"`
	UpdatedAt  *gtime.Time `json:"updatedAt"  dc:"更新时间"`
}

type GetFlowOfFundsInp struct {
	Symbol     string `json:"symbol"     dc:"股票代码 (如: 000001.SZ)"`
	Interval   string `json:"interval"  dc:"分时级别 (例如: d)"`
	AdjustType string `json:"adjustType"  dc:"除权类型 (例如: n)"`
	Token      string `json:"token" dc:"token证书"`
	StartTime  string `json:"st" dc:"开始时间"`
	EndTime    string `json:"et" dc:"结束时间"`
	Limit      int    `json:"lt" dc:"最新条数"`
}
