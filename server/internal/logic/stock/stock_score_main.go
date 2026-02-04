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
	"hotgo/internal/model/input/form"
	"hotgo/internal/model/input/stockin"
	"hotgo/internal/service"
	"hotgo/utility/convert"
	"hotgo/utility/excel"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/os/gctx"
	"github.com/gogf/gf/v2/util/gconv"
)

type sStockScoreMain struct{}

func NewStockScoreMain() *sStockScoreMain {
	return &sStockScoreMain{}
}

func init() {
	service.RegisterStockScoreMain(NewStockScoreMain())
}

// Model 股票评分主表ORM模型
func (s *sStockScoreMain) Model(ctx context.Context, option ...*handler.Option) *gdb.Model {
	return handler.Model(dao.StockScoreMain.Ctx(ctx), option...)
}

// List 获取股票评分主表列表
func (s *sStockScoreMain) List(ctx context.Context, in *stockin.StockScoreMainListInp) (list []*stockin.StockScoreMainListModel, totalCount int, err error) {
	mod := s.Model(ctx)

	// 字段过滤
	mod = mod.Fields(stockin.StockScoreMainListModel{})

	// 查询主键ID
	if in.Id > 0 {
		mod = mod.Where(dao.StockScoreMain.Columns().Id, in.Id)
	}

	// 分页
	mod = mod.Page(in.Page, in.PerPage)

	// 排序
	mod = mod.OrderDesc(dao.StockScoreMain.Columns().Id)

	// 查询数据
	if err = mod.ScanAndCount(&list, &totalCount, false); err != nil {
		err = gerror.Wrap(err, "获取股票评分主表列表失败，请稍后重试！")
		return
	}
	return
}

// Export 导出股票评分主表
func (s *sStockScoreMain) Export(ctx context.Context, in *stockin.StockScoreMainListInp) (err error) {
	list, totalCount, err := s.List(ctx, in)
	if err != nil {
		return
	}

	// 字段的排序是依据tags的字段顺序，如果你不想使用默认的排序方式，可以直接定义 tags = []string{"字段名称", "字段名称2", ...}
	tags, err := convert.GetEntityDescTags(stockin.StockScoreMainExportModel{})
	if err != nil {
		return
	}

	var (
		fileName  = "导出股票评分主表-" + gctx.CtxId(ctx)
		sheetName = fmt.Sprintf("索引条件共%v行,共%v页,当前导出是第%v页,本页共%v行", totalCount, form.CalPageCount(totalCount, in.PerPage), in.Page, len(list))
		exports   []stockin.StockScoreMainExportModel
	)

	if err = gconv.Scan(list, &exports); err != nil {
		return
	}

	err = excel.ExportByStructs(ctx, tags, exports, fileName, sheetName)
	return
}

// Edit 修改/新增股票评分主表
func (s *sStockScoreMain) Edit(ctx context.Context, in *stockin.StockScoreMainEditInp) (err error) {
	// 验证'StockCode'唯一
	return
}

// Delete 删除股票评分主表
func (s *sStockScoreMain) Delete(ctx context.Context, in *stockin.StockScoreMainDeleteInp) (err error) {

	if _, err = s.Model(ctx).WherePri(in.Id).Unscoped().Delete(); err != nil {
		err = gerror.Wrap(err, "删除股票评分主表失败，请稍后重试！")
		return
	}
	return
}

// View 获取股票评分主表指定信息
func (s *sStockScoreMain) View(ctx context.Context, in *stockin.StockScoreMainViewInp) (res *stockin.StockScoreMainViewModel, err error) {
	if err = s.Model(ctx).WherePri(in.Id).Scan(&res); err != nil {
		err = gerror.Wrap(err, "获取股票评分主表信息，请稍后重试！")
		return
	}
	return
}
