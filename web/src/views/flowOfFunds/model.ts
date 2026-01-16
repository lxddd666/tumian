import { cloneDeep } from 'lodash-es';
import { FormSchema } from '@/components/Form';
import { defRangeShortcuts } from '@/utils/dateUtil';

export class State {
  public id = 0; // 主键ID

  public symbol = ''; // 股票代码 (如: 000001.SZ)
  public t = 0; // 交易时间 (通常为HHMMSS格式的整数)
  public zmbzds = 0; // 主买单总单数
  public zmszds = 0; // 主卖单总单数
  public dddx = null; // 大单动向
  public zddy = null; // 涨跌动因
  public ddcf = null; // 大单差分
  public zmbzdszl = 0; // 主买单总单数增量
  public zmszdszl = 0; // 主卖单总单数增量
  public cjbszl = 0; // 成交笔数增量
  public zmbtdcje = null; // 主买特大单成交额
  public zmbddcje = null; // 主买大单成交额
  public zmbzdcje = null; // 主买中单成交额
  public zmbxdcje = null; // 主买小单成交额
  public zmstdcje = null; // 主卖特大单成交额
  public zmsddcje = null; // 主卖大单成交额
  public zmszdcje = null; // 主卖中单成交额
  public zmsxdcje = null; // 主卖小单成交额
  public bdmbtdcje = null; // 被动买特大单成交额
  public bdmbddcje = null; // 被动买大单成交额
  public bdmbzdcje = null; // 被动买中单成交额
  public bdmbxdcje = null; // 被动买小单成交额
  public bdmstdcje = null; // 被动卖特大单成交额
  public bdmsddcje = null; // 被动卖大单成交额
  public bdmszdcje = null; // 被动卖中单成交额
  public bdmsxdcje = null; // 被动卖小单成交额
  public zmbtdcjl = 0; // 主买特大单成交量
  public zmbddcjl = 0; // 主买大单成交量
  public zmbzdcjl = 0; // 主买中单成交量
  public zmbxdcjl = 0; // 主买小单成交量
  public zmstdcjl = 0; // 主卖特大单成交量
  public zmsddcjl = 0; // 主卖大单成交量
  public zmszdcjl = 0; // 主卖中单成交量
  public zmsxdcjl = 0; // 主卖小单成交量
  public bdmbtdcjl = 0; // 被动买特大单成交量
  public bdmbddcjl = 0; // 被动买大单成交量
  public bdmbzdcjl = 0; // 被动买中单成交量
  public bdmbxdcjl = 0; // 被动买小单成交量
  public bdmstdcjl = 0; // 被动卖特大单成交量
  public bdmsddcjl = 0; // 被动卖大单成交量
  public bdmszdcjl = 0; // 被动卖中单成交量
  public bdmsxdcjl = 0; // 被动卖小单成交量
  public zmbtdcjzl = null; // 主买特大单成交额增量
  public zmbddcjzl = null; // 主买大单成交额增量
  public zmbtdcjzlv = 0; // 主买特大单成交量增量
  public zmbddcjzlv = 0; // 主买大单成交量增量
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
  t: {
    required: true,
    trigger: ['blur', 'input'],
    type: 'number',
    message: '请输入交易时间 (通常为HHMMSS格式的整数)',
  },
};

// 表格搜索表单
export const schemas = ref<FormSchema[]>([
  {
    field: 'id',
    component: 'NInputNumber',
    label: '主键ID',
    componentProps: {
      placeholder: '请输入主键ID',
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
    title: '主键ID',
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
    title: '交易时间 (通常为HHMMSS格式的整数)',
    key: 't',
    align: 'left',
    width: -1,
  },
  {
    title: '主买单总单数',
    key: 'zmbzds',
    align: 'left',
    width: -1,
  },
  {
    title: '主卖单总单数',
    key: 'zmszds',
    align: 'left',
    width: -1,
  },
  {
    title: '大单动向',
    key: 'dddx',
    align: 'left',
    width: -1,
  },
  {
    title: '涨跌动因',
    key: 'zddy',
    align: 'left',
    width: -1,
  },
  {
    title: '大单差分',
    key: 'ddcf',
    align: 'left',
    width: -1,
  },
  {
    title: '主买单总单数增量',
    key: 'zmbzdszl',
    align: 'left',
    width: -1,
  },
  {
    title: '主卖单总单数增量',
    key: 'zmszdszl',
    align: 'left',
    width: -1,
  },
  {
    title: '成交笔数增量',
    key: 'cjbszl',
    align: 'left',
    width: -1,
  },
  {
    title: '主买特大单成交额',
    key: 'zmbtdcje',
    align: 'left',
    width: -1,
  },
  {
    title: '主买大单成交额',
    key: 'zmbddcje',
    align: 'left',
    width: -1,
  },
  {
    title: '主买中单成交额',
    key: 'zmbzdcje',
    align: 'left',
    width: -1,
  },
  {
    title: '主买小单成交额',
    key: 'zmbxdcje',
    align: 'left',
    width: -1,
  },
  {
    title: '主卖特大单成交额',
    key: 'zmstdcje',
    align: 'left',
    width: -1,
  },
  {
    title: '主卖大单成交额',
    key: 'zmsddcje',
    align: 'left',
    width: -1,
  },
  {
    title: '主卖中单成交额',
    key: 'zmszdcje',
    align: 'left',
    width: -1,
  },
  {
    title: '主卖小单成交额',
    key: 'zmsxdcje',
    align: 'left',
    width: -1,
  },
  {
    title: '被动买特大单成交额',
    key: 'bdmbtdcje',
    align: 'left',
    width: -1,
  },
  {
    title: '被动买大单成交额',
    key: 'bdmbddcje',
    align: 'left',
    width: -1,
  },
  {
    title: '被动买中单成交额',
    key: 'bdmbzdcje',
    align: 'left',
    width: -1,
  },
  {
    title: '被动买小单成交额',
    key: 'bdmbxdcje',
    align: 'left',
    width: -1,
  },
  {
    title: '被动卖特大单成交额',
    key: 'bdmstdcje',
    align: 'left',
    width: -1,
  },
  {
    title: '被动卖大单成交额',
    key: 'bdmsddcje',
    align: 'left',
    width: -1,
  },
  {
    title: '被动卖中单成交额',
    key: 'bdmszdcje',
    align: 'left',
    width: -1,
  },
  {
    title: '被动卖小单成交额',
    key: 'bdmsxdcje',
    align: 'left',
    width: -1,
  },
  {
    title: '主买特大单成交量',
    key: 'zmbtdcjl',
    align: 'left',
    width: -1,
  },
  {
    title: '主买大单成交量',
    key: 'zmbddcjl',
    align: 'left',
    width: -1,
  },
  {
    title: '主买中单成交量',
    key: 'zmbzdcjl',
    align: 'left',
    width: -1,
  },
  {
    title: '主买小单成交量',
    key: 'zmbxdcjl',
    align: 'left',
    width: -1,
  },
  {
    title: '主卖特大单成交量',
    key: 'zmstdcjl',
    align: 'left',
    width: -1,
  },
  {
    title: '主卖大单成交量',
    key: 'zmsddcjl',
    align: 'left',
    width: -1,
  },
  {
    title: '主卖中单成交量',
    key: 'zmszdcjl',
    align: 'left',
    width: -1,
  },
  {
    title: '主卖小单成交量',
    key: 'zmsxdcjl',
    align: 'left',
    width: -1,
  },
  {
    title: '被动买特大单成交量',
    key: 'bdmbtdcjl',
    align: 'left',
    width: -1,
  },
  {
    title: '被动买大单成交量',
    key: 'bdmbddcjl',
    align: 'left',
    width: -1,
  },
  {
    title: '被动买中单成交量',
    key: 'bdmbzdcjl',
    align: 'left',
    width: -1,
  },
  {
    title: '被动买小单成交量',
    key: 'bdmbxdcjl',
    align: 'left',
    width: -1,
  },
  {
    title: '被动卖特大单成交量',
    key: 'bdmstdcjl',
    align: 'left',
    width: -1,
  },
  {
    title: '被动卖大单成交量',
    key: 'bdmsddcjl',
    align: 'left',
    width: -1,
  },
  {
    title: '被动卖中单成交量',
    key: 'bdmszdcjl',
    align: 'left',
    width: -1,
  },
  {
    title: '被动卖小单成交量',
    key: 'bdmsxdcjl',
    align: 'left',
    width: -1,
  },
  {
    title: '主买特大单成交额增量',
    key: 'zmbtdcjzl',
    align: 'left',
    width: -1,
  },
  {
    title: '主买大单成交额增量',
    key: 'zmbddcjzl',
    align: 'left',
    width: -1,
  },
  {
    title: '主买特大单成交量增量',
    key: 'zmbtdcjzlv',
    align: 'left',
    width: -1,
  },
  {
    title: '主买大单成交量增量',
    key: 'zmbddcjzlv',
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