// Package flowoffunds
// @Link  https://github.com/bufanyun/hotgo
// @Copyright  Copyright (c) 2026 HotGo CLI
// @Author  Ms <133814250@qq.com>
// @License  https://github.com/bufanyun/hotgo/blob/master/LICENSE
// @AutoGenerate Version 2.18.6
package flowoffunds

import (
	"hotgo/internal/model/entity"
	"hotgo/internal/model/input/form"
	"hotgo/internal/model/input/stockin"

	"github.com/gogf/gf/v2/frame/g"
)

// ListReq 查询资金流向明细表列表
type ListReq struct {
	g.Meta `path:"/flowOfFunds/list" method:"get" tags:"资金流向明细表" summary:"获取资金流向明细表列表"`
	stockin.FlowOfFundsListInp
}

type ListRes struct {
	form.PageRes
	List []*stockin.FlowOfFundsListModel `json:"list"   dc:"数据列表"`
}

// ExportReq 导出资金流向明细表列表
type ExportReq struct {
	g.Meta `path:"/flowOfFunds/export" method:"get" tags:"资金流向明细表" summary:"导出资金流向明细表列表"`
	stockin.FlowOfFundsListInp
}

type ExportRes struct{}

// ViewReq 获取资金流向明细表指定信息
type ViewReq struct {
	g.Meta `path:"/flowOfFunds/view" method:"get" tags:"资金流向明细表" summary:"获取资金流向明细表指定信息"`
	stockin.FlowOfFundsViewInp
}

type ViewRes struct {
	*stockin.FlowOfFundsViewModel
}

// EditReq 修改/新增资金流向明细表
type EditReq struct {
	g.Meta `path:"/flowOfFunds/edit" method:"post" tags:"资金流向明细表" summary:"修改/新增资金流向明细表"`
	stockin.FlowOfFundsEditInp
}

type EditRes struct{}

// DeleteReq 删除资金流向明细表
type DeleteReq struct {
	g.Meta `path:"/flowOfFunds/delete" method:"post" tags:"资金流向明细表" summary:"删除资金流向明细表"`
	stockin.FlowOfFundsDeleteInp
}

type DeleteRes struct{}

// GetFlowOfFundsReq 获取资金流向明细表数据
type GetFlowOfFundsReq struct {
	g.Meta `path:"/bollData/getFlowOfFunds" method:"get" tags:"获取资金流向明细表数据" summary:"获取资金流向明细表数据"`
	stockin.GetFlowOfFundsInp
}

type GetFlowOfFundsRes struct {
	Data []*entity.FlowOfFunds `json:"data" dc:"返回数据"`
}
