// Package stockscoremain
// @Link  https://github.com/bufanyun/hotgo
// @Copyright  Copyright (c) 2026 HotGo CLI
// @Author  Ms <133814250@qq.com>
// @License  https://github.com/bufanyun/hotgo/blob/master/LICENSE
// @AutoGenerate Version 2.18.6
package stockscoremain

import (
	"hotgo/internal/model/input/form"
	"hotgo/internal/model/input/stockin"

	"github.com/gogf/gf/v2/frame/g"
)

// ListReq 查询股票评分主表列表
type ListReq struct {
	g.Meta `path:"/stockScoreMain/list" method:"get" tags:"股票评分主表" summary:"获取股票评分主表列表"`
	stockin.StockScoreMainListInp
}

type ListRes struct {
	form.PageRes
	List []*stockin.StockScoreMainListModel `json:"list"   dc:"数据列表"`
}

// ExportReq 导出股票评分主表列表
type ExportReq struct {
	g.Meta `path:"/stockScoreMain/export" method:"get" tags:"股票评分主表" summary:"导出股票评分主表列表"`
	stockin.StockScoreMainListInp
}

type ExportRes struct{}

// ViewReq 获取股票评分主表指定信息
type ViewReq struct {
	g.Meta `path:"/stockScoreMain/view" method:"get" tags:"股票评分主表" summary:"获取股票评分主表指定信息"`
	stockin.StockScoreMainViewInp
}

type ViewRes struct {
	*stockin.StockScoreMainViewModel
}

// EditReq 修改/新增股票评分主表
type EditReq struct {
	g.Meta `path:"/stockScoreMain/edit" method:"post" tags:"股票评分主表" summary:"修改/新增股票评分主表"`
	stockin.StockScoreMainEditInp
}

type EditRes struct{}

// DeleteReq 删除股票评分主表
type DeleteReq struct {
	g.Meta `path:"/stockScoreMain/delete" method:"post" tags:"股票评分主表" summary:"删除股票评分主表"`
	stockin.StockScoreMainDeleteInp
}

type DeleteRes struct{}
