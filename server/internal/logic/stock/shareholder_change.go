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

type sStockShareholderChange struct{}

func NewStockShareholderChange() *sStockShareholderChange {
	return &sStockShareholderChange{}
}

func init() {
	service.RegisterStockShareholderChange(NewStockShareholderChange())
}

// Model 股东户数变化记录表 (记录相邻报告期的户数变化)ORM模型
func (s *sStockShareholderChange) Model(ctx context.Context, option ...*handler.Option) *gdb.Model {
	return handler.Model(dao.ShareholderChange.Ctx(ctx), option...)
}

// List 获取股东户数变化记录表 (记录相邻报告期的户数变化)列表
func (s *sStockShareholderChange) List(ctx context.Context, in *stockin.ShareholderChangeListInp) (list []*stockin.ShareholderChangeListModel, totalCount int, err error) {
	mod := s.Model(ctx)

	// 字段过滤
	mod = mod.Fields(stockin.ShareholderChangeListModel{})

	// 查询自增主键
	if in.Id > 0 {
		mod = mod.Where(dao.ShareholderChange.Columns().Id, in.Id)
	}

	// 查询创建时间
	if len(in.CreatedAt) == 2 {
		mod = mod.WhereBetween(dao.ShareholderChange.Columns().CreatedAt, in.CreatedAt[0], in.CreatedAt[1])
	}

	// 分页
	mod = mod.Page(in.Page, in.PerPage)

	// 排序
	mod = mod.OrderDesc(dao.ShareholderChange.Columns().Id)

	// 查询数据
	if err = mod.ScanAndCount(&list, &totalCount, false); err != nil {
		err = gerror.Wrap(err, "获取股东户数变化记录表 (记录相邻报告期的户数变化)列表失败，请稍后重试！")
		return
	}
	return
}

// Export 导出股东户数变化记录表 (记录相邻报告期的户数变化)
func (s *sStockShareholderChange) Export(ctx context.Context, in *stockin.ShareholderChangeListInp) (err error) {
	list, totalCount, err := s.List(ctx, in)
	if err != nil {
		return
	}

	// 字段的排序是依据tags的字段顺序，如果你不想使用默认的排序方式，可以直接定义 tags = []string{"字段名称", "字段名称2", ...}
	tags, err := convert.GetEntityDescTags(stockin.ShareholderChangeExportModel{})
	if err != nil {
		return
	}

	var (
		fileName  = "导出股东户数变化记录表 (记录相邻报告期的户数变化)-" + gctx.CtxId(ctx)
		sheetName = fmt.Sprintf("索引条件共%v行,共%v页,当前导出是第%v页,本页共%v行", totalCount, form.CalPageCount(totalCount, in.PerPage), in.Page, len(list))
		exports   []stockin.ShareholderChangeExportModel
	)

	if err = gconv.Scan(list, &exports); err != nil {
		return
	}

	err = excel.ExportByStructs(ctx, tags, exports, fileName, sheetName)
	return
}

// Edit 修改/新增股东户数变化记录表 (记录相邻报告期的户数变化)
func (s *sStockShareholderChange) Edit(ctx context.Context, in *stockin.ShareholderChangeEditInp) (err error) {
	return g.DB().Transaction(ctx, func(ctx context.Context, tx gdb.TX) (err error) {

		// 修改
		if in.Id > 0 {
			if _, err = s.Model(ctx).
				Fields(stockin.ShareholderChangeUpdateFields{}).
				WherePri(in.Id).Data(in).Update(); err != nil {
				err = gerror.Wrap(err, "修改股东户数变化记录表 (记录相邻报告期的户数变化)失败，请稍后重试！")
			}
			return
		}

		// 新增
		if _, err = s.Model(ctx, &handler.Option{FilterAuth: false}).
			Fields(stockin.ShareholderChangeInsertFields{}).
			Data(in).OmitEmptyData().Insert(); err != nil {
			err = gerror.Wrap(err, "新增股东户数变化记录表 (记录相邻报告期的户数变化)失败，请稍后重试！")
		}
		return
	})
}

// Delete 删除股东户数变化记录表 (记录相邻报告期的户数变化)
func (s *sStockShareholderChange) Delete(ctx context.Context, in *stockin.ShareholderChangeDeleteInp) (err error) {

	if _, err = s.Model(ctx).WherePri(in.Id).Unscoped().Delete(); err != nil {
		err = gerror.Wrap(err, "删除股东户数变化记录表 (记录相邻报告期的户数变化)失败，请稍后重试！")
		return
	}
	return
}

// View 获取股东户数变化记录表 (记录相邻报告期的户数变化)指定信息
func (s *sStockShareholderChange) View(ctx context.Context, in *stockin.ShareholderChangeViewInp) (res *stockin.ShareholderChangeViewModel, err error) {
	if err = s.Model(ctx).WherePri(in.Id).Scan(&res); err != nil {
		err = gerror.Wrap(err, "获取股东户数变化记录表 (记录相邻报告期的户数变化)信息，请稍后重试！")
		return
	}
	return
}

// GetShareholderChange 获取股东户数变化记录表数据
func (s *sStockShareholderChange) GetShareholderChange(ctx context.Context, in *stockin.ShareholderChangeGetShareholderChangeInp) (data []*entity.ShareholderChange, err error) {
	//flag, err := s.Model(ctx).Where(dao.ShareholderChange.Columns().T, GetNowDate()).Exist()
	//if err != nil {
	//	return
	//}
	//if flag {
	//	return
	//}

	// 构建 API URL
	// https://api.zhituapi.com/hs/gs/gdbh/股票代码?token=token证书
	apiUrl := fmt.Sprintf("https://api.zhituapi.com/hs/gs/gdbh/%s", in.Symbol)

	// 构建查询参数
	params := g.Map{
		"token": global.StockToken,
	}

	// 发送 GET 请求并解析为 entity.ShareholderChange
	var result []*entity.ShareholderChange
	err = g.Client().GetVar(ctx, apiUrl, params).Scan(&result)
	if err != nil {
		err = gerror.Wrap(err, "调用外部API获取股东户数变化记录表数据失败，请稍后重试！")
		return
	}

	data = result
	for _, re := range result {
		re.Symbol = in.Symbol
		dateStr := re.Jzrq.Format("Y-m-d")
		re.Jzrq = gtime.NewFromStr(dateStr)
	}
	data = result
	if len(result) > 0 {
		_, err = s.Model(ctx).InsertIgnore(result)
	}
	return
}
