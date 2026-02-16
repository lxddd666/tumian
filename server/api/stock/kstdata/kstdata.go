// Package kstdata
// @Link  https://github.com/bufanyun/hotgo
// @Copyright  Copyright (c) 2026 HotGo CLI
// @Author  Ms <133814250@qq.com>
// @License  https://github.com/bufanyun/hotgo/blob/master/LICENSE
// @AutoGenerate Version 2.18.6
package kstdata

import (
	"hotgo/internal/model/input/form"
	"hotgo/internal/model/input/stockin"

	"github.com/gogf/gf/v2/frame/g"
)

// ListReq 查询kst指标数据表列表
type ListReq struct {
	g.Meta `path:"/kstData/list" method:"get" tags:"kst指标数据表" summary:"获取kst指标数据表列表"`
	stockin.KstDataListInp
}

type ListRes struct {
	form.PageRes
	List []*stockin.KstDataListModel `json:"list"   dc:"数据列表"`
}

// ExportReq 导出kst指标数据表列表
type ExportReq struct {
	g.Meta `path:"/kstData/export" method:"get" tags:"kst指标数据表" summary:"导出kst指标数据表列表"`
	stockin.KstDataListInp
}

type ExportRes struct{}

// ViewReq 获取kst指标数据表指定信息
type ViewReq struct {
	g.Meta `path:"/kstData/view" method:"get" tags:"kst指标数据表" summary:"获取kst指标数据表指定信息"`
	stockin.KstDataViewInp
}

type ViewRes struct {
	*stockin.KstDataViewModel
}

// EditReq 修改/新增kst指标数据表
type EditReq struct {
	g.Meta `path:"/kstData/edit" method:"post" tags:"kst指标数据表" summary:"修改/新增kst指标数据表"`
	stockin.KstDataEditInp
}

type EditRes struct{}

// DeleteReq 删除kst指标数据表
type DeleteReq struct {
	g.Meta `path:"/kstData/delete" method:"post" tags:"kst指标数据表" summary:"删除kst指标数据表"`
	stockin.KstDataDeleteInp
}

type DeleteRes struct{}
