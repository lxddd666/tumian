import { http, jumpExport } from '@/utils/http/axios';

// 获取移动平均线(MA)指标数据表列表
export function List(params) {
  return http.request({
    url: '/maData/list',
    method: 'get',
    params,
  });
}

// 删除/批量删除移动平均线(MA)指标数据表
export function Delete(params) {
  return http.request({
    url: '/maData/delete',
    method: 'POST',
    params,
  });
}

// 添加/编辑移动平均线(MA)指标数据表
export function Edit(params) {
  return http.request({
    url: '/maData/edit',
    method: 'POST',
    params,
  });
}

// 获取移动平均线(MA)指标数据表指定详情
export function View(params) {
  return http.request({
    url: '/maData/view',
    method: 'GET',
    params,
  });
}

// 导出移动平均线(MA)指标数据表
export function Export(params) {
  jumpExport('/maData/export', params);
}