import { cloneDeep } from 'lodash-es';
import { FormSchema } from '@/components/Form';
import { defRangeShortcuts } from '@/utils/dateUtil';

export class State {
  public id = 0; // 自增主键

  public symbol = ''; // 股票或标的代码 (例如: AAPL, 000001.SZ)
  public t = ''; // 时间
  public aiId = 0; // ai id
  public indicatorsJudgment = ''; // 指标判断
  public indicatorsFlag = 0; // 指标判断 1是2否
  public financialJudgment = ''; // 财务判断
  public financialFlag = 0; // 财务判断 1是2否
  public comprehensiveJudgment = ''; // 综合判断
  public comprehensiveFlag = 0; // 综合判断 1是2否
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
    message: '请输入时间',
  },
  aiId: {
    required: true,
    trigger: ['blur', 'input'],
    type: 'number',
    message: '请输入ai id',
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
    title: '股票或标的代码 (例如: AAPL, 000001.SZ)',
    key: 'symbol',
    align: 'left',
    width: -1,
  },
  {
    title: '时间',
    key: 't',
    align: 'left',
    width: -1,
  },
  {
    title: 'ai id',
    key: 'aiId',
    align: 'left',
    width: -1,
  },
  {
    title: '指标判断 1是2否',
    key: 'indicatorsFlag',
    align: 'left',
    width: -1,
  },
  {
    title: '财务判断 1是2否',
    key: 'financialFlag',
    align: 'left',
    width: -1,
  },
  {
    title: '综合判断 1是2否',
    key: 'comprehensiveFlag',
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