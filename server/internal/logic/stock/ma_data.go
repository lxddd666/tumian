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

type sStockMaData struct{}

func NewStockMaData() *sStockMaData {
	return &sStockMaData{}
}

func init() {
	service.RegisterStockMaData(NewStockMaData())
}

// Model 移动平均线(MA)指标数据表ORM模型
func (s *sStockMaData) Model(ctx context.Context, option ...*handler.Option) *gdb.Model {
	return handler.Model(dao.MaData.Ctx(ctx), option...)
}

// List 获取移动平均线(MA)指标数据表列表
func (s *sStockMaData) List(ctx context.Context, in *stockin.MaDataListInp) (list []*stockin.MaDataListModel, totalCount int, err error) {
	mod := s.Model(ctx)

	// 字段过滤
	mod = mod.Fields(stockin.MaDataListModel{})

	// 查询自增主键
	if in.Id > 0 {
		mod = mod.Where(dao.MaData.Columns().Id, in.Id)
	}

	// 查询数据创建时间
	if len(in.CreatedAt) == 2 {
		mod = mod.WhereBetween(dao.MaData.Columns().CreatedAt, in.CreatedAt[0], in.CreatedAt[1])
	}

	// 分页
	mod = mod.Page(in.Page, in.PerPage)

	// 排序
	mod = mod.OrderDesc(dao.MaData.Columns().Id)

	// 查询数据
	if err = mod.ScanAndCount(&list, &totalCount, false); err != nil {
		err = gerror.Wrap(err, "获取移动平均线(MA)指标数据表列表失败，请稍后重试！")
		return
	}
	return
}

// Export 导出移动平均线(MA)指标数据表
func (s *sStockMaData) Export(ctx context.Context, in *stockin.MaDataListInp) (err error) {
	list, totalCount, err := s.List(ctx, in)
	if err != nil {
		return
	}

	// 字段的排序是依据tags的字段顺序，如果你不想使用默认的排序方式，可以直接定义 tags = []string{"字段名称", "字段名称2", ...}
	tags, err := convert.GetEntityDescTags(stockin.MaDataExportModel{})
	if err != nil {
		return
	}

	var (
		fileName  = "导出移动平均线(MA)指标数据表-" + gctx.CtxId(ctx)
		sheetName = fmt.Sprintf("索引条件共%v行,共%v页,当前导出是第%v页,本页共%v行", totalCount, form.CalPageCount(totalCount, in.PerPage), in.Page, len(list))
		exports   []stockin.MaDataExportModel
	)

	if err = gconv.Scan(list, &exports); err != nil {
		return
	}

	err = excel.ExportByStructs(ctx, tags, exports, fileName, sheetName)
	return
}

// Edit 修改/新增移动平均线(MA)指标数据表
func (s *sStockMaData) Edit(ctx context.Context, in *stockin.MaDataEditInp) (err error) {
	return g.DB().Transaction(ctx, func(ctx context.Context, tx gdb.TX) (err error) {

		// 修改
		if in.Id > 0 {
			if _, err = s.Model(ctx).
				Fields(stockin.MaDataUpdateFields{}).
				WherePri(in.Id).Data(in).Update(); err != nil {
				err = gerror.Wrap(err, "修改移动平均线(MA)指标数据表失败，请稍后重试！")
			}
			return
		}

		// 新增
		if _, err = s.Model(ctx, &handler.Option{FilterAuth: false}).
			Fields(stockin.MaDataInsertFields{}).
			Data(in).OmitEmptyData().Insert(); err != nil {
			err = gerror.Wrap(err, "新增移动平均线(MA)指标数据表失败，请稍后重试！")
		}
		return
	})
}

// Delete 删除移动平均线(MA)指标数据表
func (s *sStockMaData) Delete(ctx context.Context, in *stockin.MaDataDeleteInp) (err error) {

	if _, err = s.Model(ctx).WherePri(in.Id).Unscoped().Delete(); err != nil {
		err = gerror.Wrap(err, "删除移动平均线(MA)指标数据表失败，请稍后重试！")
		return
	}
	return
}

// View 获取移动平均线(MA)指标数据表指定信息
func (s *sStockMaData) View(ctx context.Context, in *stockin.MaDataViewInp) (res *stockin.MaDataViewModel, err error) {
	if err = s.Model(ctx).WherePri(in.Id).Scan(&res); err != nil {
		err = gerror.Wrap(err, "获取移动平均线(MA)指标数据表信息，请稍后重试！")
		return
	}
	return
}

// GetMa 获取移动平均线(MA)指标数据
func (s *sStockMaData) GetMa(ctx context.Context, in *stockin.MaDataGetMaInp) (data *entity.MaData, err error) {
	// 构建 API URL
	// https://api.zhituapi.com/hs/history/ma/股票代码(如000001.SZ)/分时级别(如d)/除权类型(如n)?token=token证书&st=开始时间&et=结束时间&lt=最新条数
	apiUrl := fmt.Sprintf("https://api.zhituapi.com/hs/history/ma/%s/%s/%s", in.Symbol, in.Interval, in.AdjustType)

	// 构建查询参数
	params := g.Map{
		"token": in.Token,
	}

	// 添加可选参数
	if in.StartTime != "" {
		params["st"] = in.StartTime
	}
	if in.EndTime != "" {
		params["et"] = in.EndTime
	}
	if in.Limit > 0 {
		params["lt"] = in.Limit
	}

	// 发送 GET 请求并解析为 entity.MaData
	var result entity.MaData
	err = g.Client().GetVar(ctx, apiUrl, params).Scan(&result)
	if err != nil {
		err = gerror.Wrap(err, "调用外部API获取移动平均线(MA)指标数据失败，请稍后重试！")
		return
	}

	data = &result
	return
}
