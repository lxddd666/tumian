// Package toptencirculatingholders
// @Link  https://github.com/bufanyun/hotgo
// @Copyright  Copyright (c) 2026 HotGo CLI
// @Author  Ms <133814250@qq.com>
// @License  https://github.com/bufanyun/hotgo/blob/master/LICENSE
// @AutoGenerate Version 2.18.6
package toptencirculatingholders

import (
	"hotgo/internal/model/entity"
	"hotgo/internal/model/input/form"
	"hotgo/internal/model/input/stockin"

	"github.com/gogf/gf/v2/frame/g"
)

// ListReq 查询公司十大流通股东表 (数据来源于定期报告)[citation:4]列表
type ListReq struct {
	g.Meta `path:"/topTenCirculatingHolders/list" method:"get" tags:"公司十大流通股东表 (数据来源于定期报告)[citation:4]" summary:"获取公司十大流通股东表 (数据来源于定期报告)[citation:4]列表"`
	stockin.TopTenCirculatingHoldersListInp
}

type ListRes struct {
	form.PageRes
	List []*stockin.TopTenCirculatingHoldersListModel `json:"list"   dc:"数据列表"`
}

// ExportReq 导出公司十大流通股东表 (数据来源于定期报告)[citation:4]列表
type ExportReq struct {
	g.Meta `path:"/topTenCirculatingHolders/export" method:"get" tags:"公司十大流通股东表 (数据来源于定期报告)[citation:4]" summary:"导出公司十大流通股东表 (数据来源于定期报告)[citation:4]列表"`
	stockin.TopTenCirculatingHoldersListInp
}

type ExportRes struct{}

// ViewReq 获取公司十大流通股东表 (数据来源于定期报告)[citation:4]指定信息
type ViewReq struct {
	g.Meta `path:"/topTenCirculatingHolders/view" method:"get" tags:"公司十大流通股东表 (数据来源于定期报告)[citation:4]" summary:"获取公司十大流通股东表 (数据来源于定期报告)[citation:4]指定信息"`
	stockin.TopTenCirculatingHoldersViewInp
}

type ViewRes struct {
	*stockin.TopTenCirculatingHoldersViewModel
}

// EditReq 修改/新增公司十大流通股东表 (数据来源于定期报告)[citation:4]
type EditReq struct {
	g.Meta `path:"/topTenCirculatingHolders/edit" method:"post" tags:"公司十大流通股东表 (数据来源于定期报告)[citation:4]" summary:"修改/新增公司十大流通股东表 (数据来源于定期报告)[citation:4]"`
	stockin.TopTenCirculatingHoldersEditInp
}

type EditRes struct{}

// DeleteReq 删除公司十大流通股东表 (数据来源于定期报告)[citation:4]
type DeleteReq struct {
	g.Meta `path:"/topTenCirculatingHolders/delete" method:"post" tags:"公司十大流通股东表 (数据来源于定期报告)[citation:4]" summary:"删除公司十大流通股东表 (数据来源于定期报告)[citation:4]"`
	stockin.TopTenCirculatingHoldersDeleteInp
}

type DeleteRes struct{}

// GetTopTenCirculatingHoldersReq 获取公司十大流通股东表数据
type GetTopTenCirculatingHoldersReq struct {
	g.Meta `path:"/topTenCirculatingHolders/getTopTenCirculatingHolders" method:"get" tags:"公司十大流通股东表 (数据来源于定期报告)[citation:4]" summary:"获取公司十大流通股东表数据"`
	stockin.TopTenCirculatingHoldersGetTopTenCirculatingHoldersInp
}

type GetTopTenCirculatingHoldersRes struct {
	Data *entity.TopTenCirculatingHolders `json:"data" dc:"返回数据"`
}
