// Package stockbasicinfo
// @Link  https://github.com/bufanyun/hotgo
// @Copyright  Copyright (c) 2026 HotGo CLI
// @Author  Ms <133814250@qq.com>
// @License  https://github.com/bufanyun/hotgo/blob/master/LICENSE
// @AutoGenerate Version 2.18.6
package stockbasicinfo

import (
	"hotgo/internal/model/input/form"
	"hotgo/internal/model/input/stockin"

	"github.com/gogf/gf/v2/frame/g"
)

// ListReq 查询股票基础信息表列表
type ListReq struct {
	g.Meta `path:"/stockBasicInfo/list" method:"get" tags:"股票基础信息表" summary:"获取股票基础信息表列表"`
	stockin.StockBasicInfoListInp
}

type ListRes struct {
	form.PageRes
	List []*stockin.StockBasicInfoListModel `json:"list"   dc:"数据列表"`
}

// ExportReq 导出股票基础信息表列表
type ExportReq struct {
	g.Meta `path:"/stockBasicInfo/export" method:"get" tags:"股票基础信息表" summary:"导出股票基础信息表列表"`
	stockin.StockBasicInfoListInp
}

type ExportRes struct{}

// ViewReq 获取股票基础信息表指定信息
type ViewReq struct {
	g.Meta `path:"/stockBasicInfo/view" method:"get" tags:"股票基础信息表" summary:"获取股票基础信息表指定信息"`
	stockin.StockBasicInfoViewInp
}

type ViewRes struct {
	*stockin.StockBasicInfoViewModel
}

// EditReq 修改/新增股票基础信息表
type EditReq struct {
	g.Meta `path:"/stockBasicInfo/edit" method:"post" tags:"股票基础信息表" summary:"修改/新增股票基础信息表"`
	stockin.StockBasicInfoEditInp
}

type EditRes struct{}

// DeleteReq 删除股票基础信息表
type DeleteReq struct {
	g.Meta `path:"/stockBasicInfo/delete" method:"post" tags:"股票基础信息表" summary:"删除股票基础信息表"`
	stockin.StockBasicInfoDeleteInp
}

type DeleteRes struct{}

// GetStockBasicInfoReq 获取股票基础信息表数据
type GetStockBasicInfoReq struct {
	g.Meta `path:"/stockBasicInfo/getStockBasicInfo" method:"get" tags:"股票基础信息表" summary:"获取股票基础信息表数据"`
	stockin.StockBasicInfoGetStockBasicInfoInp
}

type GetStockBasicInfoRes struct {
	Data interface{} `json:"data" dc:"返回数据"`
}
