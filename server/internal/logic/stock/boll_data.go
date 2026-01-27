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
	"github.com/gogf/gf/v2/os/gtime"
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
	"time"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gctx"
	"github.com/gogf/gf/v2/util/gconv"
)

type sStockBollData struct{}

func NewStockBollData() *sStockBollData {
	return &sStockBollData{}
}

func init() {
	service.RegisterStockBollData(NewStockBollData())
}

// Model 布林带(BOLL)指标数据表ORM模型
func (s *sStockBollData) Model(ctx context.Context, option ...*handler.Option) *gdb.Model {
	return handler.Model(dao.BollData.Ctx(ctx), option...)
}

// List 获取布林带(BOLL)指标数据表列表
func (s *sStockBollData) List(ctx context.Context, in *stockin.BollDataListInp) (list []*stockin.BollDataListModel, totalCount int, err error) {
	mod := s.Model(ctx)

	// 字段过滤
	mod = mod.Fields(stockin.BollDataListModel{})

	// 查询自增主键
	if in.Id > 0 {
		mod = mod.Where(dao.BollData.Columns().Id, in.Id)
	}

	// 查询数据创建时间
	if len(in.CreatedAt) == 2 {
		mod = mod.WhereBetween(dao.BollData.Columns().CreatedAt, in.CreatedAt[0], in.CreatedAt[1])
	}

	// 分页
	mod = mod.Page(in.Page, in.PerPage)

	// 排序
	mod = mod.OrderDesc(dao.BollData.Columns().Id)

	// 查询数据
	if err = mod.ScanAndCount(&list, &totalCount, false); err != nil {
		err = gerror.Wrap(err, "获取布林带(BOLL)指标数据表列表失败，请稍后重试！")
		return
	}
	return
}

// Export 导出布林带(BOLL)指标数据表
func (s *sStockBollData) Export(ctx context.Context, in *stockin.BollDataListInp) (err error) {
	list, totalCount, err := s.List(ctx, in)
	if err != nil {
		return
	}

	// 字段的排序是依据tags的字段顺序，如果你不想使用默认的排序方式，可以直接定义 tags = []string{"字段名称", "字段名称2", ...}
	tags, err := convert.GetEntityDescTags(stockin.BollDataExportModel{})
	if err != nil {
		return
	}

	var (
		fileName  = "导出布林带(BOLL)指标数据表-" + gctx.CtxId(ctx)
		sheetName = fmt.Sprintf("索引条件共%v行,共%v页,当前导出是第%v页,本页共%v行", totalCount, form.CalPageCount(totalCount, in.PerPage), in.Page, len(list))
		exports   []stockin.BollDataExportModel
	)

	if err = gconv.Scan(list, &exports); err != nil {
		return
	}

	err = excel.ExportByStructs(ctx, tags, exports, fileName, sheetName)
	return
}

// Edit 修改/新增布林带(BOLL)指标数据表
func (s *sStockBollData) Edit(ctx context.Context, in *stockin.BollDataEditInp) (err error) {
	return g.DB().Transaction(ctx, func(ctx context.Context, tx gdb.TX) (err error) {

		// 修改
		if in.Id > 0 {
			if _, err = s.Model(ctx).
				Fields(stockin.BollDataUpdateFields{}).
				WherePri(in.Id).Data(in).Update(); err != nil {
				err = gerror.Wrap(err, "修改布林带(BOLL)指标数据表失败，请稍后重试！")
			}
			return
		}

		// 新增
		if _, err = s.Model(ctx, &handler.Option{FilterAuth: false}).
			Fields(stockin.BollDataInsertFields{}).
			Data(in).OmitEmptyData().Insert(); err != nil {
			err = gerror.Wrap(err, "新增布林带(BOLL)指标数据表失败，请稍后重试！")
		}
		return
	})
}

// Delete 删除布林带(BOLL)指标数据表
func (s *sStockBollData) Delete(ctx context.Context, in *stockin.BollDataDeleteInp) (err error) {

	if _, err = s.Model(ctx).WherePri(in.Id).Unscoped().Delete(); err != nil {
		err = gerror.Wrap(err, "删除布林带(BOLL)指标数据表失败，请稍后重试！")
		return
	}
	return
}

// View 获取布林带(BOLL)指标数据表指定信息
func (s *sStockBollData) View(ctx context.Context, in *stockin.BollDataViewInp) (res *stockin.BollDataViewModel, err error) {
	if err = s.Model(ctx).WherePri(in.Id).Scan(&res); err != nil {
		err = gerror.Wrap(err, "获取布林带(BOLL)指标数据表信息，请稍后重试！")
		return
	}
	return
}

// GetBoll 获取布林带(BOLL)指标数据
func (s *sStockBollData) GetBoll(ctx context.Context, in *stockin.BollDataGetBollInp) (data []*entity.BollData, err error) {
	flag, err := s.Model(ctx).Where(dao.BollData.Columns().T, GetRecentWeekday()).Where(dao.BollData.Columns().Symbol, in.Symbol).Exist()
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
	// https://api.zhituapi.com/hs/history/boll/股票代码(如000001.SZ)/分时级别(如d)/除权类型(如n)?token=token证书&st=开始时间&et=结束时间&lt=最新条数
	apiUrl := fmt.Sprintf("https://api.zhituapi.com/hs/history/boll/%s/%s/%s", in.Symbol, in.Interval, in.AdjustType)

	// 构建查询参数
	params := g.Map{
		"token": global.GetToken(),
	}

	// 添加可选参数
	if in.StartTime == "" {
		params["st"] = GetYewBefore(1)
	}
	if in.EndTime == "" {
		params["et"] = GetNowDate()
	}
	if in.Limit <= 0 {
		params["lt"] = consts.StockLimitPiecesDefault
	}

	// 发送 GET 请求并解析为 entity.BollData
	var result []map[string]interface{}

	err = g.Client().GetVar(ctx, apiUrl, params).Scan(&result)
	if err != nil {
		err = gerror.Wrap(err, "调用外部API获取布林带(BOLL)指标数据失败，请稍后重试！")
		return
	}
	for _, re := range result {
		re["symbol"] = in.Symbol
		re["IntervalType"] = in.Interval

	}
	if len(result) > 0 {
		_, err = s.Model(ctx).InsertIgnore(result)
	}
	return
}

func GetNowDate() string {
	return gtime.Now().Format("Ymd")
}

func GetYewBefore(year int) string {
	if year <= 0 {
		year = 1
	}
	return gtime.Now().AddDate(-year, 0, 0).Format("Ymd")
}

// GetRecentWeekday 返回最近的工作日（周一到周五），时间部分清零
// 如果当天是工作日，返回当天；否则返回上一个工作日
func GetRecentWeekday() string {
	now := time.Now()
	weekday := now.Weekday()

	// 计算需要回溯的天数
	daysToSubtract := 0

	switch weekday {
	case time.Saturday:
		// 周六：回溯1天到周五
		daysToSubtract = 1
	case time.Sunday:
		// 周日：回溯2天到周五
		daysToSubtract = 2
	default:
		// 周一到周五：不需要回溯
		daysToSubtract = 0
	}

	// 计算目标日期
	targetDate := now.AddDate(0, 0, -daysToSubtract)

	// 清零时间部分，只保留年月日
	dateOnly := time.Date(targetDate.Year(), targetDate.Month(), targetDate.Day(),
		0, 0, 0, 0, targetDate.Location())

	return gtime.NewFromTime(dateOnly).Format("Ymd")
}
