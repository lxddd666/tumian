// Package slowstochasticdata
// @Link  https://github.com/bufanyun/hotgo
// @Copyright  Copyright (c) 2026 HotGo CLI
// @Author  Ms <133814250@qq.com>
// @License  https://github.com/bufanyun/hotgo/blob/master/LICENSE
// @AutoGenerate Version 2.18.6
package slowstochasticdata

import (
	"hotgo/internal/model/input/form"
	"hotgo/internal/model/input/stockin"

	"github.com/gogf/gf/v2/frame/g"
)

// ListReq 查询stoch指标数据表列表
type ListReq struct {
	g.Meta `path:"/slowStochasticData/list" method:"get" tags:"stoch指标数据表" summary:"获取stoch指标数据表列表"`
	stockin.SlowStochasticDataListInp
}

type ListRes struct {
	form.PageRes
	List []*stockin.SlowStochasticDataListModel `json:"list"   dc:"数据列表"`
}

// ExportReq 导出stoch指标数据表列表
type ExportReq struct {
	g.Meta `path:"/slowStochasticData/export" method:"get" tags:"stoch指标数据表" summary:"导出stoch指标数据表列表"`
	stockin.SlowStochasticDataListInp
}

type ExportRes struct{}

// ViewReq 获取stoch指标数据表指定信息
type ViewReq struct {
	g.Meta `path:"/slowStochasticData/view" method:"get" tags:"stoch指标数据表" summary:"获取stoch指标数据表指定信息"`
	stockin.SlowStochasticDataViewInp
}

type ViewRes struct {
	*stockin.SlowStochasticDataViewModel
}

// EditReq 修改/新增stoch指标数据表
type EditReq struct {
	g.Meta `path:"/slowStochasticData/edit" method:"post" tags:"stoch指标数据表" summary:"修改/新增stoch指标数据表"`
	stockin.SlowStochasticDataEditInp
}

type EditRes struct{}

// DeleteReq 删除stoch指标数据表
type DeleteReq struct {
	g.Meta `path:"/slowStochasticData/delete" method:"post" tags:"stoch指标数据表" summary:"删除stoch指标数据表"`
	stockin.SlowStochasticDataDeleteInp
}

type DeleteRes struct{}
