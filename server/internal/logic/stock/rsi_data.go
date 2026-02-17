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
	"github.com/gogf/gf/v2/os/gtime"
	"github.com/gogf/gf/v2/text/gstr"
	"github.com/gogf/gf/v2/util/gconv"
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
	"math"
	"net/http"
	"time"
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
	flag, err := s.Model(ctx).Where(dao.MaData.Columns().T, GetRecentWeekday()).Where(dao.MaData.Columns().Symbol, in.Code).Exist()
	if err != nil {
		return
	}
	if flag {
		return
	}

	var codeList []*entity.StockSelfCode
	mod := dao.StockSelfCode.Ctx(ctx)
	if in.Code != "" {
		mod = mod.Where(dao.StockSelfCode.Columns().Dm, in.Code)
	}
	err = mod.Scan(&codeList)
	if err != nil {
		return
	}
	for _, stock := range codeList {
		rsiList := BaiduRsi(ctx, stock.Dm)
		_, _ = dao.RsiData.Ctx(ctx).InsertIgnore(rsiList)
	}

	return
}

func BaiduRsi(ctx context.Context, symbol string) (data []*entity.RsiData) {
	code := gstr.Split(symbol, ".")[0]

	// 目标URL
	url := fmt.Sprintf("https://finance.pae.baidu.com/sapi/v1/get_indicators_graph?code=%s&eventType=1003&finClientType=pc&financeType=stock&market=ab&period=dayK&finClientType=pc", code)

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
			cci := new(entity.RsiData)
			cci.Symbol = symbol
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
					cci.T = gtime.New(formatted)
				case 1:
					if gstr.Contains(c, "--") {
						continue LoopCci
					}
					cci.Rsi6 = gconv.Float64(c)

				case 2:
					if gstr.Contains(c, "--") {
						continue LoopCci
					}
					cci.Rsi14 = gconv.Float64(c)
					cci.Rsi = gconv.Float64(c)
				case 3:
					if gstr.Contains(c, "--") {
						continue LoopCci
					}
					cci.Rsi24 = gconv.Float64(c)
				}
			}
			data = append(data, cci)
		}
	}
	return
}

// CalculateRSIFromPine 严格按照 TradingView Pine Script v6 的逻辑实现
// 参数: data 必须按时间升序排列
func CalculateRSIFromPine(ctx context.Context, code string, period int) []*entity.RsiData {
	var data []*entity.EnterpriseHistoricalData
	_ = service.StockEnterpriseHistoricalData().Model(ctx).Where(dao.EnterpriseHistoricalData.Columns().Symbol, code).OrderAsc(dao.EnterpriseHistoricalData.Columns().T).Scan(&data)

	var results []*entity.RsiData
	n := len(data)

	// 检查数据长度
	if n == 0 {
		return results
	}

	// 初始化变量
	var avgGain, avgLoss float64
	var rsiValue float64

	for i := 0; i < n; i++ {
		var change, gain, loss float64

		if i == 0 {
			// 第一根K线没有变化
			change = 0
		} else {
			// 计算涨跌幅度
			change = data[i].C - data[i-1].C
		}

		// 分离涨跌幅
		if change > 0 {
			gain = change
			loss = 0
		} else {
			gain = 0
			loss = -change // 注意取正值
		}

		// 计算逻辑分叉
		if i < period {
			// --- 初始化阶段 (前 period 根K线) ---
			// 累加涨跌幅
			sumGain, sumLoss := accumulateSums(data, i)
			// 计算简单平均值 (SMA)
			avgGain = sumGain / float64(period)
			avgLoss = sumLoss / float64(period)
		} else {
			// --- 迭代阶段 (Wilder's 平滑 RMA) ---
			// 公式: (prevAvg * (n-1) + current) / n
			avgGain = (avgGain*float64(period-1) + gain) / float64(period)
			avgLoss = (avgLoss*float64(period-1) + loss) / float64(period)
		}

		// --- 计算 RSI 值 ---
		// 对应脚本逻辑: down == 0 ? 100 : up == 0 ? 0 : 100 - (100 / (1 + up / down))
		if avgLoss == 0 {
			rsiValue = 100.0
		} else if avgGain == 0 {
			rsiValue = 0.0
		} else {
			rs := avgGain / avgLoss
			rsiValue = 100 - (100 / (1 + rs))
		}

		// 保留两位小数 (对应脚本中的 precision=2)
		rsiValue = math.Round(rsiValue*100) / 100

		// 创建结果对象
		// 注意：TradingView 的 RSI 在第 14 根K线（索引13）开始出值
		// 但为了对齐时间，通常索引 i 对应 data[i] 的时间
		// 这里我们从第 'period' 根K线开始记录 (即 i >= period-1)
		if i >= period-1 {
			results = append(results, &entity.RsiData{
				Symbol: data[i].Symbol,
				T:      data[i].T,
				Rsi:    rsiValue,
			})
		}
	}

	return results
}

// accumulateSums 辅助函数：计算从开始到索引 i 的总涨幅和总跌幅 (用于初始化)
func accumulateSums(data []*entity.EnterpriseHistoricalData, endIndex int) (float64, float64) {
	var sumGain, sumLoss float64
	// 从索引 1 开始计算变化 (因为索引 0 没有前值)
	for j := 1; j <= endIndex; j++ {
		change := data[j].C - data[j-1].C
		if change > 0 {
			sumGain += change
		} else {
			sumLoss -= change // 减去负数，相当于加正数
		}
	}
	return sumGain, sumLoss
}
