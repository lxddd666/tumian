import { cloneDeep } from 'lodash-es';
import { FormSchema } from '@/components/Form';
import { defRangeShortcuts, formatToDate } from '@/utils/dateUtil';
import { renderOptionTag } from '@/utils';
import { useDictStore } from '@/store/modules/dict';

const dict = useDictStore();

export class State {
  public id = 0; // 自增主键

  public dm = ''; // 股票代码 (唯一业务标识，如: 000001)
  public mc = ''; // 股票名称 (如: 平安银行)
  public jys = ''; // 交易所代码 (如: sh, sz, bj)
  public exchangeName = ''; // 交易所全称
  public symbol = ''; // 标准股票代码 (如: 000001.SZ)
  public status = 1; // 状态: 1-正常, 0-退市
  public listDate = ''; // 上市日期
  public dataSource = ''; // 数据来源
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
  dm: {
    required: true,
    trigger: ['blur', 'input'],
    type: 'string',
    message: '请输入股票代码 (唯一业务标识，如: 000001)',
  },
  mc: {
    required: true,
    trigger: ['blur', 'input'],
    type: 'string',
    message: '请输入股票名称 (如: 平安银行)',
  },
  jys: {
    required: true,
    trigger: ['blur', 'input'],
    type: 'string',
    message: '请输入交易所代码 (如: sh, sz, bj)',
  },
  status: {
    required: true,
    trigger: ['blur', 'input'],
    type: 'number',
    message: '请输入状态: 1-正常, 0-退市',
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
    field: 'status',
    component: 'NSelect',
    label: '状态: 1-正常, 0-退市',
    defaultValue: null,
    componentProps: {
      placeholder: '请选择状态: 1-正常, 0-退市',
      options: dict.getOption('sys_normal_disable'),
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
    title: '股票代码 (唯一业务标识，如: 000001)',
    key: 'dm',
    align: 'left',
    width: -1,
  },
  {
    title: '股票名称 (如: 平安银行)',
    key: 'mc',
    align: 'left',
    width: -1,
  },
  {
    title: '交易所代码 (如: sh, sz, bj)',
    key: 'jys',
    align: 'left',
    width: -1,
  },
  {
    title: '交易所全称',
    key: 'exchangeName',
    align: 'left',
    width: -1,
  },
  {
    title: '标准股票代码 (如: 000001.SZ)',
    key: 'symbol',
    align: 'left',
    width: -1,
  },
  {
    title: '状态: 1-正常, 0-退市',
    key: 'status',
    align: 'left',
    width: -1,
    render(row: State) {
      return renderOptionTag('sys_normal_disable', row.status);
    },
  },
  {
    title: '上市日期',
    key: 'listDate',
    align: 'left',
    width: -1,
    render(row: State) {
      return formatToDate(row.listDate);
    },
  },
  {
    title: '数据来源',
    key: 'dataSource',
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

// 加载字典数据选项
export function loadOptions() {
  dict.loadOptions(['sys_normal_disable']);
}