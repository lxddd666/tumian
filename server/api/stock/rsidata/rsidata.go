// Package rsidata
// @Link  https://github.com/bufanyun/hotgo
// @Copyright  Copyright (c) 2026 HotGo CLI
// @Author  Ms <133814250@qq.com>
// @License  https://github.com/bufanyun/hotgo/blob/master/LICENSE
// @AutoGenerate Version 2.18.6
package rsidata

import (
	"hotgo/internal/model/input/form"
	"hotgo/internal/model/input/stockin"

	"github.com/gogf/gf/v2/frame/g"
)

// ListReq 查询rsi指标数据表列表
type ListReq struct {
	g.Meta `path:"/rsiData/list" method:"get" tags:"rsi指标数据表" summary:"获取rsi指标数据表列表"`
	stockin.RsiDataListInp
}

type ListRes struct {
	form.PageRes
	List []*stockin.RsiDataListModel `json:"list"   dc:"数据列表"`
}

// ExportReq 导出rsi指标数据表列表
type ExportReq struct {
	g.Meta `path:"/rsiData/export" method:"get" tags:"rsi指标数据表" summary:"导出rsi指标数据表列表"`
	stockin.RsiDataListInp
}

type ExportRes struct{}

// ViewReq 获取rsi指标数据表指定信息
type ViewReq struct {
	g.Meta `path:"/rsiData/view" method:"get" tags:"rsi指标数据表" summary:"获取rsi指标数据表指定信息"`
	stockin.RsiDataViewInp
}

type ViewRes struct {
	*stockin.RsiDataViewModel
}

// EditReq 修改/新增rsi指标数据表
type EditReq struct {
	g.Meta `path:"/rsiData/edit" method:"post" tags:"rsi指标数据表" summary:"修改/新增rsi指标数据表"`
	stockin.RsiDataEditInp
}

type EditRes struct{}

// DeleteReq 删除rsi指标数据表
type DeleteReq struct {
	g.Meta `path:"/rsiData/delete" method:"post" tags:"rsi指标数据表" summary:"删除rsi指标数据表"`
	stockin.RsiDataDeleteInp
}

type DeleteRes struct{}
