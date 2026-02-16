// Package fastkdata
// @Link  https://github.com/bufanyun/hotgo
// @Copyright  Copyright (c) 2026 HotGo CLI
// @Author  Ms <133814250@qq.com>
// @License  https://github.com/bufanyun/hotgo/blob/master/LICENSE
// @AutoGenerate Version 2.18.6
package fastkdata

import (
	"hotgo/internal/model/input/form"
	"hotgo/internal/model/input/stockin"

	"github.com/gogf/gf/v2/frame/g"
)

// ListReq 查询fastk指标数据表列表
type ListReq struct {
	g.Meta `path:"/fastkData/list" method:"get" tags:"fastk指标数据表" summary:"获取fastk指标数据表列表"`
	stockin.FastkDataListInp
}

type ListRes struct {
	form.PageRes
	List []*stockin.FastkDataListModel `json:"list"   dc:"数据列表"`
}

// ExportReq 导出fastk指标数据表列表
type ExportReq struct {
	g.Meta `path:"/fastkData/export" method:"get" tags:"fastk指标数据表" summary:"导出fastk指标数据表列表"`
	stockin.FastkDataListInp
}

type ExportRes struct{}

// ViewReq 获取fastk指标数据表指定信息
type ViewReq struct {
	g.Meta `path:"/fastkData/view" method:"get" tags:"fastk指标数据表" summary:"获取fastk指标数据表指定信息"`
	stockin.FastkDataViewInp
}

type ViewRes struct {
	*stockin.FastkDataViewModel
}

// EditReq 修改/新增fastk指标数据表
type EditReq struct {
	g.Meta `path:"/fastkData/edit" method:"post" tags:"fastk指标数据表" summary:"修改/新增fastk指标数据表"`
	stockin.FastkDataEditInp
}

type EditRes struct{}

// DeleteReq 删除fastk指标数据表
type DeleteReq struct {
	g.Meta `path:"/fastkData/delete" method:"post" tags:"fastk指标数据表" summary:"删除fastk指标数据表"`
	stockin.FastkDataDeleteInp
}

type DeleteRes struct{}
