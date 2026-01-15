import { cloneDeep } from 'lodash-es';
import { FormSchema } from '@/components/Form';
import { defRangeShortcuts, formatToDate } from '@/utils/dateUtil';

export class State {
  public id = 0; // 自增主键

  public symbol = ''; // 公司代码/股票代码 (例如: 000001.SZ)
  public jzrq = ''; // 截止日期 (报告期结束日, 如2023-09-30)[citation:4]
  public ggrq = ''; // 公告日期 (信息发布日期)
  public reportYear = 0; // 报告年度
  public reportQuarter = 0; // 报告季度 (1-4)
  public reportType = 'quarter'; // 报告类型: annual-年报, half_year-中报, quarter-季报[citation:4]
  public gdmc = ''; // 股东名称
  public gdlx = ''; // 股东类型 (如: 基金、社保、个人等)
  public gfxz = ''; // 股份性质 (如: 流通A股、限售A股等)
  public cgsl = 0; // 持股数量 (股)
  public cgbl = null; // 持股比例 (%)
  public cgpm = 0; // 持股排名 (1-10)
  public bdyy = ''; // 变动原因
  public bdType = ''; // 变动类型: increase-增持, decrease-减持, new-新进, unchanged-不变, other-其他
  public dataSource = ''; // 数据来源 (如: 交易所公告)[citation:4]
  public isLatest = 1; // 是否为该报告期最新数据: 0-历史快照, 1-最新
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
    message: '请输入公司代码/股票代码 (例如: 000001.SZ)',
  },
  jzrq: {
    required: true,
    trigger: ['blur', 'input'],
    type: 'string',
    message: '请输入截止日期 (报告期结束日, 如2023-09-30)[citation:4]',
  },
  ggrq: {
    required: true,
    trigger: ['blur', 'input'],
    type: 'string',
    message: '请输入公告日期 (信息发布日期)',
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
    message: '请输入报告类型: annual-年报, half_year-中报, quarter-季报[citation:4]',
  },
  gdmc: {
    required: true,
    trigger: ['blur', 'input'],
    type: 'string',
    message: '请输入股东名称',
  },
  cgpm: {
    required: true,
    trigger: ['blur', 'input'],
    type: 'number',
    message: '请输入持股排名 (1-10)',
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
    title: '公司代码/股票代码 (例如: 000001.SZ)',
    key: 'symbol',
    align: 'left',
    width: -1,
  },
  {
    title: '截止日期 (报告期结束日, 如2023-09-30)[citation:4]',
    key: 'jzrq',
    align: 'left',
    width: -1,
    render(row: State) {
      return formatToDate(row.jzrq);
    },
  },
  {
    title: '公告日期 (信息发布日期)',
    key: 'ggrq',
    align: 'left',
    width: -1,
    render(row: State) {
      return formatToDate(row.ggrq);
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
    title: '报告类型: annual-年报, half_year-中报, quarter-季报[citation:4]',
    key: 'reportType',
    align: 'left',
    width: -1,
  },
  {
    title: '股东名称',
    key: 'gdmc',
    align: 'left',
    width: -1,
  },
  {
    title: '股东类型 (如: 基金、社保、个人等)',
    key: 'gdlx',
    align: 'left',
    width: -1,
  },
  {
    title: '股份性质 (如: 流通A股、限售A股等)',
    key: 'gfxz',
    align: 'left',
    width: -1,
  },
  {
    title: '持股数量 (股)',
    key: 'cgsl',
    align: 'left',
    width: -1,
  },
  {
    title: '持股比例 (%)',
    key: 'cgbl',
    align: 'left',
    width: -1,
  },
  {
    title: '持股排名 (1-10)',
    key: 'cgpm',
    align: 'left',
    width: -1,
  },
  {
    title: '变动类型: increase-增持, decrease-减持, new-新进, unchanged-不变, other-其他',
    key: 'bdType',
    align: 'left',
    width: -1,
  },
  {
    title: '数据来源 (如: 交易所公告)[citation:4]',
    key: 'dataSource',
    align: 'left',
    width: -1,
  },
  {
    title: '是否为该报告期最新数据: 0-历史快照, 1-最新',
    key: 'isLatest',
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