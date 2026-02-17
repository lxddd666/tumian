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
	"github.com/gogf/gf/v2/text/gstr"
	"hotgo/internal/dao"
	"hotgo/internal/library/hgorm/handler"
	"hotgo/internal/model/entity"
	"hotgo/internal/model/input/form"
	"hotgo/internal/model/input/stockin"
	"hotgo/internal/service"
	"hotgo/utility/convert"
	"hotgo/utility/excel"
	"io"
	"log"
	"net/http"
	"time"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gctx"
	"github.com/gogf/gf/v2/util/gconv"
)

type sStockWilliamsData struct{}

func NewStockWilliamsData() *sStockWilliamsData {
	return &sStockWilliamsData{}
}

func init() {
	service.RegisterStockWilliamsData(NewStockWilliamsData())
}

// Model wmsr指标数据表ORM模型
func (s *sStockWilliamsData) Model(ctx context.Context, option ...*handler.Option) *gdb.Model {
	return handler.Model(dao.WilliamsData.Ctx(ctx), option...)
}

// List 获取wmsr指标数据表列表
func (s *sStockWilliamsData) List(ctx context.Context, in *stockin.WilliamsDataListInp) (list []*stockin.WilliamsDataListModel, totalCount int, err error) {
	mod := s.Model(ctx)

	// 字段过滤
	mod = mod.Fields(stockin.WilliamsDataListModel{})

	// 查询自增主键
	if in.Id > 0 {
		mod = mod.Where(dao.WilliamsData.Columns().Id, in.Id)
	}

	// 查询数据创建时间
	if len(in.CreatedAt) == 2 {
		mod = mod.WhereBetween(dao.WilliamsData.Columns().CreatedAt, in.CreatedAt[0], in.CreatedAt[1])
	}

	// 分页
	mod = mod.Page(in.Page, in.PerPage)

	// 排序
	mod = mod.OrderDesc(dao.WilliamsData.Columns().Id)

	// 查询数据
	if err = mod.ScanAndCount(&list, &totalCount, false); err != nil {
		err = gerror.Wrap(err, "获取wmsr指标数据表列表失败，请稍后重试！")
		return
	}
	return
}

// Export 导出wmsr指标数据表
func (s *sStockWilliamsData) Export(ctx context.Context, in *stockin.WilliamsDataListInp) (err error) {
	list, totalCount, err := s.List(ctx, in)
	if err != nil {
		return
	}

	// 字段的排序是依据tags的字段顺序，如果你不想使用默认的排序方式，可以直接定义 tags = []string{"字段名称", "字段名称2", ...}
	tags, err := convert.GetEntityDescTags(stockin.WilliamsDataExportModel{})
	if err != nil {
		return
	}

	var (
		fileName  = "导出wmsr指标数据表-" + gctx.CtxId(ctx)
		sheetName = fmt.Sprintf("索引条件共%v行,共%v页,当前导出是第%v页,本页共%v行", totalCount, form.CalPageCount(totalCount, in.PerPage), in.Page, len(list))
		exports   []stockin.WilliamsDataExportModel
	)

	if err = gconv.Scan(list, &exports); err != nil {
		return
	}

	err = excel.ExportByStructs(ctx, tags, exports, fileName, sheetName)
	return
}

// Edit 修改/新增wmsr指标数据表
func (s *sStockWilliamsData) Edit(ctx context.Context, in *stockin.WilliamsDataEditInp) (err error) {
	return g.DB().Transaction(ctx, func(ctx context.Context, tx gdb.TX) (err error) {

		// 修改
		if in.Id > 0 {
			if _, err = s.Model(ctx).
				Fields(stockin.WilliamsDataUpdateFields{}).
				WherePri(in.Id).Data(in).Update(); err != nil {
				err = gerror.Wrap(err, "修改wmsr指标数据表失败，请稍后重试！")
			}
			return
		}

		// 新增
		if _, err = s.Model(ctx, &handler.Option{FilterAuth: false}).
			Fields(stockin.WilliamsDataInsertFields{}).
			Data(in).OmitEmptyData().Insert(); err != nil {
			err = gerror.Wrap(err, "新增wmsr指标数据表失败，请稍后重试！")
		}
		return
	})
}

// Delete 删除wmsr指标数据表
func (s *sStockWilliamsData) Delete(ctx context.Context, in *stockin.WilliamsDataDeleteInp) (err error) {

	if _, err = s.Model(ctx).WherePri(in.Id).Unscoped().Delete(); err != nil {
		err = gerror.Wrap(err, "删除wmsr指标数据表失败，请稍后重试！")
		return
	}
	return
}

// View 获取wmsr指标数据表指定信息
func (s *sStockWilliamsData) View(ctx context.Context, in *stockin.WilliamsDataViewInp) (res *stockin.WilliamsDataViewModel, err error) {
	if err = s.Model(ctx).WherePri(in.Id).Scan(&res); err != nil {
		err = gerror.Wrap(err, "获取wmsr指标数据表信息，请稍后重试！")
		return
	}
	return
}

// GetWmsr 获取GetWmsr指标
func (s *sStockWilliamsData) GetWmsr(ctx context.Context, in *stockin.GetWilliamsDataInp) (data []*entity.WilliamsData, err error) {
	code := gstr.Split(in.Symbol, ".")[0]

	flag, err := s.Model(ctx).Where(dao.MacdData.Columns().T, GetRecentWeekday()).Where(dao.WilliamsData.Columns().Symbol, in.Symbol).Exist()
	if err != nil {
		return
	}
	if flag {
		return
	}

	// 目标URL
	url := fmt.Sprintf("https://finance.pae.baidu.com/sapi/v1/get_indicators_graph?code=%s&eventType=1010&finClientType=pc&financeType=stock&market=ab&period=dayK&finClientType=pc", code)

	// 发起GET请求
	resp, err := http.Get(url)
	if err != nil {
		log.Fatal("请求失败:", err)
		return
	}
	defer resp.Body.Close()

	// 读取响应体
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatal("读取数据失败:", err)
		return
	}

	var result map[string]interface{}
	err = gconv.Scan(body, &result)
	if err != nil {
		err = gerror.Wrap(err, "解析body失败")
		return
	}
	// 打印JSON结果
	if result["Result"] != nil {
		resultMap := result["Result"].(map[string]interface{})
		if resultMap == nil {
			return
		}
		marketData := resultMap["market_data"].(string)
		list := gstr.Split(marketData, ";")
	LoopCci:
		for _, l := range list {
			wmsr := new(entity.WilliamsData)
			wmsr.Symbol = in.Symbol
			cciData := gstr.Split(l, ",")
			for i, c := range cciData {
				// 20171116,18.040,17.580
				switch i {
				case 0:
					// 时间
					t, tErr := time.Parse("20060102", c)
					if tErr != nil {
						err = tErr
						return
					}

					// 格式化为目标格式
					formatted := t.Format("2006-01-02")
					wmsr.T = gtime.New(formatted)
				case 1:
					if gstr.Contains(c, "--") {
						continue LoopCci
					}
					wmsr.R = gconv.Float64(c)
				}
			}
			data = append(data, wmsr)
		}
	}
	if len(data) > 0 {
		_, _ = s.Model(ctx).InsertIgnore(data)
	}
	return
}
