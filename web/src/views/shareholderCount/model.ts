import { cloneDeep } from 'lodash-es';
import { FormSchema } from '@/components/Form';
import { defRangeShortcuts, formatToDate } from '@/utils/dateUtil';

export class State {
  public id = 0; // 自增主键

  public symbol = ''; // 公司代码/股票代码 (例如: 000001.SZ, 600000.SS)
  public jzrq = ''; // 截止日期 (报告期结束日, 如2023-09-30)
  public reportYear = 0; // 报告年度
  public reportQuarter = 0; // 报告季度 (1-4)
  public reportType = 'quarter'; // 报告类型: annual-年报, half_year-中报, quarter-季报
  public gdzs = 0; // 股东总数 (户)
  public agdhs = 0; // A股东户数 (户)
  public bgdhs = 0; // B股东户数 (户)
  public hgdhs = 0; // H股东户数 (户)
  public yltgdhs = 0; // 已流通股东户数 (户)
  public wltgdhs = 0; // 未流通股东户数 (户)
  public agRatio = null; // A股股东占比(%)
  public yltRatio = null; // 已流通股东占比(%)
  public dataSource = ''; // 数据来源
  public isLatest = 1; // 是否为该报告期最新数据
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
    message: '请输入公司代码/股票代码 (例如: 000001.SZ, 600000.SS)',
  },
  jzrq: {
    required: true,
    trigger: ['blur', 'input'],
    type: 'string',
    message: '请输入截止日期 (报告期结束日, 如2023-09-30)',
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
    message: '请输入报告类型: annual-年报, half_year-中报, quarter-季报',
  },
  gdzs: {
    required: true,
    trigger: ['blur', 'input'],
    type: 'number',
    message: '请输入股东总数 (户)',
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
    title: '公司代码/股票代码 (例如: 000001.SZ, 600000.SS)',
    key: 'symbol',
    align: 'left',
    width: -1,
  },
  {
    title: '截止日期 (报告期结束日, 如2023-09-30)',
    key: 'jzrq',
    align: 'left',
    width: -1,
    render(row: State) {
      return formatToDate(row.jzrq);
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
    title: '报告类型: annual-年报, half_year-中报, quarter-季报',
    key: 'reportType',
    align: 'left',
    width: -1,
  },
  {
    title: '股东总数 (户)',
    key: 'gdzs',
    align: 'left',
    width: -1,
  },
  {
    title: 'A股东户数 (户)',
    key: 'agdhs',
    align: 'left',
    width: -1,
  },
  {
    title: 'B股东户数 (户)',
    key: 'bgdhs',
    align: 'left',
    width: -1,
  },
  {
    title: 'H股东户数 (户)',
    key: 'hgdhs',
    align: 'left',
    width: -1,
  },
  {
    title: '已流通股东户数 (户)',
    key: 'yltgdhs',
    align: 'left',
    width: -1,
  },
  {
    title: '未流通股东户数 (户)',
    key: 'wltgdhs',
    align: 'left',
    width: -1,
  },
  {
    title: 'A股股东占比(%)',
    key: 'agRatio',
    align: 'left',
    width: -1,
  },
  {
    title: '已流通股东占比(%)',
    key: 'yltRatio',
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
    title: '是否为该报告期最新数据',
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