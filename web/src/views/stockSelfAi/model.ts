import { cloneDeep } from 'lodash-es';
import { FormSchema } from '@/components/Form';
import { defRangeShortcuts } from '@/utils/dateUtil';

export class State {
  public id = 0; // 自增主键
  public name = ''; // ai名称 例如deepseek 千问
  public model = ''; // ai model
  public baseUrl = ''; // ai base url
  public apiKey = ''; // ai api key
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
  name: {
    required: true,
    trigger: ['blur', 'input'],
    type: 'string',
    message: '请输入ai名称 例如deepseek 千问',
  },
  model: {
    required: true,
    trigger: ['blur', 'input'],
    type: 'string',
    message: '请输入ai model',
  },
  baseUrl: {
    required: true,
    trigger: ['blur', 'input'],
    type: 'string',
    message: '请输入ai base url',
  },
  apiKey: {
    required: true,
    trigger: ['blur', 'input'],
    type: 'string',
    message: '请输入ai api key',
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
    title: 'ai名称 例如deepseek 千问',
    key: 'name',
    align: 'left',
    width: -1,
  },
  {
    title: 'ai model',
    key: 'model',
    align: 'left',
    width: -1,
  },
  {
    title: 'ai base url',
    key: 'baseUrl',
    align: 'left',
    width: -1,
  },
  {
    title: 'ai api key',
    key: 'apiKey',
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