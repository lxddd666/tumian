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

type sStockIncomeStatement struct{}

func NewStockIncomeStatement() *sStockIncomeStatement {
	return &sStockIncomeStatement{}
}

func init() {
	service.RegisterStockIncomeStatement(NewStockIncomeStatement())
}

// Model 利润表 (Income Statement)ORM模型
func (s *sStockIncomeStatement) Model(ctx context.Context, option ...*handler.Option) *gdb.Model {
	return handler.Model(dao.IncomeStatement.Ctx(ctx), option...)
}

// List 获取利润表 (Income Statement)列表
func (s *sStockIncomeStatement) List(ctx context.Context, in *stockin.IncomeStatementListInp) (list []*stockin.IncomeStatementListModel, totalCount int, err error) {
	mod := s.Model(ctx)

	// 字段过滤
	mod = mod.Fields(stockin.IncomeStatementListModel{})

	// 查询自增主键
	if in.Id > 0 {
		mod = mod.Where(dao.IncomeStatement.Columns().Id, in.Id)
	}

	// 查询创建时间
	if len(in.CreatedAt) == 2 {
		mod = mod.WhereBetween(dao.IncomeStatement.Columns().CreatedAt, in.CreatedAt[0], in.CreatedAt[1])
	}

	// 分页
	mod = mod.Page(in.Page, in.PerPage)

	// 排序
	mod = mod.OrderDesc(dao.IncomeStatement.Columns().Id)

	// 查询数据
	if err = mod.ScanAndCount(&list, &totalCount, false); err != nil {
		err = gerror.Wrap(err, "获取利润表 (Income Statement)列表失败，请稍后重试！")
		return
	}
	return
}

// Export 导出利润表 (Income Statement)
func (s *sStockIncomeStatement) Export(ctx context.Context, in *stockin.IncomeStatementListInp) (err error) {
	list, totalCount, err := s.List(ctx, in)
	if err != nil {
		return
	}

	// 字段的排序是依据tags的字段顺序，如果你不想使用默认的排序方式，可以直接定义 tags = []string{"字段名称", "字段名称2", ...}
	tags, err := convert.GetEntityDescTags(stockin.IncomeStatementExportModel{})
	if err != nil {
		return
	}

	var (
		fileName  = "导出利润表 (Income Statement)-" + gctx.CtxId(ctx)
		sheetName = fmt.Sprintf("索引条件共%v行,共%v页,当前导出是第%v页,本页共%v行", totalCount, form.CalPageCount(totalCount, in.PerPage), in.Page, len(list))
		exports   []stockin.IncomeStatementExportModel
	)

	if err = gconv.Scan(list, &exports); err != nil {
		return
	}

	err = excel.ExportByStructs(ctx, tags, exports, fileName, sheetName)
	return
}

// Edit 修改/新增利润表 (Income Statement)
func (s *sStockIncomeStatement) Edit(ctx context.Context, in *stockin.IncomeStatementEditInp) (err error) {
	return g.DB().Transaction(ctx, func(ctx context.Context, tx gdb.TX) (err error) {

		// 修改
		if in.Id > 0 {
			if _, err = s.Model(ctx).
				Fields(stockin.IncomeStatementUpdateFields{}).
				WherePri(in.Id).Data(in).Update(); err != nil {
				err = gerror.Wrap(err, "修改利润表 (Income Statement)失败，请稍后重试！")
			}
			return
		}

		// 新增
		if _, err = s.Model(ctx, &handler.Option{FilterAuth: false}).
			Fields(stockin.IncomeStatementInsertFields{}).
			Data(in).OmitEmptyData().Insert(); err != nil {
			err = gerror.Wrap(err, "新增利润表 (Income Statement)失败，请稍后重试！")
		}
		return
	})
}

// Delete 删除利润表 (Income Statement)
func (s *sStockIncomeStatement) Delete(ctx context.Context, in *stockin.IncomeStatementDeleteInp) (err error) {

	if _, err = s.Model(ctx).WherePri(in.Id).Unscoped().Delete(); err != nil {
		err = gerror.Wrap(err, "删除利润表 (Income Statement)失败，请稍后重试！")
		return
	}
	return
}

// View 获取利润表 (Income Statement)指定信息
func (s *sStockIncomeStatement) View(ctx context.Context, in *stockin.IncomeStatementViewInp) (res *stockin.IncomeStatementViewModel, err error) {
	if err = s.Model(ctx).WherePri(in.Id).Scan(&res); err != nil {
		err = gerror.Wrap(err, "获取利润表 (Income Statement)信息，请稍后重试！")
		return
	}
	return
}

// GetIncomeStatement 获取利润表数据
func (s *sStockIncomeStatement) GetIncomeStatement(ctx context.Context, in *stockin.IncomeStatementGetIncomeStatementInp) (data []*entity.IncomeStatement, err error) {
	//flag, err := s.Model(ctx).Where(dao.IncomeStatement.Columns().T, GetNowDate()).Exist()
	//if err != nil {
	//	return
	//}
	//if flag {
	//	return
	//}

	// 构建 API URL
	// https://api.zhituapi.com/hs/fin/income/股票代码（如000001.SZ）?token=token证书&st=开始时间&et=结束时间
	apiUrl := fmt.Sprintf("https://api.zhituapi.com/hs/fin/income/%s", in.Symbol)

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

	// 发送 GET 请求并解析为 entity.IncomeStatement
	var result []*entity.IncomeStatement
	var resultMap []map[string]interface{}

	err = g.Client().GetVar(ctx, apiUrl, params).Scan(&resultMap)
	if err != nil {
		err = gerror.Wrap(err, "调用外部API获取利润表数据失败，请稍后重试！")
		return
	}

	data = result
	for _, re := range resultMap {
		re["symbol"] = in.Symbol
	}
	if len(resultMap) > 0 {
		_, err = s.Model(ctx).InsertIgnore(resultMap)
	}
	return
}
