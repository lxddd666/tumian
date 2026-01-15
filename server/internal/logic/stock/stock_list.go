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
	"hotgo/internal/library/hgorm"
	"hotgo/internal/library/hgorm/handler"
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

type sStockList struct{}

func NewStockList() *sStockList {
	return &sStockList{}
}

func init() {
	service.RegisterStockList(NewStockList())
}

// Model 股票列表核心表ORM模型
func (s *sStockList) Model(ctx context.Context, option ...*handler.Option) *gdb.Model {
	return handler.Model(dao.StockList.Ctx(ctx), option...)
}

// List 获取股票列表核心表列表
func (s *sStockList) List(ctx context.Context, in *stockin.StockListListInp) (list []*stockin.StockListListModel, totalCount int, err error) {
	mod := s.Model(ctx)

	// 字段过滤
	mod = mod.Fields(stockin.StockListListModel{})

	// 查询自增主键
	if in.Id > 0 {
		mod = mod.Where(dao.StockList.Columns().Id, in.Id)
	}

	// 查询状态: 1-正常, 0-退市
	if in.Status > 0 {
		mod = mod.Where(dao.StockList.Columns().Status, in.Status)
	}

	// 查询创建时间
	if len(in.CreatedAt) == 2 {
		mod = mod.WhereBetween(dao.StockList.Columns().CreatedAt, in.CreatedAt[0], in.CreatedAt[1])
	}

	// 分页
	mod = mod.Page(in.Page, in.PerPage)

	// 排序
	mod = mod.OrderDesc(dao.StockList.Columns().Id)

	// 查询数据
	if err = mod.ScanAndCount(&list, &totalCount, false); err != nil {
		err = gerror.Wrap(err, "获取股票列表核心表列表失败，请稍后重试！")
		return
	}
	return
}

// Export 导出股票列表核心表
func (s *sStockList) Export(ctx context.Context, in *stockin.StockListListInp) (err error) {
	list, totalCount, err := s.List(ctx, in)
	if err != nil {
		return
	}

	// 字段的排序是依据tags的字段顺序，如果你不想使用默认的排序方式，可以直接定义 tags = []string{"字段名称", "字段名称2", ...}
	tags, err := convert.GetEntityDescTags(stockin.StockListExportModel{})
	if err != nil {
		return
	}

	var (
		fileName  = "导出股票列表核心表-" + gctx.CtxId(ctx)
		sheetName = fmt.Sprintf("索引条件共%v行,共%v页,当前导出是第%v页,本页共%v行", totalCount, form.CalPageCount(totalCount, in.PerPage), in.Page, len(list))
		exports   []stockin.StockListExportModel
	)

	if err = gconv.Scan(list, &exports); err != nil {
		return
	}

	err = excel.ExportByStructs(ctx, tags, exports, fileName, sheetName)
	return
}

// Edit 修改/新增股票列表核心表
func (s *sStockList) Edit(ctx context.Context, in *stockin.StockListEditInp) (err error) {
	// 验证'Dm'唯一
	if err = hgorm.IsUnique(ctx, &dao.StockList, g.Map{dao.StockList.Columns().Dm: in.Dm}, "股票代码 (唯一业务标识，如: 000001)已存在", in.Id); err != nil {
		return
	}
	// 验证'Symbol'唯一
	if err = hgorm.IsUnique(ctx, &dao.StockList, g.Map{dao.StockList.Columns().Symbol: in.Symbol}, "标准股票代码 (如: 000001.SZ)已存在", in.Id); err != nil {
		return
	}
	return g.DB().Transaction(ctx, func(ctx context.Context, tx gdb.TX) (err error) {

		// 修改
		if in.Id > 0 {
			if _, err = s.Model(ctx).
				Fields(stockin.StockListUpdateFields{}).
				WherePri(in.Id).Data(in).Update(); err != nil {
				err = gerror.Wrap(err, "修改股票列表核心表失败，请稍后重试！")
			}
			return
		}

		// 新增
		if _, err = s.Model(ctx, &handler.Option{FilterAuth: false}).
			Fields(stockin.StockListInsertFields{}).
			Data(in).OmitEmptyData().Insert(); err != nil {
			err = gerror.Wrap(err, "新增股票列表核心表失败，请稍后重试！")
		}
		return
	})
}

// Delete 删除股票列表核心表
func (s *sStockList) Delete(ctx context.Context, in *stockin.StockListDeleteInp) (err error) {

	if _, err = s.Model(ctx).WherePri(in.Id).Unscoped().Delete(); err != nil {
		err = gerror.Wrap(err, "删除股票列表核心表失败，请稍后重试！")
		return
	}
	return
}

// View 获取股票列表核心表指定信息
func (s *sStockList) View(ctx context.Context, in *stockin.StockListViewInp) (res *stockin.StockListViewModel, err error) {
	if err = s.Model(ctx).WherePri(in.Id).Scan(&res); err != nil {
		err = gerror.Wrap(err, "获取股票列表核心表信息，请稍后重试！")
		return
	}
	return
}

// Status 更新股票列表核心表状态
func (s *sStockList) Status(ctx context.Context, in *stockin.StockListStatusInp) (err error) {
	if _, err = s.Model(ctx).WherePri(in.Id).Data(g.Map{
		dao.StockList.Columns().Status: in.Status,
	}).Update(); err != nil {
		err = gerror.Wrap(err, "更新股票列表核心表状态失败，请稍后重试！")
		return
	}
	return
}
