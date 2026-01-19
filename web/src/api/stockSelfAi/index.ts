import { http, jumpExport } from '@/utils/http/axios';

// 获取ai 基本信息列表
export function List(params) {
  return http.request({
    url: '/stockSelfAi/list',
    method: 'get',
    params,
  });
}

// 删除/批量删除ai 基本信息
export function Delete(params) {
  return http.request({
    url: '/stockSelfAi/delete',
    method: 'POST',
    params,
  });
}

// 添加/编辑ai 基本信息
export function Edit(params) {
  return http.request({
    url: '/stockSelfAi/edit',
    method: 'POST',
    params,
  });
}

// 获取ai 基本信息指定详情
export function View(params) {
  return http.request({
    url: '/stockSelfAi/view',
    method: 'GET',
    params,
  });
}

// 导出ai 基本信息
export function Export(params) {
  jumpExport('/stockSelfAi/export', params);
}