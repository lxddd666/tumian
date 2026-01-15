import { http, jumpExport } from '@/utils/http/axios';

// 获取利润表 (Income Statement)列表
export function List(params) {
  return http.request({
    url: '/incomeStatement/list',
    method: 'get',
    params,
  });
}

// 删除/批量删除利润表 (Income Statement)
export function Delete(params) {
  return http.request({
    url: '/incomeStatement/delete',
    method: 'POST',
    params,
  });
}

// 添加/编辑利润表 (Income Statement)
export function Edit(params) {
  return http.request({
    url: '/incomeStatement/edit',
    method: 'POST',
    params,
  });
}

// 获取利润表 (Income Statement)指定详情
export function View(params) {
  return http.request({
    url: '/incomeStatement/view',
    method: 'GET',
    params,
  });
}

// 导出利润表 (Income Statement)
export function Export(params) {
  jumpExport('/incomeStatement/export', params);
}