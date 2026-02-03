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
	"github.com/openai/openai-go"
	"hotgo/internal/dao"
	"hotgo/internal/library/hgorm/handler"
	"hotgo/internal/model/entity"
	"hotgo/internal/model/input/form"
	"hotgo/internal/model/input/stockin"
	"hotgo/internal/service"
	"hotgo/utility/convert"
	"hotgo/utility/excel"
	"regexp"

	"github.com/openai/openai-go/option"
)

type sStockSelfAi struct{}

func NewStockSelfAi() *sStockSelfAi {
	return &sStockSelfAi{}
}

func init() {
	service.RegisterStockSelfAi(NewStockSelfAi())
}

// Model ai 基本信息ORM模型
func (s *sStockSelfAi) Model(ctx context.Context, option ...*handler.Option) *gdb.Model {
	return handler.Model(dao.StockSelfAi.Ctx(ctx), option...)
}

// List 获取ai 基本信息列表
func (s *sStockSelfAi) List(ctx context.Context, in *stockin.StockSelfAiListInp) (list []*stockin.StockSelfAiListModel, totalCount int, err error) {
	mod := s.Model(ctx)

	// 字段过滤
	mod = mod.Fields(stockin.StockSelfAiListModel{})

	// 查询自增主键
	if in.Id > 0 {
		mod = mod.Where(dao.StockSelfAi.Columns().Id, in.Id)
	}

	// 查询创建时间
	if len(in.CreatedAt) == 2 {
		mod = mod.WhereBetween(dao.StockSelfAi.Columns().CreatedAt, in.CreatedAt[0], in.CreatedAt[1])
	}

	// 分页
	mod = mod.Page(in.Page, in.PerPage)

	// 排序
	mod = mod.OrderDesc(dao.StockSelfAi.Columns().Id)

	// 查询数据
	if err = mod.ScanAndCount(&list, &totalCount, false); err != nil {
		err = gerror.Wrap(err, "获取ai 基本信息列表失败，请稍后重试！")
		return
	}
	return
}

// Export 导出ai 基本信息
func (s *sStockSelfAi) Export(ctx context.Context, in *stockin.StockSelfAiListInp) (err error) {
	list, totalCount, err := s.List(ctx, in)
	if err != nil {
		return
	}

	// 字段的排序是依据tags的字段顺序，如果你不想使用默认的排序方式，可以直接定义 tags = []string{"字段名称", "字段名称2", ...}
	tags, err := convert.GetEntityDescTags(stockin.StockSelfAiExportModel{})
	if err != nil {
		return
	}

	var (
		fileName  = "导出ai 基本信息-" + gctx.CtxId(ctx)
		sheetName = fmt.Sprintf("索引条件共%v行,共%v页,当前导出是第%v页,本页共%v行", totalCount, form.CalPageCount(totalCount, in.PerPage), in.Page, len(list))
		exports   []stockin.StockSelfAiExportModel
	)

	if err = gconv.Scan(list, &exports); err != nil {
		return
	}

	err = excel.ExportByStructs(ctx, tags, exports, fileName, sheetName)
	return
}

// Edit 修改/新增ai 基本信息
func (s *sStockSelfAi) Edit(ctx context.Context, in *stockin.StockSelfAiEditInp) (err error) {
	return g.DB().Transaction(ctx, func(ctx context.Context, tx gdb.TX) (err error) {

		// 修改
		if in.Id > 0 {
			if _, err = s.Model(ctx).
				Fields(stockin.StockSelfAiUpdateFields{}).
				WherePri(in.Id).Data(in).Update(); err != nil {
				err = gerror.Wrap(err, "修改ai 基本信息失败，请稍后重试！")
			}
			return
		}

		// 新增
		if _, err = s.Model(ctx, &handler.Option{FilterAuth: false}).
			Fields(stockin.StockSelfAiInsertFields{}).
			Data(in).OmitEmptyData().Insert(); err != nil {
			err = gerror.Wrap(err, "新增ai 基本信息失败，请稍后重试！")
		}
		return
	})
}

// Delete 删除ai 基本信息
func (s *sStockSelfAi) Delete(ctx context.Context, in *stockin.StockSelfAiDeleteInp) (err error) {

	if _, err = s.Model(ctx).WherePri(in.Id).Unscoped().Delete(); err != nil {
		err = gerror.Wrap(err, "删除ai 基本信息失败，请稍后重试！")
		return
	}
	return
}

// View 获取ai 基本信息指定信息
func (s *sStockSelfAi) View(ctx context.Context, in *stockin.StockSelfAiViewInp) (res *stockin.StockSelfAiViewModel, err error) {
	if err = s.Model(ctx).WherePri(in.Id).Scan(&res); err != nil {
		err = gerror.Wrap(err, "获取ai 基本信息信息，请稍后重试！")
		return
	}
	return
}

// GetAllAi 获取ai
func (s *sStockSelfAi) GetAllAi(ctx context.Context, aiModel string) (list []*entity.StockSelfAi, err error) {
	mod := s.Model(ctx)
	if aiModel != "" {
		mod = mod.Where(dao.StockSelfAi.Columns().AiModel, aiModel)
	}
	err = mod.Scan(&list)
	return
}

// InvokeAi ai调用
func (s *sStockSelfAi) InvokeAi(ctx context.Context, aiModel *entity.StockSelfAi, scripts []string) (res string, err error) {
	if aiModel == nil {
		return
	}

	switch aiModel.AiModel {
	case "qianwen": //千问
		res, err = s.InvokeQianWen(ctx, aiModel, scripts)
	case "bailing":
		res, err = s.InvokeBailin(ctx, aiModel, scripts)
	case "deepseek":
		res, err = s.InvokeDeepseek(ctx, aiModel, scripts)
	case "minimax":
		res, err = s.InvokeMiniMax(ctx, aiModel, scripts)
	}

	re := regexp.MustCompile(`\{[^}]*\}`)
	match := re.FindString(res)
	if match != "" {
		res = match
	} else {
	}
	return
}

// InvokeQianWen 千问Api
func (s *sStockSelfAi) InvokeQianWen(ctx context.Context, aiModel *entity.StockSelfAi, scripts []string) (res string, err error) {
	client := openai.NewClient(
		option.WithAPIKey(aiModel.ApiKey),
		option.WithBaseURL(aiModel.BaseUrl),
	)

	if len(scripts) == 0 {
		return
	}

	var openaiChatCompletionNewParams openai.ChatCompletionNewParams
	openaiChatCompletionNewParams.Model = aiModel.Model
	for _, script := range scripts {
		openaiChatCompletionNewParams.Messages = append(
			openaiChatCompletionNewParams.Messages,
			openai.UserMessage(script),
		)
	}
	chatCompletion, err := client.Chat.Completions.New(
		context.TODO(), openaiChatCompletionNewParams,
	)

	//chatCompletion, err := client.Chat.Completions.New(
	//	context.TODO(), openai.ChatCompletionNewParams{
	//		Messages: []openai.ChatCompletionMessageParamUnion{
	//			openai.UserMessage("aaa"),
	//			openai.UserMessage("bbb"),
	//		},
	//		Model: aiModel.Model,
	//	},
	//)

	if err != nil {
		panic(err.Error())
	}
	if len(chatCompletion.Choices) == 0 {
		err = gerror.New("查无数据信息")
		return
	}

	return chatCompletion.Choices[0].Message.Content, nil

}

// InvokeBailin 百灵APi
func (s *sStockSelfAi) InvokeBailin(ctx context.Context, aiModel *entity.StockSelfAi, scripts []string) (res string, err error) {
	client := openai.NewClient(
		option.WithAPIKey(aiModel.ApiKey),
		option.WithBaseURL(aiModel.BaseUrl),
	)
	//chatCompletion, err := client.Chat.Completions.New(
	//	context.TODO(), openai.ChatCompletionNewParams{
	//		Messages: []openai.ChatCompletionMessageParamUnion{
	//			openai.UserMessage(script),
	//		},
	//		Model: aiModel.Model,
	//	},
	//)

	if len(scripts) == 0 {
		return
	}

	var openaiChatCompletionNewParams openai.ChatCompletionNewParams
	openaiChatCompletionNewParams.Model = aiModel.Model
	for _, script := range scripts {
		openaiChatCompletionNewParams.Messages = append(
			openaiChatCompletionNewParams.Messages,
			openai.UserMessage(script),
		)
	}
	chatCompletion, err := client.Chat.Completions.New(
		context.TODO(), openaiChatCompletionNewParams,
	)

	if err != nil {
		return
	}
	if len(chatCompletion.Choices) == 0 {
		err = gerror.New("查无数据信息")
		return
	}

	return chatCompletion.Choices[0].Message.Content, nil

}

// InvokeDeepseek deepseek
func (s *sStockSelfAi) InvokeDeepseek(ctx context.Context, aiModel *entity.StockSelfAi, scripts []string) (res string, err error) {
	client := openai.NewClient(
		option.WithAPIKey(aiModel.ApiKey),
		option.WithBaseURL(aiModel.BaseUrl),
	)

	var openaiChatCompletionNewParams openai.ChatCompletionNewParams
	openaiChatCompletionNewParams.Model = aiModel.Model
	for _, script := range scripts {
		openaiChatCompletionNewParams.Messages = append(
			openaiChatCompletionNewParams.Messages,
			openai.UserMessage(script),
		)
	}
	chatCompletion, err := client.Chat.Completions.New(
		context.TODO(), openaiChatCompletionNewParams,
	)
	//chatCompletion, err := client.Chat.Completions.New(
	//	context.TODO(), openai.ChatCompletionNewParams{
	//		Messages: []openai.ChatCompletionMessageParamUnion{
	//			openai.UserMessage(script),
	//		},
	//		Model: aiModel.Model,
	//	},
	//)

	if err != nil {
		return
	}

	println(chatCompletion.Model)

	if len(chatCompletion.Choices) > 0 {
		res = chatCompletion.Choices[0].Message.Content
	}
	return
}

// InvokeMiniMax minimax
func (s *sStockSelfAi) InvokeMiniMax(ctx context.Context, aiModel *entity.StockSelfAi, scripts []string) (res string, err error) {
	client := openai.NewClient(
		option.WithAPIKey(aiModel.ApiKey),
		option.WithBaseURL(aiModel.BaseUrl),
	)

	var openaiChatCompletionNewParams openai.ChatCompletionNewParams
	openaiChatCompletionNewParams.Model = aiModel.Model
	for _, script := range scripts {
		openaiChatCompletionNewParams.Messages = append(
			openaiChatCompletionNewParams.Messages,
			openai.UserMessage(script),
		)
	}
	chatCompletion, err := client.Chat.Completions.New(
		context.TODO(), openaiChatCompletionNewParams,
	)
	//chatCompletion, err := client.Chat.Completions.New(
	//	context.TODO(), openai.ChatCompletionNewParams{
	//		Messages: []openai.ChatCompletionMessageParamUnion{
	//			openai.UserMessage(script),
	//		},
	//		Model: aiModel.Model,
	//	},
	//)

	if err != nil {
		return
	}

	println(chatCompletion.Model)

	if len(chatCompletion.Choices) > 0 {
		res = chatCompletion.Choices[0].Message.Content
	}
	return
}
