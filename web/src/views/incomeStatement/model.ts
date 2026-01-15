import { cloneDeep } from 'lodash-es';
import { FormSchema } from '@/components/Form';
import { defRangeShortcuts, formatToDate } from '@/utils/dateUtil';

export class State {
  public id = 0; // 自增主键

  public symbol = ''; // 股票代码 (如: 000001.SZ)
  public jzrq = ''; // 截止日期 (报告期截止日，如2025-12-31)
  public plrq = ''; // 披露日期 (财报实际发布日期)
  public reportYear = 0; // 报告年度
  public reportQuarter = 0; // 报告季度 (1-4)
  public reportType = 'quarter'; // 报告类型: annual-年报, quarter-季报
  public fiscalPeriod = ''; // 会计期间
  public yysr = null; // 营业收入
  public yzbf = null; // 已赚保费
  public fdczssr = null; // 房地产销售收入
  public qtywsr = null; // 其他业务收入
  public yyzsr = null; // 营业总收入
  public lxsr = null; // 利息收入
  public sxfjyjsr = null; // 手续费及佣金收入
  public btsr = null; // 补贴收入
  public ywsr = null; // 营业外收入
  public qtsy = null; // 其他收益
  public yycb = null; // 营业成本
  public fdczscb = null; // 房地产销售成本
  public qtywcb = null; // 其他业务成本
  public yyzcb = null; // 营业总成本
  public yysjjfj = null; // 营业税金及附加
  public xsfy = null; // 销售费用
  public glfy = null; // 管理费用
  public yffy = null; // 研发费用
  public cwfy = null; // 财务费用
  public sxfjyjzc = null; // 手续费及佣金支出
  public lxzc = null; // 利息支出
  public tbj = null; // 退保金
  public pczjje = null; // 赔付支出净额
  public tqbxhtzbjje = null; // 提取保险合同准备金净额
  public bdhlzc = null; // 保单红利支出
  public fbfy = null; // 分保费用
  public zcjzss = null; // 资产减值损失
  public ywzc = null; // 营业外支出
  public qtywlr = null; // 其他业务利润
  public yylr = null; // 营业利润
  public lrze = null; // 利润总额
  public jlr = null; // 净利润
  public jlrhfcjcx = null; // 净利润(扣除非经常性损益后)
  public gsmgsyzzdjlr = null; // 归属于母公司所有者的净利润
  public bhbfzhbqsljlr = null; // 被合并方在合并前实现净利润
  public tzsy = null; // 投资收益
  public lyqyhhhqydtzsy = null; // 联营企业和合营企业的投资收益
  public gyjzbdsy = null; // 公允价值变动收益
  public qhsy = null; // 期货损益
  public tgsy = null; // 托管收益
  public hdsy = null; // 汇兑收益
  public fldzcczsy = null; // 非流动资产处置收益
  public sdsfy = null; // 所得税费用
  public ssgdsy = null; // 少数股东损益
  public wqrtzss = null; // 未确认投资损失
  public jbmgsy = null; // 基本每股收益
  public xsmgsy = null; // 稀释每股收益
  public zhsyz = null; // 综合收益总额
  public gsssgdzhsyz = null; // 归属于少数股东的综合收益总额
  public grossMargin = null; // 毛利率(%)
  public operatingMargin = null; // 营业利润率(%)
  public netMargin = null; // 净利率(%)
  public effectiveTaxRate = null; // 实际税率(%)
  public dataSource = ''; // 数据来源
  public currency = 'CNY'; // 货币单位
  public unit = 'yuan'; // 单位: yuan-元, wan-万元
  public accountingStandard = ''; // 会计准则 (如: CAS, IFRS)
  public isAudited = 0; // 是否审计: 0-未审计, 1-已审计
  public isConsolidated = 1; // 是否合并报表: 1-合并, 0-母公司
  public createdAt = ''; // 创建时间
  public updatedAt = ''; // 更新时间

  constructor(state?: Partial<State>) {
    if (state) {
      Object.assign(this, state);
    }
  }
}

export function newState(state: State | Record<string, any> | null): State {
  if (state !== null) {
    if (state instanceof State) {
      return cloneDeep(state);
    }
    return new State(state);
  }
  return new State();
}

// 表单验证规则
export const rules = {
  symbol: {
    required: true,
    trigger: ['blur', 'input'],
    type: 'string',
    message: '请输入股票代码 (如: 000001.SZ)',
  },
  jzrq: {
    required: true,
    trigger: ['blur', 'input'],
    type: 'string',
    message: '请输入截止日期 (报告期截止日，如2025-12-31)',
  },
  plrq: {
    required: true,
    trigger: ['blur', 'input'],
    type: 'string',
    message: '请输入披露日期 (财报实际发布日期)',
  },
  reportType: {
    required: true,
    trigger: ['blur', 'input'],
    type: 'string',
    message: '请输入报告类型: annual-年报, quarter-季报',
  },
};

// 表格搜索表单
export const schemas = ref<FormSchema[]>([
  {
    field: 'id',
    component: 'NInputNumber',
    label: '自增主键',
    componentProps: {
      placeholder: '请输入自增主键',
      onUpdateValue: (e: any) => {
        console.log(e);
      },
    },
  },
  {
    field: 'createdAt',
    component: 'NDatePicker',
    label: '创建时间',
    componentProps: {
      type: 'datetimerange',
      clearable: true,
      shortcuts: defRangeShortcuts(),
      onUpdateValue: (e: any) => {
        console.log(e);
      },
    },
  },
]);

// 表格列
export const columns = [
  {
    title: '自增主键',
    key: 'id',
    align: 'left',
    width: -1,
  },
  {
    title: '股票代码 (如: 000001.SZ)',
    key: 'symbol',
    align: 'left',
    width: -1,
  },
  {
    title: '截止日期 (报告期截止日，如2025-12-31)',
    key: 'jzrq',
    align: 'left',
    width: -1,
    render(row: State) {
      return formatToDate(row.jzrq);
    },
  },
  {
    title: '披露日期 (财报实际发布日期)',
    key: 'plrq',
    align: 'left',
    width: -1,
    render(row: State) {
      return formatToDate(row.plrq);
    },
  },
  {
    title: '报告年度',
    key: 'reportYear',
    align: 'left',
    width: -1,
  },
  {
    title: '报告季度 (1-4)',
    key: 'reportQuarter',
    align: 'left',
    width: -1,
  },
  {
    title: '报告类型: annual-年报, quarter-季报',
    key: 'reportType',
    align: 'left',
    width: -1,
  },
  {
    title: '会计期间',
    key: 'fiscalPeriod',
    align: 'left',
    width: -1,
  },
  {
    title: '营业收入',
    key: 'yysr',
    align: 'left',
    width: -1,
  },
  {
    title: '已赚保费',
    key: 'yzbf',
    align: 'left',
    width: -1,
  },
  {
    title: '房地产销售收入',
    key: 'fdczssr',
    align: 'left',
    width: -1,
  },
  {
    title: '其他业务收入',
    key: 'qtywsr',
    align: 'left',
    width: -1,
  },
  {
    title: '营业总收入',
    key: 'yyzsr',
    align: 'left',
    width: -1,
  },
  {
    title: '利息收入',
    key: 'lxsr',
    align: 'left',
    width: -1,
  },
  {
    title: '手续费及佣金收入',
    key: 'sxfjyjsr',
    align: 'left',
    width: -1,
  },
  {
    title: '补贴收入',
    key: 'btsr',
    align: 'left',
    width: -1,
  },
  {
    title: '营业外收入',
    key: 'ywsr',
    align: 'left',
    width: -1,
  },
  {
    title: '其他收益',
    key: 'qtsy',
    align: 'left',
    width: -1,
  },
  {
    title: '营业成本',
    key: 'yycb',
    align: 'left',
    width: -1,
  },
  {
    title: '房地产销售成本',
    key: 'fdczscb',
    align: 'left',
    width: -1,
  },
  {
    title: '其他业务成本',
    key: 'qtywcb',
    align: 'left',
    width: -1,
  },
  {
    title: '营业总成本',
    key: 'yyzcb',
    align: 'left',
    width: -1,
  },
  {
    title: '营业税金及附加',
    key: 'yysjjfj',
    align: 'left',
    width: -1,
  },
  {
    title: '销售费用',
    key: 'xsfy',
    align: 'left',
    width: -1,
  },
  {
    title: '管理费用',
    key: 'glfy',
    align: 'left',
    width: -1,
  },
  {
    title: '研发费用',
    key: 'yffy',
    align: 'left',
    width: -1,
  },
  {
    title: '财务费用',
    key: 'cwfy',
    align: 'left',
    width: -1,
  },
  {
    title: '手续费及佣金支出',
    key: 'sxfjyjzc',
    align: 'left',
    width: -1,
  },
  {
    title: '利息支出',
    key: 'lxzc',
    align: 'left',
    width: -1,
  },
  {
    title: '退保金',
    key: 'tbj',
    align: 'left',
    width: -1,
  },
  {
    title: '赔付支出净额',
    key: 'pczjje',
    align: 'left',
    width: -1,
  },
  {
    title: '提取保险合同准备金净额',
    key: 'tqbxhtzbjje',
    align: 'left',
    width: -1,
  },
  {
    title: '保单红利支出',
    key: 'bdhlzc',
    align: 'left',
    width: -1,
  },
  {
    title: '分保费用',
    key: 'fbfy',
    align: 'left',
    width: -1,
  },
  {
    title: '资产减值损失',
    key: 'zcjzss',
    align: 'left',
    width: -1,
  },
  {
    title: '营业外支出',
    key: 'ywzc',
    align: 'left',
    width: -1,
  },
  {
    title: '其他业务利润',
    key: 'qtywlr',
    align: 'left',
    width: -1,
  },
  {
    title: '营业利润',
    key: 'yylr',
    align: 'left',
    width: -1,
  },
  {
    title: '利润总额',
    key: 'lrze',
    align: 'left',
    width: -1,
  },
  {
    title: '净利润',
    key: 'jlr',
    align: 'left',
    width: -1,
  },
  {
    title: '净利润(扣除非经常性损益后)',
    key: 'jlrhfcjcx',
    align: 'left',
    width: -1,
  },
  {
    title: '归属于母公司所有者的净利润',
    key: 'gsmgsyzzdjlr',
    align: 'left',
    width: -1,
  },
  {
    title: '被合并方在合并前实现净利润',
    key: 'bhbfzhbqsljlr',
    align: 'left',
    width: -1,
  },
  {
    title: '投资收益',
    key: 'tzsy',
    align: 'left',
    width: -1,
  },
  {
    title: '联营企业和合营企业的投资收益',
    key: 'lyqyhhhqydtzsy',
    align: 'left',
    width: -1,
  },
  {
    title: '公允价值变动收益',
    key: 'gyjzbdsy',
    align: 'left',
    width: -1,
  },
  {
    title: '期货损益',
    key: 'qhsy',
    align: 'left',
    width: -1,
  },
  {
    title: '托管收益',
    key: 'tgsy',
    align: 'left',
    width: -1,
  },
  {
    title: '汇兑收益',
    key: 'hdsy',
    align: 'left',
    width: -1,
  },
  {
    title: '非流动资产处置收益',
    key: 'fldzcczsy',
    align: 'left',
    width: -1,
  },
  {
    title: '所得税费用',
    key: 'sdsfy',
    align: 'left',
    width: -1,
  },
  {
    title: '少数股东损益',
    key: 'ssgdsy',
    align: 'left',
    width: -1,
  },
  {
    title: '未确认投资损失',
    key: 'wqrtzss',
    align: 'left',
    width: -1,
  },
  {
    title: '基本每股收益',
    key: 'jbmgsy',
    align: 'left',
    width: -1,
  },
  {
    title: '稀释每股收益',
    key: 'xsmgsy',
    align: 'left',
    width: -1,
  },
  {
    title: '综合收益总额',
    key: 'zhsyz',
    align: 'left',
    width: -1,
  },
  {
    title: '归属于少数股东的综合收益总额',
    key: 'gsssgdzhsyz',
    align: 'left',
    width: -1,
  },
  {
    title: '毛利率(%)',
    key: 'grossMargin',
    align: 'left',
    width: -1,
  },
  {
    title: '营业利润率(%)',
    key: 'operatingMargin',
    align: 'left',
    width: -1,
  },
  {
    title: '净利率(%)',
    key: 'netMargin',
    align: 'left',
    width: -1,
  },
  {
    title: '实际税率(%)',
    key: 'effectiveTaxRate',
    align: 'left',
    width: -1,
  },
  {
    title: '数据来源',
    key: 'dataSource',
    align: 'left',
    width: -1,
  },
  {
    title: '货币单位',
    key: 'currency',
    align: 'left',
    width: -1,
  },
  {
    title: '单位: yuan-元, wan-万元',
    key: 'unit',
    align: 'left',
    width: -1,
  },
  {
    title: '会计准则 (如: CAS, IFRS)',
    key: 'accountingStandard',
    align: 'left',
    width: -1,
  },
  {
    title: '是否审计: 0-未审计, 1-已审计',
    key: 'isAudited',
    align: 'left',
    width: -1,
  },
  {
    title: '是否合并报表: 1-合并, 0-母公司',
    key: 'isConsolidated',
    align: 'left',
    width: -1,
  },
  {
    title: '创建时间',
    key: 'createdAt',
    align: 'left',
    width: -1,
  },
  {
    title: '更新时间',
    key: 'updatedAt',
    align: 'left',
    width: -1,
  },
];