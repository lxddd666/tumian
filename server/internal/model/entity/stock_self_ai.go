// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// StockSelfAi is the golang structure for table stock_self_ai.
type StockSelfAi struct {
	Id        uint64      `json:"id"        orm:"id"         description:"自增主键"`
	Name      string      `json:"name"      orm:"name"       description:"ai模型全称名称 例如deepseek-plus 千问"`
	Model     string      `json:"model"     orm:"model"      description:"ai model"`
	BaseUrl   string      `json:"baseUrl"   orm:"base_url"   description:"ai base url"`
	ApiKey    string      `json:"apiKey"    orm:"api_key"    description:"ai api key"`
	CreatedAt *gtime.Time `json:"createdAt" orm:"created_at" description:"创建时间"`
	AiModel   string      `json:"aiModel"   orm:"ai_model"   description:"语言模型 qianwen deepseek"`
	Status    int         `json:"status"    orm:"status"     description:"0正常 -1不正常"`
}
