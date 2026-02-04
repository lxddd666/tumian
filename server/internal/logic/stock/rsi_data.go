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
	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gctx"
	"github.com/gogf/gf/v2/util/gconv"
	"hotgo/internal/dao"
	"hotgo/internal/library/hgorm/handler"
	"hotgo/internal/model/entity"
	"hotgo/internal/model/input/form"
	"hotgo/internal/model/input/stockin"
	"hotgo/internal/service"
	"hotgo/utility/convert"
	"hotgo/utility/excel"
)

type sStockRsiData struct{}

func NewStockRsiData() *sStockRsiData {
	return &sStockRsiData{}
}

func init() {
	service.RegisterStockRsiData(NewStockRsiData())
}

// Model rsi指标数据表ORM模型
func (s *sStockRsiData) Model(ctx context.Context, option ...*handler.Option) *gdb.Model {
	return handler.Model(dao.RsiData.Ctx(ctx), option...)
}

// List 获取rsi指标数据表列表
func (s *sStockRsiData) List(ctx context.Context, in *stockin.RsiDataListInp) (list []*stockin.RsiDataListModel, totalCount int, err error) {
	mod := s.Model(ctx)

	// 字段过滤
	mod = mod.Fields(stockin.RsiDataListModel{})

	// 查询自增主键
	if in.Id > 0 {
		mod = mod.Where(dao.RsiData.Columns().Id, in.Id)
	}

	// 查询数据创建时间
	if len(in.CreatedAt) == 2 {
		mod = mod.WhereBetween(dao.RsiData.Columns().CreatedAt, in.CreatedAt[0], in.CreatedAt[1])
	}

	// 分页
	mod = mod.Page(in.Page, in.PerPage)

	// 排序
	mod = mod.OrderDesc(dao.RsiData.Columns().Id)

	// 查询数据
	if err = mod.ScanAndCount(&list, &totalCount, false); err != nil {
		err = gerror.Wrap(err, "获取rsi指标数据表列表失败，请稍后重试！")
		return
	}
	return
}

// Export 导出rsi指标数据表
func (s *sStockRsiData) Export(ctx context.Context, in *stockin.RsiDataListInp) (err error) {
	list, totalCount, err := s.List(ctx, in)
	if err != nil {
		return
	}

	// 字段的排序是依据tags的字段顺序，如果你不想使用默认的排序方式，可以直接定义 tags = []string{"字段名称", "字段名称2", ...}
	tags, err := convert.GetEntityDescTags(stockin.RsiDataExportModel{})
	if err != nil {
		return
	}

	var (
		fileName  = "导出rsi指标数据表-" + gctx.CtxId(ctx)
		sheetName = fmt.Sprintf("索引条件共%v行,共%v页,当前导出是第%v页,本页共%v行", totalCount, form.CalPageCount(totalCount, in.PerPage), in.Page, len(list))
		exports   []stockin.RsiDataExportModel
	)

	if err = gconv.Scan(list, &exports); err != nil {
		return
	}

	err = excel.ExportByStructs(ctx, tags, exports, fileName, sheetName)
	return
}

// Edit 修改/新增rsi指标数据表
func (s *sStockRsiData) Edit(ctx context.Context, in *stockin.RsiDataEditInp) (err error) {
	return g.DB().Transaction(ctx, func(ctx context.Context, tx gdb.TX) (err error) {

		// 修改
		if in.Id > 0 {
			if _, err = s.Model(ctx).
				Fields(stockin.RsiDataUpdateFields{}).
				WherePri(in.Id).Data(in).Update(); err != nil {
				err = gerror.Wrap(err, "修改rsi指标数据表失败，请稍后重试！")
			}
			return
		}

		// 新增
		if _, err = s.Model(ctx, &handler.Option{FilterAuth: false}).
			Fields(stockin.RsiDataInsertFields{}).
			Data(in).OmitEmptyData().Insert(); err != nil {
			err = gerror.Wrap(err, "新增rsi指标数据表失败，请稍后重试！")
		}
		return
	})
}

// Delete 删除rsi指标数据表
func (s *sStockRsiData) Delete(ctx context.Context, in *stockin.RsiDataDeleteInp) (err error) {

	if _, err = s.Model(ctx).WherePri(in.Id).Unscoped().Delete(); err != nil {
		err = gerror.Wrap(err, "删除rsi指标数据表失败，请稍后重试！")
		return
	}
	return
}

// View 获取rsi指标数据表指定信息
func (s *sStockRsiData) View(ctx context.Context, in *stockin.RsiDataViewInp) (res *stockin.RsiDataViewModel, err error) {
	if err = s.Model(ctx).WherePri(in.Id).Scan(&res); err != nil {
		err = gerror.Wrap(err, "获取rsi指标数据表信息，请稍后重试！")
		return
	}
	return
}

// GetRsiData 获取rsi指标数据表指定信息
func (s *sStockRsiData) GetRsiData(ctx context.Context, in *stockin.GetRsiDataInp) (res *stockin.RsiDataViewModel, err error) {
	var codeList []*entity.StockSelfCode
	err = dao.StockSelfCode.Ctx(ctx).Scan(&codeList)
	if err != nil {
		return
	}
	for _, stock := range codeList {
		rsiList, _ := CalculateAllRSI(ctx, stock.Dm, 14)
		_, _ = dao.RsiData.Ctx(ctx).InsertIgnore(rsiList)
	}

	return
}

// CalculateAllRSI 计算所有历史数据的RSI
func CalculateAllRSI(ctx context.Context, code string, period int) ([]entity.RsiData, error) {
	var historicalData []*entity.EnterpriseHistoricalData
	_ = service.StockEnterpriseHistoricalData().Model(ctx).Where(dao.EnterpriseHistoricalData.Columns().Symbol, code).OrderAsc(dao.EnterpriseHistoricalData.Columns().T).Scan(&historicalData)
	if len(historicalData) < period+1 {
		return nil, fmt.Errorf("历史数据不足，需要至少%d条数据", period+1)
	}

	// 准备结果数组
	result := make([]entity.RsiData, 0, len(historicalData))
	symbol := historicalData[0].Symbol

	// 对于每条数据，计算其RSI值
	for i := period; i < len(historicalData); i++ {
		// 提取最近period+1天的收盘价
		prices := make([]float64, 0, period+1)
		for j := i - period; j <= i; j++ {
			prices = append(prices, historicalData[j].C)
		}

		rsiValue, err := calculateRSI(prices, period)
		if err != nil {
			// 如果计算失败，可以跳过或记录错误
			continue
		}

		// 添加到结果
		result = append(result, entity.RsiData{
			T:      historicalData[i].T,
			Rsi:    rsiValue,
			Symbol: symbol,
		})
	}

	return result, nil
}

// 计算单条RSI值
func calculateRSI(prices []float64, period int) (float64, error) {
	if len(prices) < period+1 {
		return 0, fmt.Errorf("数据不足，无法计算%d日RSI", period)
	}

	// 计算价格变化
	gains := make([]float64, 0, period)
	losses := make([]float64, 0, period)

	for i := 1; i <= period; i++ {
		change := prices[i] - prices[i-1]
		if change > 0 {
			gains = append(gains, change)
			losses = append(losses, 0)
		} else {
			gains = append(gains, 0)
			losses = append(losses, -change)
		}
	}

	// 计算平均收益和平均损失（初始平均值）
	avgGain := 0.0
	avgLoss := 0.0

	for i := 0; i < period; i++ {
		avgGain += gains[i]
		avgLoss += losses[i]
	}

	avgGain /= float64(period)
	avgLoss /= float64(period)

	// 使用平滑移动平均计算后续值（如果数据超过period+1）
	if len(prices) > period+1 {
		for i := period + 1; i < len(prices); i++ {
			change := prices[i] - prices[i-1]
			gain := 0.0
			loss := 0.0

			if change > 0 {
				gain = change
			} else {
				loss = -change
			}

			// 平滑移动平均
			avgGain = (avgGain*float64(period-1) + gain) / float64(period)
			avgLoss = (avgLoss*float64(period-1) + loss) / float64(period)
		}
	}

	// 计算RSI
	if avgLoss == 0 {
		return 100.0, nil
	}

	rs := avgGain / avgLoss
	rsi := 100 - (100 / (1 + rs))

	return rsi, nil
}
