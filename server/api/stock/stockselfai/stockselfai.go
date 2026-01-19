// Package stockselfai
// @Link  https://github.com/bufanyun/hotgo
// @Copyright  Copyright (c) 2026 HotGo CLI
// @Author  Ms <133814250@qq.com>
// @License  https://github.com/bufanyun/hotgo/blob/master/LICENSE
// @AutoGenerate Version 2.18.6
package stockselfai

import (
	"hotgo/internal/model/input/form"
	"hotgo/internal/model/input/stockin"

	"github.com/gogf/gf/v2/frame/g"
)

// ListReq 查询ai 基本信息列表
type ListReq struct {
	g.Meta `path:"/stockSelfAi/list" method:"get" tags:"ai 基本信息" summary:"获取ai 基本信息列表"`
	stockin.StockSelfAiListInp
}

type ListRes struct {
	form.PageRes
	List []*stockin.StockSelfAiListModel `json:"list"   dc:"数据列表"`
}

// ExportReq 导出ai 基本信息列表
type ExportReq struct {
	g.Meta `path:"/stockSelfAi/export" method:"get" tags:"ai 基本信息" summary:"导出ai 基本信息列表"`
	stockin.StockSelfAiListInp
}

type ExportRes struct{}

// ViewReq 获取ai 基本信息指定信息
type ViewReq struct {
	g.Meta `path:"/stockSelfAi/view" method:"get" tags:"ai 基本信息" summary:"获取ai 基本信息指定信息"`
	stockin.StockSelfAiViewInp
}

type ViewRes struct {
	*stockin.StockSelfAiViewModel
}

// EditReq 修改/新增ai 基本信息
type EditReq struct {
	g.Meta `path:"/stockSelfAi/edit" method:"post" tags:"ai 基本信息" summary:"修改/新增ai 基本信息"`
	stockin.StockSelfAiEditInp
}

type EditRes struct{}

// DeleteReq 删除ai 基本信息
type DeleteReq struct {
	g.Meta `path:"/stockSelfAi/delete" method:"post" tags:"ai 基本信息" summary:"删除ai 基本信息"`
	stockin.StockSelfAiDeleteInp
}

type DeleteRes struct{}
