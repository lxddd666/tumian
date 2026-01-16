// Package shareholderchange
// @Link  https://github.com/bufanyun/hotgo
// @Copyright  Copyright (c) 2026 HotGo CLI
// @Author  Ms <133814250@qq.com>
// @License  https://github.com/bufanyun/hotgo/blob/master/LICENSE
// @AutoGenerate Version 2.18.6
package shareholderchange

import (
	"hotgo/internal/model/entity"
	"hotgo/internal/model/input/form"
	"hotgo/internal/model/input/stockin"

	"github.com/gogf/gf/v2/frame/g"
)

// ListReq 查询股东户数变化记录表 (记录相邻报告期的户数变化)列表
type ListReq struct {
	g.Meta `path:"/shareholderChange/list" method:"get" tags:"股东户数变化记录表 (记录相邻报告期的户数变化)" summary:"获取股东户数变化记录表 (记录相邻报告期的户数变化)列表"`
	stockin.ShareholderChangeListInp
}

type ListRes struct {
	form.PageRes
	List []*stockin.ShareholderChangeListModel `json:"list"   dc:"数据列表"`
}

// ExportReq 导出股东户数变化记录表 (记录相邻报告期的户数变化)列表
type ExportReq struct {
	g.Meta `path:"/shareholderChange/export" method:"get" tags:"股东户数变化记录表 (记录相邻报告期的户数变化)" summary:"导出股东户数变化记录表 (记录相邻报告期的户数变化)列表"`
	stockin.ShareholderChangeListInp
}

type ExportRes struct{}

// ViewReq 获取股东户数变化记录表 (记录相邻报告期的户数变化)指定信息
type ViewReq struct {
	g.Meta `path:"/shareholderChange/view" method:"get" tags:"股东户数变化记录表 (记录相邻报告期的户数变化)" summary:"获取股东户数变化记录表 (记录相邻报告期的户数变化)指定信息"`
	stockin.ShareholderChangeViewInp
}

type ViewRes struct {
	*stockin.ShareholderChangeViewModel
}

// EditReq 修改/新增股东户数变化记录表 (记录相邻报告期的户数变化)
type EditReq struct {
	g.Meta `path:"/shareholderChange/edit" method:"post" tags:"股东户数变化记录表 (记录相邻报告期的户数变化)" summary:"修改/新增股东户数变化记录表 (记录相邻报告期的户数变化)"`
	stockin.ShareholderChangeEditInp
}

type EditRes struct{}

// DeleteReq 删除股东户数变化记录表 (记录相邻报告期的户数变化)
type DeleteReq struct {
	g.Meta `path:"/shareholderChange/delete" method:"post" tags:"股东户数变化记录表 (记录相邻报告期的户数变化)" summary:"删除股东户数变化记录表 (记录相邻报告期的户数变化)"`
	stockin.ShareholderChangeDeleteInp
}

type DeleteRes struct{}

// GetShareholderChangeReq 获取股东户数变化记录表数据
type GetShareholderChangeReq struct {
	g.Meta `path:"/shareholderChange/getShareholderChange" method:"get" tags:"股东户数变化记录表 (记录相邻报告期的户数变化)" summary:"获取股东户数变化记录表数据"`
	stockin.ShareholderChangeGetShareholderChangeInp
}

type GetShareholderChangeRes struct {
	Data []*entity.ShareholderChange `json:"data" dc:"返回数据"`
}
