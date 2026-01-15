import { cloneDeep } from 'lodash-es';
import { FormSchema } from '@/components/Form';
import { defRangeShortcuts, formatToDate } from '@/utils/dateUtil';

export class State {
  public id = 0; // 自增主键
  public symbol = ''; // 股票代码
  public ii = ''; // ii
  public ei = ''; // ei
  public exchange = ''; // 交易所名称
  public name = ''; // name
  public shortName = ''; // short_name
  public enName = ''; // en_name
  public od = ''; // 上市日期
  public dataUpdateDate = ''; // data_update_date
  public pc = null; // pc
  public up = null; // up
  public dp = null; // dp
  public pk = null; // pk
  public fv = 0; // fv
  public tv = 0; // tv

  public floatRatio = null; // 流通股比例 (%)
  public is = 0; // is
  public tradingStatus = ''; // 交易状态描述
  public industry = ''; // industry
  public sector = ''; // sector
  public marketType = ''; // market_type
  public dataSource = ''; // data_source
  public isActive = 1; // is_active
  public version = 1; // version
  public createdAt = ''; // created_at
  public updatedAt = ''; // updated_at

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
    message: '请输入股票代码',
  },
  ii: {
    required: true,
    trigger: ['blur', 'input'],
    type: 'string',
    message: '请输入ii',
  },
  ei: {
    required: true,
    trigger: ['blur', 'input'],
    type: 'string',
    message: '请输入ei',
  },
  name: {
    required: true,
    trigger: ['blur', 'input'],
    type: 'string',
    message: '请输入name',
  },
  od: {
    required: true,
    trigger: ['blur', 'input'],
    type: 'string',
    message: '请输入上市日期',
  },
  dataUpdateDate: {
    required: true,
    trigger: ['blur', 'input'],
    type: 'string',
    message: '请输入data_update_date',
  },
  is: {
    required: true,
    trigger: ['blur', 'input'],
    type: 'number',
    message: '请输入is',
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
    label: 'created_at',
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
    title: '股票代码',
    key: 'symbol',
    align: 'left',
    width: -1,
  },
  {
    title: 'ii',
    key: 'ii',
    align: 'left',
    width: -1,
  },
  {
    title: 'ei',
    key: 'ei',
    align: 'left',
    width: -1,
  },
  {
    title: '交易所名称',
    key: 'exchange',
    align: 'left',
    width: -1,
  },
  {
    title: 'name',
    key: 'name',
    align: 'left',
    width: -1,
  },
  {
    title: 'short_name',
    key: 'shortName',
    align: 'left',
    width: -1,
  },
  {
    title: 'en_name',
    key: 'enName',
    align: 'left',
    width: -1,
  },
  {
    title: '上市日期',
    key: 'od',
    align: 'left',
    width: -1,
    render(row: State) {
      return formatToDate(row.od);
    },
  },
  {
    title: 'data_update_date',
    key: 'dataUpdateDate',
    align: 'left',
    width: -1,
    render(row: State) {
      return formatToDate(row.dataUpdateDate);
    },
  },
  {
    title: 'pc',
    key: 'pc',
    align: 'left',
    width: -1,
  },
  {
    title: 'up',
    key: 'up',
    align: 'left',
    width: -1,
  },
  {
    title: 'dp',
    key: 'dp',
    align: 'left',
    width: -1,
  },
  {
    title: 'pk',
    key: 'pk',
    align: 'left',
    width: -1,
  },
  {
    title: 'fv',
    key: 'fv',
    align: 'left',
    width: -1,
  },
  {
    title: 'tv',
    key: 'tv',
    align: 'left',
    width: -1,
  },
  {
    title: '流通股比例 (%)',
    key: 'floatRatio',
    align: 'left',
    width: -1,
  },
  {
    title: 'is',
    key: 'is',
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
    title: 'industry',
    key: 'industry',
    align: 'left',
    width: -1,
  },
  {
    title: 'sector',
    key: 'sector',
    align: 'left',
    width: -1,
  },
  {
    title: 'market_type',
    key: 'marketType',
    align: 'left',
    width: -1,
  },
  {
    title: 'data_source',
    key: 'dataSource',
    align: 'left',
    width: -1,
  },
  {
    title: 'is_active',
    key: 'isActive',
    align: 'left',
    width: -1,
  },
  {
    title: 'version',
    key: 'version',
    align: 'left',
    width: -1,
  },
  {
    title: 'created_at',
    key: 'createdAt',
    align: 'left',
    width: -1,
  },
  {
    title: 'updated_at',
    key: 'updatedAt',
    align: 'left',
    width: -1,
  },
];