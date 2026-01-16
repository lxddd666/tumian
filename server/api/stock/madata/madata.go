// Package madata
// @Link  https://github.com/bufanyun/hotgo
// @Copyright  Copyright (c) 2026 HotGo CLI
// @Author  Ms <133814250@qq.com>
// @License  https://github.com/bufanyun/hotgo/blob/master/LICENSE
// @AutoGenerate Version 2.18.6
package madata

import (
	"hotgo/internal/model/entity"
	"hotgo/internal/model/input/form"
	"hotgo/internal/model/input/stockin"

	"github.com/gogf/gf/v2/frame/g"
)

// ListReq 查询移动平均线(MA)指标数据表列表
type ListReq struct {
	g.Meta `path:"/maData/list" method:"get" tags:"移动平均线(MA)指标数据表" summary:"获取移动平均线(MA)指标数据表列表"`
	stockin.MaDataListInp
}

type ListRes struct {
	form.PageRes
	List []*stockin.MaDataListModel `json:"list"   dc:"数据列表"`
}

// ExportReq 导出移动平均线(MA)指标数据表列表
type ExportReq struct {
	g.Meta `path:"/maData/export" method:"get" tags:"移动平均线(MA)指标数据表" summary:"导出移动平均线(MA)指标数据表列表"`
	stockin.MaDataListInp
}

type ExportRes struct{}

// ViewReq 获取移动平均线(MA)指标数据表指定信息
type ViewReq struct {
	g.Meta `path:"/maData/view" method:"get" tags:"移动平均线(MA)指标数据表" summary:"获取移动平均线(MA)指标数据表指定信息"`
	stockin.MaDataViewInp
}

type ViewRes struct {
	*stockin.MaDataViewModel
}

// EditReq 修改/新增移动平均线(MA)指标数据表
type EditReq struct {
	g.Meta `path:"/maData/edit" method:"post" tags:"移动平均线(MA)指标数据表" summary:"修改/新增移动平均线(MA)指标数据表"`
	stockin.MaDataEditInp
}

type EditRes struct{}

// DeleteReq 删除移动平均线(MA)指标数据表
type DeleteReq struct {
	g.Meta `path:"/maData/delete" method:"post" tags:"移动平均线(MA)指标数据表" summary:"删除移动平均线(MA)指标数据表"`
	stockin.MaDataDeleteInp
}

type DeleteRes struct{}

// GetMaReq 获取移动平均线(MA)指标数据
type GetMaReq struct {
	g.Meta `path:"/maData/getMa" method:"get" tags:"移动平均线(MA)指标数据表" summary:"获取移动平均线(MA)指标数据"`
	stockin.MaDataGetMaInp
}

type GetMaRes struct {
	Data []*entity.MaData `json:"data" dc:"返回数据"`
}
