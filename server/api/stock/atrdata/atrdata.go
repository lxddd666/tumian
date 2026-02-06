// Package atrdata
// @Link  https://github.com/bufanyun/hotgo
// @Copyright  Copyright (c) 2026 HotGo CLI
// @Author  Ms <133814250@qq.com>
// @License  https://github.com/bufanyun/hotgo/blob/master/LICENSE
// @AutoGenerate Version 2.18.6
package atrdata

import (
	"hotgo/internal/model/input/form"
	"hotgo/internal/model/input/stockin"

	"github.com/gogf/gf/v2/frame/g"
)

// ListReq 查询atr指标数据表列表
type ListReq struct {
	g.Meta `path:"/atrData/list" method:"get" tags:"atr指标数据表" summary:"获取atr指标数据表列表"`
	stockin.AtrDataListInp
}

type ListRes struct {
	form.PageRes
	List []*stockin.AtrDataListModel `json:"list"   dc:"数据列表"`
}

// ExportReq 导出atr指标数据表列表
type ExportReq struct {
	g.Meta `path:"/atrData/export" method:"get" tags:"atr指标数据表" summary:"导出atr指标数据表列表"`
	stockin.AtrDataListInp
}

type ExportRes struct{}

// ViewReq 获取atr指标数据表指定信息
type ViewReq struct {
	g.Meta `path:"/atrData/view" method:"get" tags:"atr指标数据表" summary:"获取atr指标数据表指定信息"`
	stockin.AtrDataViewInp
}

type ViewRes struct {
	*stockin.AtrDataViewModel
}

// EditReq 修改/新增atr指标数据表
type EditReq struct {
	g.Meta `path:"/atrData/edit" method:"post" tags:"atr指标数据表" summary:"修改/新增atr指标数据表"`
	stockin.AtrDataEditInp
}

type EditRes struct{}

// DeleteReq 删除atr指标数据表
type DeleteReq struct {
	g.Meta `path:"/atrData/delete" method:"post" tags:"atr指标数据表" summary:"删除atr指标数据表"`
	stockin.AtrDataDeleteInp
}

type DeleteRes struct{}
