import { http, jumpExport } from '@/utils/http/axios';

// 获取自选股票列表
export function List(params) {
  return http.request({
    url: '/stockSelfCode/list',
    method: 'get',
    params,
  });
}

// 删除/批量删除自选股票
export function Delete(params) {
  return http.request({
    url: '/stockSelfCode/delete',
    method: 'POST',
    params,
  });
}

// 添加/编辑自选股票
export function Edit(params) {
  return http.request({
    url: '/stockSelfCode/edit',
    method: 'POST',
    params,
  });
}

// 获取自选股票指定详情
export function View(params) {
  return http.request({
    url: '/stockSelfCode/view',
    method: 'GET',
    params,
  });
}

// 导出自选股票
export function Export(params) {
  jumpExport('/stockSelfCode/export', params);
}