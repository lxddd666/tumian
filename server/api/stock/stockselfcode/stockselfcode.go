// Package stockselfcode
// @Link  https://github.com/bufanyun/hotgo
// @Copyright  Copyright (c) 2026 HotGo CLI
// @Author  Ms <133814250@qq.com>
// @License  https://github.com/bufanyun/hotgo/blob/master/LICENSE
// @AutoGenerate Version 2.18.6
package stockselfcode

import (
	"hotgo/internal/model/input/stockin"

	"github.com/gogf/gf/v2/frame/g"
)

type SelfCodeIndicatorsApiReq struct {
	g.Meta `path:"/stockSelfCode/selfCodeIndicatorsApi" method:"Get" tags:"自选股票" summary:"技术指标获取api获取（智兔api获取）"`
	stockin.SelfCodeIndicatorsApiInp
}

type SelfCodeIndicatorsApiRes struct{}

type SelfStockWorkingCapitalInfoApiReq struct {
	g.Meta `path:"/stockSelfCode/selfStockWorkingCapitalInfoApi" method:"Get" tags:"自选股票" summary:"运营资金情况（智兔api获取）"`
	stockin.SelfCodeIndicatorsApiInp
}

type SelfStockWorkingCapitalInfoApiRes struct{}
