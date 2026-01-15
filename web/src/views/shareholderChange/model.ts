import { cloneDeep } from 'lodash-es';
import { FormSchema } from '@/components/Form';
import { defRangeShortcuts, formatToDate } from '@/utils/dateUtil';

export class State {
  public id = 0; // 自增主键

  public symbol = ''; // 股票代码 (如: 000001.SZ)
  public jzrq = ''; // 截止日期 (统计截止日，如2025-12-31)[citation:9]
  public gdhs = 0; // 股东户数 (统计截止日的总户数)[citation:3][citation:6]
  public bh = null; // 比上期变化情况 (通常为百分比，如 -6.23 表示减少6.23%)[citation:9]
  public changeDirection = ''; // 变化方向 (衍生字段)
  public dataSource = ''; // 数据来源 (如: 交易所公告[citation:6]、互动平台[citation:3])
  public annDate = ''; // 公告日期 (信息发布日期)[citation:1]
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
    message: '请输入截止日期 (统计截止日，如2025-12-31)[citation:9]',
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
    title: '截止日期 (统计截止日，如2025-12-31)[citation:9]',
    key: 'jzrq',
    align: 'left',
    width: -1,
    render(row: State) {
      return formatToDate(row.jzrq);
    },
  },
  {
    title: '股东户数 (统计截止日的总户数)[citation:3][citation:6]',
    key: 'gdhs',
    align: 'left',
    width: -1,
  },
  {
    title: '比上期变化情况 (通常为百分比，如 -6.23 表示减少6.23%)[citation:9]',
    key: 'bh',
    align: 'left',
    width: -1,
  },
  {
    title: '变化方向 (衍生字段)',
    key: 'changeDirection',
    align: 'left',
    width: -1,
  },
  {
    title: '数据来源 (如: 交易所公告[citation:6]、互动平台[citation:3])',
    key: 'dataSource',
    align: 'left',
    width: -1,
  },
  {
    title: '公告日期 (信息发布日期)[citation:1]',
    key: 'annDate',
    align: 'left',
    width: -1,
    render(row: State) {
      return formatToDate(row.annDate);
    },
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