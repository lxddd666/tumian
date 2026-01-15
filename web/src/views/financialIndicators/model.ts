import { cloneDeep } from 'lodash-es';
import { FormSchema } from '@/components/Form';
import { defRangeShortcuts, formatToDate } from '@/utils/dateUtil';

export class State {
  public id = 0; // 自增主键

  public symbol = ''; // 公司代码/股票代码 (例如: 000001.SZ, AAPL)
  public jzrq = ''; // 截止日期 (报告期结束日)
  public plrq = ''; // 披露日期
  public reportYear = 0; // 报告年度
  public reportQuarter = 0; // 报告季度 (1-4)
  public reportType = 'quarter'; // 报告类型: annual-年报, quarter-季报
  public fiscalPeriod = ''; // 会计期间 (衍生字段，如2023Q1)
  public mgzbgjj = null; // 每股资本公积金
  public mgjyhdxjl = null; // 每股经营活动现金流量
  public mgjzc = null; // 每股净资产
  public jbmgsy = null; // 基本每股收益
  public xsmgsy = null; // 稀释每股收益
  public mgwfplr = null; // 每股未分配利润
  public kfmgsy = null; // 扣非每股收益
  public jzcsyl = null; // 净资产收益率(%)
  public jqjzcsyl = null; // 加权净资产收益率(%)
  public tbjzcsyl = null; // 摊薄净资产收益率(%)
  public tbzzcsyl = null; // 摊薄总资产收益率(%)
  public xsmlv = null; // 销售毛利率(%)
  public mlv = null; // 毛利率(%)
  public jlv = null; // 净利率(%)
  public sjslv = null; // 实际税率(%)
  public zyyrsrzz = null; // 主营收入同比增长(%)
  public jlrzz = null; // 净利润同比增长(%)
  public gsmgsyzzdjlrzz = null; // 归属于母公司所有者的净利润同比增长(%)
  public kfjlrzz = null; // 扣非净利润同比增长(%)
  public yyzsrgdhbzz = null; // 营业总收入滚动环比增长(%)
  public sljlrjqhbzz = null; // 归属净利润滚动环比增长(%)
  public kfjlrgdhbzz = null; // 扣非净利润滚动环比增长(%)
  public yskyysr = null; // 预收款/营业收入
  public xsxjlyysr = null; // 销售现金流/营业收入
  public zcfzl = null; // 资产负债比率(%)
  public chzzl = null; // 存货周转率(次)
  public dataSource = ''; // 数据来源
  public currency = 'CNY'; // 货币单位
  public unit = 'yuan'; // 单位: yuan-元
  public isCalculated = 0; // 是否为计算指标: 0-原始数据, 1-计算得出
  public calcVersion = '1.0'; // 计算版本
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
    message: '请输入公司代码/股票代码 (例如: 000001.SZ, AAPL)',
  },
  jzrq: {
    required: true,
    trigger: ['blur', 'input'],
    type: 'string',
    message: '请输入截止日期 (报告期结束日)',
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
    title: '公司代码/股票代码 (例如: 000001.SZ, AAPL)',
    key: 'symbol',
    align: 'left',
    width: -1,
  },
  {
    title: '截止日期 (报告期结束日)',
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
    title: '会计期间 (衍生字段，如2023Q1)',
    key: 'fiscalPeriod',
    align: 'left',
    width: -1,
  },
  {
    title: '每股资本公积金',
    key: 'mgzbgjj',
    align: 'left',
    width: -1,
  },
  {
    title: '每股经营活动现金流量',
    key: 'mgjyhdxjl',
    align: 'left',
    width: -1,
  },
  {
    title: '每股净资产',
    key: 'mgjzc',
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
    title: '每股未分配利润',
    key: 'mgwfplr',
    align: 'left',
    width: -1,
  },
  {
    title: '扣非每股收益',
    key: 'kfmgsy',
    align: 'left',
    width: -1,
  },
  {
    title: '净资产收益率(%)',
    key: 'jzcsyl',
    align: 'left',
    width: -1,
  },
  {
    title: '加权净资产收益率(%)',
    key: 'jqjzcsyl',
    align: 'left',
    width: -1,
  },
  {
    title: '摊薄净资产收益率(%)',
    key: 'tbjzcsyl',
    align: 'left',
    width: -1,
  },
  {
    title: '摊薄总资产收益率(%)',
    key: 'tbzzcsyl',
    align: 'left',
    width: -1,
  },
  {
    title: '销售毛利率(%)',
    key: 'xsmlv',
    align: 'left',
    width: -1,
  },
  {
    title: '毛利率(%)',
    key: 'mlv',
    align: 'left',
    width: -1,
  },
  {
    title: '净利率(%)',
    key: 'jlv',
    align: 'left',
    width: -1,
  },
  {
    title: '实际税率(%)',
    key: 'sjslv',
    align: 'left',
    width: -1,
  },
  {
    title: '主营收入同比增长(%)',
    key: 'zyyrsrzz',
    align: 'left',
    width: -1,
  },
  {
    title: '净利润同比增长(%)',
    key: 'jlrzz',
    align: 'left',
    width: -1,
  },
  {
    title: '归属于母公司所有者的净利润同比增长(%)',
    key: 'gsmgsyzzdjlrzz',
    align: 'left',
    width: -1,
  },
  {
    title: '扣非净利润同比增长(%)',
    key: 'kfjlrzz',
    align: 'left',
    width: -1,
  },
  {
    title: '营业总收入滚动环比增长(%)',
    key: 'yyzsrgdhbzz',
    align: 'left',
    width: -1,
  },
  {
    title: '归属净利润滚动环比增长(%)',
    key: 'sljlrjqhbzz',
    align: 'left',
    width: -1,
  },
  {
    title: '扣非净利润滚动环比增长(%)',
    key: 'kfjlrgdhbzz',
    align: 'left',
    width: -1,
  },
  {
    title: '预收款/营业收入',
    key: 'yskyysr',
    align: 'left',
    width: -1,
  },
  {
    title: '销售现金流/营业收入',
    key: 'xsxjlyysr',
    align: 'left',
    width: -1,
  },
  {
    title: '资产负债比率(%)',
    key: 'zcfzl',
    align: 'left',
    width: -1,
  },
  {
    title: '存货周转率(次)',
    key: 'chzzl',
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
    title: '单位: yuan-元',
    key: 'unit',
    align: 'left',
    width: -1,
  },
  {
    title: '是否为计算指标: 0-原始数据, 1-计算得出',
    key: 'isCalculated',
    align: 'left',
    width: -1,
  },
  {
    title: '计算版本',
    key: 'calcVersion',
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