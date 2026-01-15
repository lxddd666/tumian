// ================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// You can delete these comments if you wish manually maintain this interface file.
// ================================================================================

package service

import (
	"context"
	"hotgo/internal/library/hgorm/handler"
	"hotgo/internal/model/entity"
	"hotgo/internal/model/input/stockin"

	"github.com/gogf/gf/v2/database/gdb"
)

type (
	IStockBollData interface {
		// Model 布林带(BOLL)指标数据表ORM模型
		Model(ctx context.Context, option ...*handler.Option) *gdb.Model
		// List 获取布林带(BOLL)指标数据表列表
		List(ctx context.Context, in *stockin.BollDataListInp) (list []*stockin.BollDataListModel, totalCount int, err error)
		// Export 导出布林带(BOLL)指标数据表
		Export(ctx context.Context, in *stockin.BollDataListInp) (err error)
		// Edit 修改/新增布林带(BOLL)指标数据表
		Edit(ctx context.Context, in *stockin.BollDataEditInp) (err error)
		// Delete 删除布林带(BOLL)指标数据表
		Delete(ctx context.Context, in *stockin.BollDataDeleteInp) (err error)
		// View 获取布林带(BOLL)指标数据表指定信息
		View(ctx context.Context, in *stockin.BollDataViewInp) (res *stockin.BollDataViewModel, err error)
		// GetBoll 获取布林带(BOLL)指标数据
		GetBoll(ctx context.Context, in *stockin.BollDataGetBollInp) (data *entity.BollData, err error)
	}
	IStockEnterpriseHistoricalData interface {
		// Model 企业级历史行情数据表 (K线数据)ORM模型
		Model(ctx context.Context, option ...*handler.Option) *gdb.Model
		// List 获取企业级历史行情数据表 (K线数据)列表
		List(ctx context.Context, in *stockin.EnterpriseHistoricalDataListInp) (list []*stockin.EnterpriseHistoricalDataListModel, totalCount int, err error)
		// Export 导出企业级历史行情数据表 (K线数据)
		Export(ctx context.Context, in *stockin.EnterpriseHistoricalDataListInp) (err error)
		// Edit 修改/新增企业级历史行情数据表 (K线数据)
		Edit(ctx context.Context, in *stockin.EnterpriseHistoricalDataEditInp) (err error)
		// Delete 删除企业级历史行情数据表 (K线数据)
		Delete(ctx context.Context, in *stockin.EnterpriseHistoricalDataDeleteInp) (err error)
		// View 获取企业级历史行情数据表 (K线数据)指定信息
		View(ctx context.Context, in *stockin.EnterpriseHistoricalDataViewInp) (res *stockin.EnterpriseHistoricalDataViewModel, err error)
		// GetEnterpriseHistoricalData 获取企业级历史行情数据表数据
		GetEnterpriseHistoricalData(ctx context.Context, in *stockin.EnterpriseHistoricalDataGetEnterpriseHistoricalDataInp) (data *entity.EnterpriseHistoricalData, err error)
	}
	IStockFinancialIndicators interface {
		// Model 财务指标分析表ORM模型
		Model(ctx context.Context, option ...*handler.Option) *gdb.Model
		// List 获取财务指标分析表列表
		List(ctx context.Context, in *stockin.FinancialIndicatorsListInp) (list []*stockin.FinancialIndicatorsListModel, totalCount int, err error)
		// Export 导出财务指标分析表
		Export(ctx context.Context, in *stockin.FinancialIndicatorsListInp) (err error)
		// Edit 修改/新增财务指标分析表
		Edit(ctx context.Context, in *stockin.FinancialIndicatorsEditInp) (err error)
		// Delete 删除财务指标分析表
		Delete(ctx context.Context, in *stockin.FinancialIndicatorsDeleteInp) (err error)
		// View 获取财务指标分析表指定信息
		View(ctx context.Context, in *stockin.FinancialIndicatorsViewInp) (res *stockin.FinancialIndicatorsViewModel, err error)
		// GetFinancialIndicators 获取财务指标分析表数据
		GetFinancialIndicators(ctx context.Context, in *stockin.FinancialIndicatorsGetFinancialIndicatorsInp) (data *entity.FinancialIndicators, err error)
	}
	IStockFundStockHolding interface {
		// Model 基金持股明细表 (来源于基金定期报告)ORM模型
		Model(ctx context.Context, option ...*handler.Option) *gdb.Model
		// List 获取基金持股明细表 (来源于基金定期报告)列表
		List(ctx context.Context, in *stockin.FundStockHoldingListInp) (list []*stockin.FundStockHoldingListModel, totalCount int, err error)
		// Export 导出基金持股明细表 (来源于基金定期报告)
		Export(ctx context.Context, in *stockin.FundStockHoldingListInp) (err error)
		// Edit 修改/新增基金持股明细表 (来源于基金定期报告)
		Edit(ctx context.Context, in *stockin.FundStockHoldingEditInp) (err error)
		// Delete 删除基金持股明细表 (来源于基金定期报告)
		Delete(ctx context.Context, in *stockin.FundStockHoldingDeleteInp) (err error)
		// View 获取基金持股明细表 (来源于基金定期报告)指定信息
		View(ctx context.Context, in *stockin.FundStockHoldingViewInp) (res *stockin.FundStockHoldingViewModel, err error)
		// GetFundStockHolding 获取基金持股明细表数据
		GetFundStockHolding(ctx context.Context, in *stockin.FundStockHoldingGetFundStockHoldingInp) (data *entity.FundStockHolding, err error)
	}
	IStockIncomeStatement interface {
		// Model 利润表 (Income Statement)ORM模型
		Model(ctx context.Context, option ...*handler.Option) *gdb.Model
		// List 获取利润表 (Income Statement)列表
		List(ctx context.Context, in *stockin.IncomeStatementListInp) (list []*stockin.IncomeStatementListModel, totalCount int, err error)
		// Export 导出利润表 (Income Statement)
		Export(ctx context.Context, in *stockin.IncomeStatementListInp) (err error)
		// Edit 修改/新增利润表 (Income Statement)
		Edit(ctx context.Context, in *stockin.IncomeStatementEditInp) (err error)
		// Delete 删除利润表 (Income Statement)
		Delete(ctx context.Context, in *stockin.IncomeStatementDeleteInp) (err error)
		// View 获取利润表 (Income Statement)指定信息
		View(ctx context.Context, in *stockin.IncomeStatementViewInp) (res *stockin.IncomeStatementViewModel, err error)
		// GetIncomeStatement 获取利润表数据
		GetIncomeStatement(ctx context.Context, in *stockin.IncomeStatementGetIncomeStatementInp) (data *entity.IncomeStatement, err error)
	}
	IStockKdjData interface {
		// Model KDJ随机指标数据表ORM模型
		Model(ctx context.Context, option ...*handler.Option) *gdb.Model
		// List 获取KDJ随机指标数据表列表
		List(ctx context.Context, in *stockin.KdjDataListInp) (list []*stockin.KdjDataListModel, totalCount int, err error)
		// Export 导出KDJ随机指标数据表
		Export(ctx context.Context, in *stockin.KdjDataListInp) (err error)
		// Edit 修改/新增KDJ随机指标数据表
		Edit(ctx context.Context, in *stockin.KdjDataEditInp) (err error)
		// Delete 删除KDJ随机指标数据表
		Delete(ctx context.Context, in *stockin.KdjDataDeleteInp) (err error)
		// View 获取KDJ随机指标数据表指定信息
		View(ctx context.Context, in *stockin.KdjDataViewInp) (res *stockin.KdjDataViewModel, err error)
		// GetKdj 获取KDJ随机指标数据
		GetKdj(ctx context.Context, in *stockin.KdjDataGetKdjInp) (data *entity.KdjData, err error)
	}
	IStockMaData interface {
		// Model 移动平均线(MA)指标数据表ORM模型
		Model(ctx context.Context, option ...*handler.Option) *gdb.Model
		// List 获取移动平均线(MA)指标数据表列表
		List(ctx context.Context, in *stockin.MaDataListInp) (list []*stockin.MaDataListModel, totalCount int, err error)
		// Export 导出移动平均线(MA)指标数据表
		Export(ctx context.Context, in *stockin.MaDataListInp) (err error)
		// Edit 修改/新增移动平均线(MA)指标数据表
		Edit(ctx context.Context, in *stockin.MaDataEditInp) (err error)
		// Delete 删除移动平均线(MA)指标数据表
		Delete(ctx context.Context, in *stockin.MaDataDeleteInp) (err error)
		// View 获取移动平均线(MA)指标数据表指定信息
		View(ctx context.Context, in *stockin.MaDataViewInp) (res *stockin.MaDataViewModel, err error)
		// GetMa 获取移动平均线(MA)指标数据
		GetMa(ctx context.Context, in *stockin.MaDataGetMaInp) (data *entity.MaData, err error)
	}
	IStockMacdData interface {
		// Model MACD指标数据表ORM模型
		Model(ctx context.Context, option ...*handler.Option) *gdb.Model
		// List 获取MACD指标数据表列表
		List(ctx context.Context, in *stockin.MacdDataListInp) (list []*stockin.MacdDataListModel, totalCount int, err error)
		// Export 导出MACD指标数据表
		Export(ctx context.Context, in *stockin.MacdDataListInp) (err error)
		// Edit 修改/新增MACD指标数据表
		Edit(ctx context.Context, in *stockin.MacdDataEditInp) (err error)
		// Delete 删除MACD指标数据表
		Delete(ctx context.Context, in *stockin.MacdDataDeleteInp) (err error)
		// View 获取MACD指标数据表指定信息
		View(ctx context.Context, in *stockin.MacdDataViewInp) (res *stockin.MacdDataViewModel, err error)
		// GetMacd 获取MACD指标数据
		GetMacd(ctx context.Context, in *stockin.MacdDataGetMacdInp) (data *entity.MacdData, err error)
	}
	IStockQuarterlyProfit interface {
		// Model 季度利润数据表 (近一年各季度)ORM模型
		Model(ctx context.Context, option ...*handler.Option) *gdb.Model
		// List 获取季度利润数据表 (近一年各季度)列表
		List(ctx context.Context, in *stockin.QuarterlyProfitListInp) (list []*stockin.QuarterlyProfitListModel, totalCount int, err error)
		// Export 导出季度利润数据表 (近一年各季度)
		Export(ctx context.Context, in *stockin.QuarterlyProfitListInp) (err error)
		// Edit 修改/新增季度利润数据表 (近一年各季度)
		Edit(ctx context.Context, in *stockin.QuarterlyProfitEditInp) (err error)
		// Delete 删除季度利润数据表 (近一年各季度)
		Delete(ctx context.Context, in *stockin.QuarterlyProfitDeleteInp) (err error)
		// View 获取季度利润数据表 (近一年各季度)指定信息
		View(ctx context.Context, in *stockin.QuarterlyProfitViewInp) (res *stockin.QuarterlyProfitViewModel, err error)
		// GetQuarterlyProfit 获取季度利润数据表数据
		GetQuarterlyProfit(ctx context.Context, in *stockin.QuarterlyProfitGetQuarterlyProfitInp) (data *entity.QuarterlyProfit, err error)
	}
	IStockShareholderChange interface {
		// Model 股东户数变化记录表 (记录相邻报告期的户数变化)ORM模型
		Model(ctx context.Context, option ...*handler.Option) *gdb.Model
		// List 获取股东户数变化记录表 (记录相邻报告期的户数变化)列表
		List(ctx context.Context, in *stockin.ShareholderChangeListInp) (list []*stockin.ShareholderChangeListModel, totalCount int, err error)
		// Export 导出股东户数变化记录表 (记录相邻报告期的户数变化)
		Export(ctx context.Context, in *stockin.ShareholderChangeListInp) (err error)
		// Edit 修改/新增股东户数变化记录表 (记录相邻报告期的户数变化)
		Edit(ctx context.Context, in *stockin.ShareholderChangeEditInp) (err error)
		// Delete 删除股东户数变化记录表 (记录相邻报告期的户数变化)
		Delete(ctx context.Context, in *stockin.ShareholderChangeDeleteInp) (err error)
		// View 获取股东户数变化记录表 (记录相邻报告期的户数变化)指定信息
		View(ctx context.Context, in *stockin.ShareholderChangeViewInp) (res *stockin.ShareholderChangeViewModel, err error)
		// GetShareholderChange 获取股东户数变化记录表数据
		GetShareholderChange(ctx context.Context, in *stockin.ShareholderChangeGetShareholderChangeInp) (data *entity.ShareholderChange, err error)
	}
	IStockShareholderCount interface {
		// Model 公司股东户数统计表 (按报告期统计)ORM模型
		Model(ctx context.Context, option ...*handler.Option) *gdb.Model
		// List 获取公司股东户数统计表 (按报告期统计)列表
		List(ctx context.Context, in *stockin.ShareholderCountListInp) (list []*stockin.ShareholderCountListModel, totalCount int, err error)
		// Export 导出公司股东户数统计表 (按报告期统计)
		Export(ctx context.Context, in *stockin.ShareholderCountListInp) (err error)
		// Edit 修改/新增公司股东户数统计表 (按报告期统计)
		Edit(ctx context.Context, in *stockin.ShareholderCountEditInp) (err error)
		// Delete 删除公司股东户数统计表 (按报告期统计)
		Delete(ctx context.Context, in *stockin.ShareholderCountDeleteInp) (err error)
		// View 获取公司股东户数统计表 (按报告期统计)指定信息
		View(ctx context.Context, in *stockin.ShareholderCountViewInp) (res *stockin.ShareholderCountViewModel, err error)
		// GetShareholderCount 获取公司股东户数统计表数据
		GetShareholderCount(ctx context.Context, in *stockin.ShareholderCountGetShareholderCountInp) (data *entity.ShareholderCount, err error)
	}
	IStockBasicInfo interface {
		// Model 股票基础信息表ORM模型
		Model(ctx context.Context, option ...*handler.Option) *gdb.Model
		// List 获取股票基础信息表列表
		List(ctx context.Context, in *stockin.StockBasicInfoListInp) (list []*stockin.StockBasicInfoListModel, totalCount int, err error)
		// Export 导出股票基础信息表
		Export(ctx context.Context, in *stockin.StockBasicInfoListInp) (err error)
		// Edit 修改/新增股票基础信息表
		Edit(ctx context.Context, in *stockin.StockBasicInfoEditInp) (err error)
		// Delete 删除股票基础信息表
		Delete(ctx context.Context, in *stockin.StockBasicInfoDeleteInp) (err error)
		// View 获取股票基础信息表指定信息
		View(ctx context.Context, in *stockin.StockBasicInfoViewInp) (res *stockin.StockBasicInfoViewModel, err error)
		// GetStockBasicInfo 获取股票基础信息表数据
		GetStockBasicInfo(ctx context.Context, in *stockin.StockBasicInfoGetStockBasicInfoInp) (data interface{}, err error)
	}
	IStockList interface {
		// Model 股票列表核心表ORM模型
		Model(ctx context.Context, option ...*handler.Option) *gdb.Model
		// List 获取股票列表核心表列表
		List(ctx context.Context, in *stockin.StockListListInp) (list []*stockin.StockListListModel, totalCount int, err error)
		// Export 导出股票列表核心表
		Export(ctx context.Context, in *stockin.StockListListInp) (err error)
		// Edit 修改/新增股票列表核心表
		Edit(ctx context.Context, in *stockin.StockListEditInp) (err error)
		// Delete 删除股票列表核心表
		Delete(ctx context.Context, in *stockin.StockListDeleteInp) (err error)
		// View 获取股票列表核心表指定信息
		View(ctx context.Context, in *stockin.StockListViewInp) (res *stockin.StockListViewModel, err error)
		// Status 更新股票列表核心表状态
		Status(ctx context.Context, in *stockin.StockListStatusInp) (err error)
		// GetStockList 获取股票列表核心表数据
		GetStockList(ctx context.Context, in *stockin.StockListGetStockListInp) (data *entity.StockList, err error)
	}
	IStockTopTenCirculatingHolders interface {
		// Model 公司十大流通股东表 (数据来源于定期报告)[citation:4]ORM模型
		Model(ctx context.Context, option ...*handler.Option) *gdb.Model
		// List 获取公司十大流通股东表 (数据来源于定期报告)[citation:4]列表
		List(ctx context.Context, in *stockin.TopTenCirculatingHoldersListInp) (list []*stockin.TopTenCirculatingHoldersListModel, totalCount int, err error)
		// Export 导出公司十大流通股东表 (数据来源于定期报告)[citation:4]
		Export(ctx context.Context, in *stockin.TopTenCirculatingHoldersListInp) (err error)
		// Edit 修改/新增公司十大流通股东表 (数据来源于定期报告)[citation:4]
		Edit(ctx context.Context, in *stockin.TopTenCirculatingHoldersEditInp) (err error)
		// Delete 删除公司十大流通股东表 (数据来源于定期报告)[citation:4]
		Delete(ctx context.Context, in *stockin.TopTenCirculatingHoldersDeleteInp) (err error)
		// View 获取公司十大流通股东表 (数据来源于定期报告)[citation:4]指定信息
		View(ctx context.Context, in *stockin.TopTenCirculatingHoldersViewInp) (res *stockin.TopTenCirculatingHoldersViewModel, err error)
		// GetTopTenCirculatingHolders 获取公司十大流通股东表数据
		GetTopTenCirculatingHolders(ctx context.Context, in *stockin.TopTenCirculatingHoldersGetTopTenCirculatingHoldersInp) (data *entity.TopTenCirculatingHolders, err error)
	}
)

var (
	localStockBollData                 IStockBollData
	localStockEnterpriseHistoricalData IStockEnterpriseHistoricalData
	localStockFinancialIndicators      IStockFinancialIndicators
	localStockFundStockHolding         IStockFundStockHolding
	localStockIncomeStatement          IStockIncomeStatement
	localStockKdjData                  IStockKdjData
	localStockMaData                   IStockMaData
	localStockMacdData                 IStockMacdData
	localStockQuarterlyProfit          IStockQuarterlyProfit
	localStockShareholderChange        IStockShareholderChange
	localStockShareholderCount         IStockShareholderCount
	localStockBasicInfo                IStockBasicInfo
	localStockList                     IStockList
	localStockTopTenCirculatingHolders IStockTopTenCirculatingHolders
)

func StockBollData() IStockBollData {
	if localStockBollData == nil {
		panic("implement not found for interface IStockBollData, forgot register?")
	}
	return localStockBollData
}

func RegisterStockBollData(i IStockBollData) {
	localStockBollData = i
}

func StockEnterpriseHistoricalData() IStockEnterpriseHistoricalData {
	if localStockEnterpriseHistoricalData == nil {
		panic("implement not found for interface IStockEnterpriseHistoricalData, forgot register?")
	}
	return localStockEnterpriseHistoricalData
}

func RegisterStockEnterpriseHistoricalData(i IStockEnterpriseHistoricalData) {
	localStockEnterpriseHistoricalData = i
}

func StockFinancialIndicators() IStockFinancialIndicators {
	if localStockFinancialIndicators == nil {
		panic("implement not found for interface IStockFinancialIndicators, forgot register?")
	}
	return localStockFinancialIndicators
}

func RegisterStockFinancialIndicators(i IStockFinancialIndicators) {
	localStockFinancialIndicators = i
}

func StockFundStockHolding() IStockFundStockHolding {
	if localStockFundStockHolding == nil {
		panic("implement not found for interface IStockFundStockHolding, forgot register?")
	}
	return localStockFundStockHolding
}

func RegisterStockFundStockHolding(i IStockFundStockHolding) {
	localStockFundStockHolding = i
}

func StockIncomeStatement() IStockIncomeStatement {
	if localStockIncomeStatement == nil {
		panic("implement not found for interface IStockIncomeStatement, forgot register?")
	}
	return localStockIncomeStatement
}

func RegisterStockIncomeStatement(i IStockIncomeStatement) {
	localStockIncomeStatement = i
}

func StockKdjData() IStockKdjData {
	if localStockKdjData == nil {
		panic("implement not found for interface IStockKdjData, forgot register?")
	}
	return localStockKdjData
}

func RegisterStockKdjData(i IStockKdjData) {
	localStockKdjData = i
}

func StockMaData() IStockMaData {
	if localStockMaData == nil {
		panic("implement not found for interface IStockMaData, forgot register?")
	}
	return localStockMaData
}

func RegisterStockMaData(i IStockMaData) {
	localStockMaData = i
}

func StockMacdData() IStockMacdData {
	if localStockMacdData == nil {
		panic("implement not found for interface IStockMacdData, forgot register?")
	}
	return localStockMacdData
}

func RegisterStockMacdData(i IStockMacdData) {
	localStockMacdData = i
}

func StockQuarterlyProfit() IStockQuarterlyProfit {
	if localStockQuarterlyProfit == nil {
		panic("implement not found for interface IStockQuarterlyProfit, forgot register?")
	}
	return localStockQuarterlyProfit
}

func RegisterStockQuarterlyProfit(i IStockQuarterlyProfit) {
	localStockQuarterlyProfit = i
}

func StockShareholderChange() IStockShareholderChange {
	if localStockShareholderChange == nil {
		panic("implement not found for interface IStockShareholderChange, forgot register?")
	}
	return localStockShareholderChange
}

func RegisterStockShareholderChange(i IStockShareholderChange) {
	localStockShareholderChange = i
}

func StockShareholderCount() IStockShareholderCount {
	if localStockShareholderCount == nil {
		panic("implement not found for interface IStockShareholderCount, forgot register?")
	}
	return localStockShareholderCount
}

func RegisterStockShareholderCount(i IStockShareholderCount) {
	localStockShareholderCount = i
}

func StockBasicInfo() IStockBasicInfo {
	if localStockBasicInfo == nil {
		panic("implement not found for interface IStockBasicInfo, forgot register?")
	}
	return localStockBasicInfo
}

func RegisterStockBasicInfo(i IStockBasicInfo) {
	localStockBasicInfo = i
}

func StockList() IStockList {
	if localStockList == nil {
		panic("implement not found for interface IStockList, forgot register?")
	}
	return localStockList
}

func RegisterStockList(i IStockList) {
	localStockList = i
}

func StockTopTenCirculatingHolders() IStockTopTenCirculatingHolders {
	if localStockTopTenCirculatingHolders == nil {
		panic("implement not found for interface IStockTopTenCirculatingHolders, forgot register?")
	}
	return localStockTopTenCirculatingHolders
}

func RegisterStockTopTenCirculatingHolders(i IStockTopTenCirculatingHolders) {
	localStockTopTenCirculatingHolders = i
}
