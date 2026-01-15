// Package incomestatement
// @Link  https://github.com/bufanyun/hotgo
// @Copyright  Copyright (c) 2026 HotGo CLI
// @Author  Ms <133814250@qq.com>
// @License  https://github.com/bufanyun/hotgo/blob/master/LICENSE
// @AutoGenerate Version 2.18.6
package incomestatement

import (
	"hotgo/internal/model/entity"
	"hotgo/internal/model/input/form"
	"hotgo/internal/model/input/stockin"

	"github.com/gogf/gf/v2/frame/g"
)

// ListReq 查询利润表 (Income Statement)列表
type ListReq struct {
	g.Meta `path:"/incomeStatement/list" method:"get" tags:"利润表 (Income Statement)" summary:"获取利润表 (Income Statement)列表"`
	stockin.IncomeStatementListInp
}

type ListRes struct {
	form.PageRes
	List []*stockin.IncomeStatementListModel `json:"list"   dc:"数据列表"`
}

// ExportReq 导出利润表 (Income Statement)列表
type ExportReq struct {
	g.Meta `path:"/incomeStatement/export" method:"get" tags:"利润表 (Income Statement)" summary:"导出利润表 (Income Statement)列表"`
	stockin.IncomeStatementListInp
}

type ExportRes struct{}

// ViewReq 获取利润表 (Income Statement)指定信息
type ViewReq struct {
	g.Meta `path:"/incomeStatement/view" method:"get" tags:"利润表 (Income Statement)" summary:"获取利润表 (Income Statement)指定信息"`
	stockin.IncomeStatementViewInp
}

type ViewRes struct {
	*stockin.IncomeStatementViewModel
}

// EditReq 修改/新增利润表 (Income Statement)
type EditReq struct {
	g.Meta `path:"/incomeStatement/edit" method:"post" tags:"利润表 (Income Statement)" summary:"修改/新增利润表 (Income Statement)"`
	stockin.IncomeStatementEditInp
}

type EditRes struct{}

// DeleteReq 删除利润表 (Income Statement)
type DeleteReq struct {
	g.Meta `path:"/incomeStatement/delete" method:"post" tags:"利润表 (Income Statement)" summary:"删除利润表 (Income Statement)"`
	stockin.IncomeStatementDeleteInp
}

type DeleteRes struct{}

// GetIncomeStatementReq 获取利润表数据
type GetIncomeStatementReq struct {
	g.Meta `path:"/incomeStatement/getIncomeStatement" method:"get" tags:"利润表 (Income Statement)" summary:"获取利润表数据"`
	stockin.IncomeStatementGetIncomeStatementInp
}

type GetIncomeStatementRes struct {
	Data *entity.IncomeStatement `json:"data" dc:"返回数据"`
}
