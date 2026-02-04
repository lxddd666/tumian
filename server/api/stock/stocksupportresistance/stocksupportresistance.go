// Package stocksupportresistance
// @Link  https://github.com/bufanyun/hotgo
// @Copyright  Copyright (c) 2026 HotGo CLI
// @Author  Ms <133814250@qq.com>
// @License  https://github.com/bufanyun/hotgo/blob/master/LICENSE
// @AutoGenerate Version 2.18.6
package stocksupportresistance

import (
	"hotgo/internal/model/input/form"
	"hotgo/internal/model/input/stockin"

	"github.com/gogf/gf/v2/frame/g"
)

// ListReq 查询股票支撑阻力表列表
type ListReq struct {
	g.Meta `path:"/stockSupportResistance/list" method:"get" tags:"股票支撑阻力表" summary:"获取股票支撑阻力表列表"`
	stockin.StockSupportResistanceListInp
}

type ListRes struct {
	form.PageRes
	List []*stockin.StockSupportResistanceListModel `json:"list"   dc:"数据列表"`
}

// ExportReq 导出股票支撑阻力表列表
type ExportReq struct {
	g.Meta `path:"/stockSupportResistance/export" method:"get" tags:"股票支撑阻力表" summary:"导出股票支撑阻力表列表"`
	stockin.StockSupportResistanceListInp
}

type ExportRes struct{}

// ViewReq 获取股票支撑阻力表指定信息
type ViewReq struct {
	g.Meta `path:"/stockSupportResistance/view" method:"get" tags:"股票支撑阻力表" summary:"获取股票支撑阻力表指定信息"`
	stockin.StockSupportResistanceViewInp
}

type ViewRes struct {
	*stockin.StockSupportResistanceViewModel
}

// EditReq 修改/新增股票支撑阻力表
type EditReq struct {
	g.Meta `path:"/stockSupportResistance/edit" method:"post" tags:"股票支撑阻力表" summary:"修改/新增股票支撑阻力表"`
	stockin.StockSupportResistanceEditInp
}

type EditRes struct{}

// DeleteReq 删除股票支撑阻力表
type DeleteReq struct {
	g.Meta `path:"/stockSupportResistance/delete" method:"post" tags:"股票支撑阻力表" summary:"删除股票支撑阻力表"`
	stockin.StockSupportResistanceDeleteInp
}

type DeleteRes struct{}
