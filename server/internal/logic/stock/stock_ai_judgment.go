// Package stock
// @Link  https://github.com/bufanyun/hotgo
// @Copyright  Copyright (c) 2026 HotGo CLI
// @Author  Ms <133814250@qq.com>
// @License  https://github.com/bufanyun/hotgo/blob/master/LICENSE
// @AutoGenerate Version 2.18.6
package stock

import (
	"context"
	"github.com/gogf/gf/v2/database/gdb"
	"hotgo/internal/dao"
	"hotgo/internal/library/hgorm/handler"
	"hotgo/internal/service"
)

type sStockAiJudgment struct{}

func NewStockAiJudgment() *sStockAiJudgment {
	return &sStockAiJudgment{}
}

func init() {
	service.RegisterStockAiJudgment(NewStockAiJudgment())
}

// Model ai 选股判断ORM模型
func (s *sStockAiJudgment) Model(ctx context.Context, option ...*handler.Option) *gdb.Model {
	return handler.Model(dao.StockAiJudgment.Ctx(ctx), option...)
}

// AITemp ai 提示模板
var AITemp = `交易对:%s 现价:%s
[1D趋势]
K线:%s
MACD中，diff/dea/macd/ema12/ema26:%d/%d/%d/%d%d
MA中，m3/m5/m10/m15/m20/m30/m60:%d/%d/%d/%d/%d/%d/%d
BOLL中 上轨/下轨/中轨:%d/%d/%d
KDJ中: K值/D值/J值:%d/%d/%d

[上市公司详情-]
近一年各季度利润，净利润/营业收入/基本每股收益/毛利率/营业利润/综合收益总额:%d/%d/%d/%d/%d/%d

请按此JSON格式输出决策(无Markdown):
{{
 "symbol": "BTC_USDT|ETH_USDT|null",
 "action": "Open(Short|Long)|Close(Short|Long)|HOLD",
 "confidence": 0.0-1.0,
 "rationale": "理由<50字",
 "target": "止盈价/0",
 "stop": "止损价/0"
}}
要求:
1.仅输出JSON
2.输出一个中短期的策略
3.开仓必填止盈损
`

const SystemMessage = `你资深量化分析师。根据1d的K线及MACD/MA/BOLL/KDJ指标以及财报和公司分析行情，输出JSON交易决策。`
