// Package stock
// @Link  https://github.com/bufanyun/hotgo
// @Copyright  Copyright (c) 2026 HotGo CLI
// @Author  Ms <133814250@qq.com>
// @License  https://github.com/bufanyun/hotgo/blob/master/LICENSE
// @AutoGenerate Version 2.18.6
package stock

import (
	"context"
	"fmt"
	"hotgo/internal/dao"
	"hotgo/internal/library/hgorm/handler"
	"hotgo/internal/model/entity"
	"hotgo/internal/model/input/form"
	"hotgo/internal/model/input/stockin"
	"hotgo/internal/service"
	"hotgo/utility/convert"
	"hotgo/utility/excel"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gctx"
	"github.com/gogf/gf/v2/util/gconv"
)

type sStockTopTenCirculatingHolders struct{}

func NewStockTopTenCirculatingHolders() *sStockTopTenCirculatingHolders {
	return &sStockTopTenCirculatingHolders{}
}

func init() {
	service.RegisterStockTopTenCirculatingHolders(NewStockTopTenCirculatingHolders())
}

// Model 公司十大流通股东表 (数据来源于定期报告)[citation:4]ORM模型
func (s *sStockTopTenCirculatingHolders) Model(ctx context.Context, option ...*handler.Option) *gdb.Model {
	return handler.Model(dao.TopTenCirculatingHolders.Ctx(ctx), option...)
}

// List 获取公司十大流通股东表 (数据来源于定期报告)[citation:4]列表
func (s *sStockTopTenCirculatingHolders) List(ctx context.Context, in *stockin.TopTenCirculatingHoldersListInp) (list []*stockin.TopTenCirculatingHoldersListModel, totalCount int, err error) {
	mod := s.Model(ctx)

	// 字段过滤
	mod = mod.Fields(stockin.TopTenCirculatingHoldersListModel{})

	// 查询自增主键
	if in.Id > 0 {
		mod = mod.Where(dao.TopTenCirculatingHolders.Columns().Id, in.Id)
	}

	// 查询创建时间
	if len(in.CreatedAt) == 2 {
		mod = mod.WhereBetween(dao.TopTenCirculatingHolders.Columns().CreatedAt, in.CreatedAt[0], in.CreatedAt[1])
	}

	// 分页
	mod = mod.Page(in.Page, in.PerPage)

	// 排序
	mod = mod.OrderDesc(dao.TopTenCirculatingHolders.Columns().Id)

	// 查询数据
	if err = mod.ScanAndCount(&list, &totalCount, false); err != nil {
		err = gerror.Wrap(err, "获取公司十大流通股东表 (数据来源于定期报告)[citation:4]列表失败，请稍后重试！")
		return
	}
	return
}

// Export 导出公司十大流通股东表 (数据来源于定期报告)[citation:4]
func (s *sStockTopTenCirculatingHolders) Export(ctx context.Context, in *stockin.TopTenCirculatingHoldersListInp) (err error) {
	list, totalCount, err := s.List(ctx, in)
	if err != nil {
		return
	}

	// 字段的排序是依据tags的字段顺序，如果你不想使用默认的排序方式，可以直接定义 tags = []string{"字段名称", "字段名称2", ...}
	tags, err := convert.GetEntityDescTags(stockin.TopTenCirculatingHoldersExportModel{})
	if err != nil {
		return
	}

	var (
		fileName  = "导出公司十大流通股东表 (数据来源于定期报告)[citation:4]-" + gctx.CtxId(ctx)
		sheetName = fmt.Sprintf("索引条件共%v行,共%v页,当前导出是第%v页,本页共%v行", totalCount, form.CalPageCount(totalCount, in.PerPage), in.Page, len(list))
		exports   []stockin.TopTenCirculatingHoldersExportModel
	)

	if err = gconv.Scan(list, &exports); err != nil {
		return
	}

	err = excel.ExportByStructs(ctx, tags, exports, fileName, sheetName)
	return
}

// Edit 修改/新增公司十大流通股东表 (数据来源于定期报告)[citation:4]
func (s *sStockTopTenCirculatingHolders) Edit(ctx context.Context, in *stockin.TopTenCirculatingHoldersEditInp) (err error) {
	return g.DB().Transaction(ctx, func(ctx context.Context, tx gdb.TX) (err error) {

		// 修改
		if in.Id > 0 {
			if _, err = s.Model(ctx).
				Fields(stockin.TopTenCirculatingHoldersUpdateFields{}).
				WherePri(in.Id).Data(in).Update(); err != nil {
				err = gerror.Wrap(err, "修改公司十大流通股东表 (数据来源于定期报告)[citation:4]失败，请稍后重试！")
			}
			return
		}

		// 新增
		if _, err = s.Model(ctx, &handler.Option{FilterAuth: false}).
			Fields(stockin.TopTenCirculatingHoldersInsertFields{}).
			Data(in).OmitEmptyData().Insert(); err != nil {
			err = gerror.Wrap(err, "新增公司十大流通股东表 (数据来源于定期报告)[citation:4]失败，请稍后重试！")
		}
		return
	})
}

// Delete 删除公司十大流通股东表 (数据来源于定期报告)[citation:4]
func (s *sStockTopTenCirculatingHolders) Delete(ctx context.Context, in *stockin.TopTenCirculatingHoldersDeleteInp) (err error) {

	if _, err = s.Model(ctx).WherePri(in.Id).Unscoped().Delete(); err != nil {
		err = gerror.Wrap(err, "删除公司十大流通股东表 (数据来源于定期报告)[citation:4]失败，请稍后重试！")
		return
	}
	return
}

// View 获取公司十大流通股东表 (数据来源于定期报告)[citation:4]指定信息
func (s *sStockTopTenCirculatingHolders) View(ctx context.Context, in *stockin.TopTenCirculatingHoldersViewInp) (res *stockin.TopTenCirculatingHoldersViewModel, err error) {
	if err = s.Model(ctx).WherePri(in.Id).Scan(&res); err != nil {
		err = gerror.Wrap(err, "获取公司十大流通股东表 (数据来源于定期报告)[citation:4]信息，请稍后重试！")
		return
	}
	return
}

// GetTopTenCirculatingHolders 获取公司十大流通股东表数据
func (s *sStockTopTenCirculatingHolders) GetTopTenCirculatingHolders(ctx context.Context, in *stockin.TopTenCirculatingHoldersGetTopTenCirculatingHoldersInp) (data *entity.TopTenCirculatingHolders, err error) {
	// 构建 API URL
	// https://api.zhituapi.com/hs/gs/sdgd/股票代码?token=token证书
	apiUrl := fmt.Sprintf("https://api.zhituapi.com/hs/gs/sdgd/%s", in.Symbol)

	// 构建查询参数
	params := g.Map{
		"token": in.Token,
	}

	// 发送 GET 请求并解析为 entity.TopTenCirculatingHolders
	var result entity.TopTenCirculatingHolders
	err = g.Client().GetVar(ctx, apiUrl, params).Scan(&result)
	if err != nil {
		err = gerror.Wrap(err, "调用外部API获取公司十大流通股东表数据失败，请稍后重试！")
		return
	}

	data = &result
	return
}
