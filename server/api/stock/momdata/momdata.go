// Package momdata
// @Link  https://github.com/bufanyun/hotgo
// @Copyright  Copyright (c) 2026 HotGo CLI
// @Author  Ms <133814250@qq.com>
// @License  https://github.com/bufanyun/hotgo/blob/master/LICENSE
// @AutoGenerate Version 2.18.6
package momdata

import (
	"hotgo/internal/model/input/form"
	"hotgo/internal/model/input/stockin"

	"github.com/gogf/gf/v2/frame/g"
)

// ListReq 查询mom指标数据表列表
type ListReq struct {
	g.Meta `path:"/momData/list" method:"get" tags:"mom指标数据表" summary:"获取mom指标数据表列表"`
	stockin.MomDataListInp
}

type ListRes struct {
	form.PageRes
	List []*stockin.MomDataListModel `json:"list"   dc:"数据列表"`
}

// ExportReq 导出mom指标数据表列表
type ExportReq struct {
	g.Meta `path:"/momData/export" method:"get" tags:"mom指标数据表" summary:"导出mom指标数据表列表"`
	stockin.MomDataListInp
}

type ExportRes struct{}

// ViewReq 获取mom指标数据表指定信息
type ViewReq struct {
	g.Meta `path:"/momData/view" method:"get" tags:"mom指标数据表" summary:"获取mom指标数据表指定信息"`
	stockin.MomDataViewInp
}

type ViewRes struct {
	*stockin.MomDataViewModel
}

// EditReq 修改/新增mom指标数据表
type EditReq struct {
	g.Meta `path:"/momData/edit" method:"post" tags:"mom指标数据表" summary:"修改/新增mom指标数据表"`
	stockin.MomDataEditInp
}

type EditRes struct{}

// DeleteReq 删除mom指标数据表
type DeleteReq struct {
	g.Meta `path:"/momData/delete" method:"post" tags:"mom指标数据表" summary:"删除mom指标数据表"`
	stockin.MomDataDeleteInp
}

type DeleteRes struct{}
