import { cloneDeep } from 'lodash-es';
import { FormSchema } from '@/components/Form';
import { defRangeShortcuts } from '@/utils/dateUtil';

export class State {
  public id = 0; // 自增主键

  public symbol = ''; // 股票或标的代码 (例如: AAPL, 000001.SZ)
  public t = ''; // 交易时间 (统一为datetime类型，日线数据时间部分设为00:00:00)
  public intervalType = 'day'; // 数据间隔: minute-短分时, day-日线
  public ma3 = null; // MA3值
  public ma5 = null; // MA5值
  public ma10 = null; // MA10值
  public ma15 = null; // MA15值
  public ma20 = null; // MA20值
  public ma30 = null; // MA30值
  public ma60 = null; // MA60值
  public ma120 = null; // MA120值
  public ma200 = null; // MA200值
  public ma250 = null; // MA250值
  public createdAt = ''; // 数据创建时间
  public updatedAt = ''; // 数据更新时间

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
    message: '请输入股票或标的代码 (例如: AAPL, 000001.SZ)',
  },
  t: {
    required: true,
    trigger: ['blur', 'input'],
    type: 'string',
    message: '请输入交易时间 (统一为datetime类型，日线数据时间部分设为00:00:00)',
  },
  intervalType: {
    required: true,
    trigger: ['blur', 'input'],
    type: 'string',
    message: '请输入数据间隔: minute-短分时, day-日线',
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
    label: '数据创建时间',
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
    title: '股票或标的代码 (例如: AAPL, 000001.SZ)',
    key: 'symbol',
    align: 'left',
    width: -1,
  },
  {
    title: '交易时间 (统一为datetime类型，日线数据时间部分设为00:00:00)',
    key: 't',
    align: 'left',
    width: -1,
  },
  {
    title: '数据间隔: minute-短分时, day-日线',
    key: 'intervalType',
    align: 'left',
    width: -1,
  },
  {
    title: 'MA3值',
    key: 'ma3',
    align: 'left',
    width: -1,
  },
  {
    title: 'MA5值',
    key: 'ma5',
    align: 'left',
    width: -1,
  },
  {
    title: 'MA10值',
    key: 'ma10',
    align: 'left',
    width: -1,
  },
  {
    title: 'MA15值',
    key: 'ma15',
    align: 'left',
    width: -1,
  },
  {
    title: 'MA20值',
    key: 'ma20',
    align: 'left',
    width: -1,
  },
  {
    title: 'MA30值',
    key: 'ma30',
    align: 'left',
    width: -1,
  },
  {
    title: 'MA60值',
    key: 'ma60',
    align: 'left',
    width: -1,
  },
  {
    title: 'MA120值',
    key: 'ma120',
    align: 'left',
    width: -1,
  },
  {
    title: 'MA200值',
    key: 'ma200',
    align: 'left',
    width: -1,
  },
  {
    title: 'MA250值',
    key: 'ma250',
    align: 'left',
    width: -1,
  },
  {
    title: '数据创建时间',
    key: 'createdAt',
    align: 'left',
    width: -1,
  },
  {
    title: '数据更新时间',
    key: 'updatedAt',
    align: 'left',
    width: -1,
  },
];