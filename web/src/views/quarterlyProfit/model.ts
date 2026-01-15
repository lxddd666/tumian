import { cloneDeep } from 'lodash-es';
import { FormSchema } from '@/components/Form';
import { defRangeShortcuts, formatToDate } from '@/utils/dateUtil';

export class State {
  public id = 0; // 自增主键

  public symbol = ''; // 股票代码 (如: 000001.SZ)
  public date = ''; // 截止日期 (报告期截止日，如2025-03-31)
  public reportYear = 0; // 报告年度
  public reportQuarter = 0; // 报告季度 (1-4)
  public fiscalPeriod = ''; // 会计期间 (如2025Q1)
  public income = null; // 营业收入（万元）
  public expend = null; // 营业支出（万元）
  public profit = null; // 营业利润（万元）
  public totalp = null; // 利润总额（万元）
  public reprofit = null; // 净利润（万元）
  public basege = null; // 基本每股收益(元/股)
  public ettege = null; // 稀释每股收益(元/股)
  public otherp = null; // 其他综合收益（万元）
  public totalcp = null; // 综合收益总额（万元）
  public grossProfitMargin = null; // 毛利率(%)
  public netProfitMargin = null; // 净利率(%)
  public operatingProfitRatio = null; // 营业利润率(%)
  public reportType = '季报'; // 报告类型: 一季报, 中报, 三季报, 年报
  public dataSource = ''; // 数据来源 (如: 交易所财报)
  public currency = 'CNY'; // 货币单位
  public isAudited = 0; // 是否审计: 0-未审计, 1-已审计
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
  date: {
    required: true,
    trigger: ['blur', 'input'],
    type: 'string',
    message: '请输入截止日期 (报告期截止日，如2025-03-31)',
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
    title: '截止日期 (报告期截止日，如2025-03-31)',
    key: 'date',
    align: 'left',
    width: -1,
    render(row: State) {
      return formatToDate(row.date);
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
    title: '会计期间 (如2025Q1)',
    key: 'fiscalPeriod',
    align: 'left',
    width: -1,
  },
  {
    title: '营业收入（万元）',
    key: 'income',
    align: 'left',
    width: -1,
  },
  {
    title: '营业支出（万元）',
    key: 'expend',
    align: 'left',
    width: -1,
  },
  {
    title: '营业利润（万元）',
    key: 'profit',
    align: 'left',
    width: -1,
  },
  {
    title: '利润总额（万元）',
    key: 'totalp',
    align: 'left',
    width: -1,
  },
  {
    title: '净利润（万元）',
    key: 'reprofit',
    align: 'left',
    width: -1,
  },
  {
    title: '基本每股收益(元/股)',
    key: 'basege',
    align: 'left',
    width: -1,
  },
  {
    title: '稀释每股收益(元/股)',
    key: 'ettege',
    align: 'left',
    width: -1,
  },
  {
    title: '其他综合收益（万元）',
    key: 'otherp',
    align: 'left',
    width: -1,
  },
  {
    title: '综合收益总额（万元）',
    key: 'totalcp',
    align: 'left',
    width: -1,
  },
  {
    title: '毛利率(%)',
    key: 'grossProfitMargin',
    align: 'left',
    width: -1,
  },
  {
    title: '净利率(%)',
    key: 'netProfitMargin',
    align: 'left',
    width: -1,
  },
  {
    title: '营业利润率(%)',
    key: 'operatingProfitRatio',
    align: 'left',
    width: -1,
  },
  {
    title: '报告类型: 一季报, 中报, 三季报, 年报',
    key: 'reportType',
    align: 'left',
    width: -1,
  },
  {
    title: '数据来源 (如: 交易所财报)',
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
    title: '是否审计: 0-未审计, 1-已审计',
    key: 'isAudited',
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