// Package stocklist
// @Link  https://github.com/bufanyun/hotgo
// @Copyright  Copyright (c) 2026 HotGo CLI
// @Author  Ms <133814250@qq.com>
// @License  https://github.com/bufanyun/hotgo/blob/master/LICENSE
// @AutoGenerate Version 2.18.6
package stocklist

import (
	"hotgo/internal/model/entity"
	"hotgo/internal/model/input/form"
	"hotgo/internal/model/input/stockin"

	"github.com/gogf/gf/v2/frame/g"
)

// ListReq 查询股票列表核心表列表
type ListReq struct {
	g.Meta `path:"/stockList/list" method:"get" tags:"股票列表核心表" summary:"获取股票列表核心表列表"`
	stockin.StockListListInp
}

type ListRes struct {
	form.PageRes
	List []*stockin.StockListListModel `json:"list"   dc:"数据列表"`
}

// ExportReq 导出股票列表核心表列表
type ExportReq struct {
	g.Meta `path:"/stockList/export" method:"get" tags:"股票列表核心表" summary:"导出股票列表核心表列表"`
	stockin.StockListListInp
}

type ExportRes struct{}

// ViewReq 获取股票列表核心表指定信息
type ViewReq struct {
	g.Meta `path:"/stockList/view" method:"get" tags:"股票列表核心表" summary:"获取股票列表核心表指定信息"`
	stockin.StockListViewInp
}

type ViewRes struct {
	*stockin.StockListViewModel
}

// EditReq 修改/新增股票列表核心表
type EditReq struct {
	g.Meta `path:"/stockList/edit" method:"post" tags:"股票列表核心表" summary:"修改/新增股票列表核心表"`
	stockin.StockListEditInp
}

type EditRes struct{}

// DeleteReq 删除股票列表核心表
type DeleteReq struct {
	g.Meta `path:"/stockList/delete" method:"post" tags:"股票列表核心表" summary:"删除股票列表核心表"`
	stockin.StockListDeleteInp
}

type DeleteRes struct{}

// StatusReq 更新股票列表核心表状态
type StatusReq struct {
	g.Meta `path:"/stockList/status" method:"post" tags:"股票列表核心表" summary:"更新股票列表核心表状态"`
	stockin.StockListStatusInp
}

type StatusRes struct{}

// GetStockListReq 获取股票列表核心表数据
type GetStockListReq struct {
	g.Meta `path:"/stockList/getStockList" method:"get" tags:"股票列表核心表" summary:"获取股票列表核心表数据"`
	stockin.StockListGetStockListInp
}

type GetStockListRes struct {
	Data []*entity.StockList `json:"data" dc:"返回数据"`
}
