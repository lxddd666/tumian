import { http, jumpExport } from '@/utils/http/axios';

// 获取资金流向明细表列表
export function List(params) {
  return http.request({
    url: '/flowOfFunds/list',
    method: 'get',
    params,
  });
}

// 删除/批量删除资金流向明细表
export function Delete(params) {
  return http.request({
    url: '/flowOfFunds/delete',
    method: 'POST',
    params,
  });
}

// 添加/编辑资金流向明细表
export function Edit(params) {
  return http.request({
    url: '/flowOfFunds/edit',
    method: 'POST',
    params,
  });
}

// 获取资金流向明细表指定详情
export function View(params) {
  return http.request({
    url: '/flowOfFunds/view',
    method: 'GET',
    params,
  });
}

// 导出资金流向明细表
export function Export(params) {
  jumpExport('/flowOfFunds/export', params);
}