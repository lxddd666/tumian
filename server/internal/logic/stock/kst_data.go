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
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gctx"
	"github.com/gogf/gf/v2/util/gconv"
)

type sStockKstData struct{}

func NewStockKstData() *sStockKstData {
	return &sStockKstData{}
}

func init() {
	service.RegisterStockKstData(NewStockKstData())
}

// Model kst指标数据表ORM模型
func (s *sStockKstData) Model(ctx context.Context, option ...*handler.Option) *gdb.Model {
	return handler.Model(dao.KstData.Ctx(ctx), option...)
}

// List 获取kst指标数据表列表
func (s *sStockKstData) List(ctx context.Context, in *stockin.KstDataListInp) (list []*stockin.KstDataListModel, totalCount int, err error) {
	mod := s.Model(ctx)

	// 字段过滤
	mod = mod.Fields(stockin.KstDataListModel{})

	// 查询自增主键
	if in.Id > 0 {
		mod = mod.Where(dao.KstData.Columns().Id, in.Id)
	}

	// 查询数据创建时间
	if len(in.CreatedAt) == 2 {
		mod = mod.WhereBetween(dao.KstData.Columns().CreatedAt, in.CreatedAt[0], in.CreatedAt[1])
	}

	// 分页
	mod = mod.Page(in.Page, in.PerPage)

	// 排序
	mod = mod.OrderDesc(dao.KstData.Columns().Id)

	// 查询数据
	if err = mod.ScanAndCount(&list, &totalCount, false); err != nil {
		err = gerror.Wrap(err, "获取kst指标数据表列表失败，请稍后重试！")
		return
	}
	return
}

// Export 导出kst指标数据表
func (s *sStockKstData) Export(ctx context.Context, in *stockin.KstDataListInp) (err error) {
	list, totalCount, err := s.List(ctx, in)
	if err != nil {
		return
	}

	// 字段的排序是依据tags的字段顺序，如果你不想使用默认的排序方式，可以直接定义 tags = []string{"字段名称", "字段名称2", ...}
	tags, err := convert.GetEntityDescTags(stockin.KstDataExportModel{})
	if err != nil {
		return
	}

	var (
		fileName  = "导出kst指标数据表-" + gctx.CtxId(ctx)
		sheetName = fmt.Sprintf("索引条件共%v行,共%v页,当前导出是第%v页,本页共%v行", totalCount, form.CalPageCount(totalCount, in.PerPage), in.Page, len(list))
		exports   []stockin.KstDataExportModel
	)

	if err = gconv.Scan(list, &exports); err != nil {
		return
	}

	err = excel.ExportByStructs(ctx, tags, exports, fileName, sheetName)
	return
}

// Edit 修改/新增kst指标数据表
func (s *sStockKstData) Edit(ctx context.Context, in *stockin.KstDataEditInp) (err error) {
	return g.DB().Transaction(ctx, func(ctx context.Context, tx gdb.TX) (err error) {

		// 修改
		if in.Id > 0 {
			if _, err = s.Model(ctx).
				Fields(stockin.KstDataUpdateFields{}).
				WherePri(in.Id).Data(in).Update(); err != nil {
				err = gerror.Wrap(err, "修改kst指标数据表失败，请稍后重试！")
			}
			return
		}

		// 新增
		if _, err = s.Model(ctx, &handler.Option{FilterAuth: false}).
			Fields(stockin.KstDataInsertFields{}).
			Data(in).OmitEmptyData().Insert(); err != nil {
			err = gerror.Wrap(err, "新增kst指标数据表失败，请稍后重试！")
		}
		return
	})
}

// Delete 删除kst指标数据表
func (s *sStockKstData) Delete(ctx context.Context, in *stockin.KstDataDeleteInp) (err error) {

	if _, err = s.Model(ctx).WherePri(in.Id).Unscoped().Delete(); err != nil {
		err = gerror.Wrap(err, "删除kst指标数据表失败，请稍后重试！")
		return
	}
	return
}

// View 获取kst指标数据表指定信息
func (s *sStockKstData) View(ctx context.Context, in *stockin.KstDataViewInp) (res *stockin.KstDataViewModel, err error) {
	if err = s.Model(ctx).WherePri(in.Id).Scan(&res); err != nil {
		err = gerror.Wrap(err, "获取kst指标数据表信息，请稍后重试！")
		return
	}
	return
}
