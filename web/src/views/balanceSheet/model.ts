import { cloneDeep } from 'lodash-es';
import { FormSchema } from '@/components/Form';
import { defRangeShortcuts, formatToDate } from '@/utils/dateUtil';
import { renderPopoverMemberSumma, MemberSumma } from '@/utils';

export class State {
  public id = 0; // 自增主键

  public symbol = ''; // 公司代码/股票代码 (例如: 000001.SZ, AAPL)
  public jzrq = ''; // 截止日期 (会计期间结束日)
  public plrq = ''; // 披露日期
  public reportYear = 0; // 报告年度
  public reportQuarter = 0; // 报告季度 (1-4, 年报为NULL)
  public reportType = 'annual'; // 报告类型: annual-年报, quarter-季报, interim-中报
  public hbzj = null; // 货币资金
  public jyxjrzc = null; // 交易性金融资产
  public yspj = null; // 应收票据
  public yszk = null; // 应收账款
  public yfkx = null; // 预付款项
  public yslx = null; // 应收利息
  public ysgl = null; // 应收股利
  public qtysk = null; // 其他应收款
  public ch = null; // 存货
  public dfy = null; // 待摊费用
  public ynndqdfldzc = null; // 一年内到期的非流动资产
  public qtldzc = null; // 其他流动资产
  public ldzchj = null; // 流动资产合计
  public cqgqtz = null; // 长期股权投资
  public cqysk = null; // 长期应收款
  public gdzc = null; // 固定资产
  public zjgc = null; // 在建工程
  public wxzc = null; // 无形资产
  public sy = null; // 商誉
  public cqdtfy = null; // 长期待摊费用
  public dysdszc = null; // 递延所得税资产
  public qtfldzc = null; // 其他非流动资产
  public fldzchj = null; // 非流动资产合计
  public zczj = null; // 资产总计
  public dqjk = null; // 短期借款
  public jyxjrfz = null; // 交易性金融负债
  public yfpj = null; // 应付票据
  public yfzk = null; // 应付账款
  public ysk = null; // 预收账款
  public yfgzxc = null; // 应付职工薪酬
  public yjsf = null; // 应交税费
  public yflx = null; // 应付利息
  public yfgl = null; // 应付股利
  public qtfzk = null; // 其他应付款
  public ynndqdfldfz = null; // 一年内到期的非流动负债
  public qtldfz = null; // 其他流动负债
  public ldfzhj = null; // 流动负债合计
  public cqjk = null; // 长期借款
  public yfzq = null; // 应付债券
  public cqyfk = null; // 长期应付款
  public dysdsfz = null; // 递延所得税负债
  public qtfldfz = null; // 其他非流动负债
  public fldfzhj = null; // 非流动负债合计
  public fzhj = null; // 负债合计
  public sszb = null; // 实收资本(或股本)
  public zbgj = null; // 资本公积
  public ylgj = null; // 盈余公积
  public wfplr = null; // 未分配利润
  public gsmgdqsyhj = null; // 归属于母公司股东权益合计
  public ssgdqy = null; // 少数股东权益
  public syzqyhj = null; // 所有者权益合计
  public fzhgdqyzj = null; // 负债和股东权益总计
  public dataSource = ''; // 数据来源
  public currency = 'CNY'; // 货币单位 (CNY, USD等)
  public unit = 'yuan'; // 单位: yuan-元, wan-万元, qianwan-千万元
  public isAudited = 0; // 是否审计: 0-未审计, 1-已审计
  public version = 1; // 数据版本
  public createdAt = ''; // 创建时间
  public updatedAt = ''; // 更新时间
  public createdBy = ''; // 创建人
  public createdBySumma?: null | MemberSumma = null; // 创建人摘要信息
  public updatedBy = ''; // 更新人
  public updatedBySumma?: null | MemberSumma = null; // 更新人摘要信息

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
    message: '请输入公司代码/股票代码 (例如: 000001.SZ, AAPL)',
  },
  jzrq: {
    required: true,
    trigger: ['blur', 'input'],
    type: 'string',
    message: '请输入截止日期 (会计期间结束日)',
  },
  plrq: {
    required: true,
    trigger: ['blur', 'input'],
    type: 'string',
    message: '请输入披露日期',
  },
  reportYear: {
    required: true,
    trigger: ['blur', 'input'],
    type: 'number',
    message: '请输入报告年度',
  },
  reportType: {
    required: true,
    trigger: ['blur', 'input'],
    type: 'string',
    message: '请输入报告类型: annual-年报, quarter-季报, interim-中报',
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
    title: '公司代码/股票代码 (例如: 000001.SZ, AAPL)',
    key: 'symbol',
    align: 'left',
    width: -1,
  },
  {
    title: '截止日期 (会计期间结束日)',
    key: 'jzrq',
    align: 'left',
    width: -1,
    render(row: State) {
      return formatToDate(row.jzrq);
    },
  },
  {
    title: '披露日期',
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
    title: '报告季度 (1-4, 年报为NULL)',
    key: 'reportQuarter',
    align: 'left',
    width: -1,
  },
  {
    title: '报告类型: annual-年报, quarter-季报, interim-中报',
    key: 'reportType',
    align: 'left',
    width: -1,
  },
  {
    title: '货币资金',
    key: 'hbzj',
    align: 'left',
    width: -1,
  },
  {
    title: '交易性金融资产',
    key: 'jyxjrzc',
    align: 'left',
    width: -1,
  },
  {
    title: '应收票据',
    key: 'yspj',
    align: 'left',
    width: -1,
  },
  {
    title: '应收账款',
    key: 'yszk',
    align: 'left',
    width: -1,
  },
  {
    title: '预付款项',
    key: 'yfkx',
    align: 'left',
    width: -1,
  },
  {
    title: '应收利息',
    key: 'yslx',
    align: 'left',
    width: -1,
  },
  {
    title: '应收股利',
    key: 'ysgl',
    align: 'left',
    width: -1,
  },
  {
    title: '其他应收款',
    key: 'qtysk',
    align: 'left',
    width: -1,
  },
  {
    title: '存货',
    key: 'ch',
    align: 'left',
    width: -1,
  },
  {
    title: '待摊费用',
    key: 'dfy',
    align: 'left',
    width: -1,
  },
  {
    title: '一年内到期的非流动资产',
    key: 'ynndqdfldzc',
    align: 'left',
    width: -1,
  },
  {
    title: '其他流动资产',
    key: 'qtldzc',
    align: 'left',
    width: -1,
  },
  {
    title: '流动资产合计',
    key: 'ldzchj',
    align: 'left',
    width: -1,
  },
  {
    title: '长期股权投资',
    key: 'cqgqtz',
    align: 'left',
    width: -1,
  },
  {
    title: '长期应收款',
    key: 'cqysk',
    align: 'left',
    width: -1,
  },
  {
    title: '固定资产',
    key: 'gdzc',
    align: 'left',
    width: -1,
  },
  {
    title: '在建工程',
    key: 'zjgc',
    align: 'left',
    width: -1,
  },
  {
    title: '无形资产',
    key: 'wxzc',
    align: 'left',
    width: -1,
  },
  {
    title: '商誉',
    key: 'sy',
    align: 'left',
    width: -1,
  },
  {
    title: '长期待摊费用',
    key: 'cqdtfy',
    align: 'left',
    width: -1,
  },
  {
    title: '递延所得税资产',
    key: 'dysdszc',
    align: 'left',
    width: -1,
  },
  {
    title: '其他非流动资产',
    key: 'qtfldzc',
    align: 'left',
    width: -1,
  },
  {
    title: '非流动资产合计',
    key: 'fldzchj',
    align: 'left',
    width: -1,
  },
  {
    title: '资产总计',
    key: 'zczj',
    align: 'left',
    width: -1,
  },
  {
    title: '短期借款',
    key: 'dqjk',
    align: 'left',
    width: -1,
  },
  {
    title: '交易性金融负债',
    key: 'jyxjrfz',
    align: 'left',
    width: -1,
  },
  {
    title: '应付票据',
    key: 'yfpj',
    align: 'left',
    width: -1,
  },
  {
    title: '应付账款',
    key: 'yfzk',
    align: 'left',
    width: -1,
  },
  {
    title: '预收账款',
    key: 'ysk',
    align: 'left',
    width: -1,
  },
  {
    title: '应付职工薪酬',
    key: 'yfgzxc',
    align: 'left',
    width: -1,
  },
  {
    title: '应交税费',
    key: 'yjsf',
    align: 'left',
    width: -1,
  },
  {
    title: '应付利息',
    key: 'yflx',
    align: 'left',
    width: -1,
  },
  {
    title: '应付股利',
    key: 'yfgl',
    align: 'left',
    width: -1,
  },
  {
    title: '其他应付款',
    key: 'qtfzk',
    align: 'left',
    width: -1,
  },
  {
    title: '一年内到期的非流动负债',
    key: 'ynndqdfldfz',
    align: 'left',
    width: -1,
  },
  {
    title: '其他流动负债',
    key: 'qtldfz',
    align: 'left',
    width: -1,
  },
  {
    title: '流动负债合计',
    key: 'ldfzhj',
    align: 'left',
    width: -1,
  },
  {
    title: '长期借款',
    key: 'cqjk',
    align: 'left',
    width: -1,
  },
  {
    title: '应付债券',
    key: 'yfzq',
    align: 'left',
    width: -1,
  },
  {
    title: '长期应付款',
    key: 'cqyfk',
    align: 'left',
    width: -1,
  },
  {
    title: '递延所得税负债',
    key: 'dysdsfz',
    align: 'left',
    width: -1,
  },
  {
    title: '其他非流动负债',
    key: 'qtfldfz',
    align: 'left',
    width: -1,
  },
  {
    title: '非流动负债合计',
    key: 'fldfzhj',
    align: 'left',
    width: -1,
  },
  {
    title: '负债合计',
    key: 'fzhj',
    align: 'left',
    width: -1,
  },
  {
    title: '实收资本(或股本)',
    key: 'sszb',
    align: 'left',
    width: -1,
  },
  {
    title: '资本公积',
    key: 'zbgj',
    align: 'left',
    width: -1,
  },
  {
    title: '盈余公积',
    key: 'ylgj',
    align: 'left',
    width: -1,
  },
  {
    title: '未分配利润',
    key: 'wfplr',
    align: 'left',
    width: -1,
  },
  {
    title: '归属于母公司股东权益合计',
    key: 'gsmgdqsyhj',
    align: 'left',
    width: -1,
  },
  {
    title: '少数股东权益',
    key: 'ssgdqy',
    align: 'left',
    width: -1,
  },
  {
    title: '所有者权益合计',
    key: 'syzqyhj',
    align: 'left',
    width: -1,
  },
  {
    title: '负债和股东权益总计',
    key: 'fzhgdqyzj',
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
    title: '货币单位 (CNY, USD等)',
    key: 'currency',
    align: 'left',
    width: -1,
  },
  {
    title: '单位: yuan-元, wan-万元, qianwan-千万元',
    key: 'unit',
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
    title: '数据版本',
    key: 'version',
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
  {
    title: '创建人',
    key: 'createdBy',
    align: 'left',
    width: -1,
    render(row: State) {
      return renderPopoverMemberSumma(row.createdBySumma);
    },
  },
  {
    title: '更新人',
    key: 'updatedBy',
    align: 'left',
    width: -1,
    render(row: State) {
      return renderPopoverMemberSumma(row.updatedBySumma);
    },
  },
];