// Package bolldata
// @Link  https://github.com/bufanyun/hotgo
// @Copyright  Copyright (c) 2026 HotGo CLI
// @Author  Ms <133814250@qq.com>
// @License  https://github.com/bufanyun/hotgo/blob/master/LICENSE
// @AutoGenerate Version 2.18.6
package bolldata

import (
	"hotgo/internal/model/input/form"
	"hotgo/internal/model/input/stockin"

	"github.com/gogf/gf/v2/frame/g"
)

// ListReq 查询布林带(BOLL)指标数据表列表
type ListReq struct {
	g.Meta `path:"/bollData/list" method:"get" tags:"布林带(BOLL)指标数据表" summary:"获取布林带(BOLL)指标数据表列表"`
	stockin.BollDataListInp
}

type ListRes struct {
	form.PageRes
	List []*stockin.BollDataListModel `json:"list"   dc:"数据列表"`
}

// ExportReq 导出布林带(BOLL)指标数据表列表
type ExportReq struct {
	g.Meta `path:"/bollData/export" method:"get" tags:"布林带(BOLL)指标数据表" summary:"导出布林带(BOLL)指标数据表列表"`
	stockin.BollDataListInp
}

type ExportRes struct{}

// ViewReq 获取布林带(BOLL)指标数据表指定信息
type ViewReq struct {
	g.Meta `path:"/bollData/view" method:"get" tags:"布林带(BOLL)指标数据表" summary:"获取布林带(BOLL)指标数据表指定信息"`
	stockin.BollDataViewInp
}

type ViewRes struct {
	*stockin.BollDataViewModel
}

// EditReq 修改/新增布林带(BOLL)指标数据表
type EditReq struct {
	g.Meta `path:"/bollData/edit" method:"post" tags:"布林带(BOLL)指标数据表" summary:"修改/新增布林带(BOLL)指标数据表"`
	stockin.BollDataEditInp
}

type EditRes struct{}

// DeleteReq 删除布林带(BOLL)指标数据表
type DeleteReq struct {
	g.Meta `path:"/bollData/delete" method:"post" tags:"布林带(BOLL)指标数据表" summary:"删除布林带(BOLL)指标数据表"`
	stockin.BollDataDeleteInp
}

type DeleteRes struct{}
