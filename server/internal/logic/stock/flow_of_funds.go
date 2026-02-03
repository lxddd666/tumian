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

type sStockFlowOfFunds struct{}

func NewStockFlowOfFunds() *sStockFlowOfFunds {
	return &sStockFlowOfFunds{}
}

func init() {
	service.RegisterStockFlowOfFunds(NewStockFlowOfFunds())
}

// Model 资金流向明细表ORM模型
func (s *sStockFlowOfFunds) Model(ctx context.Context, option ...*handler.Option) *gdb.Model {
	return handler.Model(dao.FlowOfFunds.Ctx(ctx), option...)
}

// List 获取资金流向明细表列表
func (s *sStockFlowOfFunds) List(ctx context.Context, in *stockin.FlowOfFundsListInp) (list []*stockin.FlowOfFundsListModel, totalCount int, err error) {
	mod := s.Model(ctx)

	// 字段过滤
	mod = mod.Fields(stockin.FlowOfFundsListModel{})

	// 查询主键ID
	if in.Id > 0 {
		mod = mod.Where(dao.FlowOfFunds.Columns().Id, in.Id)
	}

	// 查询创建时间
	if len(in.CreatedAt) == 2 {
		mod = mod.WhereBetween(dao.FlowOfFunds.Columns().CreatedAt, in.CreatedAt[0], in.CreatedAt[1])
	}

	// 分页
	mod = mod.Page(in.Page, in.PerPage)

	// 排序
	mod = mod.OrderDesc(dao.FlowOfFunds.Columns().Id)

	// 查询数据
	if err = mod.ScanAndCount(&list, &totalCount, false); err != nil {
		err = gerror.Wrap(err, "获取资金流向明细表列表失败，请稍后重试！")
		return
	}
	return
}

// Export 导出资金流向明细表
func (s *sStockFlowOfFunds) Export(ctx context.Context, in *stockin.FlowOfFundsListInp) (err error) {
	list, totalCount, err := s.List(ctx, in)
	if err != nil {
		return
	}

	// 字段的排序是依据tags的字段顺序，如果你不想使用默认的排序方式，可以直接定义 tags = []string{"字段名称", "字段名称2", ...}
	tags, err := convert.GetEntityDescTags(stockin.FlowOfFundsExportModel{})
	if err != nil {
		return
	}

	var (
		fileName  = "导出资金流向明细表-" + gctx.CtxId(ctx)
		sheetName = fmt.Sprintf("索引条件共%v行,共%v页,当前导出是第%v页,本页共%v行", totalCount, form.CalPageCount(totalCount, in.PerPage), in.Page, len(list))
		exports   []stockin.FlowOfFundsExportModel
	)

	if err = gconv.Scan(list, &exports); err != nil {
		return
	}

	err = excel.ExportByStructs(ctx, tags, exports, fileName, sheetName)
	return
}

// Edit 修改/新增资金流向明细表
func (s *sStockFlowOfFunds) Edit(ctx context.Context, in *stockin.FlowOfFundsEditInp) (err error) {
	return g.DB().Transaction(ctx, func(ctx context.Context, tx gdb.TX) (err error) {

		// 修改
		if in.Id > 0 {
			if _, err = s.Model(ctx).
				Fields(stockin.FlowOfFundsUpdateFields{}).
				WherePri(in.Id).Data(in).Update(); err != nil {
				err = gerror.Wrap(err, "修改资金流向明细表失败，请稍后重试！")
			}
			return
		}

		// 新增
		if _, err = s.Model(ctx, &handler.Option{FilterAuth: false}).
			Fields(stockin.FlowOfFundsInsertFields{}).
			Data(in).OmitEmptyData().Insert(); err != nil {
			err = gerror.Wrap(err, "新增资金流向明细表失败，请稍后重试！")
		}
		return
	})
}

// Delete 删除资金流向明细表
func (s *sStockFlowOfFunds) Delete(ctx context.Context, in *stockin.FlowOfFundsDeleteInp) (err error) {

	if _, err = s.Model(ctx).WherePri(in.Id).Unscoped().Delete(); err != nil {
		err = gerror.Wrap(err, "删除资金流向明细表失败，请稍后重试！")
		return
	}
	return
}

// View 获取资金流向明细表指定信息
func (s *sStockFlowOfFunds) View(ctx context.Context, in *stockin.FlowOfFundsViewInp) (res *stockin.FlowOfFundsViewModel, err error) {
	if err = s.Model(ctx).WherePri(in.Id).Scan(&res); err != nil {
		err = gerror.Wrap(err, "获取资金流向明细表信息，请稍后重试！")
		return
	}
	return
}

// GetFlowOfFunds 获取资金流向明细表数据
func (s *sStockFlowOfFunds) GetFlowOfFunds(ctx context.Context, in *stockin.GetFlowOfFundsInp) (data []*entity.FlowOfFunds, err error) {
	flag, err := s.Model(ctx).Where(dao.FlowOfFunds.Columns().Symbol, in.Symbol).Where(dao.FlowOfFunds.Columns().T, GetRecentWeekday()).Exist()
	if err != nil {
		return
	}
	if flag {
		return
	}

	// 构建 API URL
	// https://api.zhituapi.com/hs/history/transaction/股票代码?token=token证书&st=开始时间&et=结束时间&lt=最新条数
	apiUrl := fmt.Sprintf("https://api.zhituapi.com/hs/history/transaction/%s", in.Symbol)

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

	// 发送 GET 请求并解析为 entity.FlowOfFunds

	var result []map[string]interface{}

	err = g.Client().GetVar(ctx, apiUrl, params).Scan(&result)
	if err != nil {
		err = gerror.Wrap(err, "调用外部API获取资金流向明细表数据失败，请稍后重试！")
		return
	}
	for _, re := range result {
		re["symbol"] = in.Symbol
	}
	// 处理数据，只保留年月日部分
	if len(result) > 0 {
		_, err = s.Model(ctx).InsertIgnore(result)
	}
	return
}
