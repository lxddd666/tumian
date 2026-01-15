import { http, jumpExport } from '@/utils/http/axios';

// 获取资产负债表列表
export function List(params) {
  return http.request({
    url: '/balanceSheet/list',
    method: 'get',
    params,
  });
}

// 删除/批量删除资产负债表
export function Delete(params) {
  return http.request({
    url: '/balanceSheet/delete',
    method: 'POST',
    params,
  });
}

// 添加/编辑资产负债表
export function Edit(params) {
  return http.request({
    url: '/balanceSheet/edit',
    method: 'POST',
    params,
  });
}

// 获取资产负债表指定详情
export function View(params) {
  return http.request({
    url: '/balanceSheet/view',
    method: 'GET',
    params,
  });
}

// 导出资产负债表
export function Export(params) {
  jumpExport('/balanceSheet/export', params);
}