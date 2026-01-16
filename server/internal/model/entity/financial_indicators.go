// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// FinancialIndicators is the golang structure for table financial_indicators.
type FinancialIndicators struct {
	Id        int64       `json:"id"        orm:"id"         description:"主键ID"`
	Date      *gtime.Time `json:"date"      orm:"date"       description:"报告日期 yyyy-MM-dd"`
	Symbol    string      `json:"symbol"    orm:"symbol"     description:"股票或标的代码 (例如: AAPL, 000001.SZ)"`
	Tbmg      float64     `json:"tbmg"      orm:"tbmg"       description:"摊薄每股收益(元)"`
	Jqmg      float64     `json:"jqmg"      orm:"jqmg"       description:"加权每股收益(元)"`
	Mgsy      float64     `json:"mgsy"      orm:"mgsy"       description:"每股收益_调整后(元)"`
	Kfmg      float64     `json:"kfmg"      orm:"kfmg"       description:"扣除非经常性损益后的每股收益(元)"`
	Mgjz      float64     `json:"mgjz"      orm:"mgjz"       description:"每股净资产_调整前(元)"`
	Mgjzad    float64     `json:"mgjzad"    orm:"mgjzad"     description:"每股净资产_调整后(元)"`
	Mgjy      float64     `json:"mgjy"      orm:"mgjy"       description:"每股经营性现金流(元)"`
	Mggjj     float64     `json:"mggjj"     orm:"mggjj"      description:"每股资本公积金(元)"`
	Mgwly     float64     `json:"mgwly"     orm:"mgwly"      description:"每股未分配利润(元)"`
	Zclr      float64     `json:"zclr"      orm:"zclr"       description:"总资产利润率(%)"`
	Zylr      float64     `json:"zylr"      orm:"zylr"       description:"主营业务利润率(%)"`
	Zzlr      float64     `json:"zzlr"      orm:"zzlr"       description:"总资产净利润率(%)"`
	Cblr      float64     `json:"cblr"      orm:"cblr"       description:"成本费用利润率(%)"`
	Yylr      float64     `json:"yylr"      orm:"yylr"       description:"营业利润率(%)"`
	Zycb      float64     `json:"zycb"      orm:"zycb"       description:"主营业务成本率(%)"`
	Xsjl      float64     `json:"xsjl"      orm:"xsjl"       description:"销售净利率(%)"`
	Gbbc      float64     `json:"gbbc"      orm:"gbbc"       description:"股本报酬率(%)"`
	Jzbc      float64     `json:"jzbc"      orm:"jzbc"       description:"净资产报酬率(%)"`
	Zcbc      float64     `json:"zcbc"      orm:"zcbc"       description:"资产报酬率(%)"`
	Xsml      float64     `json:"xsml"      orm:"xsml"       description:"销售毛利率(%)"`
	Xxbz      float64     `json:"xxbz"      orm:"xxbz"       description:"三项费用比重"`
	Fzy       float64     `json:"fzy"       orm:"fzy"        description:"非主营比重"`
	Zybz      float64     `json:"zybz"      orm:"zybz"       description:"主营利润比重"`
	Gxff      float64     `json:"gxff"      orm:"gxff"       description:"股息发放率(%)"`
	Tzsy      float64     `json:"tzsy"      orm:"tzsy"       description:"投资收益率(%)"`
	Zyyw      float64     `json:"zyyw"      orm:"zyyw"       description:"主营业务利润(元)"`
	Jzsy      float64     `json:"jzsy"      orm:"jzsy"       description:"净资产收益率(%)"`
	Jqjz      float64     `json:"jqjz"      orm:"jqjz"       description:"加权净资产收益率(%)"`
	Kflr      float64     `json:"kflr"      orm:"kflr"       description:"扣除非经常性损益后的净利润(元)"`
	Zysr      float64     `json:"zysr"      orm:"zysr"       description:"主营业务收入增长率(%)"`
	Jlzz      float64     `json:"jlzz"      orm:"jlzz"       description:"净利润增长率(%)"`
	Jzzz      float64     `json:"jzzz"      orm:"jzzz"       description:"净资产增长率(%)"`
	Zzzz      float64     `json:"zzzz"      orm:"zzzz"       description:"总资产增长率(%)"`
	Yszz      float64     `json:"yszz"      orm:"yszz"       description:"应收账款周转率(次)"`
	Yszzt     float64     `json:"yszzt"     orm:"yszzt"      description:"应收账款周转天数(天)"`
	Chzz      float64     `json:"chzz"      orm:"chzz"       description:"存货周转天数(天)"`
	Chzzl     float64     `json:"chzzl"     orm:"chzzl"      description:"存货周转率(次)"`
	Gzzz      float64     `json:"gzzz"      orm:"gzzz"       description:"固定资产周转率(次)"`
	Zzzzl     float64     `json:"zzzzl"     orm:"zzzzl"      description:"总资产周转率(次)"`
	Zzzzt     float64     `json:"zzzzt"     orm:"zzzzt"      description:"总资产周转天数(天)"`
	Ldzz      float64     `json:"ldzz"      orm:"ldzz"       description:"流动资产周转率(次)"`
	Ldzzt     float64     `json:"ldzzt"     orm:"ldzzt"      description:"流动资产周转天数(天)"`
	Gdzz      float64     `json:"gdzz"      orm:"gdzz"       description:"股东权益周转率(次)"`
	Ldbl      float64     `json:"ldbl"      orm:"ldbl"       description:"流动比率"`
	Sdbl      float64     `json:"sdbl"      orm:"sdbl"       description:"速动比率"`
	Xjbl      float64     `json:"xjbl"      orm:"xjbl"       description:"现金比率(%)"`
	Lxzf      float64     `json:"lxzf"      orm:"lxzf"       description:"利息支付倍数"`
	Zjbl      float64     `json:"zjbl"      orm:"zjbl"       description:"长期债务与营运资金比率(%)"`
	Gdqy      float64     `json:"gdqy"      orm:"gdqy"       description:"股东权益比率(%)"`
	Cqfz      float64     `json:"cqfz"      orm:"cqfz"       description:"长期负债比率(%)"`
	Gdgd      float64     `json:"gdgd"      orm:"gdgd"       description:"股东权益与固定资产比率(%)"`
	Fzqy      float64     `json:"fzqy"      orm:"fzqy"       description:"负债与所有者权益比率(%)"`
	Zczjbl    float64     `json:"zczjbl"    orm:"zczjbl"     description:"长期资产与长期资金比率(%)"`
	Zblv      float64     `json:"zblv"      orm:"zblv"       description:"资本化比率(%)"`
	Gdzcjz    float64     `json:"gdzcjz"    orm:"gdzcjz"     description:"固定资产净值率(%)"`
	Zbgdh     float64     `json:"zbgdh"     orm:"zbgdh"      description:"资本固定化比率(%)"`
	Cqbl      float64     `json:"cqbl"      orm:"cqbl"       description:"产权比率(%)"`
	Qxjzb     float64     `json:"qxjzb"     orm:"qxjzb"      description:"清算价值比率(%)"`
	Gdzcbz    float64     `json:"gdzcbz"    orm:"gdzcbz"     description:"固定资产比重(%)"`
	Zcfzl     float64     `json:"zcfzl"     orm:"zcfzl"      description:"资产负债率(%)"`
	Zzc       float64     `json:"zzc"       orm:"zzc"        description:"总资产(元)"`
	Jyxj      float64     `json:"jyxj"      orm:"jyxj"       description:"经营现金净流量对销售收入比率(%)"`
	Zcjyxj    float64     `json:"zcjyxj"    orm:"zcjyxj"     description:"资产的经营现金流量回报率(%)"`
	Jylrb     float64     `json:"jylrb"     orm:"jylrb"      description:"经营现金净流量与净利润的比率(%)"`
	Jyfzl     float64     `json:"jyfzl"     orm:"jyfzl"      description:"经营现金净流量对负债比率(%)"`
	Xjlbl     float64     `json:"xjlbl"     orm:"xjlbl"      description:"现金流量比率(%)"`
	Dqgptz    float64     `json:"dqgptz"    orm:"dqgptz"     description:"短期股票投资(元)"`
	Dqzctz    float64     `json:"dqzctz"    orm:"dqzctz"     description:"短期债券投资(元)"`
	Dqjytz    float64     `json:"dqjytz"    orm:"dqjytz"     description:"短期其它经营性投资(元)"`
	Qcgptz    float64     `json:"qcgptz"    orm:"qcgptz"     description:"长期股票投资(元)"`
	Cqzqtz    float64     `json:"cqzqtz"    orm:"cqzqtz"     description:"长期债券投资(元)"`
	Cqjyxtz   float64     `json:"cqjyxtz"   orm:"cqjyxtz"    description:"长期其它经营性投资(元)"`
	Yszk1     float64     `json:"yszk1"     orm:"yszk1"      description:"1年以内应收帐款(元)"`
	Yszk12    float64     `json:"yszk12"    orm:"yszk12"     description:"1-2年以内应收帐款(元)"`
	Yszk23    float64     `json:"yszk23"    orm:"yszk23"     description:"2-3年以内应收帐款(元)"`
	Yszk3     float64     `json:"yszk3"     orm:"yszk3"      description:"3年以内应收帐款(元)"`
	Yfhk1     float64     `json:"yfhk1"     orm:"yfhk1"      description:"1年以内预付货款(元)"`
	Yfhk12    float64     `json:"yfhk12"    orm:"yfhk12"     description:"1-2年以内预付货款(元)"`
	Yfhk23    float64     `json:"yfhk23"    orm:"yfhk23"     description:"2-3年以内预付货款(元)"`
	Yfhk3     float64     `json:"yfhk3"     orm:"yfhk3"      description:"3年以内预付货款(元)"`
	Ysk1      float64     `json:"ysk1"      orm:"ysk1"       description:"1年以内其它应收款(元)"`
	Ysk12     float64     `json:"ysk12"     orm:"ysk12"      description:"1-2年以内其它应收款(元)"`
	Ysk23     float64     `json:"ysk23"     orm:"ysk23"      description:"2-3年以内其它应收款(元)"`
	Ysk3      float64     `json:"ysk3"      orm:"ysk3"       description:"3年以内其它应收款(元)"`
	CreatedAt *gtime.Time `json:"createdAt" orm:"created_at" description:"数据创建时间"`
	UpdatedAt *gtime.Time `json:"updatedAt" orm:"updated_at" description:"数据更新时间"`
}
