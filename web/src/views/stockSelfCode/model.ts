import { cloneDeep } from 'lodash-es';
import { FormSchema } from '@/components/Form';
import { defRangeShortcuts } from '@/utils/dateUtil';

export class State {
  public dm = ''; // 股票代码
  public mc = ''; // 股票名称
  public jys = ''; // 交易所
  public createdAt = ''; // 创建时间

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
  mc: {
    required: true,
    trigger: ['blur', 'input'],
    type: 'string',
    message: '请输入股票名称',
  },
};

// 表格搜索表单
export const schemas = ref<FormSchema[]>([
  {
    field: 'dm',
    component: 'NInput',
    label: '股票代码',
    componentProps: {
      placeholder: '请输入股票代码',
      onUpdateValue: (e: any) => {
        console.log(e);
      },
    },
  },
  {
    field: 'jys',
    component: 'NInput',
    label: '交易所',
    componentProps: {
      placeholder: '请输入交易所',
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
    title: '股票代码',
    key: 'dm',
    align: 'left',
    width: -1,
  },
  {
    title: '股票名称',
    key: 'mc',
    align: 'left',
    width: -1,
  },
  {
    title: '交易所',
    key: 'jys',
    align: 'left',
    width: -1,
  },
  {
    title: '创建时间',
    key: 'createdAt',
    align: 'left',
    width: -1,
  },
];