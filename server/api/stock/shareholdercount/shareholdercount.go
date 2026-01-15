// Package shareholdercount
// @Link  https://github.com/bufanyun/hotgo
// @Copyright  Copyright (c) 2026 HotGo CLI
// @Author  Ms <133814250@qq.com>
// @License  https://github.com/bufanyun/hotgo/blob/master/LICENSE
// @AutoGenerate Version 2.18.6
package shareholdercount

import (
	"hotgo/internal/model/entity"
	"hotgo/internal/model/input/form"
	"hotgo/internal/model/input/stockin"

	"github.com/gogf/gf/v2/frame/g"
)

// ListReq 查询公司股东户数统计表 (按报告期统计)列表
type ListReq struct {
	g.Meta `path:"/shareholderCount/list" method:"get" tags:"公司股东户数统计表 (按报告期统计)" summary:"获取公司股东户数统计表 (按报告期统计)列表"`
	stockin.ShareholderCountListInp
}

type ListRes struct {
	form.PageRes
	List []*stockin.ShareholderCountListModel `json:"list"   dc:"数据列表"`
}

// ExportReq 导出公司股东户数统计表 (按报告期统计)列表
type ExportReq struct {
	g.Meta `path:"/shareholderCount/export" method:"get" tags:"公司股东户数统计表 (按报告期统计)" summary:"导出公司股东户数统计表 (按报告期统计)列表"`
	stockin.ShareholderCountListInp
}

type ExportRes struct{}

// ViewReq 获取公司股东户数统计表 (按报告期统计)指定信息
type ViewReq struct {
	g.Meta `path:"/shareholderCount/view" method:"get" tags:"公司股东户数统计表 (按报告期统计)" summary:"获取公司股东户数统计表 (按报告期统计)指定信息"`
	stockin.ShareholderCountViewInp
}

type ViewRes struct {
	*stockin.ShareholderCountViewModel
}

// EditReq 修改/新增公司股东户数统计表 (按报告期统计)
type EditReq struct {
	g.Meta `path:"/shareholderCount/edit" method:"post" tags:"公司股东户数统计表 (按报告期统计)" summary:"修改/新增公司股东户数统计表 (按报告期统计)"`
	stockin.ShareholderCountEditInp
}

type EditRes struct{}

// DeleteReq 删除公司股东户数统计表 (按报告期统计)
type DeleteReq struct {
	g.Meta `path:"/shareholderCount/delete" method:"post" tags:"公司股东户数统计表 (按报告期统计)" summary:"删除公司股东户数统计表 (按报告期统计)"`
	stockin.ShareholderCountDeleteInp
}

type DeleteRes struct{}

// GetShareholderCountReq 获取公司股东户数统计表数据
type GetShareholderCountReq struct {
	g.Meta `path:"/shareholderCount/getShareholderCount" method:"get" tags:"公司股东户数统计表 (按报告期统计)" summary:"获取公司股东户数统计表数据"`
	stockin.ShareholderCountGetShareholderCountInp
}

type GetShareholderCountRes struct {
	Data *entity.ShareholderCount `json:"data" dc:"返回数据"`
}
