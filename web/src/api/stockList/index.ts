import { http, jumpExport } from '@/utils/http/axios';

// 获取股票列表核心表列表
export function List(params) {
  return http.request({
    url: '/stockList/list',
    method: 'get',
    params,
  });
}

// 删除/批量删除股票列表核心表
export function Delete(params) {
  return http.request({
    url: '/stockList/delete',
    method: 'POST',
    params,
  });
}

// 添加/编辑股票列表核心表
export function Edit(params) {
  return http.request({
    url: '/stockList/edit',
    method: 'POST',
    params,
  });
}

// 修改股票列表核心表状态
export function Status(params) {
  return http.request({
    url: '/stockList/status',
    method: 'POST',
    params,
  });
}

// 获取股票列表核心表指定详情
export function View(params) {
  return http.request({
    url: '/stockList/view',
    method: 'GET',
    params,
  });
}

// 导出股票列表核心表
export function Export(params) {
  jumpExport('/stockList/export', params);
}