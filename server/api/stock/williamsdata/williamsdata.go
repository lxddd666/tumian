// Package williamsdata
// @Link  https://github.com/bufanyun/hotgo
// @Copyright  Copyright (c) 2026 HotGo CLI
// @Author  Ms <133814250@qq.com>
// @License  https://github.com/bufanyun/hotgo/blob/master/LICENSE
// @AutoGenerate Version 2.18.6
package williamsdata

import (
	"hotgo/internal/model/input/form"
	"hotgo/internal/model/input/stockin"

	"github.com/gogf/gf/v2/frame/g"
)

// ListReq 查询wmsr指标数据表列表
type ListReq struct {
	g.Meta `path:"/williamsData/list" method:"get" tags:"wmsr指标数据表" summary:"获取wmsr指标数据表列表"`
	stockin.WilliamsDataListInp
}

type ListRes struct {
	form.PageRes
	List []*stockin.WilliamsDataListModel `json:"list"   dc:"数据列表"`
}

// ExportReq 导出wmsr指标数据表列表
type ExportReq struct {
	g.Meta `path:"/williamsData/export" method:"get" tags:"wmsr指标数据表" summary:"导出wmsr指标数据表列表"`
	stockin.WilliamsDataListInp
}

type ExportRes struct{}

// ViewReq 获取wmsr指标数据表指定信息
type ViewReq struct {
	g.Meta `path:"/williamsData/view" method:"get" tags:"wmsr指标数据表" summary:"获取wmsr指标数据表指定信息"`
	stockin.WilliamsDataViewInp
}

type ViewRes struct {
	*stockin.WilliamsDataViewModel
}

// EditReq 修改/新增wmsr指标数据表
type EditReq struct {
	g.Meta `path:"/williamsData/edit" method:"post" tags:"wmsr指标数据表" summary:"修改/新增wmsr指标数据表"`
	stockin.WilliamsDataEditInp
}

type EditRes struct{}

// DeleteReq 删除wmsr指标数据表
type DeleteReq struct {
	g.Meta `path:"/williamsData/delete" method:"post" tags:"wmsr指标数据表" summary:"删除wmsr指标数据表"`
	stockin.WilliamsDataDeleteInp
}

type DeleteRes struct{}

// GetWilliamsReq 获取Williams指标数据表
type GetWilliamsReq struct {
	g.Meta `path:"/williamsData/getWilliams" method:"get" tags:"wmsr指标数据表" summary:"获取Williams指标数据表(百度财经)"`
	stockin.GetWilliamsDataInp
}

type GetWilliamsRes struct{}
