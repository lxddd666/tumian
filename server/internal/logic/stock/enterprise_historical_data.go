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
	"hotgo/internal/consts"
	"hotgo/internal/dao"
	"hotgo/internal/global"
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

type sStockEnterpriseHistoricalData struct{}

func NewStockEnterpriseHistoricalData() *sStockEnterpriseHistoricalData {
	return &sStockEnterpriseHistoricalData{}
}

func init() {
	service.RegisterStockEnterpriseHistoricalData(NewStockEnterpriseHistoricalData())
}

// Model 企业级历史行情数据表 (K线数据)ORM模型
func (s *sStockEnterpriseHistoricalData) Model(ctx context.Context, option ...*handler.Option) *gdb.Model {
	return handler.Model(dao.EnterpriseHistoricalData.Ctx(ctx), option...)
}

// List 获取企业级历史行情数据表 (K线数据)列表
func (s *sStockEnterpriseHistoricalData) List(ctx context.Context, in *stockin.EnterpriseHistoricalDataListInp) (list []*stockin.EnterpriseHistoricalDataListModel, totalCount int, err error) {
	mod := s.Model(ctx)

	// 字段过滤
	mod = mod.Fields(stockin.EnterpriseHistoricalDataListModel{})

	// 查询自增主键
	if in.Id > 0 {
		mod = mod.Where(dao.EnterpriseHistoricalData.Columns().Id, in.Id)
	}

	// 查询创建时间
	if len(in.CreatedAt) == 2 {
		mod = mod.WhereBetween(dao.EnterpriseHistoricalData.Columns().CreatedAt, in.CreatedAt[0], in.CreatedAt[1])
	}

	// 分页
	mod = mod.Page(in.Page, in.PerPage)

	// 排序
	mod = mod.OrderDesc(dao.EnterpriseHistoricalData.Columns().Id)

	// 查询数据
	if err = mod.ScanAndCount(&list, &totalCount, false); err != nil {
		err = gerror.Wrap(err, "获取企业级历史行情数据表 (K线数据)列表失败，请稍后重试！")
		return
	}
	return
}

// Export 导出企业级历史行情数据表 (K线数据)
func (s *sStockEnterpriseHistoricalData) Export(ctx context.Context, in *stockin.EnterpriseHistoricalDataListInp) (err error) {
	list, totalCount, err := s.List(ctx, in)
	if err != nil {
		return
	}

	// 字段的排序是依据tags的字段顺序，如果你不想使用默认的排序方式，可以直接定义 tags = []string{"字段名称", "字段名称2", ...}
	tags, err := convert.GetEntityDescTags(stockin.EnterpriseHistoricalDataExportModel{})
	if err != nil {
		return
	}

	var (
		fileName  = "导出企业级历史行情数据表 (K线数据)-" + gctx.CtxId(ctx)
		sheetName = fmt.Sprintf("索引条件共%v行,共%v页,当前导出是第%v页,本页共%v行", totalCount, form.CalPageCount(totalCount, in.PerPage), in.Page, len(list))
		exports   []stockin.EnterpriseHistoricalDataExportModel
	)

	if err = gconv.Scan(list, &exports); err != nil {
		return
	}

	err = excel.ExportByStructs(ctx, tags, exports, fileName, sheetName)
	return
}

// Edit 修改/新增企业级历史行情数据表 (K线数据)
func (s *sStockEnterpriseHistoricalData) Edit(ctx context.Context, in *stockin.EnterpriseHistoricalDataEditInp) (err error) {
	return g.DB().Transaction(ctx, func(ctx context.Context, tx gdb.TX) (err error) {

		// 修改
		if in.Id > 0 {
			if _, err = s.Model(ctx).
				Fields(stockin.EnterpriseHistoricalDataUpdateFields{}).
				WherePri(in.Id).Data(in).Update(); err != nil {
				err = gerror.Wrap(err, "修改企业级历史行情数据表 (K线数据)失败，请稍后重试！")
			}
			return
		}

		// 新增
		if _, err = s.Model(ctx, &handler.Option{FilterAuth: false}).
			Fields(stockin.EnterpriseHistoricalDataInsertFields{}).
			Data(in).OmitEmptyData().Insert(); err != nil {
			err = gerror.Wrap(err, "新增企业级历史行情数据表 (K线数据)失败，请稍后重试！")
		}
		return
	})
}

// Delete 删除企业级历史行情数据表 (K线数据)
func (s *sStockEnterpriseHistoricalData) Delete(ctx context.Context, in *stockin.EnterpriseHistoricalDataDeleteInp) (err error) {

	if _, err = s.Model(ctx).WherePri(in.Id).Unscoped().Delete(); err != nil {
		err = gerror.Wrap(err, "删除企业级历史行情数据表 (K线数据)失败，请稍后重试！")
		return
	}
	return
}

// View 获取企业级历史行情数据表 (K线数据)指定信息
func (s *sStockEnterpriseHistoricalData) View(ctx context.Context, in *stockin.EnterpriseHistoricalDataViewInp) (res *stockin.EnterpriseHistoricalDataViewModel, err error) {
	if err = s.Model(ctx).WherePri(in.Id).Scan(&res); err != nil {
		err = gerror.Wrap(err, "获取企业级历史行情数据表 (K线数据)信息，请稍后重试！")
		return
	}
	return
}

// GetEnterpriseHistoricalData 获取企业级历史行情数据表数据
func (s *sStockEnterpriseHistoricalData) GetEnterpriseHistoricalData(ctx context.Context, in *stockin.EnterpriseHistoricalDataGetEnterpriseHistoricalDataInp) (data []*entity.EnterpriseHistoricalData, err error) {
	flag, err := s.Model(ctx).Where(dao.EnterpriseHistoricalData.Columns().T, GetNowDate()).Exist()
	if err != nil {
		return
	}
	if flag {
		return
	}
	if in.Interval == "" {
		in.Interval = "d"
	}
	if in.AdjustType == "" {
		in.AdjustType = "n"
	}

	// 构建 API URL
	// http://专属子域名.zhituapi.com/hs/hsstock/vip/股票代码.市场（如000001.SZ）/分时级别(如d)/除权方式(如n)?token=token证书&st=开始时间(如20240601)&et=结束时间(如20250430)&lt=最新条数(如100)
	// 注意：专属子域名需要根据实际情况配置，这里使用 api 作为默认值
	apiUrl := fmt.Sprintf("http://api.zhituapi.com/hs/hsstock/vip/%s/%s/%s", in.Symbol, in.Interval, in.AdjustType)

	// 构建查询参数
	params := g.Map{
		"token": global.StockToken,
	}

	// 添加可选参数
	if in.StartTime != "" {
		params["st"] = GetYewBefore(1)
	}
	if in.EndTime != "" {
		params["et"] = GetNowDate()
	}
	if in.Limit <= 0 {
		params["lt"] = consts.StockLimitPiecesDefault
	}

	// 发送 GET 请求并解析为 entity.EnterpriseHistoricalData
	var result []*entity.EnterpriseHistoricalData
	err = g.Client().GetVar(ctx, apiUrl, params).Scan(&result)
	if err != nil {
		err = gerror.Wrap(err, "调用外部API获取企业级历史行情数据表数据失败，请稍后重试！")
		return
	}

	data = result
	return
}
