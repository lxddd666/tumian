// Package ccidata
// @Link  https://github.com/bufanyun/hotgo
// @Copyright  Copyright (c) 2026 HotGo CLI
// @Author  Ms <133814250@qq.com>
// @License  https://github.com/bufanyun/hotgo/blob/master/LICENSE
// @AutoGenerate Version 2.18.6
package ccidata

import (
	"hotgo/internal/model/input/form"
	"hotgo/internal/model/input/stockin"

	"github.com/gogf/gf/v2/frame/g"
)

// ListReq 查询cci指标数据表列表
type ListReq struct {
	g.Meta `path:"/cciData/list" method:"get" tags:"cci指标数据表" summary:"获取cci指标数据表列表"`
	stockin.CciDataListInp
}

type ListRes struct {
	form.PageRes
	List []*stockin.CciDataListModel `json:"list"   dc:"数据列表"`
}

// ExportReq 导出cci指标数据表列表
type ExportReq struct {
	g.Meta `path:"/cciData/export" method:"get" tags:"cci指标数据表" summary:"导出cci指标数据表列表"`
	stockin.CciDataListInp
}

type ExportRes struct{}

// ViewReq 获取cci指标数据表指定信息
type ViewReq struct {
	g.Meta `path:"/cciData/view" method:"get" tags:"cci指标数据表" summary:"获取cci指标数据表指定信息"`
	stockin.CciDataViewInp
}

type ViewRes struct {
	*stockin.CciDataViewModel
}

// EditReq 修改/新增cci指标数据表
type EditReq struct {
	g.Meta `path:"/cciData/edit" method:"post" tags:"cci指标数据表" summary:"修改/新增cci指标数据表"`
	stockin.CciDataEditInp
}

type EditRes struct{}

// DeleteReq 删除cci指标数据表
type DeleteReq struct {
	g.Meta `path:"/cciData/delete" method:"post" tags:"cci指标数据表" summary:"删除cci指标数据表"`
	stockin.CciDataDeleteInp
}

type DeleteRes struct{}
