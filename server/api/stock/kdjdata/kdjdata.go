// Package kdjdata
// @Link  https://github.com/bufanyun/hotgo
// @Copyright  Copyright (c) 2026 HotGo CLI
// @Author  Ms <133814250@qq.com>
// @License  https://github.com/bufanyun/hotgo/blob/master/LICENSE
// @AutoGenerate Version 2.18.6
package kdjdata

import (
	"hotgo/internal/model/input/form"
	"hotgo/internal/model/input/stockin"

	"github.com/gogf/gf/v2/frame/g"
)

// ListReq 查询KDJ随机指标数据表列表
type ListReq struct {
	g.Meta `path:"/kdjData/list" method:"get" tags:"KDJ随机指标数据表" summary:"获取KDJ随机指标数据表列表"`
	stockin.KdjDataListInp
}

type ListRes struct {
	form.PageRes
	List []*stockin.KdjDataListModel `json:"list"   dc:"数据列表"`
}

// ExportReq 导出KDJ随机指标数据表列表
type ExportReq struct {
	g.Meta `path:"/kdjData/export" method:"get" tags:"KDJ随机指标数据表" summary:"导出KDJ随机指标数据表列表"`
	stockin.KdjDataListInp
}

type ExportRes struct{}

// ViewReq 获取KDJ随机指标数据表指定信息
type ViewReq struct {
	g.Meta `path:"/kdjData/view" method:"get" tags:"KDJ随机指标数据表" summary:"获取KDJ随机指标数据表指定信息"`
	stockin.KdjDataViewInp
}

type ViewRes struct {
	*stockin.KdjDataViewModel
}

// EditReq 修改/新增KDJ随机指标数据表
type EditReq struct {
	g.Meta `path:"/kdjData/edit" method:"post" tags:"KDJ随机指标数据表" summary:"修改/新增KDJ随机指标数据表"`
	stockin.KdjDataEditInp
}

type EditRes struct{}

// DeleteReq 删除KDJ随机指标数据表
type DeleteReq struct {
	g.Meta `path:"/kdjData/delete" method:"post" tags:"KDJ随机指标数据表" summary:"删除KDJ随机指标数据表"`
	stockin.KdjDataDeleteInp
}

type DeleteRes struct{}

type GetKdjReq struct {
	g.Meta `path:"/kdjData/getKdj" method:"post" tags:"KDJ随机指标数据表" summary:"获取kdj api"`
	stockin.KdjDataEditInp
}
