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

type sStockFinancialIndicators struct{}

func NewStockFinancialIndicators() *sStockFinancialIndicators {
	return &sStockFinancialIndicators{}
}

func init() {
	service.RegisterStockFinancialIndicators(NewStockFinancialIndicators())
}

// Model 财务指标分析表ORM模型
func (s *sStockFinancialIndicators) Model(ctx context.Context, option ...*handler.Option) *gdb.Model {
	return handler.Model(dao.FinancialIndicators.Ctx(ctx), option...)
}

// List 获取财务指标分析表列表
func (s *sStockFinancialIndicators) List(ctx context.Context, in *stockin.FinancialIndicatorsListInp) (list []*stockin.FinancialIndicatorsListModel, totalCount int, err error) {
	mod := s.Model(ctx)

	// 字段过滤
	mod = mod.Fields(stockin.FinancialIndicatorsListModel{})

	// 查询自增主键
	if in.Id > 0 {
		mod = mod.Where(dao.FinancialIndicators.Columns().Id, in.Id)
	}

	// 查询创建时间
	if len(in.CreatedAt) == 2 {
		mod = mod.WhereBetween(dao.FinancialIndicators.Columns().CreatedAt, in.CreatedAt[0], in.CreatedAt[1])
	}

	// 分页
	mod = mod.Page(in.Page, in.PerPage)

	// 排序
	mod = mod.OrderDesc(dao.FinancialIndicators.Columns().Id)

	// 查询数据
	if err = mod.ScanAndCount(&list, &totalCount, false); err != nil {
		err = gerror.Wrap(err, "获取财务指标分析表列表失败，请稍后重试！")
		return
	}
	return
}

// Export 导出财务指标分析表
func (s *sStockFinancialIndicators) Export(ctx context.Context, in *stockin.FinancialIndicatorsListInp) (err error) {
	list, totalCount, err := s.List(ctx, in)
	if err != nil {
		return
	}

	// 字段的排序是依据tags的字段顺序，如果你不想使用默认的排序方式，可以直接定义 tags = []string{"字段名称", "字段名称2", ...}
	tags, err := convert.GetEntityDescTags(stockin.FinancialIndicatorsExportModel{})
	if err != nil {
		return
	}

	var (
		fileName  = "导出财务指标分析表-" + gctx.CtxId(ctx)
		sheetName = fmt.Sprintf("索引条件共%v行,共%v页,当前导出是第%v页,本页共%v行", totalCount, form.CalPageCount(totalCount, in.PerPage), in.Page, len(list))
		exports   []stockin.FinancialIndicatorsExportModel
	)

	if err = gconv.Scan(list, &exports); err != nil {
		return
	}

	err = excel.ExportByStructs(ctx, tags, exports, fileName, sheetName)
	return
}

// Edit 修改/新增财务指标分析表
func (s *sStockFinancialIndicators) Edit(ctx context.Context, in *stockin.FinancialIndicatorsEditInp) (err error) {
	return g.DB().Transaction(ctx, func(ctx context.Context, tx gdb.TX) (err error) {

		// 修改
		if in.Id > 0 {
			if _, err = s.Model(ctx).
				Fields(stockin.FinancialIndicatorsUpdateFields{}).
				WherePri(in.Id).Data(in).Update(); err != nil {
				err = gerror.Wrap(err, "修改财务指标分析表失败，请稍后重试！")
			}
			return
		}

		// 新增
		if _, err = s.Model(ctx, &handler.Option{FilterAuth: false}).
			Fields(stockin.FinancialIndicatorsInsertFields{}).
			Data(in).OmitEmptyData().Insert(); err != nil {
			err = gerror.Wrap(err, "新增财务指标分析表失败，请稍后重试！")
		}
		return
	})
}

// Delete 删除财务指标分析表
func (s *sStockFinancialIndicators) Delete(ctx context.Context, in *stockin.FinancialIndicatorsDeleteInp) (err error) {

	if _, err = s.Model(ctx).WherePri(in.Id).Unscoped().Delete(); err != nil {
		err = gerror.Wrap(err, "删除财务指标分析表失败，请稍后重试！")
		return
	}
	return
}

// View 获取财务指标分析表指定信息
func (s *sStockFinancialIndicators) View(ctx context.Context, in *stockin.FinancialIndicatorsViewInp) (res *stockin.FinancialIndicatorsViewModel, err error) {
	if err = s.Model(ctx).WherePri(in.Id).Scan(&res); err != nil {
		err = gerror.Wrap(err, "获取财务指标分析表信息，请稍后重试！")
		return
	}
	return
}

// GetFinancialIndicators 获取财务指标分析表数据
func (s *sStockFinancialIndicators) GetFinancialIndicators(ctx context.Context, in *stockin.FinancialIndicatorsGetFinancialIndicatorsInp) (data *entity.FinancialIndicators, err error) {
	// 构建 API URL
	// https://api.zhituapi.com/hs/gs/cwzb/股票代码?token=token证书
	apiUrl := fmt.Sprintf("https://api.zhituapi.com/hs/gs/cwzb/%s", in.Symbol)

	// 构建查询参数
	params := g.Map{
		"token": in.Token,
	}

	// 发送 GET 请求并解析为 entity.FinancialIndicators
	var result entity.FinancialIndicators
	err = g.Client().GetVar(ctx, apiUrl, params).Scan(&result)
	if err != nil {
		err = gerror.Wrap(err, "调用外部API获取财务指标分析表数据失败，请稍后重试！")
		return
	}

	data = &result
	return
}
