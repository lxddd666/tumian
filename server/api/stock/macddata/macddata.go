// Package macddata
// @Link  https://github.com/bufanyun/hotgo
// @Copyright  Copyright (c) 2026 HotGo CLI
// @Author  Ms <133814250@qq.com>
// @License  https://github.com/bufanyun/hotgo/blob/master/LICENSE
// @AutoGenerate Version 2.18.6
package macddata

import (
	"hotgo/internal/model/entity"
	"hotgo/internal/model/input/form"
	"hotgo/internal/model/input/stockin"

	"github.com/gogf/gf/v2/frame/g"
)

// ListReq 查询MACD指标数据表列表
type ListReq struct {
	g.Meta `path:"/macdData/list" method:"get" tags:"MACD指标数据表" summary:"获取MACD指标数据表列表"`
	stockin.MacdDataListInp
}

type ListRes struct {
	form.PageRes
	List []*stockin.MacdDataListModel `json:"list"   dc:"数据列表"`
}

// ExportReq 导出MACD指标数据表列表
type ExportReq struct {
	g.Meta `path:"/macdData/export" method:"get" tags:"MACD指标数据表" summary:"导出MACD指标数据表列表"`
	stockin.MacdDataListInp
}

type ExportRes struct{}

// ViewReq 获取MACD指标数据表指定信息
type ViewReq struct {
	g.Meta `path:"/macdData/view" method:"get" tags:"MACD指标数据表" summary:"获取MACD指标数据表指定信息"`
	stockin.MacdDataViewInp
}

type ViewRes struct {
	*stockin.MacdDataViewModel
}

// EditReq 修改/新增MACD指标数据表
type EditReq struct {
	g.Meta `path:"/macdData/edit" method:"post" tags:"MACD指标数据表" summary:"修改/新增MACD指标数据表"`
	stockin.MacdDataEditInp
}

type EditRes struct{}

// DeleteReq 删除MACD指标数据表
type DeleteReq struct {
	g.Meta `path:"/macdData/delete" method:"post" tags:"MACD指标数据表" summary:"删除MACD指标数据表"`
	stockin.MacdDataDeleteInp
}

type DeleteRes struct{}

// GetMacdReq 获取MACD指标数据
type GetMacdReq struct {
	g.Meta `path:"/macdData/getMacd" method:"get" tags:"MACD指标数据表" summary:"获取MACD指标数据"`
	stockin.MacdDataGetMacdInp
}

type GetMacdRes struct {
	Data []*entity.MacdData `json:"data" dc:"返回数据"`
}
