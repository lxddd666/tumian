import { cloneDeep } from 'lodash-es';
import { FormSchema } from '@/components/Form';
import { defRangeShortcuts, formatToDate } from '@/utils/dateUtil';

export class State {
  public id = 0; // 自增主键

  public jzrq = ''; // 截止日期 (报告期，如2025-12-31)
  public t = ''; // 交易时间 (衍生自jzrq，兼容时间序列查询)
  public jjmc = ''; // 基金名称
  public jjdm = ''; // 基金代码
  public symbol = ''; // 股票代码 (如: 000001.SZ)
  public ccsl = 0; // 持仓数量(股)
  public ltbl = null; // 占流通股比例(%)
  public cgsz = null; // 持股市值（元）
  public jzbl = null; // 占净值比例（%）
  public avgCost = null; // 估算持仓成本 (元/股)
  public dataSource = ''; // 数据来源 (如: 基金季报)
  public reportType = '季报'; // 报告类型: 季报, 中报, 年报
  public isLatest = 0; // 是否为该基金对该股票的最新持仓: 0-否, 1-是
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
  jzrq: {
    required: true,
    trigger: ['blur', 'input'],
    type: 'string',
    message: '请输入截止日期 (报告期，如2025-12-31)',
  },
  jjmc: {
    required: true,
    trigger: ['blur', 'input'],
    type: 'string',
    message: '请输入基金名称',
  },
  jjdm: {
    required: true,
    trigger: ['blur', 'input'],
    type: 'string',
    message: '请输入基金代码',
  },
  symbol: {
    required: true,
    trigger: ['blur', 'input'],
    type: 'string',
    message: '请输入股票代码 (如: 000001.SZ)',
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
    title: '截止日期 (报告期，如2025-12-31)',
    key: 'jzrq',
    align: 'left',
    width: -1,
    render(row: State) {
      return formatToDate(row.jzrq);
    },
  },
  {
    title: '交易时间 (衍生自jzrq，兼容时间序列查询)',
    key: 't',
    align: 'left',
    width: -1,
  },
  {
    title: '基金名称',
    key: 'jjmc',
    align: 'left',
    width: -1,
  },
  {
    title: '基金代码',
    key: 'jjdm',
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
    title: '持仓数量(股)',
    key: 'ccsl',
    align: 'left',
    width: -1,
  },
  {
    title: '占流通股比例(%)',
    key: 'ltbl',
    align: 'left',
    width: -1,
  },
  {
    title: '持股市值（元）',
    key: 'cgsz',
    align: 'left',
    width: -1,
  },
  {
    title: '占净值比例（%）',
    key: 'jzbl',
    align: 'left',
    width: -1,
  },
  {
    title: '估算持仓成本 (元/股)',
    key: 'avgCost',
    align: 'left',
    width: -1,
  },
  {
    title: '数据来源 (如: 基金季报)',
    key: 'dataSource',
    align: 'left',
    width: -1,
  },
  {
    title: '报告类型: 季报, 中报, 年报',
    key: 'reportType',
    align: 'left',
    width: -1,
  },
  {
    title: '是否为该基金对该股票的最新持仓: 0-否, 1-是',
    key: 'isLatest',
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