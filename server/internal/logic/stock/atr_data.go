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
	"math"
)

type sStockAtrData struct{}

func NewStockAtrData() *sStockAtrData {
	return &sStockAtrData{}
}

func init() {
	service.RegisterStockAtrData(NewStockAtrData())
}

// Model atr指标数据表ORM模型
func (s *sStockAtrData) Model(ctx context.Context, option ...*handler.Option) *gdb.Model {
	return handler.Model(dao.AtrData.Ctx(ctx), option...)
}

// List 获取atr指标数据表列表
func (s *sStockAtrData) List(ctx context.Context, in *stockin.AtrDataListInp) (list []*stockin.AtrDataListModel, totalCount int, err error) {
	mod := s.Model(ctx)

	// 字段过滤
	mod = mod.Fields(stockin.AtrDataListModel{})

	// 查询自增主键
	if in.Id > 0 {
		mod = mod.Where(dao.AtrData.Columns().Id, in.Id)
	}

	// 查询数据创建时间
	if len(in.CreatedAt) == 2 {
		mod = mod.WhereBetween(dao.AtrData.Columns().CreatedAt, in.CreatedAt[0], in.CreatedAt[1])
	}

	// 分页
	mod = mod.Page(in.Page, in.PerPage)

	// 排序
	mod = mod.OrderDesc(dao.AtrData.Columns().Id)

	// 查询数据
	if err = mod.ScanAndCount(&list, &totalCount, false); err != nil {
		err = gerror.Wrap(err, "获取atr指标数据表列表失败，请稍后重试！")
		return
	}
	return
}

// Export 导出atr指标数据表
func (s *sStockAtrData) Export(ctx context.Context, in *stockin.AtrDataListInp) (err error) {
	list, totalCount, err := s.List(ctx, in)
	if err != nil {
		return
	}

	// 字段的排序是依据tags的字段顺序，如果你不想使用默认的排序方式，可以直接定义 tags = []string{"字段名称", "字段名称2", ...}
	tags, err := convert.GetEntityDescTags(stockin.AtrDataExportModel{})
	if err != nil {
		return
	}

	var (
		fileName  = "导出atr指标数据表-" + gctx.CtxId(ctx)
		sheetName = fmt.Sprintf("索引条件共%v行,共%v页,当前导出是第%v页,本页共%v行", totalCount, form.CalPageCount(totalCount, in.PerPage), in.Page, len(list))
		exports   []stockin.AtrDataExportModel
	)

	if err = gconv.Scan(list, &exports); err != nil {
		return
	}

	err = excel.ExportByStructs(ctx, tags, exports, fileName, sheetName)
	return
}

// Edit 修改/新增atr指标数据表
func (s *sStockAtrData) Edit(ctx context.Context, in *stockin.AtrDataEditInp) (err error) {
	return g.DB().Transaction(ctx, func(ctx context.Context, tx gdb.TX) (err error) {

		// 修改
		if in.Id > 0 {
			if _, err = s.Model(ctx).
				Fields(stockin.AtrDataUpdateFields{}).
				WherePri(in.Id).Data(in).Update(); err != nil {
				err = gerror.Wrap(err, "修改atr指标数据表失败，请稍后重试！")
			}
			return
		}

		// 新增
		if _, err = s.Model(ctx, &handler.Option{FilterAuth: false}).
			Fields(stockin.AtrDataInsertFields{}).
			Data(in).OmitEmptyData().Insert(); err != nil {
			err = gerror.Wrap(err, "新增atr指标数据表失败，请稍后重试！")
		}
		return
	})
}

// Delete 删除atr指标数据表
func (s *sStockAtrData) Delete(ctx context.Context, in *stockin.AtrDataDeleteInp) (err error) {

	if _, err = s.Model(ctx).WherePri(in.Id).Unscoped().Delete(); err != nil {
		err = gerror.Wrap(err, "删除atr指标数据表失败，请稍后重试！")
		return
	}
	return
}

// View 获取atr指标数据表指定信息
func (s *sStockAtrData) View(ctx context.Context, in *stockin.AtrDataViewInp) (res *stockin.AtrDataViewModel, err error) {
	if err = s.Model(ctx).WherePri(in.Id).Scan(&res); err != nil {
		err = gerror.Wrap(err, "获取atr指标数据表信息，请稍后重试！")
		return
	}
	return
}

// GetAtrData 获取atr数据指标
func (s *sStockAtrData) GetAtrData(ctx context.Context, in *stockin.GetAtrDataInp) (res *stockin.AtrDataViewModel, err error) {
	flag, err := s.Model(ctx).Where(dao.AtrData.Columns().T, GetRecentWeekday()).Where(dao.AtrData.Columns().Symbol, in.Symbol).Exist()
	if err != nil {
		return
	}
	if flag {
		return
	}

	var data []*entity.EnterpriseHistoricalData
	_ = dao.EnterpriseHistoricalData.Ctx(ctx).Where(dao.EnterpriseHistoricalData.Columns().Symbol, in.Symbol).OrderAsc(dao.AtrData.Columns().T).Scan(&data)
	artList := CalculateATR(data, 14, "RMA")
	if len(artList) > 0 {
		dao.AtrData.Ctx(ctx).InsertIgnore(artList)
	}
	return
}

// CalculateATR 计算 ATR 数据
func CalculateATR(data []*entity.EnterpriseHistoricalData, length int, smoothing string) []*entity.AtrData {
	if len(data) == 0 || length <= 0 {
		return nil
	}

	n := len(data)
	trList := make([]float64, n)
	atrList := make([]float64, n)

	// Step 1: 计算 True Range (TR)
	for i := 0; i < n; i++ {
		high := data[i].H
		low := data[i].L
		//closePrice = data[i].C

		var tr float64
		if i == 0 {
			tr = high - low // 第一天没有前收盘，只用 H-L
		} else {
			prevClose := data[i-1].C
			tr1 := high - low
			tr2 := math.Abs(high - prevClose)
			tr3 := math.Abs(low - prevClose)
			tr = math.Max(tr1, math.Max(tr2, tr3))
		}
		trList[i] = tr
	}

	// Step 2: 根据 smoothing 计算 ATR
	switch smoothing {
	case "RMA":
		// Wilder's Smoothing (默认)
		if n < length {
			// 不足 length，无法计算完整 ATR，可选择返回空或部分
			// 这里我们仍计算可用部分（前 length 个用 SMA 初始化）
			sum := 0.0
			for i := 0; i < n && i < length; i++ {
				sum += trList[i]
				atrList[i] = sum / float64(i+1)
			}
			for i := length; i < n; i++ {
				atrList[i] = (atrList[i-1]*(float64(length)-1) + trList[i]) / float64(length)
			}
		} else {
			// 先计算前 length 个的 SMA 作为初始 ATR
			sum := 0.0
			for i := 0; i < length; i++ {
				sum += trList[i]
			}
			initialATR := sum / float64(length)
			for i := 0; i < length; i++ {
				atrList[i] = initialATR // 或者也可以只从第 length-1 开始赋值
			}
			// 从第 length 个开始用 RMA 公式
			for i := length; i < n; i++ {
				atrList[i] = (atrList[i-1]*(float64(length)-1) + trList[i]) / float64(length)
			}
		}

	case "SMA":
		for i := 0; i < n; i++ {
			if i < length-1 {
				atrList[i] = 0 // 或 NaN，但 Go 中用 0 表示无效
			} else {
				sum := 0.0
				for j := i - length + 1; j <= i; j++ {
					sum += trList[j]
				}
				atrList[i] = sum / float64(length)
			}
		}

	case "EMA":
		alpha := 2.0 / float64(length+1)
		for i := 0; i < n; i++ {
			if i == 0 {
				atrList[i] = trList[0]
			} else {
				atrList[i] = alpha*trList[i] + (1-alpha)*atrList[i-1]
			}
		}

	case "WMA":
		for i := 0; i < n; i++ {
			if i < length-1 {
				atrList[i] = 0
			} else {
				sum := 0.0
				weightSum := 0.0
				for j := 0; j < length; j++ {
					weight := float64(j + 1)
					sum += weight * trList[i-length+1+j]
					weightSum += weight
				}
				atrList[i] = sum / weightSum
			}
		}

	default:
		// 默认使用 RMA
		return CalculateATR(data, length, "RMA")
	}

	// Step 3: 构建 AtrData 列表
	result := make([]*entity.AtrData, n)
	for i := 0; i < n; i++ {
		result[i] = &entity.AtrData{
			Symbol: data[i].Symbol,
			T:      data[i].T,
			Tr:     trList[i],
			Atr:    atrList[i],
			// Id, CreatedAt, UpdatedAt 可由数据库自动填充，此处忽略
		}
	}

	return result
}
