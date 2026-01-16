// Package fundstockholding
// @Link  https://github.com/bufanyun/hotgo
// @Copyright  Copyright (c) 2026 HotGo CLI
// @Author  Ms <133814250@qq.com>
// @License  https://github.com/bufanyun/hotgo/blob/master/LICENSE
// @AutoGenerate Version 2.18.6
package fundstockholding

import (
	"hotgo/internal/model/entity"
	"hotgo/internal/model/input/form"
	"hotgo/internal/model/input/stockin"

	"github.com/gogf/gf/v2/frame/g"
)

// ListReq 查询基金持股明细表 (来源于基金定期报告)列表
type ListReq struct {
	g.Meta `path:"/fundStockHolding/list" method:"get" tags:"基金持股明细表 (来源于基金定期报告)" summary:"获取基金持股明细表 (来源于基金定期报告)列表"`
	stockin.FundStockHoldingListInp
}

type ListRes struct {
	form.PageRes
	List []*stockin.FundStockHoldingListModel `json:"list"   dc:"数据列表"`
}

// ExportReq 导出基金持股明细表 (来源于基金定期报告)列表
type ExportReq struct {
	g.Meta `path:"/fundStockHolding/export" method:"get" tags:"基金持股明细表 (来源于基金定期报告)" summary:"导出基金持股明细表 (来源于基金定期报告)列表"`
	stockin.FundStockHoldingListInp
}

type ExportRes struct{}

// ViewReq 获取基金持股明细表 (来源于基金定期报告)指定信息
type ViewReq struct {
	g.Meta `path:"/fundStockHolding/view" method:"get" tags:"基金持股明细表 (来源于基金定期报告)" summary:"获取基金持股明细表 (来源于基金定期报告)指定信息"`
	stockin.FundStockHoldingViewInp
}

type ViewRes struct {
	*stockin.FundStockHoldingViewModel
}

// EditReq 修改/新增基金持股明细表 (来源于基金定期报告)
type EditReq struct {
	g.Meta `path:"/fundStockHolding/edit" method:"post" tags:"基金持股明细表 (来源于基金定期报告)" summary:"修改/新增基金持股明细表 (来源于基金定期报告)"`
	stockin.FundStockHoldingEditInp
}

type EditRes struct{}

// DeleteReq 删除基金持股明细表 (来源于基金定期报告)
type DeleteReq struct {
	g.Meta `path:"/fundStockHolding/delete" method:"post" tags:"基金持股明细表 (来源于基金定期报告)" summary:"删除基金持股明细表 (来源于基金定期报告)"`
	stockin.FundStockHoldingDeleteInp
}

type DeleteRes struct{}

// GetFundStockHoldingReq 获取基金持股明细表数据
type GetFundStockHoldingReq struct {
	g.Meta `path:"/fundStockHolding/getFundStockHolding" method:"get" tags:"基金持股明细表 (来源于基金定期报告)" summary:"获取基金持股明细表数据"`
	stockin.FundStockHoldingGetFundStockHoldingInp
}

type GetFundStockHoldingRes struct {
	Data []*entity.FundStockHolding `json:"data" dc:"返回数据"`
}
