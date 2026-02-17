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

type sStockMomData struct{}

func NewStockMomData() *sStockMomData {
	return &sStockMomData{}
}

func init() {
	service.RegisterStockMomData(NewStockMomData())
}

// Model mom指标数据表ORM模型
func (s *sStockMomData) Model(ctx context.Context, option ...*handler.Option) *gdb.Model {
	return handler.Model(dao.MomData.Ctx(ctx), option...)
}

// List 获取mom指标数据表列表
func (s *sStockMomData) List(ctx context.Context, in *stockin.MomDataListInp) (list []*stockin.MomDataListModel, totalCount int, err error) {
	mod := s.Model(ctx)

	// 字段过滤
	mod = mod.Fields(stockin.MomDataListModel{})

	// 查询自增主键
	if in.Id > 0 {
		mod = mod.Where(dao.MomData.Columns().Id, in.Id)
	}

	// 查询数据创建时间
	if len(in.CreatedAt) == 2 {
		mod = mod.WhereBetween(dao.MomData.Columns().CreatedAt, in.CreatedAt[0], in.CreatedAt[1])
	}

	// 分页
	mod = mod.Page(in.Page, in.PerPage)

	// 排序
	mod = mod.OrderDesc(dao.MomData.Columns().Id)

	// 查询数据
	if err = mod.ScanAndCount(&list, &totalCount, false); err != nil {
		err = gerror.Wrap(err, "获取mom指标数据表列表失败，请稍后重试！")
		return
	}
	return
}

// Export 导出mom指标数据表
func (s *sStockMomData) Export(ctx context.Context, in *stockin.MomDataListInp) (err error) {
	list, totalCount, err := s.List(ctx, in)
	if err != nil {
		return
	}

	// 字段的排序是依据tags的字段顺序，如果你不想使用默认的排序方式，可以直接定义 tags = []string{"字段名称", "字段名称2", ...}
	tags, err := convert.GetEntityDescTags(stockin.MomDataExportModel{})
	if err != nil {
		return
	}

	var (
		fileName  = "导出mom指标数据表-" + gctx.CtxId(ctx)
		sheetName = fmt.Sprintf("索引条件共%v行,共%v页,当前导出是第%v页,本页共%v行", totalCount, form.CalPageCount(totalCount, in.PerPage), in.Page, len(list))
		exports   []stockin.MomDataExportModel
	)

	if err = gconv.Scan(list, &exports); err != nil {
		return
	}

	err = excel.ExportByStructs(ctx, tags, exports, fileName, sheetName)
	return
}

// Edit 修改/新增mom指标数据表
func (s *sStockMomData) Edit(ctx context.Context, in *stockin.MomDataEditInp) (err error) {
	return g.DB().Transaction(ctx, func(ctx context.Context, tx gdb.TX) (err error) {

		// 修改
		if in.Id > 0 {
			if _, err = s.Model(ctx).
				Fields(stockin.MomDataUpdateFields{}).
				WherePri(in.Id).Data(in).Update(); err != nil {
				err = gerror.Wrap(err, "修改mom指标数据表失败，请稍后重试！")
			}
			return
		}

		// 新增
		if _, err = s.Model(ctx, &handler.Option{FilterAuth: false}).
			Fields(stockin.MomDataInsertFields{}).
			Data(in).OmitEmptyData().Insert(); err != nil {
			err = gerror.Wrap(err, "新增mom指标数据表失败，请稍后重试！")
		}
		return
	})
}

// Delete 删除mom指标数据表
func (s *sStockMomData) Delete(ctx context.Context, in *stockin.MomDataDeleteInp) (err error) {

	if _, err = s.Model(ctx).WherePri(in.Id).Unscoped().Delete(); err != nil {
		err = gerror.Wrap(err, "删除mom指标数据表失败，请稍后重试！")
		return
	}
	return
}

// View 获取mom指标数据表指定信息
func (s *sStockMomData) View(ctx context.Context, in *stockin.MomDataViewInp) (res *stockin.MomDataViewModel, err error) {
	if err = s.Model(ctx).WherePri(in.Id).Scan(&res); err != nil {
		err = gerror.Wrap(err, "获取mom指标数据表信息，请稍后重试！")
		return
	}
	return
}

// GetMom 获取mom指标
func (s *sStockMomData) GetMom(ctx context.Context, in *stockin.GetMomDataInp) (data []*entity.MomData, err error) {
	code := gstr.Split(in.Symbol, ".")[0]

	flag, err := s.Model(ctx).Where(dao.MacdData.Columns().T, GetRecentWeekday()).Where(dao.SlowStochasticData.Columns().Symbol, in.Symbol).Exist()
	if err != nil {
		return
	}
	if flag {
		return
	}

	// 目标URL
	url := fmt.Sprintf("https://finance.pae.baidu.com/sapi/v1/get_indicators_graph?code=%s&eventType=1005&finClientType=pc&financeType=stock&market=ab&period=dayK&finClientType=pc", code)

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
			mom := new(entity.MomData)
			mom.Symbol = in.Symbol
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
					mom.T = gtime.New(formatted)
				case 1:
					if gstr.Contains(c, "--") {
						continue LoopCci
					}
					mom.Mon = gconv.Float64(c)

				}
			}
			data = append(data, mom)
		}
		if len(data) > 0 {
			_, _ = s.Model(ctx).InsertIgnore(data)
		}
	}
	return

}
