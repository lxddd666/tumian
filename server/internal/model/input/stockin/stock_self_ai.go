// Package stockin
// @Link  https://github.com/bufanyun/hotgo
// @Copyright  Copyright (c) 2026 HotGo CLI
// @Author  Ms <133814250@qq.com>
// @License  https://github.com/bufanyun/hotgo/blob/master/LICENSE
// @AutoGenerate Version 2.18.6
package stockin

import (
	"context"
	"hotgo/internal/model/entity"
	"hotgo/internal/model/input/form"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// StockSelfAiUpdateFields 修改ai 基本信息字段过滤
type StockSelfAiUpdateFields struct {
	Name    string `json:"name"    dc:"ai名称 例如deepseek 千问"`
	Model   string `json:"model"   dc:"ai model"`
	BaseUrl string `json:"baseUrl" dc:"ai base url"`
	ApiKey  string `json:"apiKey"  dc:"ai api key"`
}

// StockSelfAiInsertFields 新增ai 基本信息字段过滤
type StockSelfAiInsertFields struct {
	Name    string `json:"name"    dc:"ai名称 例如deepseek 千问"`
	Model   string `json:"model"   dc:"ai model"`
	BaseUrl string `json:"baseUrl" dc:"ai base url"`
	ApiKey  string `json:"apiKey"  dc:"ai api key"`
}

// StockSelfAiEditInp 修改/新增ai 基本信息
type StockSelfAiEditInp struct {
	entity.StockSelfAi
}

func (in *StockSelfAiEditInp) Filter(ctx context.Context) (err error) {
	// 验证ai名称 例如deepseek 千问
	if err := g.Validator().Rules("required").Data(in.Name).Messages("ai名称 例如deepseek 千问不能为空").Run(ctx); err != nil {
		return err.Current()
	}
	// 验证ai model
	if err := g.Validator().Rules("required").Data(in.Model).Messages("ai model不能为空").Run(ctx); err != nil {
		return err.Current()
	}
	// 验证ai base url
	if err := g.Validator().Rules("required").Data(in.BaseUrl).Messages("ai base url不能为空").Run(ctx); err != nil {
		return err.Current()
	}
	// 验证ai api key
	if err := g.Validator().Rules("required").Data(in.ApiKey).Messages("ai api key不能为空").Run(ctx); err != nil {
		return err.Current()
	}
	return
}

type StockSelfAiEditModel struct{}

// StockSelfAiDeleteInp 删除ai 基本信息
type StockSelfAiDeleteInp struct {
	Id interface{} `json:"id" v:"required#自增主键不能为空" dc:"自增主键"`
}

func (in *StockSelfAiDeleteInp) Filter(ctx context.Context) (err error) {
	return
}

type StockSelfAiDeleteModel struct{}

// StockSelfAiViewInp 获取指定ai 基本信息信息
type StockSelfAiViewInp struct {
	Id int64 `json:"id" v:"required#自增主键不能为空" dc:"自增主键"`
}

func (in *StockSelfAiViewInp) Filter(ctx context.Context) (err error) {
	return
}

type StockSelfAiViewModel struct {
	entity.StockSelfAi
}

// StockSelfAiListInp 获取ai 基本信息列表
type StockSelfAiListInp struct {
	form.PageReq
	Id        int64         `json:"id"        dc:"自增主键"`
	CreatedAt []*gtime.Time `json:"createdAt" dc:"创建时间"`
}

func (in *StockSelfAiListInp) Filter(ctx context.Context) (err error) {
	return
}

type StockSelfAiListModel struct {
	Id        int64       `json:"id"        dc:"自增主键"`
	Name      string      `json:"name"      dc:"ai名称 例如deepseek 千问"`
	Model     string      `json:"model"     dc:"ai model"`
	BaseUrl   string      `json:"baseUrl"   dc:"ai base url"`
	ApiKey    string      `json:"apiKey"    dc:"ai api key"`
	CreatedAt *gtime.Time `json:"createdAt" dc:"创建时间"`
}

// StockSelfAiExportModel 导出ai 基本信息
type StockSelfAiExportModel struct {
	Id        int64       `json:"id"        dc:"自增主键"`
	Name      string      `json:"name"      dc:"ai名称 例如deepseek 千问"`
	Model     string      `json:"model"     dc:"ai model"`
	BaseUrl   string      `json:"baseUrl"   dc:"ai base url"`
	ApiKey    string      `json:"apiKey"    dc:"ai api key"`
	CreatedAt *gtime.Time `json:"createdAt" dc:"创建时间"`
}
