import { cloneDeep } from 'lodash-es';
import { FormSchema } from '@/components/Form';
import { defRangeShortcuts, formatToDate } from '@/utils/dateUtil';

export class State {
  public id = 0; // 自增主键

  public symbol = ''; // 证券代码 (如: 000001.SZ, AAPL)
  public t = ''; // 交易时间 (精确到分钟或日)
  public date = ''; // 交易日期 (衍生字段)
  public year = 0; // 交易年份
  public month = 0; // 交易月份
  public weekday = 0; // 星期几 (1=周日,7=周六)
  public o = null; // 开盘价
  public h = null; // 最高价
  public l = null; // 最低价
  public c = null; // 收盘价
  public pc = null; // 前收盘价
  public v = 0; // 成交量 (股/手)
  public a = null; // 成交额 (元)
  public change = null; // 涨跌额
  public changePct = null; // 涨跌幅 (%)
  public amplitude = null; // 振幅 (%)
  public sf = 0; // 停牌标志: 0-正常, 1-停牌
  public tradingStatus = ''; // 交易状态描述
  public period = 'day'; // 数据周期: 1min, 5min, 15min, 30min, 60min, day, week, month
  public isAdjusted = 0; // 是否复权: 0-不复权, 1-前复权, 2-后复权
  public dataQuality = 1; // 数据质量: 0-异常, 1-正常, 2-补全
  public isVerified = 0; // 是否已验证: 0-未验证, 1-已验证
  public dataSource = ''; // 数据来源
  public version = 1; // 数据版本
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
    message: '请输入证券代码 (如: 000001.SZ, AAPL)',
  },
  t: {
    required: true,
    trigger: ['blur', 'input'],
    type: 'string',
    message: '请输入交易时间 (精确到分钟或日)',
  },
  sf: {
    required: true,
    trigger: ['blur', 'input'],
    type: 'number',
    message: '请输入停牌标志: 0-正常, 1-停牌',
  },
  period: {
    required: true,
    trigger: ['blur', 'input'],
    type: 'string',
    message: '请输入数据周期: 1min, 5min, 15min, 30min, 60min, day, week, month',
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
    title: '证券代码 (如: 000001.SZ, AAPL)',
    key: 'symbol',
    align: 'left',
    width: -1,
  },
  {
    title: '交易时间 (精确到分钟或日)',
    key: 't',
    align: 'left',
    width: -1,
  },
  {
    title: '交易日期 (衍生字段)',
    key: 'date',
    align: 'left',
    width: -1,
    render(row: State) {
      return formatToDate(row.date);
    },
  },
  {
    title: '交易年份',
    key: 'year',
    align: 'left',
    width: -1,
  },
  {
    title: '交易月份',
    key: 'month',
    align: 'left',
    width: -1,
  },
  {
    title: '星期几 (1=周日,7=周六)',
    key: 'weekday',
    align: 'left',
    width: -1,
  },
  {
    title: '开盘价',
    key: 'o',
    align: 'left',
    width: -1,
  },
  {
    title: '最高价',
    key: 'h',
    align: 'left',
    width: -1,
  },
  {
    title: '最低价',
    key: 'l',
    align: 'left',
    width: -1,
  },
  {
    title: '收盘价',
    key: 'c',
    align: 'left',
    width: -1,
  },
  {
    title: '前收盘价',
    key: 'pc',
    align: 'left',
    width: -1,
  },
  {
    title: '成交量 (股/手)',
    key: 'v',
    align: 'left',
    width: -1,
  },
  {
    title: '成交额 (元)',
    key: 'a',
    align: 'left',
    width: -1,
  },
  {
    title: '涨跌额',
    key: 'change',
    align: 'left',
    width: -1,
  },
  {
    title: '涨跌幅 (%)',
    key: 'changePct',
    align: 'left',
    width: -1,
  },
  {
    title: '振幅 (%)',
    key: 'amplitude',
    align: 'left',
    width: -1,
  },
  {
    title: '停牌标志: 0-正常, 1-停牌',
    key: 'sf',
    align: 'left',
    width: -1,
  },
  {
    title: '交易状态描述',
    key: 'tradingStatus',
    align: 'left',
    width: -1,
  },
  {
    title: '数据周期: 1min, 5min, 15min, 30min, 60min, day, week, month',
    key: 'period',
    align: 'left',
    width: -1,
  },
  {
    title: '是否复权: 0-不复权, 1-前复权, 2-后复权',
    key: 'isAdjusted',
    align: 'left',
    width: -1,
  },
  {
    title: '数据质量: 0-异常, 1-正常, 2-补全',
    key: 'dataQuality',
    align: 'left',
    width: -1,
  },
  {
    title: '是否已验证: 0-未验证, 1-已验证',
    key: 'isVerified',
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
];